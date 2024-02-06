// Package coreFirebase
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
package coreFirebase

import (
	"runtime"
	"testing"

	"albert/constants"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
)

var (
//goland:noinspection ALL
)

func TestFindFirebaseAuthUser(tPtr *testing.T) {

	var (
		tAuthPtr           *auth.Client
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		errorInfo          cpi.ErrorInfo
	)

	_, tAuthPtr, _ = GetFirebaseAppAuthConnection(rcv.TEST_FIREBASE_CREDENTIALS)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		if _, errorInfo = FindFirebaseAuthUser(tAuthPtr, rcv.TEST_USERNAME_SAVUP_REQUESTOR_ID); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, cpi.ERROR, "nil")
		}
	})

}

func TestGetIdTokenPayload(tPtr *testing.T) {

	var (
		tAuthPtr           *auth.Client
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tTokenPayload      = make(map[any]interface{})
	)

	_, tAuthPtr, _ = GetFirebaseAppAuthConnection(rcv.TEST_FIREBASE_CREDENTIALS)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		if tTokenPayload, _ = GetFirebaseIdTokenPayload(tAuthPtr, TEST_FIREBASE_IDTOKEN_VALID); len(tTokenPayload) == 0 {
			tPtr.Errorf("%v Failed: Was expecting the JWT payload to be populated.", tFunctionName)
		}
	})
}

func TestGetIdTokenPtr(tPtr *testing.T) {

	var (
		tAuthPtr           *auth.Client
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tIdTokenPtr        *auth.Token
	)

	_, tAuthPtr, _ = GetFirebaseAppAuthConnection(rcv.TEST_FIREBASE_CREDENTIALS)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		if tIdTokenPtr, _ = GetIdTokenPtr(tAuthPtr, TEST_FIREBASE_IDTOKEN_VALID); tIdTokenPtr == nil {
			tPtr.Errorf("%v Failed: No Token was return. Make sure the tIdTokenValid is a valid and recent JWT.", tFunctionName)
		}
	})
}

func TestIsFirebaseIdTokenValid(tPtr *testing.T) {

	var (
		tAuthPtr           *auth.Client
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tValid             bool
	)

	_, tAuthPtr, _ = GetFirebaseAppAuthConnection(rcv.TEST_FIREBASE_CREDENTIALS)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if tValid = IsFirebaseIdTokenValid(tAuthPtr, TEST_FIREBASE_IDTOKEN_INVALID); tValid == true {
			tPtr.Errorf("%v Failed: Token is should be invalid. Valid: %v", tFunctionName, tValid)
		}
		if tValid = IsFirebaseIdTokenValid(tAuthPtr, TEST_FIREBASE_IDTOKEN_VALID); tValid == false {
			tPtr.Errorf("%v Failed: Token is should be valid. Valid: %v", tFunctionName, tValid)
		}
	})
}

func TestNewFirebaseApp(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tAppPtr            *firebase.App
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if tAppPtr, errorInfo = NewFirebaseApp(rcv.TEST_FIREBASE_CREDENTIALS); tAppPtr == nil || errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Firebase app was not created.", tFunctionName)
		}
	})
}

func TestValidateFirebaseJWTPayload(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tAuthPtr           *auth.Client
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tTokenPayload      = make(map[any]interface{})
		tValid             bool
	)

	_, tAuthPtr, _ = GetFirebaseAppAuthConnection(rcv.TEST_FIREBASE_CREDENTIALS)
	tTokenPayload, _ = GetFirebaseIdTokenPayload(tAuthPtr, TEST_FIREBASE_IDTOKEN_VALID)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if errorInfo = ValidateFirebaseJWTPayload(tTokenPayload, rcv.CERT_SAVUPDEV_AUDIENCE, rcv.CERT_DEV_ID_TOEKN_ISSUER); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Token payload should be valid. Valid: %v", tFunctionName, tValid)
		}
	})
}
