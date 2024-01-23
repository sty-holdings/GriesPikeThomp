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
	"encoding/json"
	"fmt"
	"os"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	cv "GriesPikeThomp/shared-services/src/coreValidators"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// Configuration is a generic config file structure for application servers.
type Configuration struct {
	ConfigFileName         string
	SkeletonConfigFilename string                            `json:"skeleton_config_filename"`
	DebugModeOn            bool                              `json:"debug_mode_on"`
	Environment            string                            `json:"environment"`
	LogDirectory           string                            `json:"log_directory"`
	MaxThreads             int                               `json:"max_threads"`
	PIDDirectory           string                            `json:"pid_directory"`
	Extensions             map[string]map[string]interface{} `json:"extensions"`
}

// GenerateConfigFileSkeleton will output to the console a skeleton file with notes.
//
//	Customer Messages: None
//	Errors: ErrConfigFileMissing
//	Verifications: None
func GenerateConfigFileSkeleton(serverName, SkeletonConfigDirectory, SkeletonConfigFilenameNoSuffix string) {

	var (
		errorInfo                   cpi.ErrorInfo
		tSkeletonConfigData         []byte
		tSkeletonConfigNoteData     []byte
		tSkeletonConfigFilename     string
		tSkeletonConfigNoteFilename string
	)

	if serverName == rcv.VAL_EMPTY {
		cpi.PrintError(cpi.ErrMissingServerName, fmt.Sprintf("%v %v", rcv.TXT_SERVER_NAME, serverName))
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

// ValidateConfiguration -checks the values in the configuration file are valid. ValidateConfiguration doesn't
// test if the configuration file exists, readable, or parsable. LogDirectory, MaxThreads, and PIDDirectory will be
// set to '/var/log/nats-connect', 1, and '/var/run/nats-connect', respectively.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrDirectoryMissing, ErrMaxThreadsInvalid
//	Verifications: None
func ValidateConfiguration(config Configuration) (errorInfo cpi.ErrorInfo) {

	if cv.IsEnvironmentValid(config.Environment) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.Environment))
		return
	}
	if cv.DoesDirectoryExist(config.LogDirectory) == false {
		cpi.PrintError(cpi.ErrDirectoryMissing, fmt.Sprintf("%v%v - Default Set: %v", rcv.TXT_DIRECTORY, config.LogDirectory, DEFAULT_LOG_DIRECTORY))
		config.LogDirectory = DEFAULT_LOG_DIRECTORY
	}
	if config.MaxThreads < 1 || config.MaxThreads > THREAD_CAP {
		cpi.PrintError(cpi.ErrMaxThreadsInvalid, fmt.Sprintf("%v%v - Default Set: %v", rcv.TXT_MAX_THREADS, config.LogDirectory, DEFAULT_MAX_THREADS))
		config.MaxThreads = DEFAULT_MAX_THREADS
	}
	if cv.DoesDirectoryExist(config.PIDDirectory) == false {
		cpi.PrintError(cpi.ErrDirectoryMissing, fmt.Sprintf("%v%v - Default Set: %v", rcv.TXT_DIRECTORY, config.LogDirectory, DEFAULT_PID_DIRECTORY))
		config.PIDDirectory = DEFAULT_PID_DIRECTORY
	}

	return
}
