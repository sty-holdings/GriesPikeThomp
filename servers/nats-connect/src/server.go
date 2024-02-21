// Package src
/*
This is code for STY-Holdings NATS Connect

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

	ext "GriesPikeThomp/servers/nats-connect/loadExtensions"
	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cns "GriesPikeThomp/shared-services/src/coreNATS"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/nats-io/nats.go"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type Auth struct {
	authenticatorService string
}

// Shorten the name to resolve naming conflict
type extInstance struct {
	instanceName      string
	natsConnectionPtr *nats.Conn
	processChannel    chan string
	subscriptionPtrs  map[string]*nats.Subscription
	waitGroup         sync.WaitGroup
}

// Instance - Some of these values can change over the life of the instance.
type Instance struct {
	baseURL            string
	debugModeOn        bool
	extensionInstances map[string]extInstance
	hostname           string
	logFileHandlerPtr  *os.File
	logFQN             string
	messageHandlers    map[string]nats.MsgHandler
	mu                 sync.RWMutex
	numberCPUS         int
	outputMode         string
	pid                int
	pidFQN             string
	processChannel     chan string
	running            bool
	runStartTime       time.Time
	serverName         string
	testingOn          bool
	threadsAssigned    uint
	version            string
	waitGroupCreated   bool
	workingDirectory   string
}

type Server struct {
	config           cc.BaseConfiguration
	instance         Instance
	extensionConfigs map[string]ext.ExtensionConfiguration
}

// RunServer will set up a new server instance after parsing the configuration defined
// in the supplied configuration file. If no configuration file is provide, an
// error will be returned. If the configuration is invalid, an error will be returned.
// After the server is created, RunServer will block waiting for messages.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RunServer(
	configFileFQN, serverName, version string,
	testingOn bool,
) (returnCode int) {

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
	tLogFQD = chv.PrependWorkingDirectoryWithEndingSlash(tConfig.LogDirectory)
	if tLogFileHandlerPtr, tlogFQN, errorInfo = chv.CreateAndRedirectLogOutput(tLogFQD, rcv.MODE_OUTPUT_LOG_DISPLAY); errorInfo.Error != nil {
		cpi.PrintErrorInfo(errorInfo)
		return
	}

	log.Printf("Creating a new instance of %v server.\n", serverName)

	if serverPtr, errorInfo = initializeServer(tConfig, serverName, version, tlogFQN, tLogFileHandlerPtr, testingOn); errorInfo.Error != nil {
		cpi.PrintErrorInfo(errorInfo)
		return 1
	}

	serverPtr.run()

	return
}

// Shutdown - unsubscribes the server to all subjects and removes the pid file.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) Shutdown() {

	shutdown(serverPtr.instance.serverName, serverPtr.instance.pidFQN, serverPtr.instance.testingOn)

	if serverPtr.instance.waitGroupCreated {
		for _, tExtensionInstance := range serverPtr.instance.extensionInstances {
			tExtensionInstance.waitGroup.Done() // This must be the last statement in the Shutdown process.
		}
	}
}

// Private Functions

// extensionHandler - starts extensions in goroutine.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) extensionHandler(extensionKey string) (errorInfo cpi.ErrorInfo) {

	log.Println("key: ", extensionKey)

	if extensionKey == NC_INTERNAL {
		errorInfo = serverPtr.buildNCIExtension()
	}

	return
}

// initializeServer - create an instance and loads loadExtensions.
//
//	Customer Messages: None
//	Errors: ErrPIDFileExists
//	Verifications: None
func initializeServer(
	config cc.BaseConfiguration,
	serverName, version, logFQN string,
	logFileHandlerPtr *os.File,
	testingOn bool,
) (
	serverPtr *Server,
	errorInfo cpi.ErrorInfo,
) {

	log.Printf("Initializing instance of %v server.\n", serverName)

	if serverPtr, errorInfo = newServer(config, serverName, version, logFQN, logFileHandlerPtr, testingOn); errorInfo.Error != nil {
		return
	}

	// Check if a server.pid exists, if so shutdown
	if chv.DoesFileExist(serverPtr.instance.pidFQN) {
		errorInfo = cpi.NewErrorInfo(cpi.ErrPIDFileExists, fmt.Sprintf("PID Directory: %v", serverPtr.instance.pidFQN))
		return nil, errorInfo
	}

	if testingOn == false {
		if errorInfo = chv.WritePidFile(serverPtr.instance.pidFQN, serverPtr.instance.pid); errorInfo.Error != nil {
			return nil, errorInfo
		}
	}

	// Avoid RACE between Start() and Shutdown(), running is set below.
	serverPtr.instance.mu.Lock()
	serverPtr.instance.running = true
	serverPtr.instance.mu.Unlock()

	log.Printf("Instance of %v server has been initialized.\n", serverName)

	if len(config.Extensions) == 0 {
		log.Println("No loadExtensions defined in the configuration file.")
	} else {
		log.Println("Loading extension configs.")
		for _, values := range config.Extensions {
			if serverPtr.extensionConfigs[strings.ToLower(values.Name)], errorInfo = ext.LoadExtensionConfig(values.ConfigFilename); errorInfo.
				Error != nil {
				return
			}
			log.Printf("%v configuration is loaded.", values.Name)
		}
	}

	return
}

// messageHandler - subscribes subjects to handlers.
//
//	Customer Messages: None
//	Errors: ErrSignalUnknown
//	Verifications: None
func (serverPtr *Server) messageHandler() {

	// Use a WaitGroup to wait for a message to arrive
	// serverPtr.instance.waitGroup = sync.WaitGroup{}
	// serverPtr.instance.waitGroup.Add(1)
	// serverPtr.instance.waitGroupCreated = true
	//
	// for serviceName, serviceInfo := range serverPtr.loadExtensions {
	// 	switch serviceName {
	// 	case NC_INTERNAL:
	// 		retrievedService := serviceInfo.(cn.NATSService)
	// serverPtr.getNATSHandlers(retrievedService)
	// 	case STRIPE:
	// 		retrievedService := serviceInfo.(cn.NATSService)
	// 		serverPtr.getNATSHandlers(retrievedService)
	// 	}
	// }
	//
	// // Waiting for a message to come in for processing.
	// serverPtr.instance.waitGroup.Wait()

	return
}

// newServer - creates an instance by setting the values for the Server struct
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func newServer(
	config cc.BaseConfiguration,
	serverName, version, logFQN string,
	logFileHandlerPtr *os.File,
	testingOn bool,
) (
	serverPtr *Server,
	errorInfo cpi.ErrorInfo,
) {

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
	serverPtr.extensionConfigs = make(map[string]ext.ExtensionConfiguration)
	serverPtr.instance.extensionInstances = make(map[string]extInstance)
	serverPtr.instance.hostname, _ = os.Hostname()
	serverPtr.instance.messageHandlers = make(map[string]nats.MsgHandler)
	serverPtr.instance.pidFQN = chv.PrependWorkingDirectoryWithEndingSlash(config.PIDDirectory) + rcv.PID_FILENAME
	serverPtr.instance.workingDirectory, _ = os.Getwd()

	return
}

// Run - blocks for requests.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) run() {

	// capture signals
	go initializeSignals(serverPtr)

	// start extensions
	// for key, _ := range serverPtr.extensionConfigs {
	// 	log.Printf("Key: %v", key)
	// 	go serverPtr.extensionHandler(key)
	serverPtr.instance.processChannel = make(chan string)
	go func() {
		serverPtr.extensionHandler(NC_INTERNAL)
	}()
	select {
	case <-serverPtr.instance.processChannel:
	}
	// }

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

// extensionHandler - starts extensions in goroutine.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) buildNCIExtension() (
	errorInfo cpi.ErrorInfo,
) {

	var (
		tExtensionInstance extInstance
		tSubscriptionPtr   *nats.Subscription
		tSubscriptionPtrs  = make(map[string]*nats.Subscription)
	)

	if tExtensionInstance.instanceName, errorInfo = cns.BuildInstanceName(cns.METHOD_BLANK, NC_INTERNAL); errorInfo.Error != nil {
		return
	}
	if tExtensionInstance.natsConnectionPtr, errorInfo = cns.GetConnection(tExtensionInstance.instanceName,
		serverPtr.extensionConfigs[NC_INTERNAL]); errorInfo.Error != nil {
		return
	}

	// Use a WaitGroup to wait for a message to arrive
	tExtensionInstance.waitGroup = sync.WaitGroup{}
	tExtensionInstance.waitGroup.Add(1)
	for subject, handler := range serverPtr.nciMessageHandles() {
		if tSubscriptionPtr, errorInfo.Error = tExtensionInstance.natsConnectionPtr.Subscribe(subject,
			handler.Handler); errorInfo.Error != nil {
			log.Printf("Subscribe failed on subject: %v", subject)
			return
		}
		tSubscriptionPtrs[subject] = tSubscriptionPtr
	}

	tExtensionInstance.subscriptionPtrs = tSubscriptionPtrs
	serverPtr.instance.extensionInstances[NC_INTERNAL] = tExtensionInstance

	if serverPtr.instance.testingOn {
		tExtensionInstance.waitGroup.Done()
	} else {
		tExtensionInstance.waitGroup.Wait()
	}

	return
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
func shutdown(
	serverName, pidFQN string,
	testingOn bool,
) {
	var (
		errorInfo cpi.ErrorInfo
	)

	// Remove pid file
	if testingOn == false {
		if errorInfo = chv.RemovePidFile(pidFQN); errorInfo.Error != nil {
			cpi.PrintError(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_FILENAME, pidFQN))
		}
	}

	log.Println(rcv.LINE_SHORT)
	log.Printf("The pid file (%v) has been removed", pidFQN)
	log.Printf("The %v server has shutdown gracefully.", serverName)

}
