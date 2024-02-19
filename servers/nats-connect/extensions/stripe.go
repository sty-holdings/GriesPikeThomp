// Package extensions
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
package extensions

import (
	"encoding/json"
	"fmt"

	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cj "GriesPikeThomp/shared-services/src/coreJWT"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type NCInternalConfiguration struct {
	MessageEnvironment      string        `json:"message_environment"`
	NATSCredentialsFilename string        `json:"credentials_filename"`
	NATSPort                int           `json:"port"`
	NATSTLSInfo             cj.TLSInfo    `json:"tls_info"`
	NATSURL                 string        `json:"url"`
	RequestedThreads        uint          `json:"requested_threads"`
	SubjectRegistry         []SubjectInfo `json:"subject_registry"`
}

type SubjectInfo struct {
	Namespace   string `json:"namespace"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

// LoadNCInternal - reads, validates, and loads into service extension map
//
//	Customer Messages: None
//	Errors: error returned by validateConfiguration
//	Verifications: validateConfiguration
func LoadNCInternal(configFilename string) (ncInternal interface{}, errorInfo cpi.ErrorInfo) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v%v", rcv.TXT_FILENAME, configFilename)
		tConfig         NCInternalConfiguration
		tConfigData     []byte
	)

	if tConfigData, errorInfo = cc.ReadConfigFile(chv.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &tConfig); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	if errorInfo = validateConfiguration(tConfig); errorInfo.Error != nil {
		return
	}

	ncInternal = interface{}(tConfig)

	return
}

//  Private Functions

// validateConfiguration - checks the NATS service configuration is valid.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrMessageNamespaceInvalid, ErrDomainInvalid, error returned from DoesFileExistsAndReadable, ErrSubjectsMissing
//	Verifications: None
func validateConfiguration(config NCInternalConfiguration) (errorInfo cpi.ErrorInfo) {

	if errorInfo = chv.DoesFileExistsAndReadable(config.NATSCredentialsFilename, rcv.TXT_FILENAME); errorInfo.Error != nil {
		cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.NATSCredentialsFilename))
		return
	}
	if chv.IsEnvironmentValid(config.MessageEnvironment) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.MessageEnvironment))
		return
	}
	if chv.IsDomainValid(config.NATSURL) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrDomainInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.NATSURL))
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
	if len(config.SubjectRegistry) == rcv.VAL_ZERO {
		cpi.NewErrorInfo(cpi.ErrSubjectsMissing, rcv.VAL_EMPTY)
	}

	return
}
