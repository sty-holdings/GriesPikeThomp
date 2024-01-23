// Package sharedServices
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc.
	Use is subject to license terms.

	Unauthorized copying of this file, via any medium is strictly prohibited.

	Proprietary and confidential

	Written by <Replace with FULL_NAME> / syacko
	STY-Holdings, Inc.
	support@sty-holdings.com
	www.sty-holdings.com

	01-2024
	USA

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package sharedServices

import (
	"fmt"
	"os"
	"regexp"
	"strings"

	ch "GriesPikeThomp/shared-services/src/coreHelpers"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// AreMapKeysPopulated - will test to make sure all map keys are set to anything other than nil or empty.
// func AreMapKeysPopulated(myMap map[any]interface{}) bool {
//
// 	if IsMapPopulated(myMap) {
// 		for key, _ := range myMap {
// 			if key == nil || key == rcv.TXT_EMPTY {
// 				return false
// 			}
// 		}
// 	} else {
// 		return false
// 	}
//
// 	return true
// }

// AreMapValuesPopulated - will test to make sure all map values are set to anything other than nil or empty.
// func AreMapValuesPopulated(myMap map[any]interface{}) bool {
//
// 	if IsMapPopulated(myMap) {
// 		for _, value := range myMap {
// 			if value == nil || value == rcv.VAL_EMPTY {
// 				return false
// 			}
// 		}
// 	} else {
// 		return false
// 	}
//
// 	return true
// }

// AreMapKeysValuesPopulated - check keys and value for missing values. Findings are rcv.GOOD, rcv.MISSING_VALUE,
// rcv.MISSING_KEY, or rcv.VAL_EMPTY_WORD.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: IsMapPopulated, AreMapKeysPopulated, AreMapValuesPopulated
// func AreMapKeysValuesPopulated(myMap map[any]interface{}) (finding string) {
//
// 	if IsMapPopulated(myMap) {
// 		if AreMapKeysPopulated(myMap) {
// 			if AreMapValuesPopulated(myMap) {
// 				finding = rcv.TXT_GOOD
// 			} else {
// 				finding = rcv.TXT_MISSING_VALUE
// 			}
// 		} else {
// 			finding = rcv.TXT_MISSING_KEY
// 		}
// 	} else {
// 		finding = rcv.TXT_EMPTY
// 	}
//
// 	return
// }

// DoesFileExistsAndReadable - works on any file. If the filename is not fully qualified
// the working directory will be prepended to the filename.
//
//	Customer Messages: None
//	Errors: ErrFileMissing, ErrFileUnreadable
//	Verifications: None
func DoesFileExistsAndReadable(filename, fileLabel string) (errorInfo cpi.ErrorInfo) {

	var (
		fqn = ch.PrependWorkingDirectory(filename)
	)

	if fileLabel == rcv.VAL_EMPTY {
		fileLabel = rcv.TXT_NO_LABEL_PROVIDED
	}
	errorInfo.AdditionalInfo = fmt.Sprintf("File: %v  Config File Label: %v", filename, fileLabel)

	if filename == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrFileMissing, errorInfo.AdditionalInfo)
		return
	}
	if DoesFileExist(fqn) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrFileMissing, errorInfo.AdditionalInfo)
		return
	}
	if IsFileReadable(fqn) == false { // File is not readable
		errorInfo = cpi.NewErrorInfo(cpi.ErrFileUnreadable, errorInfo.AdditionalInfo)
	}

	return
}

// CheckFileValidJSON - reads the file and checks the contents
// func CheckFileValidJSON(FQN, fileLabel string) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		jsonData           []byte
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
// 	errorInfo = cpi.GetFunctionInfo()
// 	errorInfo.AdditionalInfo = fmt.Sprintf("File: %v  Config File Label: %v", FQN, fileLabel)
//
// 	if jsonData, errorInfo.Error = os.ReadFile(FQN); errorInfo.Error != nil {
// 		errorInfo.Error = cpi.ErrFileUnreadable
// 		errorInfo.AdditionalInfo = fmt.Sprintf("FQN: %v File Label: %v", FQN, fileLabel)
// 		cpi.PrintError(errorInfo)
// 	} else {
// 		if _isJSON := IsJSONValid(jsonData); _isJSON == false {
// 			errorInfo.Error = cpi.ErrFileUnreadable
// 			errorInfo.AdditionalInfo = fmt.Sprintf("FQN: %v File Label: %v", FQN, fileLabel)
// 			cpi.PrintError(errorInfo)
// 		}
// 	}
//
// 	return
// }

// DoesDirectoryExist - checks is the directory exists
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DoesDirectoryExist(directoryName string) bool {

	return DoesFileExist(directoryName)
}

// DoesFileExist - does the value exist on the file system
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func DoesFileExist(fileName string) bool {

	if _, err := os.Stat(fileName); err == nil {
		return true
	}

	return false
}

// IsDomainValid - checks if domain naming is followed
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsDomainValid(domain string) bool {

	if strings.ToLower(domain) == rcv.LOCAL_HOST {
		return true
	} else {
		regex := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
		if regex.MatchString(domain) {
			return true
		}
	}

	return false
}

// IsEnvironmentValid - checks that the value is valid. This function input is case-insensitive
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsEnvironmentValid(environment string) bool {

	switch strings.ToLower(environment) {
	case rcv.ENVIRONMENT_LOCAL:
	case rcv.ENVIRONMENT_DEVELOPMENT:
	case rcv.ENVIRONMENT_PRODUCTION:
	default:
		return false
	}

	return true
}

