// Package coreNATS
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc..
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
package coreNATS

import (
	"fmt"
	"log"
	"runtime"

	"albert/constants"
	"albert/core/coreError"
	"albert/core/coreJWT"
	"github.com/nats-io/nats.go"
)

type NatsHelper struct {
	NatsConnPtr        *nats.Conn
	NatsCredentialsFQN string
	NatsURL            string
}

// GetNATSConnection
func GetNATSConnection(natsURL, natsCredentialsFQN string, tls coreJWT.TLSInfo) (natsConnPtr *nats.Conn, errorInfo coreError.ErrorInfo) {

	var (
		tConnectionType    string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	if tls.TLSCert == constants.EMPTY || tls.TLSKey == constants.EMPTY || tls.TLSCABundle == constants.EMPTY {
		tConnectionType = constants.NATS_NON_TLS_CONNECTION
		if natsConnPtr, errorInfo.Error = nats.Connect(natsURL, nats.UserCredentials(natsCredentialsFQN)); errorInfo.Error != nil {
			coreError.PrintError(errorInfo)
		}
	} else {
		tConnectionType = constants.NATS_TLS_CONNECTION
		if natsConnPtr, errorInfo.Error = nats.Connect(natsURL, nats.UserCredentials(natsCredentialsFQN), nats.RootCAs(tls.TLSCABundle), nats.ClientCert(tls.TLSCert, tls.TLSKey)); errorInfo.Error != nil {
			coreError.PrintError(errorInfo)
		}
	}

	if errorInfo.Error == nil {
		log.Printf("A %v connection has been established with the NATS server.", tConnectionType)
	} else {
		errorInfo.Error = coreError.ErrNATSConnectionFailed
		errorInfo.AdditionalInfo = fmt.Sprintf("Connection Type: %v", tConnectionType)
		coreError.PrintError(errorInfo)
	}

	return
}
