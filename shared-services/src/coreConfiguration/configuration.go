// Package shared_services
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
package shared_services

import (
	"encoding/json"
	"fmt"
	"os"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// Configuration is a generic config file structure for application servers.
type Configuration struct {
	ConfigFileName         string
	SkeletonConfigFilename string         `json:"skeleton_config_filename"`
	DebugModeOn            bool           `json:"debug_mode_on"`
	Environment            string         `json:"environment"`
	LogDirectory           string         `json:"log_directory"`
	PIDDirectory           string         `json:"pid_directory"`
	Extensions             map[string]any `json:"extensions"`
}

// GenerateConfigFileSkeleton will output to the console a skeleton file with notes.
//
//	Customer Messages: None
//	Errors: ErrConfigFileMissing
//	Verifications: None
func GenerateConfigFileSkeleton(serverName, SkeletonConfigDirectory, SkeletonConfigFilenameNoSuffix string) (errorInfo cpi.ErrorInfo) {

	var (
		tSkeletonConfigData         []byte
		tSkeletonConfigNoteData     []byte
		tSkeletonConfigFilename     string
		tSkeletonConfigNoteFilename string
	)

	if serverName == rcv.VAL_EMPTY {
		cpi.PrintError(cpi.ErrMissingServerName, fmt.Sprintf("%v %v", rcv.TXT_SERVER_NAME, serverName), rcv.MODE_OUTPUT_DISPLAY)
		return
	}
	if SkeletonConfigDirectory == rcv.VAL_EMPTY || SkeletonConfigFilenameNoSuffix == rcv.VAL_EMPTY {
		tSkeletonConfigFilename = fmt.Sprintf("%v%v.json", DEFAULT_SKELETON_CONFIG_DIRECTORY, DEFAULT_SKELETON_CONFIG_FILENAME_NO_SUFFIX)
		tSkeletonConfigNoteFilename = fmt.Sprintf("%v%v.txt", DEFAULT_SKELETON_CONFIG_DIRECTORY, DEFAULT_SKELETON_CONFIG_FILENAME_NO_SUFFIX)
	} else {
		tSkeletonConfigFilename = fmt.Sprintf("%v%v.json", SkeletonConfigDirectory, SkeletonConfigFilenameNoSuffix)
		tSkeletonConfigNoteFilename = fmt.Sprintf("%v%v.txt", SkeletonConfigDirectory, SkeletonConfigFilenameNoSuffix)
	}

	if tSkeletonConfigData, errorInfo.Error = os.ReadFile(tSkeletonConfigFilename); errorInfo.Error != nil {
		cpi.PrintError(cpi.ErrFileUnreadable, fmt.Sprintf("%v %v", rcv.TXT_FILENAME, tSkeletonConfigFilename), rcv.MODE_OUTPUT_DISPLAY)
		return
	}

	if tSkeletonConfigNoteData, errorInfo.Error = os.ReadFile(tSkeletonConfigNoteFilename); errorInfo.Error != nil {
		cpi.PrintError(cpi.ErrFileUnreadable, fmt.Sprintf("%v %v", rcv.TXT_FILENAME, tSkeletonConfigNoteFilename), rcv.MODE_OUTPUT_DISPLAY)
		return
	}

	fmt.Println("\nWhen '-g' is used all other program arguments are ignored.")
	fmt.Printf("\n%v Config file Skeleton: \n%v\n", serverName, string(tSkeletonConfigData))
	fmt.Println()
	fmt.Printf("%v\n", string(tSkeletonConfigNoteData))
	fmt.Println()

	return
}

// ReadAndParseConfigFile opens the provide file, unmarshal the file and returns the Configuration object.
//
//	Customer Messages: None
//	Errors: ErrConfigFileMissing, ErrJSONInvalid
//	Verifications: None
func ReadAndParseConfigFile(configFileFQN string) (config Configuration, errorInfo cpi.ErrorInfo) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v %v", rcv.TXT_FILENAME, configFileFQN)
		tConfigFile     []byte
	)

	if tConfigFile, errorInfo.Error = os.ReadFile(configFileFQN); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(cpi.ErrConfigFileMissing, tAdditionalInfo)
		return
	}
	if errorInfo.Error = json.Unmarshal(tConfigFile, &config); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(cpi.ErrJSONInvalid, tAdditionalInfo)
		return
	}

	fileStat, _ := os.Stat(configFileFQN)
	config.ConfigFileName = fileStat.Name()

	return
}

