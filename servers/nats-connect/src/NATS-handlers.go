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

	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	ns "GriesPikeThomp/shared-services/src/coreNATS"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/nats-io/nats.go"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type natsReply struct {
	Data    string `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
	Message string `json:"message,omitempty"`
	Status  string `json:"status"`
}

// getNATSHandlers - builds the NATS message handlers
//
//	Customer Messages: None
//	Errors: ErrSubjectSubscriptionFailed
//	Verifications: None
func (serverPtr *Server) getNATSHandlers(service ns.NATSService) (errorInfo cpi.ErrorInfo) {

	var (
		connPtr *nats.Conn
	)

	connPtr = serverPtr.extensions[NATS_INTERNAL].(ns.NATSService).ConnPtr

	for _, subjectInfo := range service.Config.SubjectRegistry {
		switch strings.ToLower(subjectInfo.Subject) {
		case TURN_DEBUG_ON:
			serverPtr.instance.messageHandlers[TURN_DEBUG_ON] = serverPtr.natsTurnDebugOn()
		case TURN_DEBUG_OFF:
			serverPtr.instance.messageHandlers[TURN_DEBUG_OFF] = serverPtr.natsTurnDebugOff()
		default:
			errorInfo = cpi.NewErrorInfo(cpi.ErrSubjectInvalid, fmt.Sprintf("%v%v", rcv.TXT_SUBJECT, subjectInfo.Subject))
		}
		if errorInfo.Error == nil {
			if serverPtr.instance.subscriptionPtrs[subjectInfo.Subject], errorInfo.Error = connPtr.Subscribe(subjectInfo.Subject, serverPtr.instance.messageHandlers[subjectInfo.Subject]); errorInfo.Error != nil {
				errorInfo = cpi.NewErrorInfo(cpi.ErrSubjectSubscriptionFailed, fmt.Sprintf("%v%v", rcv.TXT_SUBJECT, subjectInfo.Subject))
			}
		}
	}

	return
}

// NATS Message Handlers go below this line.
//

// natsTurnDebugOff - removes the server out of debug mode via a nats message
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) natsTurnDebugOff() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo cpi.ErrorInfo
			tReply    natsReply
		)

		serverPtr.instance.debugModeOn = false
		tReply.Status = rcv.STATUS_SUCCESS

		if errorInfo = chv.SendReply(tReply, msg); errorInfo.Error != nil {
			cpi.PrintErrorInfo(errorInfo)
		}

		return
	}
}

// natsTurnDebugOn - puts the server into debug mode via a nats message
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) natsTurnDebugOn() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo cpi.ErrorInfo
			tReply    natsReply
		)

		serverPtr.instance.debugModeOn = true
		tReply.Status = rcv.STATUS_SUCCESS

		if errorInfo = chv.SendReply(tReply, msg); errorInfo.Error != nil {
			cpi.PrintErrorInfo(errorInfo)
		}

		return
	}
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
