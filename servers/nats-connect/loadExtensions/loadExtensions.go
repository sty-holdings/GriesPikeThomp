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
	NATSConfig       ns.NATSConfiguration `json:"nats_config"`
	RequestedThreads uint                 `json:"requested_threads"`
	SubjectRegistry  []SubjectInfo        `json:"subject_registry"`
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
	environment string,
	serverInstanceNumber string,
	extConfigPtr *BaseConfigExtensions,
) (
	extension ExtensionConfiguration,
	errorInfo pi.ErrorInfo,
) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v%v", ctv.TXT_FILENAME, extConfigPtr.ConfigFilename)
		tConfigData     []byte
		tExtension      ExtensionConfiguration
	)

	if tConfigData, errorInfo = config.ReadConfigFile(hv.PrependWorkingDirectory(extConfigPtr.ConfigFilename)); errorInfo.Error != nil {
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &tExtension); errorInfo.Error != nil {
		errorInfo = pi.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	if errorInfo = validateConfiguration(environment, serverInstanceNumber, extConfigPtr.Name, &tExtension); errorInfo.Error != nil {
		return
	}

	extension = tExtension
	return
}

// validateConfiguration - checks the NATS service configuration is valid.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrMessageNamespaceInvalid, ErrDomainInvalid, error returned from DoesFileExistsAndReadable, ErrSubjectsMissing
//	Verifications: None
func validateConfiguration(
	environment string,
	serverInstanceNumber string,
	name string,
	configPtr *ExtensionConfiguration,
) (errorInfo pi.ErrorInfo) {

	if errorInfo = hv.DoesFileExistsAndReadable(fmt.Sprintf(configPtr.NATSConfig.NATSCredentialsFilename, environment), ctv.TXT_FILENAME); errorInfo.Error != nil {
		pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, configPtr.NATSConfig.NATSCredentialsFilename))
		return
	} else {
		configPtr.NATSConfig.NATSCredentialsFilename = fmt.Sprintf(configPtr.NATSConfig.NATSCredentialsFilename, environment)
	}
	if configPtr.NATSConfig.NATSPort == ctv.VAL_ZERO {
		errorInfo = pi.NewErrorInfo(pi.ErrNatsPortInvalid, fmt.Sprintf("%v%v", ctv.TXT_NATS_PORT, configPtr.NATSConfig.NATSPort))
		return
	}
	if configPtr.NATSConfig.NATSTLSInfo.TLSCertFQN != ctv.VAL_EMPTY && configPtr.NATSConfig.NATSTLSInfo.TLSPrivateKeyFQN != ctv.VAL_EMPTY && configPtr.NATSConfig.NATSTLSInfo.
		TLSCABundleFQN != ctv.VAL_EMPTY {
		if errorInfo = hv.DoesFileExistsAndReadable(fmt.Sprintf(configPtr.NATSConfig.NATSTLSInfo.TLSCertFQN, environment), ctv.TXT_FILENAME); errorInfo.Error != nil {
			pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, configPtr.NATSConfig.NATSTLSInfo.TLSCert))
			return
		} else {
			configPtr.NATSConfig.NATSTLSInfo.TLSCertFQN = fmt.Sprintf(configPtr.NATSConfig.NATSTLSInfo.TLSCertFQN, environment)
		}
		if errorInfo = hv.DoesFileExistsAndReadable(fmt.Sprintf(configPtr.NATSConfig.NATSTLSInfo.TLSPrivateKeyFQN, environment), ctv.TXT_FILENAME); errorInfo.Error != nil {
			pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, configPtr.NATSConfig.NATSTLSInfo.TLSPrivateKey))
			return
		} else {
			configPtr.NATSConfig.NATSTLSInfo.TLSPrivateKeyFQN = fmt.Sprintf(configPtr.NATSConfig.NATSTLSInfo.TLSPrivateKeyFQN, environment)
		}
		if errorInfo = hv.DoesFileExistsAndReadable(fmt.Sprintf(configPtr.NATSConfig.NATSTLSInfo.TLSCABundleFQN, environment), ctv.TXT_FILENAME); errorInfo.Error != nil {
			pi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, configPtr.NATSConfig.NATSTLSInfo.TLSCABundle))
			return
		} else {
			configPtr.NATSConfig.NATSTLSInfo.TLSCABundleFQN = fmt.Sprintf(configPtr.NATSConfig.NATSTLSInfo.TLSCABundleFQN, environment)
		}
	}
	if hv.IsDomainValid(fmt.Sprintf(configPtr.NATSConfig.NATSURL, environment, serverInstanceNumber)) == false {
		errorInfo = pi.NewErrorInfo(pi.ErrDomainInvalid, fmt.Sprintf("%v%v", ctv.TXT_EVIRONMENT, configPtr.NATSConfig.NATSURL))
		return
	} else {
		configPtr.NATSConfig.NATSURL = fmt.Sprintf(configPtr.NATSConfig.NATSURL, environment, serverInstanceNumber)
	}
	// ToDo Add support for requested threads
	if len(configPtr.SubjectRegistry) == ctv.VAL_ZERO {
		if name == ctv.NC_INTERNAL {
			return
		}
		errorInfo = pi.NewErrorInfo(pi.ErrSubjectsMissing, ctv.VAL_EMPTY)
	}

	return
}
