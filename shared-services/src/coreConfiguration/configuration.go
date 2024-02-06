// Package sharedServices
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, Inc
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
	"encoding/json"
	"fmt"
	"os"
	"strings"

	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// Configuration is a generic config file structure for application servers.
type BaseConfiguration struct {
	ConfigFQN         string
	SkeletonConfigFQD string                 `json:"skeleton_config_fqd"`
	DebugModeOn       bool                   `json:"debug_mode_on"`
	Environment       string                 `json:"environment"`
	LogDirectory      string                 `json:"log_directory"`
	MaxThreads        int                    `json:"max_threads"`
	PIDDirectory      string                 `json:"pid_directory"`
	Extensions        []BaseConfigExtensions `json:"extensions"`
}

type BaseConfigExtensions struct {
	Name           string `json:"name"`
	ConfigFilename string `json:"config_filename"`
}

// GenerateConfigFileSkeleton will output to the console a skeleton file with notes.
//
//	Customer Messages: None
//	Errors: ErrConfigFileMissing
//	Verifications: None
func GenerateConfigFileSkeleton(serverName, SkeletonConfigFQD string) (errorInfo cpi.ErrorInfo) {

	var (
		tSkeletonConfigData         []byte
		tSkeletonConfigNoteData     []byte
		tSkeletonConfigFilename     string
		tSkeletonConfigNoteFilename string
	)

	if serverName == rcv.VAL_EMPTY {
		cpi.PrintError(cpi.ErrMissingServerName, fmt.Sprintf("%v %v", rcv.TXT_SERVER_NAME, serverName))
		return
	}
	if SkeletonConfigFQD == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrFileMissing, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, SkeletonConfigFQD))
		return
	}
	tSkeletonConfigFilename = fmt.Sprintf("%v%v", SkeletonConfigFQD, DEFAULT_SKELETON_CONFIG_FILENAME)
	tSkeletonConfigNoteFilename = fmt.Sprintf("%v%v", SkeletonConfigFQD, DEFAULT_SKELETON_CONFIG_NOTE_FILENAME)

	if tSkeletonConfigData, errorInfo.Error = os.ReadFile(tSkeletonConfigFilename); errorInfo.Error != nil {
		cpi.PrintError(cpi.ErrFileUnreadable, fmt.Sprintf("%v %v", rcv.TXT_FILENAME, tSkeletonConfigFilename))
		return
	}

	if tSkeletonConfigNoteData, errorInfo.Error = os.ReadFile(tSkeletonConfigNoteFilename); errorInfo.Error != nil {
		cpi.PrintError(cpi.ErrFileUnreadable, fmt.Sprintf("%v %v", rcv.TXT_FILENAME, tSkeletonConfigNoteFilename))
		return
	}

	fmt.Println("\nWhen '-g' is used all other program arguments are ignored.")
	fmt.Printf("\n%v Config file Skeleton: \n%v\n", serverName, string(tSkeletonConfigData))
	fmt.Println()
	fmt.Printf("%v\n", string(tSkeletonConfigNoteData))
	fmt.Println()

	return
}

// ProcessBaseConfigFile - handles the base configuration file.
//
//	Customer Messages: None
//	Errors: errors returned from ReadConfigFile, ErrJSONInvalid
//	Verifications: None
func ProcessBaseConfigFile(configFileFQN string) (config BaseConfiguration, errorInfo cpi.ErrorInfo) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v %v", rcv.TXT_FILENAME, configFileFQN)
		tConfigData     []byte
	)

	if tConfigData, errorInfo = ReadConfigFile(configFileFQN); errorInfo.Error != nil {
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &config); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	config.ConfigFQN = configFileFQN
	config.Environment = strings.ToLower(config.Environment)

	return
}

// ReadConfigFile opens the provide file, unmarshal the file and returns the Configuration object.
//
//	Customer Messages: None
//	Errors: ErrConfigFileMissing, ErrJSONInvalid
//	Verifications: None
func ReadConfigFile(configFileFQN string) (configData []byte, errorInfo cpi.ErrorInfo) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v %v", rcv.TXT_FILENAME, configFileFQN)
	)

	if configData, errorInfo.Error = os.ReadFile(configFileFQN); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(cpi.ErrConfigFileMissing, tAdditionalInfo)
	}

	return
}

// ValidateConfiguration - checks the values in the configuration file are valid. ValidateConfiguration doesn't
// test if the configuration file exists, readable, or parsable. Defaults for LogDirectory, MaxThreads, and PIDDirectory
// are '/var/log/nats-connect', 1, and '/var/run/nats-connect', respectively.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrDirectoryMissing, ErrMaxThreadsInvalid
//	Verifications: None
func ValidateConfiguration(config BaseConfiguration) (errorInfo cpi.ErrorInfo) {

	if chv.IsEnvironmentValid(config.Environment) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.Environment))
		return
	}
	if chv.DoesDirectoryExist(config.SkeletonConfigFQD) == false {
		cpi.PrintError(cpi.ErrDirectoryMissing, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.SkeletonConfigFQD))
		config.LogDirectory = DEFAULT_LOG_DIRECTORY
	}
	if chv.DoesDirectoryExist(config.LogDirectory) == false {
		cpi.PrintError(cpi.ErrDirectoryMissing, fmt.Sprintf("%v%v - Default Set: %v", rcv.TXT_DIRECTORY, config.LogDirectory, DEFAULT_LOG_DIRECTORY))
		config.LogDirectory = DEFAULT_LOG_DIRECTORY
	}
	if config.MaxThreads < 1 || config.MaxThreads > THREAD_CAP {
		cpi.PrintError(cpi.ErrMaxThreadsInvalid, fmt.Sprintf("%v%v - Default Set: %v", rcv.TXT_MAX_THREADS, config.LogDirectory, DEFAULT_MAX_THREADS))
		config.MaxThreads = DEFAULT_MAX_THREADS
	}
	if chv.DoesDirectoryExist(config.PIDDirectory) == false {
		cpi.PrintError(cpi.ErrDirectoryMissing, fmt.Sprintf("%v%v - Default Set: %v", rcv.TXT_DIRECTORY, config.LogDirectory, DEFAULT_PID_DIRECTORY))
		config.PIDDirectory = DEFAULT_PID_DIRECTORY
	}

	return
}
