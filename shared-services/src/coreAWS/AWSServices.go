// Package coreAWS
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

<<<<<<< HEAD
	Copyright (c) 2022 STY-Holdings, Inc
=======
	Copyright (c) 2022 STY-Holdings, inc
>>>>>>> fbf9762 (Fixed the label)
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
	"crypto/rsa"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"math/big"
	"net/http"
	"os"
	"runtime"
	"strings"

	"albert/constants"
	"albert/core/coreFirestore"
	"cloud.google.com/go/firestore"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	awsSession "github.com/aws/aws-sdk-go/aws/session"
	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"

	"github.com/golang-jwt/jwt/v4"
)

// Config ...
type Config struct {
	CognitoRegion     string
	CognitoUserPoolID string
}

// KeySet
type KeySet struct {
	Keys []struct {
		Alg string `json:"alg"`
		E   string `json:"e"`
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		N   string `json:"n"`
	} `json:"keys"`
}

type AWSConfig struct {
	ProfileName string `json:"Profile_Name"`
	Region      string `json:"Region"`
	UserPoolId  string `json:"User_Pool_Id"`
}

type AWSHelper struct {
	InfoFQN    string
	KeySetURL  string
	KeySet     KeySet
	AWSConfig  AWSConfig
	SessionPtr *awsSession.Session
	tokenType  string
}

type Claims struct {
	AtHash              string `json:"at_hash"`
	AuthTime            int    `json:"auth_time"`
	CognitoUsername     string `json:"cognito:username"`
	Email               string `json:"email"`
	EmailVerified       bool   `json:"email_verified"`
	PhoneNumber         string `json:"phone_number"`
	PhoneNumberVerified bool   `json:"phone_number_verified"`
	TokenUse            string `json:"token_use"`
	UserName            string `json:"username"`
	jwt.RegisteredClaims
}

var (
	tTrueString = "true"
)

