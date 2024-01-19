// Package coreAWS
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
package coreAWS

import (
	"fmt"
	"runtime"
	"testing"

	"albert/constants"
	"albert/core/coreError"
	"albert/core/coreHelpers"
)

// Part of run_AWS_No_Token_Test list
func TestNewAWSSession(tPtr *testing.T) {

	var (
		errorInfo         coreError.ErrorInfo
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if _, errorInfo = NewAWSSession(constants.TEST_AWS_INFORMATION_FQN); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		if _, errorInfo = NewAWSSession(constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})
}

// Part of run_AWS_No_Token_Test list
func TestAWSHelper_ConfirmUser(tPtr *testing.T) {

	var (
		errorInfo         coreError.ErrorInfo
		tAWSHelper        AWSHelper
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
	)

	tAWSHelper, _ = NewAWSSession(constants.TEST_AWS_INFORMATION_FQN)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if errorInfo = tAWSHelper.ConfirmUser(constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		if errorInfo = tAWSHelper.ConfirmUser(constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
func TestAWSHelper_GetRequestorEmailPhoneFromIdTokenClaims(tPtr *testing.T) {

	var (
		errorInfo         coreError.ErrorInfo
		myAWS             AWSHelper
		myFireBase        coreHelpers.FirebaseFirestoreHelper
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
	)

	myAWS, myFireBase = StartTest()

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		//  Positive Test - Successful
		if _, _, _, errorInfo = myAWS.GetRequestorEmailPhoneFromIdTokenClaims(myFireBase.FirestoreClientPtr, string(testingIdTokenValid)); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		if _, _, _, errorInfo = myAWS.GetRequestorEmailPhoneFromIdTokenClaims(myFireBase.FirestoreClientPtr, constants.TEST_TOKEN_INVALID); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		if _, _, _, errorInfo = myAWS.GetRequestorEmailPhoneFromIdTokenClaims(myFireBase.FirestoreClientPtr, constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})

	StopTest(myFireBase)
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
func TestAWSHelper_GetRequestorFromAccessTokenClaims(tPtr *testing.T) {

	var (
		errorInfo         coreError.ErrorInfo
		myAWS             AWSHelper
		myFireBase        coreHelpers.FirebaseFirestoreHelper
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
	)

	myAWS, myFireBase = StartTest()

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		//  Positive Test - Successful
		if _, errorInfo = myAWS.GetRequestorFromAccessTokenClaims(myFireBase.FirestoreClientPtr, string(testingAccessTokenValid)); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		if _, errorInfo = myAWS.GetRequestorFromAccessTokenClaims(myFireBase.FirestoreClientPtr, constants.TEST_TOKEN_INVALID); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		if _, errorInfo = myAWS.GetRequestorFromAccessTokenClaims(myFireBase.FirestoreClientPtr, constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})

	StopTest(myFireBase)
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
func TestAWSHelper_ParseAWSJWTWithClaims(tPtr *testing.T) {

	type arguments struct {
		tokenType string
		token     string
	}

	var (
		errorInfo          coreError.ErrorInfo
		gotError           bool
		myAWS              AWSHelper
		myFireBase         coreHelpers.FirebaseFirestoreHelper
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tToken             string
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful id token!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ID,
				token:     "valid",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful access token!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ACCESS,
				token:     "valid",
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing token type!",
			arguments: arguments{
				tokenType: constants.EMPTY,
				token:     "valid",
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing token!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ACCESS,
				token:     constants.EMPTY,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Invalid id token!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ID,
				token:     "invalid",
			},
			wantError: true,
		},
		{
			name: "Negative Case: Invalid access token!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ACCESS,
				token:     "invalid",
			},
			wantError: true,
		},
	}

	myAWS, myFireBase = StartTest()

	for _, ts := range tests {
		tPtr.Run(tFunctionName, func(t *testing.T) {
			tToken = getToken(ts.arguments.tokenType, ts.arguments.token)
			if _, errorInfo = myAWS.ParseAWSJWTWithClaims(ts.arguments.tokenType, tToken); errorInfo.Error != nil {
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

	StopTest(myFireBase)
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
func TestAWSHelper_ParseJWT(tPtr *testing.T) {

	var (
		errorInfo         coreError.ErrorInfo
		myAWS             AWSHelper
		myFireBase        coreHelpers.FirebaseFirestoreHelper
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
	)

	myAWS, myFireBase = StartTest()

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if _, errorInfo = myAWS.ParseAWSJWT(string(testingAccessTokenValid)); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		if _, errorInfo = myAWS.ParseAWSJWT(constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})

	StopTest(myFireBase)
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
func TestAWSHelper_PullCognitoUserInfo(tPtr *testing.T) {

	var (
		errorInfo         coreError.ErrorInfo
		myAWS             AWSHelper
		myFireBase        coreHelpers.FirebaseFirestoreHelper
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
	)

	myAWS, myFireBase = StartTest()

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if _, errorInfo = myAWS.PullCognitoUserInfo(constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		if _, errorInfo = myAWS.PullCognitoUserInfo(constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})

	StopTest(myFireBase)
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
//
//	The actual reset will be bypassed because the resetByPass is set to true
func TestAWSHelper_ResetUserPassword(tPtr *testing.T) {

	var (
		errorInfo         coreError.ErrorInfo
		myAWS             AWSHelper
		myFireBase        coreHelpers.FirebaseFirestoreHelper
		resetByPass       = true
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
	)

	myAWS, myFireBase = StartTest()

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if errorInfo = myAWS.ResetUserPassword(constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE, resetByPass); errorInfo.Error != nil {
			if errorInfo.Error.Error() == coreError.ATTEMPTS_EXCEEDED {
				fmt.Println(coreError.ATTEMPTS_EXCEEDED)
			} else {
				tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
			}
		}
		if errorInfo = myAWS.ResetUserPassword(constants.EMPTY, resetByPass); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})

	StopTest(myFireBase)
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
func TestAWSHelper_UpdateAWSEmailVerifyFlag(tPtr *testing.T) {
	//
	// NOTE: The Id and Access token must match the username in constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE
	//

	var (
		errorInfo         coreError.ErrorInfo
		myAWS             AWSHelper
		myFireBase        coreHelpers.FirebaseFirestoreHelper
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
		tUsername         = constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE
	)

	myAWS, myFireBase = StartTest()

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if errorInfo = myAWS.UpdateAWSEmailVerifyFlag(tUsername); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		if errorInfo = myAWS.UpdateAWSEmailVerifyFlag(constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})

	StopTest(myFireBase)
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
func TestAWSHelper_ValidAWSJWT(tPtr *testing.T) {

	type arguments struct {
		tokenType string
		token     string
	}

	var (
		errorInfo         coreError.ErrorInfo
		myAWS             AWSHelper
		myFireBase        coreHelpers.FirebaseFirestoreHelper
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
		tToken            string
		tValid            bool
	)

	tests := []struct {
		name          string
		arguments     arguments
		shouldBeValid bool
	}{
		{
			name: "Positive Case: Access Successful!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ACCESS,
				token:     "valid",
			},
			shouldBeValid: true,
		},
		{
			name: "Positive Case: Id Successful!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ID,
				token:     "valid",
			},
			shouldBeValid: true,
		},
		{
			name: "Negative Case: Access invalid!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ACCESS,
				token:     "invalid",
			},
			shouldBeValid: false,
		},
		{
			name: "Negative Case: Id invalid!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ID,
				token:     "invalid",
			},
			shouldBeValid: false,
		},
		{
			name: "Negative Case: Access missing!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ACCESS,
				token:     "missing",
			},
			shouldBeValid: false,
		},
		{
			name: "Negative Case: Id missing!",
			arguments: arguments{
				tokenType: constants.TOKEN_TYPE_ID,
				token:     "missing",
			},
			shouldBeValid: false,
		},
	}

	myAWS, myFireBase = StartTest()

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			tToken = getToken(ts.arguments.tokenType, ts.arguments.token)
			if tValid, errorInfo = myAWS.ValidAWSJWT(myFireBase.FirestoreClientPtr, ts.arguments.tokenType, tToken); tValid != ts.shouldBeValid {
				tPtr.Error(tFunctionName, ts.name, errorInfo, fmt.Sprintf("Expected the token to be %v and it was %v", ts.shouldBeValid, tValid))
			}
		})
	}

	StopTest(myFireBase)
}

