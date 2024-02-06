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
	"os/signal"
	"runtime"
	"sync"
	"syscall"
	"time"

	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
<<<<<<< HEAD
	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cn "GriesPikeThomp/shared-services/src/coreNATS"
=======
	h "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	ns "GriesPikeThomp/shared-services/src/coreNATS"
>>>>>>> parent of bc61635 (Working HTTP ListenAndServe)
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
	debugModeOn       bool
	hostname          string
	logFileHandlerPtr *os.File
	logFQN            string
	messageHandlers   map[string]nats.MsgHandler
	mu                sync.RWMutex
	numberCPUS        int
	outputMode        string
	pid               int
	pidFQN            string
	processChannel    chan string
	running           bool
	runStartTime      time.Time
	serverName        string
	subscriptionPtrs  map[string]*nats.Subscription
	testingOn         bool
	threadsAssigned   uint
	version           string
	waitGroup         sync.WaitGroup
	workingDirectory  string
	Namespace         string
}

type Server struct {
	config     cc.BaseConfiguration
	instance   Instance
	extensions map[string]interface{}
}

// Run - blocks for requests.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) Run() {

	// capture signals
	go initializeSignals(serverPtr)

	//
	serverPtr.instance.processChannel = make(chan string)
	go func() {
		serverPtr.instance.waitGroup = sync.WaitGroup{}
		serverPtr.instance.waitGroup.Add(1)
		serverPtr.messageHandler()
	}()
	select {
	case <-serverPtr.instance.processChannel:
	}

	return
}

// Shutdown - unsubscribes the server to all subjects and removes the pid file.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) Shutdown() {

	shutdown(serverPtr.instance.serverName, serverPtr.instance.pidFQN, serverPtr.instance.testingOn)

	serverPtr.instance.waitGroup.Done() // This must be the last statement in the Shutdown process.
}

// InitializeServer - create an instance and loads extensions.
//
//	Customer Messages: None
//	Errors: ErrPIDFileExists
//	Verifications: None
func InitializeServer(config cc.BaseConfiguration, serverName, version, logFQN string, logFileHandlerPtr *os.File, testingOn bool) (serverPtr *Server, errorInfo cpi.ErrorInfo) {

	log.Printf("Initializing instance of %v server.\n", serverName)

	if serverPtr, errorInfo = NewServer(config, serverName, version, logFQN, logFileHandlerPtr, testingOn); errorInfo.Error != nil {
		return
	}

	// Check if a server.pid exists, if so shutdown
	if h.DoesFileExist(serverPtr.instance.pidFQN) {
		errorInfo = cpi.NewErrorInfo(cpi.ErrPIDFileExists, fmt.Sprintf("PID Directory: %v", serverPtr.instance.pidFQN))
		return nil, errorInfo
	}

	if testingOn == false {
		if errorInfo = h.WritePidFile(serverPtr.instance.pidFQN, serverPtr.instance.pid); errorInfo.Error != nil {
			return nil, errorInfo
		}
	}

	// Avoid RACE between Start() and Shutdown(), running is set below.
	serverPtr.instance.mu.Lock()
	serverPtr.instance.running = true
	serverPtr.instance.mu.Unlock()

	log.Printf("Instance of %v server has been initialized.\n", serverName)

	if len(config.Extensions) == 0 {
		log.Println("No extensions defined in the configuration file.")
	} else {
		log.Println("Loading extensions.")
		errorInfo = serverPtr.HandleExtension(serverPtr.instance.hostname, config.Extensions)
	}

	return
}

// NewServer - creates an instance by setting the values for the Server struct
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewServer(config cc.BaseConfiguration, serverName, version, logFQN string, logFileHandlerPtr *os.File, testingOn bool) (serverPtr *Server, errorInfo cpi.ErrorInfo) {

	if config.MaxThreads > runtime.NumCPU() {
		errorInfo = cpi.NewErrorInfo(cpi.ErrMaxThreadsInvalid, fmt.Sprintf("%v%v", rcv.TXT_MAX_THREADS, "exceeds available threads."))
	}

	serverPtr = &Server{
		config: config,
		instance: Instance{
			debugModeOn:       config.DebugModeOn,
			logFileHandlerPtr: logFileHandlerPtr,
			logFQN:            logFQN,
			numberCPUS:        runtime.NumCPU(),
			outputMode:        rcv.MODE_OUTPUT_LOG_DISPLAY,
			pid:               os.Getppid(),
			serverName:        serverName,
			testingOn:         testingOn,
			version:           version,
		},
	}
	serverPtr.extensions = make(map[string]interface{})
	serverPtr.instance.hostname, _ = os.Hostname()
	serverPtr.instance.messageHandlers = make(map[string]nats.MsgHandler)
	serverPtr.instance.pidFQN = h.PrependWorkingDirectoryWithEndingSlash(config.PIDDirectory) + rcv.PID_FILENAME
	serverPtr.instance.subscriptionPtrs = make(map[string]*nats.Subscription)
	serverPtr.instance.workingDirectory, _ = os.Getwd()

	return
}

