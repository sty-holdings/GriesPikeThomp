// Package src
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
package src

import (
	"fmt"
	"strings"

	ns "GriesPikeThomp/shared-services/src/coreNATS"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/nats-io/nats.go"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// turnDebugOn - puts the server into debug mode
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) turnDebugOn() nats.MsgHandler {

	return func(msg *nats.Msg) {
		fmt.Println("On")

		serverPtr.instance.debugModeOn = true
		return
	}
}

// turnDebugOff - removes the server out of debug mode
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) turnDebugOff() nats.MsgHandler {

	return func(msg *nats.Msg) {
		fmt.Println("Off")

		serverPtr.instance.debugModeOn = false
		return
	}
}

func (serverPtr *Server) getHandlers(service ns.NATSService) (errorInfo cpi.ErrorInfo) {

	var (
		connection *nats.Conn
	)

	connection = serverPtr.extensions[NATS_INTERNAL].(ns.NATSService).ConnPtr

	for _, handler := range service.Config.SubjectRegistry {
		fmt.Printf("\nSubject: %v Description: %v\n", handler.Subject, handler.Description)
		switch strings.ToLower(handler.Subject) {
		case TURN_DEBUG_ON:
			serverPtr.instance.messageHandlers[TURN_DEBUG_ON] = serverPtr.turnDebugOn()
		case TURN_DEBUG_OFF:
			serverPtr.instance.messageHandlers[TURN_DEBUG_OFF] = serverPtr.turnDebugOff()
		default:
			errorInfo = cpi.NewErrorInfo(cpi.ErrSubjectInvalid, fmt.Sprintf("%v%v", rcv.TXT_SUBJECT, handler.Subject))
		}
		if errorInfo.Error == nil {
			if serverPtr.instance.subscriptionPtrs[handler.Subject], errorInfo.Error = connection.Subscribe(handler.Subject, serverPtr.instance.messageHandlers[handler.Subject]); errorInfo.Error != nil {
				errorInfo = cpi.NewErrorInfo(cpi.ErrSubjectSubscriptionFailed, fmt.Sprintf("%v%v", rcv.TXT_SUBJECT, handler.Subject))
			}
		}
	}

	return
}

// getMessagePrefix
// func (server *Server) getMessagePrefix() string {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	return server.extensionPtrs[].messagePrefix
// }

// getServerInfo
// func (myServer *Server) getBackendInfo() nats.MsgHandler {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	return func(msg *nats.Msg) {
// 		var (
// 			tFunction, _, _, _ = runtime.Caller(0)
// 			tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		)
//
// 		cpi.PrintRequestStart(2)
// 		cpi.PrintDebugTrail(tFunctionName)
//
// 		executeGetBackendInfo(myServer, msg)
//
// 		return
// 	}
// }

// userVerification
// func (myServer *Server) userVerification() nats.MsgHandler {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	return func(msg *nats.Msg) {
// 		var (
// 			tFunction, _, _, _ = runtime.Caller(0)
// 			tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		)
//
// 		cpi.PrintRequestStart(2)
// 		cpi.PrintDebugTrail(tFunctionName)
//
// 		executeUserVerification(myServer.authenticatorService, myServer.MyAWS, myServer.MyFireBase, msg)
// 	}
// }