// NewAWSSession
//
//	Customer Messages: None
//	Errors:
//	Verifications: None
func NewAWSSession(awsInfoFQN string) (awsHelper AWSHelper, errorInfo cpi.ErrorInfo) {

	var (
		tAWSConfig         AWSConfig
		tData              []byte
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if awsInfoFQN == rcv.EMPTY {
		cpi.PrintDebugLine(tFunctionName, fmt.Sprintf("awsInfoFQN: %v", awsInfoFQN))
		errorInfo.Error = cpi.ErrRequiredArgumentMissing
		errorInfo.AdditionalInfo = fmt.Sprintf("AWS Credential FQN: '%v'", awsInfoFQN)
		cpi.PrintError(errorInfo)
	} else {
		if tData, errorInfo.Error = os.ReadFile(awsInfoFQN); errorInfo.Error == nil {
			cpi.PrintDebugLine(tFunctionName, fmt.Sprintf("tData: %v", tData))
			if errorInfo.Error = json.Unmarshal(tData, &tAWSConfig); errorInfo.Error == nil {
				cpi.PrintDebugLine(tFunctionName, fmt.Sprintf("tAWSConfig: %v", tAWSConfig))
				awsHelper.SessionPtr, errorInfo.Error = awsSession.NewSessionWithOptions(awsSession.Options{
					Config: aws.Config{
						Region: aws.String(tAWSConfig.Region),
					},
					Profile: tAWSConfig.ProfileName,
				})
				if errorInfo.Error == nil {
					cpi.PrintDebugLine(tFunctionName, "AWS Session started.")
					awsHelper.InfoFQN = awsInfoFQN
					awsHelper.AWSConfig = tAWSConfig
					awsHelper.KeySetURL = fmt.Sprintf("https://cognito-idp.%s.amazonaws.com/%s/.well-known/jwks.json", awsHelper.AWSConfig.Region, awsHelper.AWSConfig.UserPoolId)
					awsHelper.KeySet, errorInfo = getPublicKeySet(awsHelper.KeySetURL)
				} else {
					cpi.PrintDebugLine(tFunctionName, "AWS Session failed.")
					errorInfo.AdditionalInfo = fmt.Sprintf("AWS Account Info JSON file: '%v'", awsInfoFQN)
					cpi.PrintError(errorInfo)
				}
			} else {
				errorInfo.Error = cpi.ErrJSONInvalid
				errorInfo.AdditionalInfo = fmt.Sprintf("AWS Account Info JSON file: '%v'", awsInfoFQN)
				cpi.PrintError(errorInfo)
			}
		} else {
			errorInfo.Error = cpi.ErrServiceFailedAWS
			errorInfo.AdditionalInfo = fmt.Sprintf("Required AWS Account Info file %v has an issue.", awsInfoFQN)
			cpi.PrintError(errorInfo)
		}
	}

	return
}

// ConfirmUser - mark the AWS user as confirmed
func (a *AWSHelper) ConfirmUser(userName string) (errorInfo cpi.ErrorInfo) {

	var (
		tAdminConfirmSignUpInput    cognito.AdminConfirmSignUpInput
		tCognitoIdentityProviderPtr *cognito.CognitoIdentityProvider
		tFunction, _, _, _          = runtime.Caller(0)
		tFunctionName               = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if userName == rcv.EMPTY {
		errorInfo.Error = cpi.ErrRequiredArgumentMissing
		log.Println(errorInfo.Error)
	} else {
		tCognitoIdentityProviderPtr = cognito.New(a.SessionPtr)
		tAdminConfirmSignUpInput.Username = &userName
		tAdminConfirmSignUpInput.UserPoolId = &a.AWSConfig.UserPoolId
		if _, errorInfo.Error = tCognitoIdentityProviderPtr.AdminConfirmSignUp(&tAdminConfirmSignUpInput); errorInfo.Error != nil {
			// If the user is already confirmed, AWS will return an error, and do not care about this error.
			if strings.Contains(errorInfo.Error.Error(), rcv.STATUS_CONFIRMED) {
				errorInfo.Error = nil
			} else {
				if strings.Contains(errorInfo.Error.Error(), cpi.USER_DOES_NOT_EXIST) {
					errorInfo.Error = cpi.ErrUserMissing
				}
			}
		}
	}

	return
}

// GetRequestorEmailPhoneFromIdTokenClaims - will validate the AWS Id JWT, check to make sure the email has been verified, and return the requestor id, email, and phone number.
func (a *AWSHelper) GetRequestorEmailPhoneFromIdTokenClaims(firestoreClientPtr *firestore.Client, token string) (requestorId, email, phoneNumber string, errorInfo cpi.ErrorInfo) {

	var (
		tClaimsPtr         *Claims
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if token == rcv.EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Token: '%v'", token))
		log.Println(errorInfo.Error)
	} else {
		if tClaimsPtr, errorInfo = getTokenClaims(a, rcv.TOKEN_TYPE_ID, token); errorInfo.Error == nil {
			if isTokenValid(firestoreClientPtr, a, rcv.TOKEN_TYPE_ID, token) {
				requestorId = tClaimsPtr.Subject
				email = tClaimsPtr.Email
				phoneNumber = tClaimsPtr.PhoneNumber
			} else {
				errorInfo.Error = cpi.ErrTokenInvalid
				log.Println(errorInfo.Error)
			}
		}
	}

	return
}

// GetRequestorFromAccessTokenClaims - will valid the AWS Access JWT, and return the requestor id.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (a *AWSHelper) GetRequestorFromAccessTokenClaims(firestoreClientPtr *firestore.Client, token string) (requestorId string, errorInfo cpi.ErrorInfo) {

	var (
		tClaimsPtr         *Claims
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if token == rcv.TEST_STRING {
		requestorId = rcv.TEST_USERNAME_SAVUP_REQUESTOR_ID
	} else {
		if token == rcv.EMPTY {
			errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Token: '%v'", token))
			log.Println(errorInfo.Error)
		} else {
			if tClaimsPtr, errorInfo = getTokenClaims(a, rcv.TOKEN_TYPE_ACCESS, token); errorInfo.Error == nil {
				if isTokenValid(firestoreClientPtr, a, rcv.TOKEN_TYPE_ACCESS, token) {
					requestorId = tClaimsPtr.Subject
				} else {
					errorInfo.Error = cpi.ErrTokenInvalid
					log.Println(errorInfo.Error)
				}
			}
		}
	}

	return
}

// ParseAWSJWTWithClaims - will return an err if the AWS JWT is invalid.
func (a *AWSHelper) ParseAWSJWTWithClaims(tokenType, tokenString string) (claimsPtr *Claims, errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if tokenString == rcv.EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Token: '%v'", tokenString))
		log.Println(errorInfo.Error)
	} else {
		if _, errorInfo.Error = jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (key interface{}, err error) {
			switch strings.ToUpper(tokenType) {
			case rcv.TOKEN_TYPE_ID:
				key, err = convertKey(a.KeySet.Keys[0].E, a.KeySet.Keys[0].N)
			case rcv.TOKEN_TYPE_ACCESS:
				key, err = convertKey(a.KeySet.Keys[1].E, a.KeySet.Keys[1].N)
			}
			claimsPtr = token.Claims.(*Claims)
			return
		}); errorInfo.Error != nil {
			log.Println(errorInfo.Error)
		}
	}

	return
}

