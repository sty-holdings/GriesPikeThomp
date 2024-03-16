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

	"github.com/nats-io/nats.go"
	cs "github.com/sty-holdings/GriesPikeThomp/servers/nats-connect/integrations/coreStripe"
	ext "github.com/sty-holdings/GriesPikeThomp/servers/nats-connect/loadExtensions"
	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	config "github.com/sty-holdings/sty-shared/v2024/configuration"
	hv "github.com/sty-holdings/sty-shared/v2024/helpersValidators"
	ns "github.com/sty-holdings/sty-shared/v2024/natsSerices"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

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
		errorInfo          pi.ErrorInfo
		serverPtr          *Server
		tConfig            BaseConfiguration
		tConfigExtensions  []ext.BaseConfigExtensions
		tConfigData        = make(map[string]interface{})
		tLogFileHandlerPtr *os.File
		tLogFQD            string
		tlogFQN            string
	)

	// See if configuration file exists and is readable, if not, return an error
	if tConfigData, errorInfo = config.GetConfigFile(configFileFQN); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return 1
	}

	if value, ok := tConfigData[ctv.FN_DEBUG_MODE_ON]; ok {
		tConfig.DebugModeOn = value.(bool)
	}
	if value, ok := tConfigData[ctv.FN_ENVIRONMENT]; ok {
		tConfig.Environment = value.(string)
	}
	if value, ok := tConfigData[ctv.FN_LOAD_EXTENSIONS]; ok {
		for _, i2 := range value.([]interface{}) {
			x := make(map[string]interface{})
			x = i2.(map[string]interface{})
			y := ext.BaseConfigExtensions{
				Name:           x["name"].(string),
				ConfigFilename: x["config_filename"].(string),
			}
			tConfigExtensions = append(tConfigExtensions, y)
		}
		tConfig.Extensions = tConfigExtensions
	}
	if value, ok := tConfigData[ctv.FN_LOG_DIRECTORY]; ok {
		tConfig.LogDirectory = value.(string)
	}
	if value, ok := tConfigData[ctv.FN_MAX_THREADS]; ok {
		tConfig.MaxThreads = int(value.(float64))
	}
	if value, ok := tConfigData[ctv.FN_PID_DIRECTORY]; ok {
		tConfig.PIDDirectory = value.(string)
	}
	if value, ok := tConfigData[ctv.FN_SKELETON_DIRECTORY]; ok {
		tConfig.SkeletonConfigDirectory = value.(string)
	}

	if errorInfo = validateConfiguration(tConfig); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return 1
	}

	// Initializing the log output.
	tLogFQD = hv.PrependWorkingDirectoryWithEndingSlash(tConfig.LogDirectory)
	if tLogFileHandlerPtr, tlogFQN, errorInfo = hv.CreateAndRedirectLogOutput(tLogFQD, ctv.MODE_OUTPUT_LOG_DISPLAY); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
		return
	}

	log.Printf("Creating a new instance of %v server.\n", serverName)

	if serverPtr, errorInfo = initializeServer(tConfig, serverName, version, tlogFQN, tLogFileHandlerPtr, testingOn); errorInfo.Error != nil {
		pi.PrintErrorInfo(errorInfo)
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
		for _, tExtInstance := range serverPtr.instance.extInstances {
			tExtInstance.WaitGroup.Done() // This must be the last statement in the Shutdown process.
		}
	}
}

// Private Functions

