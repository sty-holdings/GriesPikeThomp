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

	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cj "GriesPikeThomp/shared-services/src/coreJWT"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

const (
// Add Constants to the constants.go file
)

type ExtensionConfiguration struct {
	MessageEnvironment      string        `json:"message_environment"`
	NATSCredentialsFilename string        `json:"nats_credentials_filename"`
	NATSPort                int           `json:"nats_port"`
	NATSTLSInfo             cj.TLSInfo    `json:"nats_tls_info"`
	NATSURL                 string        `json:"nats_url"`
	RequestedThreads        uint          `json:"requested_threads"`
	SubjectRegistry         []SubjectInfo `json:"subject_registry"`
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
func LoadExtensionConfig(configFilename string) (
	extension ExtensionConfiguration,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v%v", rcv.TXT_FILENAME, configFilename)
		tConfigData     []byte
	)

	if tConfigData, errorInfo = cc.ReadConfigFile(chv.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &extension); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	if errorInfo = validateConfiguration(extension); errorInfo.Error != nil {
		return
	}

	return
}

// validateConfiguration - checks the NATS service configuration is valid.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrMessageNamespaceInvalid, ErrDomainInvalid, error returned from DoesFileExistsAndReadable, ErrSubjectsMissing
//	Verifications: None
func validateConfiguration(config ExtensionConfiguration) (errorInfo cpi.ErrorInfo) {

	if errorInfo = chv.DoesFileExistsAndReadable(config.NATSCredentialsFilename, rcv.TXT_FILENAME); errorInfo.Error != nil {
		cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.NATSCredentialsFilename))
		return
	}
	if chv.IsEnvironmentValid(config.MessageEnvironment) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.MessageEnvironment))
		return
	}
	if config.NATSPort == rcv.VAL_ZERO {
		errorInfo = cpi.NewErrorInfo(cpi.ErrNatsPortInvalid, fmt.Sprintf("%v%v", rcv.TXT_NATS_PORT, config.NATSPort))
		return
	}
	if config.NATSTLSInfo.TLSCert != rcv.VAL_EMPTY && config.NATSTLSInfo.TLSPrivateKey != rcv.VAL_EMPTY && config.NATSTLSInfo.TLSCABundle != rcv.VAL_EMPTY {
		if errorInfo = chv.DoesFileExistsAndReadable(config.NATSTLSInfo.TLSCert, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.NATSTLSInfo.TLSCert))
			return
		}
		if errorInfo = chv.DoesFileExistsAndReadable(config.NATSTLSInfo.TLSPrivateKey, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.NATSTLSInfo.TLSPrivateKey))
			return
		}
		if errorInfo = chv.DoesFileExistsAndReadable(config.NATSTLSInfo.TLSCABundle, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.NATSTLSInfo.TLSCABundle))
			return
		}
	}
	if chv.IsDomainValid(config.NATSURL) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrDomainInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.NATSURL))
		return
	}
	// ToDo Add support for requested threads
	if len(config.SubjectRegistry) == rcv.VAL_ZERO {
		cpi.NewErrorInfo(cpi.ErrSubjectsMissing, rcv.VAL_EMPTY)
	}

	return
}
