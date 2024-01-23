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
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"errors"
	"fmt"
	"log"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	jwt2 "github.com/golang-jwt/jwt/v4"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
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

// TLSInfo files
type TLSInfo struct {
	TLSCert       string `json:"TLS_Certificate_FQN"`
	TLSPrivateKey string `json:"TLS_Private_Key_FQN"`
	TLSCABundle   string `json:"TLS_CABundle_FQN"`
}

// GenerateJWT
// Create a new token object, specifying signing method and the claims
// you would like it to contain.
// func GenerateJWT(privateKey, requestorId, period string, duration int64) (jwt string, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tDuration      time.Duration
// 		tPrivateKey    *rsa.PrivateKey
// 		tRawPrivateKey []byte
// 	)
//
// 	if privateKey == rcv.VAL_EMPTY {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v'", rcv.FN_PRIVATE_KEY, rcv.VAL_EMPTY))
// 		log.Println(errorInfo.Error)
// 	} else {
// 		if requestorId == rcv.VAL_EMPTY || period == rcv.VAL_EMPTY || duration < 1 {
// 			errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v' %v: '%v' %v: '%v'", rcv.FN_REQUESTOR_ID, requestorId, rcv.FN_PERIOD, period, rcv.FN_DURATION, duration))
// 			log.Println(errorInfo.Error)
// 		} else {
// 			if cv.IsPeriodValid(period) && duration > 0 {
// 				tRawPrivateKey = []byte(privateKey)
// 				if tPrivateKey, errorInfo = ParsePrivateKey(tRawPrivateKey); errorInfo.Error == nil {
// 					switch strings.ToUpper(period) {
// 					case "M":
// 						tDuration = time.Minute * time.Duration(duration)
// 					case "H":
// 						tDuration = time.Hour * time.Duration(duration)
// 					case "D":
// 						tDuration = time.Hour * time.Duration(duration*24)
// 					default:
// 						tDuration = time.Hour * time.Duration(duration)
// 					}
// 					jwt, errorInfo.Error = jwt2.NewWithClaims(jwt2.SigningMethodRS512, jwt2.MapClaims{
// 						"requestorId": requestorId,
// 						"Issuer":      rcv.CERT_ISSUER,
// 						"Subject":     requestorId,
// 						"ExpiresAt":   time.Now().Add(tDuration).String(),
// 						"NotBefore":   time.Now(),
// 					}).SignedString(tPrivateKey)
// 				}
// 			}
// 		}
// 	}
//
// 	return
// }

// GenerateRSAKey
func GenerateRSAKey(rsaBits int) (privateKey crypto.PrivateKey, publicKey crypto.PublicKey, errorInfo cpi.ErrorInfo) {

	var (
		_PrivateKey *rsa.PrivateKey
	)

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
func ParsePrivateKey(tRawPrivateKey []byte) (privateKey *rsa.PrivateKey, errorInfo cpi.ErrorInfo) {

	if privateKey, errorInfo.Error = jwt2.ParseRSAPrivateKeyFromPEM(tRawPrivateKey); errorInfo.Error != nil {
		errorInfo.Error = errors.New("Unable to parse the private key referred to in the configuration file.")
		log.Println(errorInfo.Error)
	}

	return
}

// ValidateSavUpJWT
func ValidateSavUpJWT(privateKey []byte, jwt string) (errorInfo cpi.ErrorInfo) {

	var (
		tParsedPrivateKey *rsa.PrivateKey
		//  ToDo is this needed?
		// tToken         *jwt2.Token
	)

	if len(privateKey) == 0 || jwt == rcv.VAL_EMPTY {
		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v' %v: '%v'", rcv.FN_PRIVATE_KEY, privateKey, rcv.FN_JWT, jwt))
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
