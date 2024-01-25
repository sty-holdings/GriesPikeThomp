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

// signal handler
// func (myServer *Server) signalHandler(signal os.Signal) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintRequestStart(2)
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	log.Printf("Caught signal: %+v\n", signal)
// 	switch signal {
// 	case syscall.SIGHUP: // kill -SIGHUP XXXX
// 		fallthrough
// 	case syscall.SIGINT: // kill -SIGINT XXXX or Ctrl+c
// 		fallthrough
// 	case syscall.SIGTERM: // kill -SIGTERM XXXX
// 		fallthrough
// 	case syscall.SIGQUIT: // kill -SIGQUIT XXXX
// 		myServer.Shutdown(false)
// 	default:
// 		log.Printf("%v - unknown signal", signal)
// 	}
//
// 	fmt.Println("\nFinished server cleanup")
// 	os.Exit(0)
// }

// unsubscribeMessages - will turn off all registered messages and output a message as such to the log.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func (myServer *Server) unsubscribeMessages() {
//
// 	var (
// 		err                error
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintRequestStart(2)
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	// Unsubscribing to other messages.
// 	for subject, message := range myServer.messages {
// 		if err = message.subscriptionPtr.Unsubscribe(); err == nil {
// 			log.Printf("Unsubscribing to %v.", subject)
// 		} else {
// 			log.Println(err.Error())
// 		}
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