// IsGinModeValid validates that the Gin HTTP framework mode is correctly set.
// func IsGinModeValid(mode string) bool {
//
// 	switch strings.ToUpper(mode) {
// 	case rcv.DEBUG_MODE:
// 	case rcv.RELEASE_MODE:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// IsFileReadable - tries to open the file using 0644 permissions
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func IsFileReadable(fileName string) bool {

	if _, err := os.OpenFile(fileName, os.O_RDONLY, 0644); err == nil {
		return true
	}

	return false
}

// IsIPAddressValid - checks if the data provide is a valid IP address
// func IsIPAddressValid(ipAddress any) bool {
//
// 	// Checking if it is a valid IP addresses
// 	if IsIPv4Valid(ipAddress.(string)) || IsIPv6Valid(ipAddress.(string)) {
// 		return true
// 	}
//
// 	return false
// }

// IsIPv4Valid - checks if the data provide is a valid IPv4 address
// func IsIPv4Valid(ipAddress any) bool {
//
// 	var (
// 		tIPv4Regex = regexp.MustCompile(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`)
// 	)
//
// 	// Checking if it is a valid IPv4 addresses
// 	if tIPv4Regex.MatchString(ipAddress.(string)) {
// 		return true
// 	}
//
// 	return false
// }

// IsIPv6Valid - checks if the data provide is a valid IPv6 address
// func IsIPv6Valid(ipAddress any) bool {
//
// 	var (
// 		tIPv6Regex = regexp.MustCompile(`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`)
// 	)
//
// 	// Checking if it is a valid IPv4 addresses
// 	if tIPv6Regex.MatchString(ipAddress.(string)) {
// 		return true
// 	}
//
// 	return false
// }

// IsJSONValid - checks if the data provide is valid JSON
// func IsJSONValid(jsonIn []byte) bool {
//
// 	var (
// 		jsonString         map[string]interface{}
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	return json.Unmarshal(jsonIn, &jsonString) == nil
// }

// IsMapPopulated - will determine if the map is populated.
// func IsMapPopulated(myMap map[any]interface{}) bool {
//
// 	if len(myMap) > 0 {
// 		return true
// 	}
//
// 	return false
// }

// IsMessagePrefixValid - is case-insensitive
// func IsMessagePrefixValid(messagePrefix string) bool {
//
// 	switch strings.ToUpper(messagePrefix) {
// 	case rcv.MESSAGE_PREFIX_SAVUPPROD:
// 	case rcv.MESSAGE_PREFIX_SAVUPDEV:
// 	case rcv.MESSAGE_PREFIX_SAVUPLOCAL:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// IsPeriodValid
// func IsPeriodValid(period string) bool {
//
// 	switch strings.ToUpper(period) {
// 	case rcv.YEAR:
// 	case rcv.MONTH:
// 	case rcv.DAY:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// This will set the connection values so GetConnection can be executed.
// func IsPostgresSSLModeValid(sslMode string) bool {
//
// 	switch sslMode {
// 	case rcv.POSTGRES_SSL_MODE_ALLOW:
// 	case rcv.POSTGRES_SSL_MODE_DISABLE:
// 	case rcv.POSTGRES_SSL_MODE_PREFER:
// 	case rcv.POSTGRES_SSL_MODE_REQUIRED:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// IsUserRegisterTypedValid
// func IsUserRegisterTypedValid(period string) bool {
//
// 	switch strings.ToUpper(period) {
// 	case rcv.COLLECTION_USER_TO_DO_LIST:
// 	case rcv.COLLECTION_USER_GOALS:
// 	default:
// 		return false
// 	}
//
// 	return true
// }

// IsURLValid
// func IsURLValid(URL string) bool {
//
// 	if _, err := url.ParseRequestURI(URL); err == nil {
// 		return true
// 	}
//
// 	return false
// }

// IsUUIDValid
// func IsUUIDValid(uuid string) bool {
//
// 	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
// 	return r.MatchString(uuid)
// }

// ValidateAuthenticatorService - Firebase is not support at this time
// func ValidateAuthenticatorService(authenticatorService string) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	switch strings.ToUpper(authenticatorService) {
// 	case rcv.AUTH_COGNITO:
// 	case rcv.AUTH_FIREBASE:
// 		fallthrough // ToDo This is because AUTH_FIREBASE is not supported right now
// 	default:
// 		errorInfo.Error = errors.New(fmt.Sprintf("The supplied authenticator service is not supported! Authenticator Service: %v (Authenticator Service is case insensitive)", authenticatorService))
// 		if authenticatorService == rcv.VAL_EMPTY {
// 			errorInfo.AdditionalInfo = "Authenticator Service parameter is empty"
// 		} else {
// 			errorInfo.AdditionalInfo = "Authenticator Service: " + authenticatorService
// 		}
// 	}
//
// 	return
// }

// ValidateDirectory - validates that the directory value is not empty and the value exists on the file system
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func ValidateDirectory(directory string) (errorInfo cpi.ErrorInfo) {

	if directory == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrRequiredArgumentMissing, rcv.TXT_DIRECTORY_PARAM_EMPTY)
		return
	}
	if DoesDirectoryExist(directory) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, directory))
	}

	return
}

// ValidateTransferMethod
// func ValidateTransferMethod(transferMethod string) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	switch strings.ToUpper(transferMethod) {
// 	case rcv.TRANFER_STRIPE:
// 	case rcv.TRANFER_WIRE:
// 	case rcv.TRANFER_CHECK:
// 	case rcv.TRANFER_ZELLE:
// 	default:
// 		errorInfo.Error = cpi.ErrTransferMethodInvalid
// 		if transferMethod == rcv.VAL_EMPTY {
// 			errorInfo.AdditionalInfo = "Transfer Method parameter is empty"
// 		} else {
// 			errorInfo.AdditionalInfo = "Transfer Method: " + transferMethod
// 		}
// 	}
//
// 	return
// }
