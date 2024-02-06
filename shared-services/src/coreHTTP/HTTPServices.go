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
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	styhHTTP "GriesPikeThomp/servers/nats-connect/extensions/styh-http"
	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cj "GriesPikeThomp/shared-services/src/coreJWT"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type HTTPConfiguration struct {
	CredentialsFilename string      `json:"credentials_filename"`
	GinMode             string      `json:"gin_mode"`
	HTTPDomain          string      `json:"http_domain"`
	MessageEnvironment  string      `json:"message_environment"`
	Port                int         `json:"port"`
	RequestedThreads    uint        `json:"requested_threads"`
	RouteRegistry       []RouteInfo `json:"route_registry"`
	TLSInfo             cj.TLSInfo  `json:"tls_info"`
}

type RouteInfo struct {
	Namespace   string `json:"namespace"`
	Subject     string `json:"subject"`
	Description string `json:"description"`
}

type HTTPService struct {
	Config         HTTPConfiguration
	CredentialsFQN string
	HTTPServerPtr  *http.Server
	Secure         bool
}

// NewHTTP - creates a new styh-http service using the provided extension values.
//
//	Customer Messages: None
//	Errors: error returned by validateConfiguration
//	Verifications: validateConfiguration
func NewHTTP(configFilename string) (service HTTPService, errorInfo cpi.ErrorInfo) {

	var (
		tAdditionalInfo = fmt.Sprintf("%v%v", rcv.TXT_FILENAME, configFilename)
		tConfig         HTTPConfiguration
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

	service.Config = tConfig
	service.CredentialsFQN = chv.PrependWorkingDirectory(tConfig.CredentialsFilename)

	if tConfig.TLSInfo.TLSCert == rcv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSPrivateKey == rcv.VAL_EMPTY ||
		tConfig.TLSInfo.TLSCABundle == rcv.VAL_EMPTY {
		service.Secure = false
	} else {
		service.Secure = true
	}

	service.HTTPServerPtr, errorInfo = createServerInstance(service)

	return
}

//  Private Functions

// createServerInstance - will start listening to request and serving them.
//
//	Customer Messages: None
//	Errors: error returned by nats.Connect
//	Verifications: None
func createServerInstance(service HTTPService) (httpServerPtr *http.Server, errorInfo cpi.ErrorInfo) {

	var ()

	// Start server
	httpServerPtr = &http.Server{
		Addr:    fmt.Sprintf(":%v", service.Config.Port),
		Handler: styhHTTP.SetRoutes(service.Config.GinMode),
	}

	go func() {
		startListener(service.Secure, httpServerPtr, service.Config.TLSInfo.TLSCABundle, service.Config.TLSInfo.TLSPrivateKey)
	}()

	return
}

func startListener(secure bool, httpServerPtr *http.Server, tlsCABundle, tlsPrivateKey string) {

	var (
		errorInfo cpi.ErrorInfo
	)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	if secure {
		// The SSL Cert is from an authority, so the CA bundle must be used.
		if errorInfo.Error = httpServerPtr.ListenAndServeTLS(tlsCABundle, tlsPrivateKey); errorInfo.Error != nil {
			return
		}
	} else {
		errorInfo.Error = httpServerPtr.ListenAndServe()
	}
	// }
	log.Print("Server Started")

	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// extra handling here
		cancel()
	}()

	if err := httpServerPtr.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}

// validateConfiguration - checks the NATS service configuration is valid.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrMessageNamespaceInvalid, ErrDomainInvalid, error returned from DoesFileExistsAndReadable, ErrSubjectsMissing
//	Verifications: None
func validateConfiguration(config HTTPConfiguration) (errorInfo cpi.ErrorInfo) {

	if errorInfo = chv.DoesFileExistsAndReadable(config.CredentialsFilename, rcv.TXT_FILENAME); errorInfo.Error != nil {
		cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.CredentialsFilename))
		return
	}
	if chv.IsBase64Encode(config.CredentialsFilename) == false {
		cpi.NewErrorInfo(cpi.ErrBase64Invalid, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.CredentialsFilename))
		return
	}
	if chv.IsGinModeValid(config.GinMode) == false {
		cpi.NewErrorInfo(cpi.ErrBase64Invalid, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.CredentialsFilename))
		return
	}
	if chv.IsEnvironmentValid(config.MessageEnvironment) == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", rcv.TXT_EVIRONMENT, config.MessageEnvironment))
		return
	}
	if chv.IsGinModeValid(config.GinMode) {
		config.GinMode = strings.ToLower(config.GinMode)
	} else {
		errorInfo = cpi.NewErrorInfo(cpi.ErrGinModeInvalid, fmt.Sprintf("%v%v", rcv.TXT_GIN_MODE, config.GinMode))
		return
	}
	if config.TLSInfo.TLSCert != rcv.VAL_EMPTY && config.TLSInfo.TLSPrivateKey != rcv.VAL_EMPTY && config.TLSInfo.TLSCABundle != rcv.VAL_EMPTY {
		if errorInfo = chv.DoesFileExistsAndReadable(config.TLSInfo.TLSCert, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.TLSInfo.TLSCert))
			return
		}
		if errorInfo = chv.DoesFileExistsAndReadable(config.TLSInfo.TLSPrivateKey, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.TLSInfo.TLSPrivateKey))
			return
		}
		if errorInfo = chv.DoesFileExistsAndReadable(config.TLSInfo.TLSCABundle, rcv.TXT_FILENAME); errorInfo.Error != nil {
			cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_DIRECTORY, config.TLSInfo.TLSCABundle))
			return
		}
	}
	if len(config.RouteRegistry) == rcv.VAL_ZERO {
		cpi.NewErrorInfo(cpi.ErrSubjectsMissing, rcv.VAL_EMPTY)
	}

	return
}

// writeHTTPResponse
