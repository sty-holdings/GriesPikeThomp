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
	"log"
	"reflect"
	"strings"

	ch "GriesPikeThomp/shared-services/src/coreHelpers"
	cj "GriesPikeThomp/shared-services/src/coreJWT"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	cv "GriesPikeThomp/shared-services/src/coreValidators"
	"github.com/nats-io/nats.go"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type Configuration struct {
	credentialsFilename string
	messageEnvironment  string
	messageNamespace    string
	secure              bool
	url                 string
	port                uint
	tlsInfo             cj.TLSInfo
}

type Service struct {
	config         Configuration
	ConnPtr        *nats.Conn
	credentialsFQN string
	secure         bool
	url            string
}

// NewNATS - creates a new NATS service using the provided extension values.
//
//	Customer Messages: None
//	Errors: error returned by validateConfiguration
//	Verifications: validateConfiguration
func NewNATS(extensionValues map[string]interface{}) (natsPtr *Service, errorInfo cpi.ErrorInfo) {

	var (
		natsService = Service{
			config: populateConfiguration(extensionValues),
		}
	)

	if errorInfo = validateConfiguration(natsService.config); errorInfo.Error != nil {
		return
	}

	natsService.credentialsFQN = ch.PrependWorkingDirectory(natsService.config.credentialsFilename)
	natsService.url = natsService.config.url

	if natsService.config.tlsInfo.TLSCert == rcv.VAL_EMPTY ||
		natsService.config.tlsInfo.TLSPrivateKey == rcv.VAL_EMPTY ||
		natsService.config.tlsInfo.TLSCABundle == rcv.VAL_EMPTY {
		natsService.secure = false
	} else {
		natsService.secure = true
	}

	natsService.ConnPtr, errorInfo = getConnection(natsService)
	natsPtr = &natsService

	return
}

//  Private Functions

// getConnection - will connect to a NATS leaf server with either a ssl or non-ssl connection.
//
//	Customer Messages: None
//	Errors: error returned by nats.Connect
//	Verifications: None
func getConnection(natsService Service) (connPtr *nats.Conn, errorInfo cpi.ErrorInfo) {

	if natsService.secure {
		if connPtr, errorInfo.Error = nats.Connect(natsService.url, nats.UserCredentials(natsService.credentialsFQN), nats.RootCAs(natsService.config.tlsInfo.TLSCABundle),
			nats.ClientCert(natsService.config.tlsInfo.TLSCert, natsService.config.tlsInfo.TLSPrivateKey)); errorInfo.Error != nil {
			errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprint(rcv.TXT_SECURE_CONNECTION_FAILED))
			return
		}
	} else {
		if connPtr, errorInfo.Error = nats.Connect(natsService.url, nats.UserCredentials(natsService.credentialsFQN)); errorInfo.Error != nil {
			errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprint(rcv.TXT_NON_SECURE_CONNECTION_FAILED))
			return
		}
	}

	log.Printf("A connection has been established with the NATS server at %v.", natsService.url)

	return
}

// populateConfiguration - builds the NATS service configuration
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func populateConfiguration(extensionValues map[string]interface{}) (tConfig Configuration) {

	for fieldName, fieldValue := range extensionValues {
		switch strings.ToLower(fieldName) {
		case rcv.FN_CREDENTIALS_FILENAME:
			tConfig.credentialsFilename = fieldValue.(string)
		case rcv.FN_MESSAGE_ENVIRONMENT:
			tConfig.messageEnvironment = fieldValue.(string)
		case rcv.FN_MESSAGE_NAMESPACE:
			tConfig.messageNamespace = fieldValue.(string)
		case rcv.FN_URL:
			tConfig.url = fieldValue.(string)
		case rcv.FN_PORT:
			tConfig.port = 4222
			if reflect.TypeOf(fieldValue).String() == rcv.TXT_DATATYPE_FLOAT64 {
				tConfig.port = uint(fieldValue.(float64))
				if tConfig.port == rcv.VAL_ZERO {
					tConfig.port = 4222
				}
			}
		case rcv.FN_TLS_INFO:
			if tTLSInfo, ok := fieldValue.(map[string]interface{}); ok {
				for tFieldName, tFieldValue := range tTLSInfo {
					switch strings.ToLower(tFieldName) {
					case rcv.FN_TLS_CERTIFICATE_FILENAME:
						tConfig.tlsInfo.TLSCert = tFieldValue.(string)
					case rcv.FN_TLS_PRIVATE_KEY_FILENAME:
						tConfig.tlsInfo.TLSPrivateKey = tFieldValue.(string)
					case rcv.FN_TLS_CA_BUNDLE_FILENAME:
						tConfig.tlsInfo.TLSCABundle = tFieldValue.(string)
					}
				}
			}
		}
	}

	return
}

// validateConfiguration - checks the NATS service configuration is valid.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrDomainInvalid, error returned from DoesFileExistsAndReadable
//	Verifications: None
func validateConfiguration(config Configuration) (errorInfo cpi.ErrorInfo) {

	if errorInfo = cv.DoesFileExistsAndReadable(config.credentialsFilename, rcv.TXT_FILENAME); errorInfo.Error != nil {
		cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.credentialsFilename))
		return
	}
	if cv.IsEnvironmentValid(config.messageEnvironment) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.messageEnvironment))
		return
	}
	if config.messageNamespace == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrMessageNamespaceInvalid, fmt.Sprintf("%v%v", rcv.TXT_MESSAGE, config.messageEnvironment))
		return
	}
	if cv.IsDomainValid(config.url) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrDomainInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.messageEnvironment))
		return
	}
	if errorInfo = cv.DoesFileExistsAndReadable(config.tlsInfo.TLSCert, rcv.TXT_FILENAME); errorInfo.Error != nil {
		cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.tlsInfo.TLSCert))
		return
	}
	if errorInfo = cv.DoesFileExistsAndReadable(config.tlsInfo.TLSPrivateKey, rcv.TXT_FILENAME); errorInfo.Error != nil {
		cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.tlsInfo.TLSPrivateKey))
		return
	}
	if errorInfo = cv.DoesFileExistsAndReadable(config.tlsInfo.TLSCABundle, rcv.TXT_FILENAME); errorInfo.Error != nil {
		cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.tlsInfo.TLSCABundle))
		return
	}

	return
}