// ParseAWSJWT - will return an err if the AWS JWT is invalid.
func (a *AWSHelper) ParseAWSJWT(tokenString string) (tTokenPtr *jwt.Token, errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if tokenString == rcv.EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Token: '%v'", tokenString))
		fmt.Println(errorInfo.Error)
	} else {
		tTokenPtr, errorInfo.Error = jwt.Parse(tokenString, func(token *jwt.Token) (key interface{}, err error) {
			key, err = convertKey(a.KeySet.Keys[1].E, a.KeySet.Keys[1].N)
			return
		})
	}

	return
}

// PullCognitoUserInfo - pull user information from the Cognito user pool.
func (a *AWSHelper) PullCognitoUserInfo(username string) (userData map[string]interface{}, errorInfo cpi.ErrorInfo) {

	var (
		tAdminGetUserInput          cognito.AdminGetUserInput
		tAdminGetUserOutputPtr      *cognito.AdminGetUserOutput
		tCognitoIdentityProviderPtr *cognito.CognitoIdentityProvider
		tFunction, _, _, _          = runtime.Caller(0)
		tFunctionName               = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if username == rcv.EMPTY {
		errorInfo.Error = cpi.ErrRequiredArgumentMissing
		cpi.PrintError(errorInfo)
	} else {
		tCognitoIdentityProviderPtr = cognito.New(a.SessionPtr)
		if tCognitoIdentityProviderPtr == nil {
			errorInfo.FileName, errorInfo.ErrorLineNumber = cpi.GetFileLineNumber()
			errorInfo.Error = cpi.ErrServiceFailedAWS
			cpi.PrintError(errorInfo)
		} else {
			// Set up the request to get the user
			tAdminGetUserInput.UserPoolId = &a.AWSConfig.UserPoolId
			tAdminGetUserInput.Username = &username
			// Make the request to get the user
			if tAdminGetUserOutputPtr, errorInfo.Error = tCognitoIdentityProviderPtr.AdminGetUser(&tAdminGetUserInput); errorInfo.Error == nil {
				userData = make(map[string]interface{})
				for _, attr := range tAdminGetUserOutputPtr.UserAttributes {
					userData[*attr.Name] = *attr.Value
				}
			} else {
				errorInfo.FileName, errorInfo.ErrorLineNumber = cpi.GetFileLineNumber()
				errorInfo.Error = cpi.ErrServiceFailedAWS
				cpi.PrintError(errorInfo)
			}
		}
	}

	return
}

// ValidAWSJWT - will valid the AWS JWT and check to make sure either the phone or email has been verified.
func (a *AWSHelper) ValidAWSJWT(firestoreClientPtr *firestore.Client, tokenType, token string) (valid bool, errorInfo cpi.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	cpi.PrintDebugTrail(tFunctionName)

	if token == rcv.EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Token: '%v'", token))
		log.Println(errorInfo.Error)
	} else {
		valid = isTokenValid(firestoreClientPtr, a, tokenType, token)
	}

	return
}

