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
	"errors"
	"fmt"
	"runtime"
	"testing"

	cj "GriesPikeThomp/shared-services/src/coreJWT"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// This is needed, because GIT must have read access for push,
// and it must be the first test in this file.
func TestGetNATSConnection(tPtr *testing.T) {

	type arguments struct {
		URL                 string
		credentialsLocation string
		myTLS               cj.TLSInfo
	}

	var (
		errorInfo          cpi.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				URL:                 TEST_URL,
				credentialsLocation: TEST_CREDENTIALS_FILENAME,
				myTLS: cj.TLSInfo{
					TLSCert:       TEST_TLS_CERT,
					TLSPrivateKey: TEST_TLS_PRIVATE_KEY,
					TLSCABundle:   TEST_TLS_CA_BUNDLE,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Bad URL.",
			arguments: arguments{
				URL:                 TEST_INVALID_URL,
				credentialsLocation: TEST_CREDENTIALS_FILENAME,
				myTLS: cj.TLSInfo{
					TLSCert:       TEST_TLS_CERT,
					TLSPrivateKey: TEST_TLS_PRIVATE_KEY,
					TLSCABundle:   TEST_TLS_CA_BUNDLE,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing credentials location.",
			arguments: arguments{
				URL:                 TEST_URL,
				credentialsLocation: rcv.VAL_EMPTY,
				myTLS: cj.TLSInfo{
					TLSCert:       TEST_TLS_CERT,
					TLSPrivateKey: TEST_TLS_PRIVATE_KEY,
					TLSCABundle:   TEST_TLS_CA_BUNDLE,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing certificate FQN.",
			arguments: arguments{
				URL:                 TEST_URL,
				credentialsLocation: TEST_CREDENTIALS_FILENAME,
				myTLS: cj.TLSInfo{
					TLSCert:       rcv.VAL_EMPTY,
					TLSPrivateKey: TEST_TLS_PRIVATE_KEY,
					TLSCABundle:   TEST_TLS_CA_BUNDLE,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing private key FQN.",
			arguments: arguments{
				URL:                 TEST_URL,
				credentialsLocation: TEST_CREDENTIALS_FILENAME,
				myTLS: cj.TLSInfo{
					TLSCert:       TEST_TLS_CERT,
					TLSPrivateKey: rcv.VAL_EMPTY,
					TLSCABundle:   TEST_TLS_CA_BUNDLE,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing CA bundle FQN.",
			arguments: arguments{
				URL:                 TEST_URL,
				credentialsLocation: TEST_CREDENTIALS_FILENAME,
				myTLS: cj.TLSInfo{
					TLSCert:       TEST_TLS_CERT,
					TLSPrivateKey: TEST_TLS_PRIVATE_KEY,
					TLSCABundle:   rcv.VAL_EMPTY,
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _, errorInfo = GetConnection(ts.arguments.URL, ts.arguments.credentialsLocation, ts.arguments.myTLS); errorInfo.Error != nil {
				gotError = true
				errorInfo = cpi.ErrorInfo{
					Error: errors.New(fmt.Sprintf("Failed - NATS connection was not created for Test: %v", tFunctionName)),
				}
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}
}
