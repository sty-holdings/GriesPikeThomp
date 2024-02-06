// Package styhHTTP
/*
This is the core code for STY-Holdings SavUp NATS services

RESTRICTIONS:

	    AWS functions:
		Not used

		Google Cloud Platform:
		None

	   	Ports
		* 3000 and 8000 must be open to communicate with Plaid and the client
		* NATS standard ports

NOTES:

	None

COPYRIGHT:

	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package styhHTTP

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type httpReply struct {
	Data    string `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

type HTTPResponse struct {
	Data            string `json:"data"`
	CustomerMessage string `json:"message,omitempty"`
	Error           string `json:"error,omitempty"`
}

// SetRoutes - builds the routing for http messages
//
//	Customer Messages: None
//	Errors: ErrSubjectSubscriptionFailed
//	Verifications: None
func SetRoutes(ginMode string) (router *gin.Engine) {

	gin.SetMode(ginMode)
	router = gin.Default()

	// Register request handler function
	router.POST(fmt.Sprintf("/%v", ENDPOINT_CREATE_ACCOUNT), createAccount)
	// router.GET(fmt.Sprintf("/%v", ENDPOINT_FORGOT_USERNAME), src.forgotUserName)
	// router.GET(fmt.Sprintf("/%v", ENDPOINT_VERIFY_EMAIL), src.processEmailVerification)
	// router.GET(fmt.Sprintf("/%v", ENDPOINT_PULL_USER), src.pullUser)
	// router.GET(fmt.Sprintf("/%v", ENDPOINT_RESEND_VERIFY_EMAIL), src.resendVerifyEmail)
	// router.GET(fmt.Sprintf("/%v", ENDPOINT_RESET_USER_PASSWORD), src.resetUserPassword)
	// router.GET(fmt.Sprintf("/%v", ENDPOINT_GET_HTTP_INFO), src.getHTTPInfo)
	// router.GET(fmt.Sprintf("/%v", ENDPOINT_GET_BACKEND_INFO), src.getBackendInfo)

	// router.GET("/shutdown", func(c *gin.Context) {
	// 	// Todo run systemctl stop savup-http from shell
	//
	// 	// Write response
	// 	tResponse := fmt.Sprintf(`{"status": %v, "error": %v}`, rcv.STATUS_SUCCESS, rcv.TXT_EMPTY)
	// 	c.String(200, tResponse)
	// })
	//
	// router.GET("/listRoutes", func(c *gin.Context) {
	// 	// Todo run systemctl stop savup-http from shell
	//
	// 	// Write response
	// 	tResponse := fmt.Sprintf(`{"status": %v, "error": %v}`, rcv.STATUS_SUCCESS, rcv.TXT_EMPTY)
	// 	c.String(200, tResponse)
	// })

	return
}

// styh-http Request Handlers go below this line and they are private

// createAccount - will send out a NATS message to create the account on Cognito
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func createAccount(c *gin.Context) {

	var (
		tHTTPResponse HTTPResponse
	)

	log.Printf("YOUR ARE HERE!!!!")

	tHTTPResponse.Data = "TEST Page"
	sendMobileResponse(c, tHTTPResponse)
}

func sendMobileResponse(c *gin.Context, httpResponse HTTPResponse) {

	if httpResponse.Error == rcv.VAL_EMPTY {
		c.String(rcv.HTTP_STATUS_200, writeHTTPResponse(httpResponse))
	} else {
		c.String(rcv.HTTP_STATUS_400, writeHTTPResponse(httpResponse))
	}
}

func writeHTTPResponse(httpResponse HTTPResponse) string {

	var (
		err           error
		tJSONResponse []byte
	)

	if tJSONResponse, err = json.Marshal(httpResponse); err != nil {
		log.Println(err.Error())
	}

	return string(tJSONResponse)
}