// RunServer will set up a new server instance after parsing the configuration defined
// in the supplied configuration file. If no configuration file is provide, an
// error will be returned. If the configuration is invalid, an error will be returned.
// After the server is created, RunServer will block waiting for messages.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RunServer(configFileFQN, serverName, version string, testingOn bool) (returnCode int) {

	var (
		errorInfo          cpi.ErrorInfo
		serverPtr          *Server
		tConfig            cc.BaseConfiguration
		tLogFileHandlerPtr *os.File
		tLogFQD            string
		tlogFQN            string
	)

	// See if configuration file exists and is readable, if not, return an error
	if tConfig, errorInfo = cc.ProcessBaseConfigFile(configFileFQN); errorInfo.Error != nil {
		cpi.PrintErrorInfo(errorInfo)
		return 1
	}

	if errorInfo = cc.ValidateConfiguration(tConfig); errorInfo.Error != nil {
		cpi.PrintErrorInfo(errorInfo)
		return 1
	}

	// Initializing the log output.
	tLogFQD = h.PrependWorkingDirectoryWithEndingSlash(tConfig.LogDirectory)
	if tLogFileHandlerPtr, tlogFQN, errorInfo = h.CreateAndRedirectLogOutput(tLogFQD, rcv.MODE_OUTPUT_LOG_DISPLAY); errorInfo.Error != nil {
		cpi.PrintErrorInfo(errorInfo)
		return
	}

	log.Printf("Creating a new instance of %v server.\n", serverName)

	if serverPtr, errorInfo = InitializeServer(tConfig, serverName, version, tlogFQN, tLogFileHandlerPtr, testingOn); errorInfo.Error != nil {
		cpi.PrintErrorInfo(errorInfo)
		return 1
	}

	serverPtr.Run()

	return
}

// Private Functions

// messageHandler - subscribes subjects to handlers.
//
//	Customer Messages: None
//	Errors: ErrSignalUnknown
//	Verifications: None
func (serverPtr *Server) messageHandler() {

	// Use a WaitGroup to wait for a message to arrive
	serverPtr.instance.waitGroup = sync.WaitGroup{}
	serverPtr.instance.waitGroup.Add(1)

	for serviceName, serviceInfo := range serverPtr.extensions {
		switch serviceName {
		case NATS_INTERNAL:
<<<<<<< HEAD
			retrievedService := serviceInfo.(cn.NATSService)
			serverPtr.getNATSHandlers(retrievedService)
=======
			retrievedService := serviceInfo.(ns.NATSService)
			serverPtr.getHandlers(retrievedService)
>>>>>>> parent of bc61635 (Working HTTP ListenAndServe)
		}
	}

	// Waiting for a message to come in for processing.
	serverPtr.instance.waitGroup.Wait()

	return
}

// signalHandler - collects keyboard input and manages the server response.
//
//	Customer Messages: None
//	Errors: ErrSignalUnknown
//	Verifications: None
func (serverPtr *Server) signalHandler(signal os.Signal) {

	switch signal {
	case syscall.SIGHUP: // kill -SIGHUP XXXX
		fallthrough
	case syscall.SIGINT: // kill -SIGINT XXXX or Ctrl+c
		fallthrough
	case syscall.SIGTERM: // kill -SIGTERM XXXX
		fallthrough
	case syscall.SIGQUIT: // kill -SIGQUIT XXXX
		serverPtr.Shutdown()
	default:
		cpi.PrintError(cpi.ErrSignalUnknown, fmt.Sprintf("%v%v", rcv.TXT_SIGNAL, signal.String()))
	}

	os.Exit(0)
}

// initialize signal handler - handles signals from the console.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func initializeSignals(serverPtr *Server) {
	var (
		captureSignal = make(chan os.Signal, 1)
	)

	signal.Notify(captureSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	serverPtr.signalHandler(<-captureSignal)
}

// shutdown
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func shutdown(serverName, pidFQN string, testingOn bool) {
	var (
		errorInfo cpi.ErrorInfo
	)

	// Remove pid file
	if testingOn == false {
		if errorInfo = h.RemovePidFile(pidFQN); errorInfo.Error != nil {
			cpi.PrintError(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_FILENAME, pidFQN))
		}
	}

	log.Println(rcv.LINE_SHORT)
	log.Printf("The pid file (%v) has been removed", pidFQN)
	log.Printf("The %v server has shutdown gracefully.", serverName)

}
