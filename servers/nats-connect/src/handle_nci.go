// Package src
/*
This is code for STY-Holdings NATS Connect

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
	"github.com/nats-io/nats.go"
	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

// nciMessageHandles - builds a map of message handlers
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) loadNCIMessageHandles() (
	handlers map[string]ns.MessageHandler,
) {

	handlers = make(map[string]ns.MessageHandler)

	handlers[NCI_TURN_DEBUG_OFF] = ns.MessageHandler{
		Handler: serverPtr.nciTurnDebugOff(),
	}
	handlers[NCI_TURN_DEBUG_ON] = ns.MessageHandler{
		Handler: serverPtr.nciTurnDebugOn(),
	}

	return
}

// NATS Message Handlers go below this line.

// nciTurnDebugOff - sets serverPtr.instance.debugModeOn to false
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) nciTurnDebugOff() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo pi.ErrorInfo
			tReply    ns.NATSReply
		)

		serverPtr.instance.debugModeOn = false
		tReply.Response = "Debug mode turned off"

		if errorInfo = ns.SendReply(tReply, msg); errorInfo.Error != nil {
			pi.PrintErrorInfo(errorInfo)
		}

		return
	}
}

// nciTurnDebugOn - puts the server into debug mode via a nats message
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) nciTurnDebugOn() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo pi.ErrorInfo
			tReply    ns.NATSReply
		)

		serverPtr.instance.debugModeOn = true
		tReply.Response = "Debug mode turned on"

		if errorInfo = ns.SendReply(tReply, msg); errorInfo.Error != nil {
			pi.PrintErrorInfo(errorInfo)
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
// 	pi.PrintDebugTrail(tFunctionName)
//
// 	return func(msg *nats.Msg) {
// 		var (
// 			tFunction, _, _, _ = runtime.Caller(0)
// 			tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		)
//
// 		pi.PrintRequestStart(2)
// 		pi.PrintDebugTrail(tFunctionName)
//
// 		executeGetBackendInfo(myServer, msg)
//
// 		return
// 	}
// }
