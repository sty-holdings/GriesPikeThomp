// Package loadExtensions
/*
This is code for STY-Holdings NATS Connect

RESTRICTIONS:
	None

NOTES:

	None

COPYRIGHT:

	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package loadExtensions

import (
	"encoding/json"
	"fmt"

	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	config "github.com/sty-holdings/sty-shared/v2024/configuration"
	hv "github.com/sty-holdings/sty-shared/v2024/helpersValidators"
	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

const (
// Add Constants to the constants.go file
)

type ExtensionConfiguration struct {
	MessageEnvironment string               `json:"message_environment"`
	NATSConfig         ns.NATSConfiguration `json:"nats_config"`
	RequestedThreads   uint                 `json:"requested_threads"`
	SubjectRegistry    []SubjectInfo        `json:"subject_registry"`
}

type SubjectInfo struct {
	Subject string `json:"subject"`
}

var (
// Add Variables here for the file (Remember, they are global)
)

// Private functions

// LoadExtensionConfig - reads, validates, and returns
//
//	Customer Messages: None
//	Errors: error returned by ReadConfigFile or validateConfiguration
//	Verifications: validateConfiguration
func LoadExtensionConfig(
	extConfig config.BaseConfigExtensions,
) (
	extension ExtensionConfiguration,
	errorInfo pi.ErrorInfo,
) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v%v", ctv.TXT_FILENAME, extConfig.ConfigFilename)
		tConfigData     []byte
	)

	if tConfigData, errorInfo = config.ReadConfigFile(hv.PrependWorkingDirectory(extConfig.ConfigFilename)); errorInfo.Error != nil {
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &extension); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	if errorInfo = validateConfiguration(extConfig.Name, extension); errorInfo.Error != nil {
		return
	}

	return
}

// validateConfiguration - checks the NATS service configuration is valid.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrMessageNamespaceInvalid, ErrDomainInvalid, error returned from DoesFileExistsAndReadable, ErrSubjectsMissing
//	Verifications: None
func validateConfiguration(
	name string,
	config ExtensionConfiguration,
) (errorInfo pi.ErrorInfo) {

	if errorInfo = hv.DoesFileExistsAndReadable(config.NATSConfig.NATSCredentialsFilename, ctv.TXT_FILENAME); errorInfo.Error != nil {
		pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, config.NATSConfig.NATSCredentialsFilename))
		return
	}
	if hv.IsEnvironmentValid(config.MessageEnvironment) == false {
		errorInfo = pi.NewErrorInfo(pi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", ctv.TXT_EVIRONMENT, config.MessageEnvironment))
		return
	}
	if config.NATSConfig.NATSPort == ctv.VAL_ZERO {
		errorInfo = pi.NewErrorInfo(pi.ErrNatsPortInvalid, fmt.Sprintf("%v%v", ctv.TXT_NATS_PORT, config.NATSConfig.NATSPort))
		return
	}
	if config.NATSConfig.NATSTLSInfo.TLSCert != ctv.VAL_EMPTY && config.NATSConfig.NATSTLSInfo.TLSPrivateKey != ctv.VAL_EMPTY && config.NATSConfig.NATSTLSInfo.TLSCABundle != ctv.VAL_EMPTY {
		if errorInfo = hv.DoesFileExistsAndReadable(config.NATSConfig.NATSTLSInfo.TLSCert, ctv.TXT_FILENAME); errorInfo.Error != nil {
			pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, config.NATSConfig.NATSTLSInfo.TLSCert))
			return
		}
		if errorInfo = hv.DoesFileExistsAndReadable(config.NATSConfig.NATSTLSInfo.TLSPrivateKey, ctv.TXT_FILENAME); errorInfo.Error != nil {
			pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, config.NATSConfig.NATSTLSInfo.TLSPrivateKey))
			return
		}
		if errorInfo = hv.DoesFileExistsAndReadable(config.NATSConfig.NATSTLSInfo.TLSCABundle, ctv.TXT_FILENAME); errorInfo.Error != nil {
			pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, config.NATSConfig.NATSTLSInfo.TLSCABundle))
			return
		}
	}
	if hv.IsDomainValid(config.NATSConfig.NATSURL) == false {
		errorInfo = pi.NewErrorInfo(pi.ErrDomainInvalid, fmt.Sprintf("%v%v", ctv.TXT_EVIRONMENT, config.NATSConfig.NATSURL))
		return
	}
	// ToDo Add support for requested threads
	if len(config.SubjectRegistry) == ctv.VAL_ZERO {
		if name == ctv.NC_INTERNAL {
			return
		}
		errorInfo = pi.NewErrorInfo(pi.ErrSubjectsMissing, ctv.VAL_EMPTY)
	}

	return
}