// UpdateAWSEmailVerifyFlag - will update the email_valid field for the user in the Cognito user pool.
func (a *AWSHelper) UpdateAWSEmailVerifyFlag(username string) (errorInfo cpi.ErrorInfo) {

	var (
		tAdminUpdateUserAttributesInput cognito.AdminUpdateUserAttributesInput
		tAttributeType                  cognito.AttributeType
		tAttributeTypePtrs              []*cognito.AttributeType
		tCognitoIdentityProviderPtr     *cognito.CognitoIdentityProvider
		tFunction, _, _, _              = runtime.Caller(0)
		tFunctionName                   = runtime.FuncForPC(tFunction).Name()
		tName                           string
	)

	cpi.PrintDebugTrail(tFunctionName)

	if username == rcv.EMPTY {
		errorInfo.Error = cpi.ErrRequiredArgumentMissing
	} else {
		tCognitoIdentityProviderPtr = cognito.New(a.SessionPtr)
		tName = rcv.FN_EMAIL_VERIFIED // This is required because go doesn't support pointers to rcv.
		tAttributeType = cognito.AttributeType{
			Name:  &tName,
			Value: &tTrueString,
		}
		tAttributeTypePtrs = append(tAttributeTypePtrs, &tAttributeType)
		tAdminUpdateUserAttributesInput.UserAttributes = tAttributeTypePtrs
		tAdminUpdateUserAttributesInput.Username = &username
		tAdminUpdateUserAttributesInput.UserPoolId = &a.AWSConfig.UserPoolId
		req, _ := tCognitoIdentityProviderPtr.AdminUpdateUserAttributesRequest(&tAdminUpdateUserAttributesInput)
		errorInfo.Error = req.Send()
	}

	return
}

// ResetUserPassword - trigger one-time code to be set to user email.
func (a *AWSHelper) ResetUserPassword(userName string, test bool) (errorInfo cpi.ErrorInfo) {

	var (
		tAdminResetUserPasswordInput cognito.AdminResetUserPasswordInput
		tCognitoIdentityProviderPtr  *cognito.CognitoIdentityProvider
		tFunction, _, _, _           = runtime.Caller(0)
		tFunctionName                = runtime.FuncForPC(tFunction).Name()
		req                          *request.Request
	)

	cpi.PrintDebugTrail(tFunctionName)

	if userName == rcv.EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! AWS User Name: '%v'", userName))
		log.Println(errorInfo.Error)
	} else {
		tCognitoIdentityProviderPtr = cognito.New(a.SessionPtr)
		tAdminResetUserPasswordInput.Username = &userName
		tAdminResetUserPasswordInput.UserPoolId = &a.AWSConfig.UserPoolId
		if test == false {
			req, _ = tCognitoIdentityProviderPtr.AdminResetUserPasswordRequest(&tAdminResetUserPasswordInput)
			errorInfo.Error = req.Send()
		}
	}

	return
}

// areAWSClaimsValid - Checks if email is verified and token is either an Id or Access token.
func areAWSClaimsValid(FirestoreClientPtr *firestore.Client, subject, email, username, tokenUse string, emailVerified bool) bool {

	var (
		errorInfo          cpi.ErrorInfo
		tDocumentPtr       *firestore.DocumentSnapshot
		tEmailInterface    interface{}
		tSubjectInterface  interface{}
		tUsernameInterface interface{}
	)

	if _, tDocumentPtr, errorInfo = coreFirestore.FindDocument(FirestoreClientPtr, rcv.DATASTORE_USERS, coreFirestore.NameValueQuery{
		FieldName:  rcv.FN_REQUESTOR_ID,
		FieldValue: subject,
	}); errorInfo.Error == nil {
		switch strings.ToUpper(tokenUse) {
		case rcv.TOKEN_TYPE_ID:
			if tSubjectInterface, errorInfo.Error = tDocumentPtr.DataAt(rcv.FN_REQUESTOR_ID); errorInfo.Error == nil {
				if tUsernameInterface, errorInfo.Error = tDocumentPtr.DataAt(rcv.FN_USERNAME); errorInfo.Error == nil {
					if tEmailInterface, errorInfo.Error = tDocumentPtr.DataAt(rcv.FN_EMAIL); errorInfo.Error == nil {
						if emailVerified && tSubjectInterface.(string) == subject && tEmailInterface.(string) == email && tUsernameInterface.(string) == username {
							return true
						}
					}
				}
			}
		case rcv.TOKEN_TYPE_ACCESS:
			if tSubjectInterface, errorInfo.Error = tDocumentPtr.DataAt(rcv.FN_REQUESTOR_ID); errorInfo.Error == nil {
				if tUsernameInterface, errorInfo.Error = tDocumentPtr.DataAt(rcv.FN_USERNAME); errorInfo.Error == nil {
					if emailVerified && tSubjectInterface.(string) == subject && tUsernameInterface.(string) == username {
						return true
					}
				}
			}
		}
	}

	return false
}

