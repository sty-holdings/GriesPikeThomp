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
	"log"
	"time"

	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	ch "GriesPikeThomp/shared-services/src/coreHelpers"
	cj "GriesPikeThomp/shared-services/src/coreJWT"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	cv "GriesPikeThomp/shared-services/src/coreValidators"
	"github.com/nats-io/nats.go"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type NATSConfiguration struct {
	CredentialsFilename string        `json:"credentials_filename"`
	MessageEnvironment  string        `json:"message_environment"`
	MessageNamespace    string        `json:"message_namespace"`
	Port                int           `json:"port"`
	RequestedThreads    uint          `json:"requested_threads"`
	SubjectRegistry     []SubjectInfo `json:"subject_registry"`
	TLSInfo             cj.TLSInfo    `json:"tls_info"`
	URL                 string        `json:"url"`
}

type SubjectInfo struct {
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type NATSService struct {
	Config         NATSConfiguration
	ConnPtr        *nats.Conn
	CredentialsFQN string
	Namespace      string
	Secure         bool
	URL            string
}

// NewNATS - creates a new NATS service using the provided extension values.
//
//	Customer Messages: None
//	Errors: error returned by validateConfiguration
//	Verifications: validateConfiguration
func NewNATS(hostname string, configFilename string) (service NATSService, errorInfo cpi.ErrorInfo) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v %v", rcv.TXT_FILENAME, configFilename)
		tConfig         NATSConfiguration
		tConfigData     []byte
	)

	if tConfigData, errorInfo = cc.ReadConfigFile(ch.PrependWorkingDirectory(configFilename)); errorInfo.Error != nil {
		return
	}

	if errorInfo.Error = json.Unmarshal(tConfigData, &tConfig); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, tAdditionalInfo)
		return
	}

	if errorInfo = validateConfiguration(tConfig); errorInfo.Error != nil {
		return
	}

	service.Config = tConfig
	service.CredentialsFQN = ch.PrependWorkingDirectory(tConfig.CredentialsFilename)
	service.Namespace = tConfig.MessageNamespace
	service.URL = tConfig.URL

	if tConfig.TLSInfo.TLSCert == rcv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSPrivateKey == rcv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSCABundle == rcv.VAL_EMPTY {
		service.Secure = false
	} else {
		service.Secure = true
	}

	service.ConnPtr, errorInfo = getConnection(hostname, service)

	return
}

//  Private Functions

// getConnection - will connect to a NATS leaf server with either a ssl or non-ssl connection.
//
//	Customer Messages: None
//	Errors: error returned by nats.Connect
//	Verifications: None
func getConnection(hostname string, service NATSService) (connPtr *nats.Conn, errorInfo cpi.ErrorInfo) {

	var (
		opts []nats.Option
	)

	opts = []nats.Option{
		nats.Name(hostname),                 // Set a client name
		nats.MaxReconnects(5),               // Set maximum reconnection attempts
		nats.ReconnectWait(5 * time.Second), // Set reconnection wait time
		nats.UserCredentials(service.CredentialsFQN),
		nats.RootCAs(service.Config.TLSInfo.TLSCABundle),
		nats.ClientCert(service.Config.TLSInfo.TLSCert, service.Config.TLSInfo.TLSPrivateKey),
	}

	if service.Secure {
		if connPtr, errorInfo.Error = nats.Connect(service.URL, opts...); errorInfo.Error != nil {
			errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprint(rcv.TXT_SECURE_CONNECTION_FAILED))
			return
		}
	} else {
		if connPtr, errorInfo.Error = nats.Connect(service.URL, nats.UserCredentials(service.CredentialsFQN)); errorInfo.Error != nil {
			errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprint(rcv.TXT_NON_SECURE_CONNECTION_FAILED))
			return
		}
	}

	log.Printf("A connection has been established with the NATS server at %v.", service.URL)
	log.Printf("URL: %v Server Name: %v Server Id: %v Address: %v", connPtr.ConnectedUrl(), connPtr.ConnectedClusterName(), connPtr.ConnectedServerId(), connPtr.ConnectedAddr())

	return
}

// validateConfiguration - checks the NATS service configuration is valid.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrMessageNamespaceInvalid, ErrDomainInvalid, error returned from DoesFileExistsAndReadable, ErrSubjectsMissing
//	Verifications: None
func validateConfiguration(config NATSConfiguration) (errorInfo cpi.ErrorInfo) {

	if errorInfo = cv.DoesFileExistsAndReadable(config.CredentialsFilename, rcv.TXT_FILENAME); errorInfo.Error != nil {
		cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.CredentialsFilename))
		return
	}
	if cv.IsEnvironmentValid(config.MessageEnvironment) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.MessageEnvironment))
		return
	}
	if config.MessageNamespace == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrMessageNamespaceInvalid, fmt.Sprintf("%v%v", rcv.TXT_MESSAGE, config.MessageEnvironment))
		return
	}
	if cv.IsDomainValid(config.URL) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrDomainInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.URL))
		return
	}
	if config.TLSInfo.TLSCert != rcv.VAL_EMPTY && config.TLSInfo.TLSPrivateKey != rcv.VAL_EMPTY && config.TLSInfo.TLSCABundle != rcv.VAL_EMPTY {
		if errorInfo = cv.DoesFileExistsAndReadable(config.TLSInfo.TLSCert, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.TLSInfo.TLSCert))
			return
		}
		if errorInfo = cv.DoesFileExistsAndReadable(config.TLSInfo.TLSPrivateKey, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.TLSInfo.TLSPrivateKey))
			return
		}
		if errorInfo = cv.DoesFileExistsAndReadable(config.TLSInfo.TLSCABundle, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.TLSInfo.TLSCABundle))
			return
		}
	}
	if len(config.SubjectRegistry) == rcv.VAL_ZERO {
		cpi.NewErrorInfo(cpi.ErrSubjectsMissing, rcv.VAL_EMPTY)
	}

	return
}
