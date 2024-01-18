// Package coreJWT
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
package coreJWT

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"
	"runtime"
	"strings"
	"time"

	"albert/constants"
	"albert/core/coreError"
	"albert/core/coreValidators"
	jwt2 "github.com/golang-jwt/jwt/v4"
)

type GenerateCertificate struct {
	CertFileName       string
	Certificate        []byte
	Host               string
	PublicKey          crypto.PublicKey
	PrivateKey         crypto.PrivateKey
	PrivateKeyFileName string
	RSABits            int
	SelfCA             bool
	ValidFor           string
}

// TLS Options for SavUp server.
type TLSInfo struct {
	TLSCert     string `json:"TLS_Certificate_FQN"`
	TLSKey      string `json:"TLS_Private_Key_FQN"`
	TLSCABundle string `json:"TLS_CABundle_FQN"`
}

// GenerateJWT
// Create a new token object, specifying signing method and the claims
// you would like it to contain.
func GenerateJWT(privateKey, requestorId, period string, duration int64) (jwt string, errorInfo coreError.ErrorInfo) {

	var (
		tDuration          time.Duration
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tPrivateKey        *rsa.PrivateKey
		tRawPrivateKey     []byte
	)

	coreError.PrintDebugTrail(tFunctionName)

	if privateKey == constants.EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v'", constants.FN_PRIVATE_KEY, constants.EMPTY))
		log.Println(errorInfo.Error)
	} else {
		if requestorId == constants.EMPTY || period == constants.EMPTY || duration < 1 {
			errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v' %v: '%v' %v: '%v'", constants.FN_REQUESTOR_ID, requestorId, constants.FN_PERIOD, period, constants.FN_DURATION, duration))
			log.Println(errorInfo.Error)
		} else {
			if coreValidators.IsPeriodValid(period) && duration > 0 {
				tRawPrivateKey = []byte(privateKey)
				if tPrivateKey, errorInfo = ParsePrivateKey(tRawPrivateKey); errorInfo.Error == nil {
					switch strings.ToUpper(period) {
					case "M":
						tDuration = time.Minute * time.Duration(duration)
					case "H":
						tDuration = time.Hour * time.Duration(duration)
					case "D":
						tDuration = time.Hour * time.Duration(duration*24)
					default:
						tDuration = time.Hour * time.Duration(duration)
					}
					jwt, errorInfo.Error = jwt2.NewWithClaims(jwt2.SigningMethodRS512, jwt2.MapClaims{
						"requestorId": requestorId,
						"Issuer":      constants.CERT_ISSUER,
						"Subject":     requestorId,
						"ExpiresAt":   time.Now().Add(tDuration).String(),
						"NotBefore":   time.Now(),
					}).SignedString(tPrivateKey)
				}
			}
		}
	}

	return
}

// GenerateRSAKey
func GenerateRSAKey(rsaBits int) (privateKey crypto.PrivateKey, publicKey crypto.PublicKey, errorInfo coreError.ErrorInfo) {

	var (
		_PrivateKey        *rsa.PrivateKey
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	if _PrivateKey, errorInfo.Error = rsa.GenerateKey(rand.Reader, rsaBits); errorInfo.Error != nil {
		log.Println(errorInfo.Error)
	}

	if errorInfo.Error == nil {
		// The public key is a part of the *rsa.PrivateKey struct
		publicKey = _PrivateKey.Public()
		privateKey = _PrivateKey
	}

	return
}

// ParsePrivateKey
func ParsePrivateKey(tRawPrivateKey []byte) (privateKey *rsa.PrivateKey, errorInfo coreError.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	if privateKey, errorInfo.Error = jwt2.ParseRSAPrivateKeyFromPEM(tRawPrivateKey); errorInfo.Error != nil {
		errorInfo.Error = errors.New("Unable to parse the private key referred to in the configuration file.")
		log.Println(errorInfo.Error)
	}

	return
}

// ValidateSavUpJWT
func ValidateSavUpJWT(privateKey []byte, jwt string) (errorInfo coreError.ErrorInfo) {

	var (
		tParsedPrivateKey  *rsa.PrivateKey
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		//  ToDo is this needed?
		// tToken         *jwt2.Token
	)

	coreError.PrintDebugTrail(tFunctionName)

	if len(privateKey) == 0 || jwt == constants.EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v' %v: '%v'", constants.FN_PRIVATE_KEY, privateKey, constants.FN_JWT, jwt))
		log.Println(errorInfo.Error)
	} else {
		if tParsedPrivateKey, errorInfo = ParsePrivateKey(privateKey); errorInfo.Error == nil {
			publicKey := tParsedPrivateKey.Public()
			//  ToDo is this needed?
			// tToken, err = jwt2.Parse(jwt, func(jwtToken *jwt2.Token) (interface{}, error) {
			if _, errorInfo.Error = jwt2.Parse(jwt, func(jwtToken *jwt2.Token) (interface{}, error) {
				if _, ok := jwtToken.Method.(*jwt2.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("unexpected method: %s", jwtToken.Header["alg"])
				}
				return publicKey, nil
			}); errorInfo.Error != nil {
				log.Println(errorInfo.Error)
			}
		}
	}

	return
}