// Part of run_AWS_No_Token_Test list
func TestGetPublicKeySet(tPtr *testing.T) {

	var (
		errorInfo         coreError.ErrorInfo
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
		tKeySetURL        = fmt.Sprintf(constants.TEST_AWS_KEYSET_URL, constants.TEST_USER_POOL_ID)
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if _, errorInfo = getPublicKeySet(tKeySetURL); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		if _, errorInfo = getPublicKeySet(constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		if _, errorInfo = getPublicKeySet(constants.TEST_URL_INVALID); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, `errorInfo.Error.Error()`, "nil")
		}
	})
}

// Requires updated access token. You can use Cognito > User pools > App integration > App clients and analytics > {app name} > Hosted UI > View Hosted UI
// to login. This will output an access and id token for the user.
func TestValidAWSClaims(tPtr *testing.T) {

	type arguments struct {
		subject       string
		email         string
		username      string
		emailVerified bool // emailVerified is only checked for constants.TOKEN_TYPE_ID
		tokenUse      string
	}

	var (
		errorInfo         coreError.ErrorInfo
		myFireBase        coreHelpers.FirebaseFirestoreHelper
		function, _, _, _ = runtime.Caller(0)
		tFunctionName     = runtime.FuncForPC(function).Name()
		tValid            bool
	)

	tests := []struct {
		name          string
		arguments     arguments
		shouldBeValid bool
	}{
		{
			name: "Positive Case: Successful Id Token!",
			arguments: arguments{
				subject:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				email:         constants.TEST_USER_EMAIL,
				username:      constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE,
				emailVerified: true,
				tokenUse:      constants.TOKEN_TYPE_ID,
			},
			shouldBeValid: true,
		},
		{
			name: "Positive Case: Successful Access Token!",
			arguments: arguments{
				subject:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				email:         constants.TEST_USER_EMAIL,
				username:      constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE,
				emailVerified: true,
				tokenUse:      constants.TOKEN_TYPE_ACCESS,
			},
			shouldBeValid: true,
		},
		{
			name: "Negative Case: Email not verified!",
			arguments: arguments{
				subject:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				email:         constants.TEST_USER_EMAIL,
				username:      constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE,
				emailVerified: false,
				tokenUse:      constants.TOKEN_TYPE_ID,
			},
			shouldBeValid: false,
		},
		{
			name: "Negative Case: Token type missing!",
			arguments: arguments{
				subject:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				email:         constants.TEST_USER_EMAIL,
				username:      constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE,
				emailVerified: true,
				tokenUse:      constants.EMPTY,
			},
			shouldBeValid: false,
		},
		{
			name: "Negative Case: subject is missing!",
			arguments: arguments{
				subject:       constants.EMPTY,
				email:         constants.TEST_USER_EMAIL,
				username:      constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE,
				emailVerified: true,
				tokenUse:      constants.EMPTY,
			},
			shouldBeValid: false,
		},
		{
			name: "Negative Case: email is missing!",
			arguments: arguments{
				subject:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				email:         constants.EMPTY,
				username:      constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE,
				emailVerified: true,
				tokenUse:      constants.EMPTY,
			},
			shouldBeValid: false,
		},
		{
			name: "Negative Case: username is missing!",
			arguments: arguments{
				subject:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				email:         constants.TEST_USER_EMAIL,
				username:      constants.EMPTY,
				emailVerified: true,
				tokenUse:      constants.EMPTY,
			},
			shouldBeValid: false,
		},
	}

	_, myFireBase = StartTest()

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if tValid = areAWSClaimsValid(myFireBase.FirestoreClientPtr, ts.arguments.subject, ts.arguments.email, ts.arguments.username, ts.arguments.tokenUse, ts.arguments.emailVerified); tValid != ts.shouldBeValid {
				tPtr.Error(tFunctionName, ts.name, errorInfo, fmt.Sprintf("Expected the token to be %v and it was %v", ts.shouldBeValid, tValid))
			}
		})
	}

	StopTest(myFireBase)
}