// extensionHandler - starts extensions in goroutine.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) buildNCIExtension() (
	errorInfo pi.ErrorInfo,
) {

	var (
		tExtInstance      ExtInstance
		tSubscriptionPtrs = make(map[string]*nats.Subscription)
	)

	if tExtInstance.InstanceName, errorInfo = ns.BuildInstanceName(ns.METHOD_BLANK, ctv.NC_INTERNAL); errorInfo.Error != nil {
		return
	}
	if tExtInstance.NatsConnectionPtr, errorInfo = ns.GetConnection(
		tExtInstance.InstanceName,
		serverPtr.extensionConfigs[ctv.NC_INTERNAL].NATSConfig,
	); errorInfo.Error != nil {
		return
	}

	// Use a WaitGroup to wait for a message to arrive
	tExtInstance.WaitGroup = sync.WaitGroup{}
	tExtInstance.WaitGroup.Add(1)
	for subject, handler := range serverPtr.loadNCIMessageHandles() {
		tSubscriptionPtrs[subject], errorInfo = ns.Subscribe(tExtInstance.NatsConnectionPtr, tExtInstance.InstanceName, subject, handler.Handler)
	}

	tExtInstance.SubscriptionPtrs = tSubscriptionPtrs
	serverPtr.instance.extInstances[ctv.NC_INTERNAL] = tExtInstance

	if serverPtr.instance.testingOn {
		tExtInstance.WaitGroup.Done()
	} else {
		tExtInstance.WaitGroup.Wait()
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
	// 		retrievedService := serviceInfo.(ns.NATSService)
	// serverPtr.getNATSHandlers(retrievedService)
	// 	case STRIPE:
	// 		retrievedService := serviceInfo.(ns.NATSService)
	// 		serverPtr.getNATSHandlers(retrievedService)
	// 	}
	// }
	//
	// // Waiting for a message to come in for processing.
	// serverPtr.instance.waitGroup.Wait()

	return
}

// Run - blocks for requests.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (serverPtr *Server) run() {

	var (
		key string
	)

	key = strings.ToLower(strings.Trim(key, ctv.VAL_EMPTY))

	// capture signals
	go initializeSignals(serverPtr)

	// start extensions
	for key = range serverPtr.extensionConfigs {
		if key != ctv.NC_INTERNAL { // Skipping NC_INSTANCE so server is not block extension creation
			switch key {
			case ctv.STRIPE_EXTENSION:
				go cs.NewExtension(
					serverPtr.instance.hostname,
					serverPtr.extensionConfigs[key],
					serverPtr.instance.testingOn,
				)
			}
		}
	}

	// start nats connect internal which will block the server
	serverPtr.instance.nciProcessChannel = make(chan string) // This is for NC_INTERNAL only
	go func() {
		serverPtr.buildNCIExtension()
	}()
	select {
	case <-serverPtr.instance.nciProcessChannel:
	}

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
		pi.PrintError(pi.ErrSignalUnknown, fmt.Sprintf("%v%v", ctv.TXT_SIGNAL, signal.String()))
	}

	os.Exit(0)
}

// initializeServer - create an instance and loads loadExtensions.
//
//	Customer Messages: None
//	Errors: ErrPIDFileExists
//	Verifications: None
func initializeServer(
	config BaseConfiguration,
	serverName, version, logFQN string,
	logFileHandlerPtr *os.File,
	testingOn bool,
) (
	serverPtr *Server,
	errorInfo pi.ErrorInfo,
) {

	log.Printf("Initializing instance of %v server.\n", serverName)

	if serverPtr, errorInfo = newServer(config, serverName, version, logFQN, logFileHandlerPtr, testingOn); errorInfo.Error != nil {
		return
	}

	// Check if a server.pid exists, if so shutdown
	if hv.DoesFileExist(serverPtr.instance.pidFQN) {
		errorInfo = pi.NewErrorInfo(pi.ErrPIDFileExists, fmt.Sprintf("PID Directory: %v", serverPtr.instance.pidFQN))
		return nil, errorInfo
	}

	if testingOn == false {
		if errorInfo = hv.WritePidFile(serverPtr.instance.pidFQN, serverPtr.instance.pid); errorInfo.Error != nil {
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
			if serverPtr.extensionConfigs[strings.ToLower(values.Name)], errorInfo = ext.LoadExtensionConfig(values); errorInfo.Error != nil {
				return
			}
			log.Printf("%v configuration is loaded.", values.Name)
		}
	}

	return
}

