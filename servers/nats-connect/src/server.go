// Package src
/*
This is the core code for STY-Holdings SavUp NATS services

RESTRICTIONS:
	None

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
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"

	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	h "GriesPikeThomp/shared-services/src/coreHelpers"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/nats-io/nats.go"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type Auth struct {
	authenticatorService string
}

// Instance - Some of these values can change over the life of the instance.
type Instance struct {
	baseURL           string
	debugMode         bool
	hostname          string
	logFileHandlerPtr *os.File
	logFQN            string
	mu                sync.RWMutex
	numberCPUS        int
	pid               int
	pidFQN            string
	processChannel    chan string
	running           bool
	secured           bool
	runStartTime      time.Time
	version           string
	waitGroup         sync.WaitGroup
	workingDirectory  string
}

type NATSInfo struct {
	credentialsFilename string
	messageEnvironment  string
	messagePrefix       string
	url                 string
	port                string
}

type natsMessage struct {
	restrictedUsage bool
	handler         nats.MsgHandler
	subscriptionPtr *nats.Subscription
}

type Server struct {
	config          cc.Configuration
	instanceDetails Instance
	natsInfo        NATSInfo
	// MyAWS               coreAWS.AWSHelper
	// MyFireBase          coreHelpers.FirebaseFirestoreHelper
	// MyPlaid             PlaidHelper
	// MySendGrid          coreSendGrid.SendGridHelper
	// MyStripe            StripeHelper
	// sendGridTemplateIds SendGridTemplateIds
	// tls                 coreJWT.TLSInfo
	// verifyEmailURL      string
	// WebAssetsURL        string
}

// NewServer will set up a new server struct after parsing the configuration defined
// in the supplied configuration file. If no configuration file is provide, an
// error will be returned. If the configuration is invalid, an error will be returned.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewServer(configFileFQN, version string, test bool) (server *Server, returnCode int) {

	var (
		errorInfo cpi.ErrorInfo
		tConfig   cc.Configuration
	)

	if version == rcv.VAL_EMPTY && test == false {
		cpi.PrintError(cpi.ErrVersionInvalid, fmt.Sprintf("%v %v", rcv.TXT_SERVER_VERSION, version), rcv.MODE_OUTPUT_DISPLAY)
		return nil, 1
	}

	// See if configuration file exists and is readable, if not, return an error
	if tConfig, errorInfo = cc.ReadAndParseConfigFile(configFileFQN); errorInfo.Error != nil {
		cpi.PrintErrorInfo(errorInfo, rcv.MODE_OUTPUT_DISPLAY)
		return nil, 1
	}

	// Adjusting non-fully qualified filenames to fully qualified
	tConfig.LogDirectory = h.PrependWorkingDirectory(tConfig.LogDirectory)
	tConfig.PIDDirectory = h.PrependWorkingDirectory(tConfig.PIDDirectory)

	if errorInfo = cc.ValidateConfiguration(tConfig); errorInfo.Error != nil {
		cpi.PrintErrorInfo(errorInfo, rcv.MODE_OUTPUT_DISPLAY)
		return nil, 1
	}

	fmt.Println(tConfig)

	server = setServerValues(tConfig, version)

	// // Determine if connections are secure
	// if myServer.tls.TLSCert == rcv.EMPTY || myServer.tls.TLSKey == rcv.EMPTY || myServer.tls.TLSCABundle == rcv.EMPTY {
	// 	myServer.secured = false
	// } else {
	// 	myServer.secured = true
	// }
	// myServer.baseURL = coreHelpers.GenerateURL(myServer.environment, myServer.secured)
	// myServer.verifyEmailURL = coreHelpers.GenerateVerifyEmailURL(myServer.environment, myServer.secured)
	// //
	// Redirecting output to the log
	if test == false {
		if server.instanceDetails.logFileHandlerPtr, server.instanceDetails.logFQN, errorInfo = h.RedirectLogOutput(tConfig.LogDirectory); errorInfo.Error != nil {
			cpi.PrintErrorInfo(errorInfo, rcv.MODE_OUTPUT_DISPLAY)
			return nil, 1
		}
	}

	return
}

// setServerValues - sets the values for the Server struct
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func setServerValues(config cc.Configuration, version string) (server *Server) {
	server = &Server{
		config: cc.Configuration{
			ConfigFileName:         config.ConfigFileName,
			SkeletonConfigFilename: config.SkeletonConfigFilename,
			DebugModeOn:            config.DebugModeOn,
			Environment:            strings.ToLower(config.Environment),
			LogDirectory:           config.LogDirectory,
			MaxThreads:             config.MaxThreads,
			PIDDirectory:           config.PIDDirectory,
			Extensions:             nil,
		},
		instanceDetails: Instance{
			baseURL:           "",
			debugMode:         config.DebugModeOn,
			logFileHandlerPtr: &os.File{},
			mu:                sync.RWMutex{},
			numberCPUS:        runtime.NumCPU(),
			pid:               os.Getppid(),
			pidFQN:            config.PIDDirectory + rcv.PID_FILENAME,
			processChannel:    nil,
			running:           false,
			secured:           false,
			runStartTime:      time.Time{},
			version:           version,
			waitGroup:         sync.WaitGroup{},
		},
	}
	server.instanceDetails.hostname, _ = os.Hostname()
	server.instanceDetails.workingDirectory, _ = os.Getwd()
	server.instanceDetails.mu.Lock()
	defer server.instanceDetails.mu.Unlock()

	return
}

// Run starts the NATS server.
func (myServer *Server) Run() (returnCode int) {

	var (
	// errorInfo          cpi.ErrorInfo
	// tFunction, _, _, _ = runtime.Caller(0)
	// tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	log.Println("Starting SavUp server.")

	// if errorInfo = InitializeServer(myServer); errorInfo.Error != nil {
	// 	fmt.Println(rcv.BASH_COLOR_RED, errorInfo, rcv.BASH_COLOR_RESET)
	// 	os.Exit(1)
	// }
	//
	// // capture signals
	// go initializeSignals(myServer)
	//
	// coreError.PrintDebugLine(tFunctionName, "Blocking for NATS messages.")
	// myServer.processChannel = make(chan string)
	// go func() {
	// 	_ = myServer.messageHandler()
	// }()
	// select {
	// case <-myServer.processChannel:
	// }

	return
}

//
// func (myServer *Server) messageHandler() (err error) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tNatsMessage       natsMessage
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	// Use a WaitGroup to wait for a message to arrive
// 	myServer.waitGroup = sync.WaitGroup{}
// 	myServer.waitGroup.Add(1)
//
// 	for subject, message := range myServer.messages {
// 		if tNatsMessage.subscriptionPtr, err = myServer.MyNATS.NatsConnPtr.Subscribe(subject, message.handler); err != nil {
// 			log.Printf("Subscribe failed on subject: %v", subject)
// 			log.Fatalln(err.Error() + rcv.ENDING_EXECUTION)
// 		} else {
// 			tNatsMessage.handler = message.handler
// 			tNatsMessage.restrictedUsage = message.restrictedUsage
// 			myServer.messages[subject] = tNatsMessage
// 		}
// 	}
//
// 	// Waiting for a message to come in for processing.
// 	myServer.waitGroup.Wait()
//
// 	return
// }
//
// // debug sets the global debug variable to either true or false
// //
// //	Customer Messages: None
// //	Errors: None
// //	Verifications: None
// func (myServer *Server) debug() nats.MsgHandler {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	return func(msg *nats.Msg) {
// 		var (
// 			errorInfo          cpi.ErrorInfo
// 			tFunction, _, _, _ = runtime.Caller(0)
// 			tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 			tJSONReply         []byte
// 			tReply             DebugReply
// 			tRequest           DebugRequest
// 		)
//
// 		coreError.PrintDebugTrail(tFunctionName)
// 		if errorInfo = coreHelpers.UnmarshalMessageData(msg, &tRequest); errorInfo.Error == nil {
// 			coreError.SetDebugMode(tRequest.DebugOn)
// 			tReply.Data = fmt.Sprintf(`{"debug_on": %v}`, coreError.GetDebugMode())
// 		} else {
// 			tReply.Error = errorInfo.Error.Error()
// 		}
//
// 		tJSONReply = coreHelpers.BuildJSONReply(tReply, rcv.EMPTY, msg.Subject)
// 		_ = coreHelpers.SendReply(tFunctionName, tJSONReply, msg)
// 	}
// }
//
// // listMessages will call outputMessages and return the output as a reply to the request
// func (myServer *Server) listMessages() nats.MsgHandler {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	return func(msg *nats.Msg) {
// 		var (
// 			err error
// 		)
//
// 		tSupportedMessages := BuildListMessagesReply(myServer)
//
// 		if err = msg.Respond([]byte(tSupportedMessages)); err != nil {
// 			log.Println("Failed to create reply for the listMessages request")
// 			// ToDo Handle Error & Notification
// 		}
//
// 		return
// 	}
// }
//
// // shutdown unsubscribes the server to all subjects and removes the pid file.
// func (myServer *Server) Shutdown(test bool) {
//
// 	var (
// 		errorInfo          cpi.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	log.Println("A signal instructing shutdown has been received.")
// 	myServer.unsubscribeMessages()
// 	// Remove pid file
// 	if errorInfo = coreHelpers.RemovePidFile(myServer.pidFQN); errorInfo.Error == nil {
// 		log.Printf("The pid file (%v) has been removed", myServer.pidFQN)
// 		log.Println("The SavUp server has shutdown gracefully.")
// 	} else {
// 		log.Printf("WARNING: %v was not removed from the file system. This will need to be removed manually.", myServer.pidFQN)
// 	}
// 	if test == false {
// 		myServer.waitGroup.Done() // This must be the last statement in the Shutdown process.
// 	}
// }
//
// // displayServerInfo
// func displayServerInfo(myServer *Server) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	log.Println("SavUp Server Info:")
// 	// AWS Info
// 	log.Printf("\t%v", rcv.LINE_SEPARATOR)
// 	log.Printf("\tAWS Info")
// 	log.Printf("\t\tAWS Cognito Pool: \t%v", myServer.MyAWS.AWSConfig.UserPoolId)
// 	log.Printf("\t\tAWS Information FQN: %v", myServer.MyAWS.InfoFQN)
// 	log.Printf("\t\tAWS Profile Name: \t%v", myServer.MyAWS.AWSConfig.ProfileName)
// 	log.Printf("\t\tAWS Region: \t%v", myServer.MyAWS.AWSConfig.Region)
// 	//
// 	log.Printf("\tBase URL: \t%v", myServer.baseURL)
// 	log.Printf("\tConfig Filename: %v", myServer.opts.ConfigFileName)
// 	log.Printf("\tDebug Mode: \t%v", myServer.debugMode)
// 	log.Printf("\tEnvironment: \t%v", myServer.opts.Environment)
// 	// Firebase Info
// 	log.Printf("\t%v", rcv.LINE_SEPARATOR)
// 	log.Printf("\tFirebase Info")
// 	log.Printf("\t\tFirebase/Firestore Credentials FQN: %v", myServer.MyFireBase.CredentialsLocation)
// 	//
// 	log.Printf("\tHostname: \t%v", myServer.hostname)
// 	log.Printf("\tLog FQN: \t%v", myServer.logFQN)
// 	// NATS Info
// 	log.Printf("\t%v", rcv.LINE_SEPARATOR)
// 	log.Printf("\tNATS Info")
// 	log.Printf("\t\tNATS Message Prefix: %v", myServer.messagePrefix)
// 	log.Printf("\t\tNATS Credentials FQN: %v", myServer.MyNATS.NatsCredentialsFQN)
// 	log.Printf("\t\tNATS URL: %v", myServer.MyNATS.NatsURL)
// 	//
// 	// Plaid Info
// 	log.Printf("\t%v", rcv.LINE_SEPARATOR)
// 	log.Printf("\tPlaid Info")
// 	log.Printf("\t\tPlaid Credentials FQN: %v", myServer.MyPlaid.CredentialsLocation)
// 	//
// 	log.Printf("\tPID Directory: %v", myServer.pidDirectory)
// 	log.Printf("\tPID FQN: \t%v", myServer.pidFQN)
// 	log.Printf("\tStart Time: \t%v", myServer.startTime)
// 	// SendGrid Info
// 	log.Printf("\t%v", rcv.LINE_SEPARATOR)
// 	log.Printf("\tSendGrid Info")
// 	log.Printf("\t\tSendgrid FQN: %v", myServer.MySendGrid.SendGridCredentialsFQN)
// 	for key, value := range myServer.sendGridTemplateIds.Ids {
// 		log.Printf("\t\t\t%v Template Id: %v", key, value)
// 	}
// 	//
// 	log.Printf("\tVerify Email URL: %v", myServer.verifyEmailURL)
// 	// Stripe Info
// 	log.Printf("\t%v", rcv.LINE_SEPARATOR)
// 	log.Printf("\tStripe Info")
// 	log.Printf("\t\tStripe Credentials FQN: %v", myServer.MyFireBase.CredentialsLocation)
// 	//  TLS Info
// 	log.Printf("\t%v", rcv.LINE_SEPARATOR)
// 	if myServer.opts.TLS.TLSCert == rcv.EMPTY {
// 		log.Printf("\tTLS: \t%v", rcv.STATUS_INACTIVE)
// 	} else {
// 		log.Printf("\tTLS: \t%v", rcv.STATUS_ACTIVE)
// 		log.Printf("\t\tTLS Cert: \t%v", myServer.opts.TLS.TLSCert)
// 		log.Printf("\t\tTLS CA Bundle:  %v", myServer.opts.TLS.TLSCABundle)
// 		log.Printf("\t\tTLS Key: \t%v", myServer.opts.TLS.TLSKey)
// 	}
// 	//
// 	log.Printf("\tVersion: \t%v", myServer.version)
// 	log.Printf("\tWorking Directory: %v", myServer.workingDirectory)
// 	// End of Start Up Info
// 	log.Printf("%v", rcv.LINE_SEPARATOR)
// }
//
// func InitializeServer(myServer *Server) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	// Avoid RACE between Start() and Shutdown()
// 	myServer.mu.Lock()
// 	myServer.running = true
// 	myServer.mu.Unlock()
//
// 	// Check if a server.pid exists, if so shutdown
// 	if coreValidators.DoesFileExist(myServer.pidDirectory + rcv.PID_FILENAME) { // err of nil, means the file exists
// 		errorInfo.Error = coreError.ErrPIDFileExists
// 		errorInfo.AdditionalInfo = fmt.Sprintf("PID Directory: %v", myServer.pidDirectory+rcv.PID_FILENAME)
// 		coreError.PrintError(errorInfo)
// 	} else {
// 		errorInfo = coreHelpers.WritePidFile(myServer.pidDirectory)
// 	}
//
// 	// Setting up AWS
// 	if errorInfo.Error == nil {
// 		myServer.MyAWS, errorInfo = coreAWS.NewAWSSession(myServer.MyAWS.InfoFQN)
// 	}
//
// 	// Setting up Firebase & Firestore
// 	if errorInfo.Error == nil {
// 		if myServer.MyFireBase.AppPtr, myServer.MyFireBase.AuthPtr, errorInfo = coreFirebase.GetFirebaseAppAuthConnection(myServer.MyFireBase.CredentialsLocation); errorInfo.Error == nil {
// 			if myServer.MyFireBase.FirestoreClientPtr, errorInfo = coreFirestore.GetFirestoreClientConnection(myServer.MyFireBase.AppPtr); errorInfo.Error != nil {
// 				coreError.PrintError(errorInfo)
// 			}
// 		} else {
// 			coreError.PrintError(errorInfo)
// 		}
// 	}
//
// 	// Setting up NATS
// 	if errorInfo.Error == nil {
// 		myServer.MyNATS.NatsConnPtr, errorInfo = coreNATS.GetNATSConnection(myServer.MyNATS.NatsURL, myServer.MyNATS.NatsCredentialsFQN, myServer.tls)
// 	}
//
// 	// Setting up Plaid
// 	if errorInfo.Error == nil && myServer.MyPlaid.Keys.ClientId != rcv.EMPTY && myServer.MyPlaid.Keys.Secret != rcv.EMPTY {
// 		myServer.MyPlaid.PlaidClient, errorInfo = getPlaidClientConnection(myServer.MyPlaid.Keys)
// 	}
//
// 	// Setting up SendGrid email server
// 	if errorInfo.Error == nil {
// 		myServer.MySendGrid.EmailServerPtr, errorInfo = coreSendGrid.NewSendGridServer(coreSendGrid.DEFAULT_SENDER_ADDRESS, coreSendGrid.DEFAULT_SENDER_NAME, myServer.environment, myServer.MySendGrid.SendGridCredentialsFQN)
// 	}
//
// 	if errorInfo.Error == nil {
// 		if errorInfo = InitializeNATSMessages(myServer); errorInfo.Error != nil {
// 			coreError.PrintError(errorInfo)
// 		}
// 	}
//
// 	//
// 	// Outputting Server Info to the log
// 	displayServerInfo(myServer)
//
// 	return
// }
//
// // InitializeNATSMessages
// func InitializeNATSMessages(myServer *Server) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tNatMessage        natsMessage
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	if myServer == nil {
// 		errorInfo.Error = coreError.ErrPointerMissing
// 	} else {
// 		//
// 		// Command Messages
// 		tNatMessage.handler = myServer.debug()
// 		tNatMessage.restrictedUsage = true
// 		myServer.messages[myServer.messagePrefix+rcv.COMMAND_DEBUG] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.listMessages()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.COMMAND_LIST_MESSAGES] = tNatMessage
// 		//
// 		//
// 		// Functional Messages
// 		tNatMessage.handler = myServer.bundleAllocationAdjustment()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_BUNDLE_ALLOCATION_ADJUSTMENT] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.checkBundleExists()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_CHECK_BUNDLE_EXISTS] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.customerTransfer()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_CUSTOMER_TRANSFER] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.forgotUsername()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_FORGOT_USERNAME] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getAllBundles()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_ALL_BUNDLES] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getBundle()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_BUNDLE] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getInstitutionAccountBalances()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_INSTITUTION_BALANCES] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getInstitutionInfo()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_INSTITUTION_INFO] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getBackendInfo()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_BACKEND_INFO] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getToDoList()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_TODO_LIST] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getUserBundles()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_USER_BUNDLES] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getUserProfile()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_USER_PROFILE] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.getUserRegister()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_GET_USER_REGISTER] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.listInstitutions()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_LIST_INSTITUTIONS] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.listInstitutionAccountNames()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_LIST_INSTITUTION_ACCOUNT_NAMES] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.plaidGetLinkToken()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_PLAID_GET_LINK_TOKEN] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.userVerification()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_USER_VERIFICATION] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.pullUser()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_PULL_USER] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.pushLinkAndCreateCustomer()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_PUSH_LINK_AND_CREATE_CUSTOMER] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.resendUserVerifyEmail()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_RESEND_VERIFICATION_EMAIL] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.resetUserPassword()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_RESET_USER_PASSWORD] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.resetUserPassword()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_RESET_USER_PASSWORD] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.setFederalTaxId()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_SET_FEDERAL_TAX_ID] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.updateToDoList()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_UPDATE_TO_DO] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.updateUserFunds()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_UPDATE_USER_FUNDS] = tNatMessage
// 		//
// 		tNatMessage.handler = myServer.updateUserProfile()
// 		tNatMessage.restrictedUsage = false
// 		myServer.messages[myServer.messagePrefix+rcv.MSG_UPDATE_USER_PROFILE] = tNatMessage
// 	}
//
// 	return
// }
//
// // initialize signal handler
// func initializeSignals(myServer *Server) {
// 	var (
// 		captureSignal      = make(chan os.Signal, 1)
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	signal.Notify(captureSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
// 	myServer.signalHandler(<-captureSignal)
// }
