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
	"strings"
	"sync"
	"syscall"
	"time"

	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	ce "GriesPikeThomp/shared-services/src/coreExtensions"
	h "GriesPikeThomp/shared-services/src/coreHelpers"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	cv "GriesPikeThomp/shared-services/src/coreValidators"
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
	outputMode        string
	pid               int
	pidFQN            string
	processChannel    chan string
	running           bool
	runStartTime      time.Time
	serverName        string
	testingOn         bool
	version           string
	waitGroup         sync.WaitGroup
	workingDirectory  string
}

type Server struct {
	config        cc.Configuration
	instance      Instance
	extensionPtrs map[string]interface{}
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
		// _ = serverPtr.instance.messageHandler()
	}()
	select {
	case <-serverPtr.instance.processChannel:
	}

	return
}

// Shutdown - unsubscribes the server to all subjects and removes the pid file.
func (serverPtr *Server) Shutdown() {

	var (
		errorInfo cpi.ErrorInfo
	)

	// Remove pid file
	if serverPtr.instance.testingOn == false {
		if errorInfo = h.RemovePidFile(serverPtr.instance.pidFQN); errorInfo.Error != nil {
			cpi.PrintError(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_FILENAME, serverPtr.instance.pidFQN))
		}
	}

	log.Println(rcv.LINE_SHORT)
	log.Printf("The pid file (%v) has been removed", serverPtr.instance.pidFQN)
	log.Printf("The %v server has shutdown gracefully.", serverPtr.instance.serverName)

	serverPtr.instance.waitGroup.Done() // This must be the last statement in the Shutdown process.
}

// InitializeServer - create an instance and loads extensions.
//
//	Customer Messages: None
//	Errors: ErrPIDFileExists
//	Verifications: None
func InitializeServer(config cc.Configuration, serverName, version, logFQN string, logFileHandlerPtr *os.File, testingOn bool) (serverPtr *Server, errorInfo cpi.ErrorInfo) {

	log.Printf("Initializing instance of %v server.\n", serverName)

	serverPtr = NewServer(config, serverName, version, logFQN, logFileHandlerPtr, testingOn)

	// Check if a server.pid exists, if so shutdown
	if cv.DoesFileExist(serverPtr.instance.pidFQN) {
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
		serverPtr.extensionPtrs, errorInfo = ce.HandleExtension(config.Extensions)
	}

	return
}

// NewServer - creates an instance by setting the values for the Server struct
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewServer(config cc.Configuration, serverName, version, logFQN string, logFileHandlerPtr *os.File, testingOn bool) (server *Server) {
	server = &Server{
		config: cc.Configuration{
			ConfigFileName:         config.ConfigFileName,
			SkeletonConfigFilename: config.SkeletonConfigFilename,
			DebugModeOn:            config.DebugModeOn,
			Environment:            strings.ToLower(config.Environment),
			MaxThreads:             config.MaxThreads,
		},
		instance: Instance{
			debugMode:         config.DebugModeOn,
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
	server.instance.hostname, _ = os.Hostname()
	server.instance.workingDirectory, _ = os.Getwd()
	server.instance.pidFQN = h.PrependWorkingDirectoryWithEndingSlash(config.PIDDirectory) + rcv.PID_FILENAME

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
		tConfig            cc.Configuration
		tLogFileHandlerPtr *os.File
		tLogFQD            string
		tlogFQN            string
	)

	// See if configuration file exists and is readable, if not, return an error
	if tConfig, errorInfo = cc.ReadAndParseConfigFile(configFileFQN); errorInfo.Error != nil {
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

// initialize signal handler
func initializeSignals(serverPtr *Server) {
	var (
		captureSignal = make(chan os.Signal, 1)
	)

	signal.Notify(captureSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGABRT)
	serverPtr.signalHandler(<-captureSignal)
}
