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
	"context"
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"

	"albert/constants"
	"albert/core/coreValidators"
	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

var (
	CTXBackground = context.Background()
)

// FindFirebaseAuthUser - determines if the user exists in the Firebase Auth database. If so, then pointer to the user is return, otherwise, an error.
func FindFirebaseAuthUser(authPtr *auth.Client, requestorId string) (userRecordPtr *auth.UserRecord, errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if userRecordPtr, errorInfo.Error = authPtr.GetUser(CTXBackground, requestorId); errorInfo.Error != nil {
		errorInfo.Error = errors.New(fmt.Sprintf("An error while getting Requestor Id: %v from Firebase Auth. Error: %v", requestorId, errorInfo.Error.Error()))
		log.Println(errorInfo.Error.Error())
	}

	return
}

// GetFirebaseFirestoreConnection
func GetFirebaseAppAuthConnection(credentialsLocation string) (appPtr *firebase.App, authPtr *auth.Client, errorInfo cpi.ErrorInfo) {

	if appPtr, errorInfo = NewFirebaseApp(credentialsLocation); errorInfo.Error == nil {
		authPtr, errorInfo = GetFirebaseAuthConnection(appPtr)
	}

	return
}

// GetFirebaseIdTokenPayload
func GetFirebaseIdTokenPayload(authPtr *auth.Client, idToken string) (tokenPayload map[any]interface{}, errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tIdTokenPtr        *auth.Token
	)

	cpi.PrintDebugTrail(tFunctionName)

	tokenPayload = make(map[any]interface{})
	if tIdTokenPtr, errorInfo = GetIdTokenPtr(authPtr, idToken); errorInfo.Error == nil {
		tokenPayload[rcv.PAYLOAD_SUBJECT_FN] = tIdTokenPtr.Subject
		tokenPayload[rcv.PAYLOAD_CLAIMS_FN] = tIdTokenPtr.Claims
		tokenPayload[rcv.PAYLOAD_AUDIENCE_FN] = tIdTokenPtr.Audience
		tokenPayload[rcv.PAYLOAD_REQUESTOR_ID_FN] = tIdTokenPtr.UID
		tokenPayload[rcv.PAYLOAD_EXPIRES_FN] = tIdTokenPtr.Expires
		tokenPayload[rcv.PAYLOAD_ISSUER_FN] = tIdTokenPtr.Issuer
		tokenPayload[rcv.PAYLOAD_ISSUED_AT_FN] = tIdTokenPtr.IssuedAt
	} else {
		errorInfo.Error = errors.New(fmt.Sprintf("The provided idTokenPtr is invalid. ERROR: %v", errorInfo.Error.Error()))
	}

	return
}

// GetIdTokenPtr
func GetIdTokenPtr(authPtr *auth.Client, idToken string) (IdTokenPtr *auth.Token, errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if IdTokenPtr, errorInfo.Error = authPtr.VerifyIDToken(CTXBackground, idToken); errorInfo.Error != nil {
		log.Println(errorInfo.Error.Error())
	}

	return
}

// IsFirebaseIdTokenValid
func IsFirebaseIdTokenValid(authPtr *auth.Client, idToken string) bool {

	if _, err := authPtr.VerifyIDToken(CTXBackground, idToken); err != nil {
		log.Println(err.Error())
		return false
	}

	return true
}

// NewFirebaseApp - creates a new Firebase App
func NewFirebaseApp(credentialsLocation string) (appPtr *firebase.App, errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if appPtr, errorInfo.Error = firebase.NewApp(CTXBackground, nil, option.WithCredentialsFile(credentialsLocation)); errorInfo.Error != nil {
		log.Println(errorInfo.Error.Error(), rcv.ENDING_EXECUTION)
	}

	return
}

// GetFirebaseAuthConnection - creates a new Firebase Auth Connection
func GetFirebaseAuthConnection(appPtr *firebase.App) (authPtr *auth.Client, errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if authPtr, errorInfo.Error = appPtr.Auth(CTXBackground); errorInfo.Error != nil {
		log.Println(errorInfo.Error.Error(), rcv.ENDING_EXECUTION)
	} else {
		log.Println("The Firebase Auth client has been created.")
	}

	return
}

// SetFirebaseAuthEmailVerified - This will set the Firebase Auth email verify flag to true
func SetFirebaseAuthEmailVerified(authPtr *auth.Client, requestorId string) (errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tUserRecordPtr     *auth.UserRecord
	)

	cpi.PrintDebugTrail(tFunctionName)

	if tUserRecordPtr, errorInfo = FindFirebaseAuthUser(authPtr, requestorId); tUserRecordPtr != nil {
		params := (&auth.UserToUpdate{}).EmailVerified(true)
		if _, errorInfo.Error = authPtr.UpdateUser(CTXBackground, requestorId, params); errorInfo.Error != nil {
			errorInfo.Error = errors.New(fmt.Sprintf("Firebase Auth - Setting email verify to true, failed for Requestor Id: %v Error: %v", requestorId, errorInfo.Error))
			log.Println(errorInfo.Error.Error())
		}
	}

	return
}

// ValidateFirebaseJWTPayload - Firebase ID Token that is returned when a user logs on successfully
func ValidateFirebaseJWTPayload(tokenPayload map[any]interface{}, audience, issuer string) (errorInfo cpi.ErrorInfo) {

	var (
		tFindings          string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tRequestorId       string
		tSubject           string
	)

	cpi.PrintDebugTrail(tFunctionName)

	if tFindings = coreValidators.AreMapKeysValuesPopulated(tokenPayload); tFindings != rcv.GOOD {
		errorInfo.Error = cpi.GetMapKeyPopulatedError(tFindings)
	} else {
		if audience == rcv.EMPTY || issuer == rcv.EMPTY {
			errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v' %v: '%v'", rcv.FN_AUDIENCE, audience, rcv.FN_ISSUER, issuer))
		} else {
			for key, value := range tokenPayload {
				switch strings.ToUpper(key.(string)) {
				case rcv.PAYLOAD_AUDIENCE_FN:
					if value != audience {
						errorInfo.Error = errors.New("The audience of the ID Token is invalid.")
						log.Println(errorInfo.Error.Error())
					}
				case rcv.PAYLOAD_ISSUER_FN:
					if value != issuer {
						errorInfo.Error = errors.New("The issuer of the ID Token is invalid.")
						log.Println(errorInfo.Error.Error())
					}
				case rcv.PAYLOAD_SUBJECT_FN:
					tSubject = value.(string)
				case rcv.PAYLOAD_REQUESTOR_ID_FN:
					tRequestorId = value.(string)
				}
			}
			if tRequestorId != tSubject {
				errorInfo.Error = errors.New("The requestorId/user_id do not match the subject/sub. The ID is invalid.")
				log.Println(errorInfo.Error.Error())
			}
		}
	}

	return
}
