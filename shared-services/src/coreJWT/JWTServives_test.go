// Package coreJWT
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
package coreJWT

import (
	"runtime"
	"testing"

	"albert/constants"
	"albert/core/coreError"
)

//goland:noinspection ALL
const (
	TEST_PRIVATE_KEY = "-----BEGIN RSA PRIVATE KEY-----\n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\nYd7tZETziCa3dFTmIcKeu+EpCzX5SA==\n-----END RSA PRIVATE KEY-----\n"
)

func TestGenerateJWT(tPtr *testing.T) {

	type arguments struct {
		privateKey  string
		requestorId string
		period      string
		duration    int64
	}

	var (
		errorInfo coreError.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful - period year!",
			arguments: arguments{
				privateKey:  TEST_PRIVATE_KEY,
				requestorId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				period:      constants.YEAR,
				duration:    10,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful! - period month",
			arguments: arguments{
				privateKey:  TEST_PRIVATE_KEY,
				requestorId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				period:      constants.MONTH,
				duration:    10,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful - period day!",
			arguments: arguments{
				privateKey:  TEST_PRIVATE_KEY,
				requestorId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				period:      constants.DAY,
				duration:    10,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Empty period!",
			arguments: arguments{
				privateKey:  TEST_PRIVATE_KEY,
				requestorId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				period:      constants.EMPTY,
				duration:    10,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Empty requestor id!",
			arguments: arguments{
				privateKey:  TEST_PRIVATE_KEY,
				requestorId: constants.EMPTY,
				period:      constants.DAY,
				duration:    10,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Zero duration",
			arguments: arguments{
				privateKey:  TEST_PRIVATE_KEY,
				requestorId: constants.EMPTY,
				period:      constants.DAY,
				duration:    0,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Zero negative",
			arguments: arguments{
				privateKey:  TEST_PRIVATE_KEY,
				requestorId: constants.EMPTY,
				period:      constants.DAY,
				duration:    -1,
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _, errorInfo = GenerateJWT(ts.arguments.privateKey, ts.arguments.requestorId, ts.arguments.period, ts.arguments.duration); errorInfo.Error != nil {
				gotError = true
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

func TestGenerateRSAKey(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		errorInfo          coreError.ErrorInfo
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		if _, _, errorInfo = GenerateRSAKey(4096); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: GenerateRSAKey returned an error.", tFunctionName)
		}
		if _, _, errorInfo = GenerateRSAKey(0); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: GenerateRSAKey did not returned an error.", tFunctionName)
		}
	})

}

func TestParsePrivateKey(tPtr *testing.T) {

	var (
		errorInfo          coreError.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tRawPrivateKey     []byte
	)

	tRawPrivateKey = []byte(TEST_PRIVATE_KEY)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if _, errorInfo = ParsePrivateKey(tRawPrivateKey); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expected %v and got error.", tFunctionName, constants.STATUS_SUCCESS)
		}
		if _, errorInfo = ParsePrivateKey([]byte(constants.EMPTY)); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expected an error and got %v.", tFunctionName, constants.STATUS_SUCCESS)
		}
	})
}