// validateOptions - will validate option values and where defaults are applicable set them.
//
//	Validated Options: Authenticator Service, AWSInfoFQN, Environment, Firebase Project Id,
//	Firebase Credentials FQN, Log Directory, Message Prefix, NATS Creds FQN, NATS URL,
//	PID Directory, Plaid Key FQN, Private Key FQN, Sendgrid Key FQN, Stripe Key FQN, and TLS
// func validateOptions(opts Options) (errorInfos []coreError.ErrorInfo) {
//
// 	var (
// 		errorInfo          coreError.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	// AuthenticatorService
// 	if errorInfo = coreValidators.ValidateAuthenticatorService(opts.AuthenticatorService); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	} else {
// 		opts.AuthenticatorService = strings.ToUpper(opts.AuthenticatorService)
// 	}
//
// 	// AWSInfoFQN
// 	if errorInfo = coreValidators.CheckFileExistsAndReadable(opts.AWSInfoFQN, rcv.FN_AWS_INFO_FQN); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// Environment
// 	if coreValidators.IsEnvironmentValid(opts.Environment) == false {
// 		errorInfo.Error = errors.New(fmt.Sprintf("The Environment: '%v' is invalid.", opts.Environment))
// 		errorInfos = append(errorInfos, errorInfo)
// 	} else {
// 		opts.Environment = strings.ToUpper(opts.Environment)
// 	}
//
// 	// FirebaseProjectId
// 	if opts.FirebaseProjectId == "" {
// 		errorInfo.Error = coreError.ErrFirebaseProjectMissing
// 		log.Println(errorInfo.Error.Error())
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// FirebaseCredentialsFQN
// 	if errorInfo = coreValidators.CheckFileExistsAndReadable(opts.FirebaseCredentialsFQN, constants.FIREBASE_CREDENTIALS_FQN); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
// 	if errorInfo = coreValidators.CheckFileValidJSON(opts.FirebaseCredentialsFQN, constants.FIREBASE_CREDENTIALS_FQN); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// GCPCredentialsFQN
// 	if errorInfo = coreValidators.CheckFileExistsAndReadable(opts.GCPCredentialsFQN, constants.GCP_CREDENTIALS_FQN); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
// 	if errorInfo = coreValidators.CheckFileValidJSON(opts.GCPCredentialsFQN, constants.GCP_CREDENTIALS_FQN); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// LogDirectory
// 	if errorInfo = coreValidators.ValidateDirectory(opts.LogDirectory); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// MessagePrefix
// 	if coreValidators.IsMessagePrefixValid(opts.MessagePrefix) == false {
// 		errorInfo.Error = errors.New(fmt.Sprintf("The Message Prefix: '%v' is invalid.", opts.MessagePrefix))
// 		log.Println(errorInfo.Error.Error())
// 		errorInfos = append(errorInfos, errorInfo)
// 	} else {
// 		opts.MessagePrefix = strings.ToUpper(opts.MessagePrefix)
// 	}
//
// 	// NATSCredsFQN
// 	if errorInfo = coreValidators.CheckFileExistsAndReadable(opts.NATSCredsFQN, constants.NATS_CREDENTIALS_FQN); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// NATSURL
// 	if opts.NATSURL == "" {
// 		errorInfo.Error = errors.New(fmt.Sprintf("The NATSURL is not populated. It must be provided so the SavUp server can connection to NATS. %v%v", opts.NATSURL, constants.ENDING_EXECUTION))
// 		log.Println(errorInfo.Error.Error())
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// PIDDirectory
// 	if errorInfo = coreValidators.ValidateDirectory(opts.PIDDirectory); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// SendGridKeyFQN
// 	if errorInfo = coreValidators.CheckFileExistsAndReadable(opts.SendGridKeyFQN, constants.SENDGRID_KEY_FQN); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
// 	if errorInfo = coreValidators.CheckFileValidJSON(opts.SendGridKeyFQN, constants.SENDGRID_KEY_FQN); errorInfo.Error != nil {
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// SendGridTemplateIds
// 	if len(opts.SendGridTemplateIds.Ids) != coreSendGrid.TEMPLATE_ID_COUNT {
// 		errorInfo.Error = coreError.ErrRequiredArgumentMissing
// 		errorInfos = append(errorInfos, errorInfo)
// 	}
//
// 	// TLS
// 	if opts.TLS.TLSCert == constants.EMPTY || opts.TLS.TLSKey == constants.EMPTY || opts.TLS.TLSCABundle == constants.EMPTY {
// 		log.Println("TLS is not populated in the configuration file. You will not be able to connect to a NATS server using TLS with verify set to true.")
// 	} else {
// 		if errorInfo = coreValidators.CheckFileExistsAndReadable(opts.TLS.TLSCert, constants.TLS_CERTIFICATE_FQN); errorInfo.Error != nil {
// 			errorInfos = append(errorInfos, errorInfo)
// 		} else {
// 			log.Println("TLS Cert file exists and is readable.")
// 		}
// 		if errorInfo = coreValidators.CheckFileExistsAndReadable(opts.TLS.TLSKey, constants.TLS_PRIVATE_KEY_FQN); errorInfo.Error != nil {
// 			errorInfos = append(errorInfos, errorInfo)
// 		} else {
// 			log.Println("TLS Private key file exists and is readable.")
// 		}
// 		if errorInfo = coreValidators.CheckFileExistsAndReadable(opts.TLS.TLSCABundle, constants.TLS_PRIVATE_KEY_FQN); errorInfo.Error != nil {
// 			errorInfos = append(errorInfos, errorInfo)
// 		} else {
// 			log.Println("TLS Private key file exists and is readable.")
// 		}
// 	}
//
// 	if len(errorInfos) > 0 {
// 		GenerateConfigFileSkeleton()
// 	}
//
// 	return
// }