// convertKey - does not follow the errorInfo format because it is called by a function that only allows error to be returned
func convertKey(rawE, rawN string) (publicKey *rsa.PublicKey, err error) {

	var (
		decodedN      []byte
		decodedBase64 []byte
		ndata         []byte
	)

	decodedBase64, err = base64.RawURLEncoding.DecodeString(rawE)

	if err == nil {
		if len(decodedBase64) < 4 {
			ndata = make([]byte, 4)
			copy(ndata[4-len(decodedBase64):], decodedBase64)
			decodedBase64 = ndata
		}
		publicKey = &rsa.PublicKey{
			N: &big.Int{},
			E: int(binary.BigEndian.Uint32(decodedBase64[:])),
		}
		if decodedN, err = base64.RawURLEncoding.DecodeString(rawN); err == nil {
			publicKey.N.SetBytes(decodedN)
		}
	}

	return
}

// getPublicKeySet
func getPublicKeySet(keySetURL string) (keySet KeySet, errorInfo cpi.ErrorInfo) {

	var (
		tBody              []byte
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tKeySetPtr         *http.Response
	)

	cpi.PrintDebugTrail(tFunctionName)

	if keySetURL == rcv.EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Key Set URL: '%v'", keySetURL))
		log.Println(errorInfo.Error)
	} else {
		if tKeySetPtr, errorInfo.Error = http.Get(keySetURL); errorInfo.Error == nil {
			if tKeySetPtr.StatusCode == rcv.HTTP_STATUS_200 {
				if tBody, errorInfo.Error = io.ReadAll(tKeySetPtr.Body); errorInfo.Error == nil {
					if errorInfo.Error = json.Unmarshal(tBody, &keySet); errorInfo.Error != nil {
						errorInfo.Error = errors.New(fmt.Sprintf("Get Public Key's KeySet is corrupt. Response Body: '%v' Error: %v", tBody, errorInfo.Error))
						log.Println(errorInfo.Error)
					}
				}
			} else {
				errorInfo.Error = errors.New(fmt.Sprintf("Fetching the public key has failed! Key Set URL: '%v' Status Code: '%v'", keySetURL, tKeySetPtr.StatusCode))
				log.Println(errorInfo.Error)
			}
		}
	}

	return
}

func isTokenValid(firestoreClientPtr *firestore.Client, a *AWSHelper, tokenType, token string) bool {

	var (
		errorInfo  cpi.ErrorInfo
		tClaimsPtr *Claims
	)

	a.tokenType = tokenType
	if tClaimsPtr, errorInfo = getTokenClaims(a, tokenType, token); errorInfo.Error == nil {
		switch strings.ToUpper(tClaimsPtr.TokenUse) {
		case rcv.TOKEN_TYPE_ID:
			return areAWSClaimsValid(firestoreClientPtr, tClaimsPtr.Subject, tClaimsPtr.Email, tClaimsPtr.CognitoUsername, tClaimsPtr.TokenUse, tClaimsPtr.EmailVerified)
		case rcv.TOKEN_TYPE_ACCESS:
			return areAWSClaimsValid(firestoreClientPtr, tClaimsPtr.Subject, rcv.EMPTY, tClaimsPtr.UserName, tClaimsPtr.TokenUse, true)
		}
	}

	return false
}

func getTokenClaims(a *AWSHelper, tokenType, token string) (claimsPtr *Claims, errorInfo cpi.ErrorInfo) {

	return a.ParseAWSJWTWithClaims(tokenType, token)
}