// newServer - creates an instance by setting the values for the Server struct
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func newServer(
	config BaseConfiguration,
	serverName, version, logFQN string,
	logFileHandlerPtr *os.File,
	testingOn bool,
) (
	serverPtr *Server,
	errorInfo pi.ErrorInfo,
) {

	if config.MaxThreads > runtime.NumCPU() {
		errorInfo = pi.NewErrorInfo(pi.ErrMaxThreadsInvalid, fmt.Sprintf("%v%v", ctv.TXT_MAX_THREADS, "exceeds available threads."))
	}

	serverPtr = &Server{
		config: config,
		instance: Instance{
			debugModeOn:       config.DebugModeOn,
			logFileHandlerPtr: logFileHandlerPtr,
			logFQN:            logFQN,
			numberCPUS:        runtime.NumCPU(),
			outputMode:        ctv.MODE_OUTPUT_LOG_DISPLAY,
			pid:               os.Getppid(),
			serverName:        serverName,
			testingOn:         testingOn,
			version:           version,
		},
	}
	serverPtr.extensionConfigs = make(map[string]ext.ExtensionConfiguration)
	serverPtr.instance.extInstances = make(map[string]ExtInstance)
	serverPtr.instance.hostname, _ = os.Hostname()
	serverPtr.instance.messageHandlers = make(map[string]nats.MsgHandler)
	serverPtr.instance.pidFQN = hv.PrependWorkingDirectoryWithEndingSlash(config.PIDDirectory) + ctv.PID_FILENAME
	serverPtr.instance.workingDirectory, _ = os.Getwd()

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
		errorInfo pi.ErrorInfo
	)

	// Remove pid file
	if testingOn == false {
		if errorInfo = hv.RemovePidFile(pidFQN); errorInfo.Error != nil {
			pi.PrintError(errorInfo.Error, fmt.Sprintf("%v%v", ctv.TXT_FILENAME, pidFQN))
		}
	}

	log.Println(ctv.LINE_SHORT)
	log.Printf("The pid file (%v) has been removed", pidFQN)
	log.Printf("The %v server has shutdown gracefully.", serverName)

}

// validateConfiguration - checks the values in the configuration file are valid. ValidateConfiguration doesn't
// test if the configuration file exists, readable, or parsable.
//
//	Customer Messages: None
//	Errors: ErrEnvironmentInvalid, ErrDirectoryMissing, ErrMaxThreadsInvalid
//	Verifications: None
func validateConfiguration(config BaseConfiguration) (
	errorInfo pi.ErrorInfo,
) {

	if hv.IsEnvironmentValid(config.Environment) == false {
		errorInfo = pi.NewErrorInfo(pi.ErrEnvironmentInvalid, fmt.Sprintf("%v%v", ctv.TXT_EVIRONMENT, ctv.FN_ENVIRONMENT))
		return
	}
	if hv.DoesDirectoryExist(config.LogDirectory) == false {
		errorInfo = pi.NewErrorInfo(pi.ErrDirectoryMissing, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, ctv.FN_LOG_DIRECTORY))
		return
	}
	if config.MaxThreads < 1 {
		errorInfo = pi.NewErrorInfo(pi.ErrMaxThreadsInvalid, fmt.Sprintf("%v%v", ctv.TXT_MAX_THREADS, ctv.FN_MAX_THREADS))
		return
	}
	if hv.DoesDirectoryExist(config.PIDDirectory) == false {
		errorInfo = pi.NewErrorInfo(pi.ErrDirectoryMissing, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, ctv.FN_PID_DIRECTORY))
		return
	}
	if hv.DoesDirectoryExist(config.SkeletonConfigDirectory) == false {
		errorInfo = pi.NewErrorInfo(pi.ErrDirectoryMissing, fmt.Sprintf("%v%v", ctv.TXT_DIRECTORY, ctv.FN_SKELETON_DIRECTORY))
		return
	}

	return
}
