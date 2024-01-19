// Package coreNATS
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
package coreNATS

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	"albert/constants"
	"albert/core/coreError"
	"albert/core/coreJWT"
)

// This is needed, because GIT must have read access for push,
// and it must be the first test in this file.
func TestGetNATSConnection(tPtr *testing.T) {

	type arguments struct {
		URL                 string
		credentialsLocation string
		myTLS               coreJWT.TLSInfo
	}

	var (
		errorInfo          coreError.ErrorInfo
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
			name: "Positive Case: Connected to NATS.",
			arguments: arguments{
				URL:                 constants.TEST_NATS_URL,
				credentialsLocation: constants.TEST_NATS_CREDENTIALS,
				myTLS: coreJWT.TLSInfo{
					TLSCert:     constants.TEST_CERTIFICATE_FQN,
					TLSKey:      constants.TEST_SAVUP_PRIVATE_KEY_FQN,
					TLSCABundle: constants.TEST_CA_BUNDLE_FQN,
				},
			},
			wantError: false,
		},
		{
			name: "Negative Case: Bad URL.",
			arguments: arguments{
				URL:                 constants.TEST_URL_INVALID,
				credentialsLocation: constants.TEST_NATS_CREDENTIALS,
				myTLS: coreJWT.TLSInfo{
					TLSCert:     constants.TEST_CERTIFICATE_FQN,
					TLSKey:      constants.TEST_SAVUP_PRIVATE_KEY_FQN,
					TLSCABundle: constants.TEST_CA_BUNDLE_FQN,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing credentials location.",
			arguments: arguments{
				URL:                 constants.TEST_NATS_URL,
				credentialsLocation: constants.EMPTY,
				myTLS: coreJWT.TLSInfo{
					TLSCert:     constants.TEST_CERTIFICATE_FQN,
					TLSKey:      constants.TEST_SAVUP_PRIVATE_KEY_FQN,
					TLSCABundle: constants.TEST_CA_BUNDLE_FQN,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing certificate FQN.",
			arguments: arguments{
				URL:                 constants.TEST_NATS_URL,
				credentialsLocation: constants.TEST_NATS_CREDENTIALS,
				myTLS: coreJWT.TLSInfo{
					TLSCert:     constants.EMPTY,
					TLSKey:      constants.TEST_SAVUP_PRIVATE_KEY_FQN,
					TLSCABundle: constants.TEST_CA_BUNDLE_FQN,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing private key FQN.",
			arguments: arguments{
				URL:                 constants.TEST_NATS_URL,
				credentialsLocation: constants.TEST_NATS_CREDENTIALS,
				myTLS: coreJWT.TLSInfo{
					TLSCert:     constants.TEST_CERTIFICATE_FQN,
					TLSKey:      constants.EMPTY,
					TLSCABundle: constants.TEST_CA_BUNDLE_FQN,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing CA bundle FQN.",
			arguments: arguments{
				URL:                 constants.TEST_NATS_URL,
				credentialsLocation: constants.TEST_NATS_CREDENTIALS,
				myTLS: coreJWT.TLSInfo{
					TLSCert:     constants.TEST_CERTIFICATE_FQN,
					TLSKey:      constants.TEST_SAVUP_PRIVATE_KEY_FQN,
					TLSCABundle: constants.EMPTY,
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _, errorInfo = GetNATSConnection(ts.arguments.URL, ts.arguments.credentialsLocation, ts.arguments.myTLS); errorInfo.Error != nil {
				gotError = true
				errorInfo = coreError.ErrorInfo{
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
