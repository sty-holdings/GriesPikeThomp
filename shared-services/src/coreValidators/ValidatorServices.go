// Package coreValidators
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc..
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
package coreValidators

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"

	"albert/constants"
	"albert/core/coreError"
)

// AreMapKeysPopulated - will test to make sure all map keys are set to anything other than nil or empty.
func AreMapKeysPopulated(myMap map[any]interface{}) bool {

	if IsMapPopulated(myMap) {
		for key, _ := range myMap {
			if key == nil || key == constants.EMPTY {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

// AreMapValuesPopulated - will test to make sure all map values are set to anything other than nil or empty.
func AreMapValuesPopulated(myMap map[any]interface{}) bool {

	if IsMapPopulated(myMap) {
		for _, value := range myMap {
			if value == nil || value == constants.EMPTY {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

// AreMapKeysValuesPopulated - check keys and value for missing values. Findings are constants.GOOD, constants.MISSING_VALUE,
// constants.MISSING_KEY, or constants.EMPTY_WORD.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: IsMapPopulated, AreMapKeysPopulated, AreMapValuesPopulated
func AreMapKeysValuesPopulated(myMap map[any]interface{}) (finding string) {

	if IsMapPopulated(myMap) {
		if AreMapKeysPopulated(myMap) {
			if AreMapValuesPopulated(myMap) {
				finding = constants.GOOD
			} else {
				finding = constants.MISSING_VALUE
			}
		} else {
			finding = constants.MISSING_KEY
		}
	} else {
		finding = constants.EMPTY_WORD
	}

	return
}

// CheckFileExistsAndReadable
func CheckFileExistsAndReadable(FQN, fileLabel string) (errorInfo coreError.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)
	errorInfo = coreError.GetFunctionInfo()
	errorInfo.AdditionalInfo = fmt.Sprintf("File: %v  Config File Label: %v", FQN, fileLabel)

	if FQN == constants.EMPTY {
		errorInfo.AdditionalInfo = fileLabel + constants.IS_EMPTY
		errorInfo.Error = coreError.ErrFileMissing
		coreError.PrintError(errorInfo)
	} else if DoesFileExist(FQN) == false {
		errorInfo.Error = coreError.ErrFileMissing
		coreError.PrintError(errorInfo)
	} else {
		if IsFileReadable(FQN) == false { // File is not readable
			errorInfo.Error = coreError.ErrFileUnreadable
			coreError.PrintError(errorInfo)
		}
	}

	return
}

// CheckFileValidJSON - reads the file and checks the contents
func CheckFileValidJSON(FQN, fileLabel string) (errorInfo coreError.ErrorInfo) {

	var (
		jsonData           []byte
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)
	errorInfo = coreError.GetFunctionInfo()
	errorInfo.AdditionalInfo = fmt.Sprintf("File: %v  Config File Label: %v", FQN, fileLabel)

	if jsonData, errorInfo.Error = os.ReadFile(FQN); errorInfo.Error != nil {
		errorInfo.Error = coreError.ErrFileUnreadable
		errorInfo.AdditionalInfo = fmt.Sprintf("FQN: %v File Label: %v", FQN, fileLabel)
		coreError.PrintError(errorInfo)
	} else {
		if _isJSON := IsJSONValid(jsonData); _isJSON == false {
			errorInfo.Error = coreError.ErrFileUnreadable
			errorInfo.AdditionalInfo = fmt.Sprintf("FQN: %v File Label: %v", FQN, fileLabel)
			coreError.PrintError(errorInfo)
		}
	}

	return
}

// DoesDirectoryExist
func DoesDirectoryExist(directoryName string) bool {

	return DoesFileExist(directoryName)
}

// DoesFileExist
func DoesFileExist(fileName string) bool {

	if _, err := os.Stat(fileName); err == nil {
		return true
	}

	return false
}

// IsDomainValid
func IsDomainValid(domain string) bool {

	if strings.ToLower(domain) == constants.LOCAL_HOST {
		return true
	} else {
		regex := regexp.MustCompile(`^(([a-zA-Z]{1})|([a-zA-Z]{1}[a-zA-Z]{1})|([a-zA-Z]{1}[0-9]{1})|([0-9]{1}[a-zA-Z]{1})|([a-zA-Z0-9][a-zA-Z0-9-_]{1,61}[a-zA-Z0-9]))\.([a-zA-Z]{2,6}|[a-zA-Z0-9-]{2,30}\.[a-zA-Z]{2,3})$`)
		if regex.MatchString(domain) {
			return true
		}
	}

	return false
}

// IsEnvironmentValid - is case-insensitive
func IsEnvironmentValid(environment string) bool {

	switch strings.ToUpper(environment) {
	case constants.ENVIRONMENT_LOCAL:
	case constants.ENVIRONMENT_DEVELOPMENT:
	case constants.ENVIRONMENT_PRODUCTION:
	default:
		return false
	}

	return true
}

// IsGinModeValid validates that the Gin HTTP framework mode is correctly set.
func IsGinModeValid(mode string) bool {

	switch strings.ToUpper(mode) {
	case constants.DEBUG_MODE:
	case constants.RELEASE_MODE:
	default:
		return false
	}

	return true
}

// IsFileReadable
func IsFileReadable(fileName string) bool {

	if _, err := os.OpenFile(fileName, os.O_RDONLY, 0644); err == nil {
		return true
	}

	return false
}

// IsIPValid - checks if the data provide is a valid IP address
func IsIPValid(ipAddress any) bool {

	// Checking if it is a valid IP addresses
	if IsIPv4Valid(ipAddress.(string)) || IsIPv6Valid(ipAddress.(string)) {
		return true
	}

	return false
}

// IsIPv4Valid - checks if the data provide is a valid IPv4 address
func IsIPv4Valid(ipAddress any) bool {

	var (
		tIPv4Regex = regexp.MustCompile(`^(((25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.|$)){4})`)
	)

	// Checking if it is a valid IPv4 addresses
	if tIPv4Regex.MatchString(ipAddress.(string)) {
		return true
	}

	return false
}

// IsIPv6Valid - checks if the data provide is a valid IPv6 address
func IsIPv6Valid(ipAddress any) bool {

	var (
		tIPv6Regex = regexp.MustCompile(`^(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))$`)
	)

	// Checking if it is a valid IPv4 addresses
	if tIPv6Regex.MatchString(ipAddress.(string)) {
		return true
	}

	return false
}

// IsJSONValid - checks if the data provide is valid JSON
func IsJSONValid(jsonIn []byte) bool {

	var (
		jsonString         map[string]interface{}
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	return json.Unmarshal(jsonIn, &jsonString) == nil
}

// IsMapPopulated - will determine if the map is populated.
func IsMapPopulated(myMap map[any]interface{}) bool {

	if len(myMap) > 0 {
		return true
	}

	return false
}

// IsMessagePrefixValid - is case-insensitive
func IsMessagePrefixValid(messagePrefix string) bool {

	switch strings.ToUpper(messagePrefix) {
	case constants.MESSAGE_PREFIX_SAVUPPROD:
	case constants.MESSAGE_PREFIX_SAVUPDEV:
	case constants.MESSAGE_PREFIX_SAVUPLOCAL:
	default:
		return false
	}

	return true
}

// IsPeriodValid
func IsPeriodValid(period string) bool {

	switch strings.ToUpper(period) {
	case constants.YEAR:
	case constants.MONTH:
	case constants.DAY:
	default:
		return false
	}

	return true
}

// This will set the connection values so GetConnection can be executed.
func IsPostgresSSLModeValid(sslMode string) bool {

	switch sslMode {
	case constants.POSTGRES_SSL_MODE_ALLOW:
	case constants.POSTGRES_SSL_MODE_DISABLE:
	case constants.POSTGRES_SSL_MODE_PREFER:
	case constants.POSTGRES_SSL_MODE_REQUIRED:
	default:
		return false
	}

	return true
}

// IsUserRegisterTypedValid
func IsUserRegisterTypedValid(period string) bool {

	switch strings.ToUpper(period) {
	case constants.COLLECTION_USER_TO_DO_LIST:
	case constants.COLLECTION_USER_GOALS:
	default:
		return false
	}

	return true
}

// IsURLValid
func IsURLValid(URL string) bool {

	if _, err := url.ParseRequestURI(URL); err == nil {
		return true
	}

	return false
}

// IsUUIDValid
func IsUUIDValid(uuid string) bool {

	r := regexp.MustCompile("^[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9aAbB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}$")
	return r.MatchString(uuid)
}

// ValidateAuthenticatorService - Firebase is not support at this time
func ValidateAuthenticatorService(authenticatorService string) (errorInfo coreError.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	switch strings.ToUpper(authenticatorService) {
	case constants.AUTH_COGNITO:
	case constants.AUTH_FIREBASE:
		fallthrough // ToDo This is because AUTH_FIREBASE is not supported right now
	default:
		errorInfo.Error = errors.New(fmt.Sprintf("The supplied authenticator service is not supported! Authenticator Service: %v (Authenticator Service is case insensitive)", authenticatorService))
		if authenticatorService == constants.EMPTY {
			errorInfo.AdditionalInfo = "Authenticator Service parameter is empty"
		} else {
			errorInfo.AdditionalInfo = "Authenticator Service: " + authenticatorService
		}
	}

	return
}

// ValidateDirectory validates that the directory value is not empty and the value exists on the file system
func ValidateDirectory(directory string) (errorInfo coreError.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	if directory == constants.EMPTY {
		errorInfo.AdditionalInfo = "Directory parameter is empty"
		errorInfo.Error = coreError.ErrRequiredArgumentMissing
		coreError.PrintError(errorInfo)
	} else if DoesDirectoryExist(directory) == false {
		errorInfo.AdditionalInfo = "Directory: " + directory
		errorInfo.Error = coreError.ErrDirectoryMissing
		coreError.PrintError(errorInfo)
	}

	return
}

// ValidateTransferMethod
func ValidateTransferMethod(transferMethod string) (errorInfo coreError.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	switch strings.ToUpper(transferMethod) {
	case constants.TRANFER_STRIPE:
	case constants.TRANFER_WIRE:
	case constants.TRANFER_CHECK:
	case constants.TRANFER_ZELLE:
	default:
		errorInfo.Error = coreError.ErrTransferMethodInvalid
		if transferMethod == constants.EMPTY {
			errorInfo.AdditionalInfo = "Transfer Method parameter is empty"
		} else {
			errorInfo.AdditionalInfo = "Transfer Method: " + transferMethod
		}
	}

	return
}
