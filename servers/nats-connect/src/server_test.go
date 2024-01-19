/*
NOTES:
*/
package src

import (
	"bytes"
	"log"
	"os"
	"runtime"
	"testing"

	"albert/constants"
	"albert/core/coreError"
	"github.com/nats-io/nats-server/v2/server"
)

var (
	tTestNameValue = map[string]interface{}{
		constants.TEST_FIELD_NAME: constants.TEST_FIELD_VALUE,
	}
)

func TestInitializeNATSMessages(tPtr *testing.T) {

	var (
		errorInfo          coreError.ErrorInfo
		tServer            *Server
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		if errorInfo = InitializeNATSMessages(tServer); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, coreError.ERROR, "nil")
		}
		tServer, _ = NewServer(constants.TEST_CONFIGURATION_FQN, constants.TEST_VERSION, true)
		if errorInfo = InitializeNATSMessages(tServer); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo)
		}
	})
}

func TestInitializeServer(tPtr *testing.T) {

}

func TestListMessages(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	StartTest(tFunctionName, true, false)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		myServer, _ := NewServer(constants.TEST_CONFIGURATION_FQN, server.VERSION, true)
		myServer.listMessages()
	})
}

func TestNewServer(tPtr *testing.T) {

	var (
		errorInfos         []coreError.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if _, errorInfos = NewServer(constants.TEST_CONFIGURATION_FQN, constants.TEST_VERSION, true); len(errorInfos) > 0 {
			tPtr.Errorf("%v Failed: Server was not created using the configuration file: %v.", tFunctionName, constants.TEST_CONFIGURATION_FQN)
		}
		if _, errorInfos = NewServer(constants.TEST_CONFIGURATION_WTIH_TLS_FQN, constants.TEST_VERSION, true); len(errorInfos) > 0 {
			tPtr.Errorf("%v Failed: Server was not created using the configuration file: %v.", tFunctionName, constants.TEST_CONFIGURATION_FQN)
		}
	})
}

func TestDisplayServerInfo(tPtr *testing.T) {

	var (
		myServer           *Server
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tBuffer            bytes.Buffer
	)

	log.SetOutput(&tBuffer)
	defer func() {
		log.SetOutput(os.Stderr)
	}()

	myServer = StartTest(tFunctionName, true, false)

	displayServerInfo(myServer)
	tPtr.Log(tBuffer.String())

	if tBuffer.Len() == 0 {
		tPtr.Errorf("%v Failed: Expected output in the buffer, instead got nothing.", tFunctionName)
	}

	StopTest(myServer)

}

func TestShutdown(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		myServer, _ := NewServer(constants.TEST_CONFIGURATION_FQN, server.VERSION, true)
		myServer.Shutdown(true)
	})
}
