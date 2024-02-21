/*
NOTES:
*/
package src

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	le "GriesPikeThomp/servers/nats-connect/loadExtensions"
	cj "GriesPikeThomp/shared-services/src/coreJWT"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

func TestBuildExtension(tPtr *testing.T) {

	type arguments struct {
		extensionKey string
		config       le.ExtensionConfiguration
	}

	var (
		errorInfo          cpi.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		// Missing extension key is not tested. No way for the program to get to this code without a key.
		// Empty subject registry is not tested. No way for the program to get to this code without a populated Subject Registry.
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				extensionKey: NC_INTERNAL,
				config: le.ExtensionConfiguration{
					MessageEnvironment:      rcv.ENVIRONMENT_LOCAL,
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL:         "savup-local-0030.savup.com",
					SubjectRegistry: buildSubjectRegistry(),
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing Credential filename.",
			arguments: arguments{
				extensionKey: NC_INTERNAL,
				config: le.ExtensionConfiguration{
					NATSCredentialsFilename: rcv.VAL_EMPTY,
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL:         "savup-local-0030.savup.com",
					SubjectRegistry: buildSubjectRegistry(),
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Port is zero.",
			arguments: arguments{
				extensionKey: NC_INTERNAL,
				config: le.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                0,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL:         "savup-local-0030.savup.com",
					SubjectRegistry: buildSubjectRegistry(),
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing certificate FQN.",
			arguments: arguments{
				extensionKey: NC_INTERNAL,
				config: le.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       rcv.VAL_EMPTY,
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL:         "savup-local-0030.savup.com",
					SubjectRegistry: buildSubjectRegistry(),
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing private key FQN.",
			arguments: arguments{
				extensionKey: NC_INTERNAL,
				config: le.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: rcv.VAL_EMPTY,
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL:         "savup-local-0030.savup.com",
					SubjectRegistry: buildSubjectRegistry(),
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing CA bundle FQN.",
			arguments: arguments{
				extensionKey: NC_INTERNAL,
				config: le.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   rcv.VAL_EMPTY,
					},
					NATSURL:         "savup-local-0030.savup.com",
					SubjectRegistry: buildSubjectRegistry(),
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing URL.",
			arguments: arguments{
				extensionKey: NC_INTERNAL,
				config: le.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL:         rcv.VAL_EMPTY,
					SubjectRegistry: buildSubjectRegistry(),
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _, errorInfo = buildExtension(ts.arguments.extensionKey, ts.arguments.config, true); errorInfo.Error != nil {
				gotError = true
				errorInfo = cpi.ErrorInfo{
					Error: errors.New(fmt.Sprintf("Failed - NATS connection was not created for Test: %v", tFunctionName)),
				}
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}
}

// func TestInitializeServer(tPtr *testing.T) {
//
// }

// func TestNewServer(tPtr *testing.T) {
//
// 	var (
// 		errorInfos         []cpi.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if _, errorInfos = NewServer(rcv.TEST_CONFIGURATION_FQN, rcv.TEST_VERSION, true); len(errorInfos) > 0 {
// 			tPtr.Errorf("%v Failed: Server was not created using the configuration file: %v.", tFunctionName, rcv.TEST_CONFIGURATION_FQN)
// 		}
// 		if _, errorInfos = NewServer(rcv.TEST_CONFIGURATION_WTIH_TLS_FQN, rcv.TEST_VERSION, true); len(errorInfos) > 0 {
// 			tPtr.Errorf("%v Failed: Server was not created using the configuration file: %v.", tFunctionName, rcv.TEST_CONFIGURATION_FQN)
// 		}
// 	})
// }

// func TestDisplayServerInfo(tPtr *testing.T) {
//
// 	var (
// 		myServer           *Server
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tBuffer            bytes.Buffer
// 	)
//
// 	log.SetOutput(&tBuffer)
// 	defer func() {
// 		log.SetOutput(os.Stderr)
// 	}()
//
// 	myServer = StartTest(tFunctionName, true, false)
//
// 	displayServerInfo(myServer)
// 	tPtr.Log(tBuffer.String())
//
// 	if tBuffer.Len() == 0 {
// 		tPtr.Errorf("%v Failed: Expected output in the buffer, instead got nothing.", tFunctionName)
// 	}
//
// 	StopTest(myServer)
//
// }

// func TestShutdown(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		myServer, _ := NewServer(rcv.TEST_CONFIGURATION_FQN, server.VERSION, true)
// 		myServer.Shutdown(true)
// 	})
// }

func buildSubjectRegistry() (subjectRegistry []le.SubjectInfo) {
	var (
		tSubjectRegistry le.SubjectInfo
	)

	tSubjectRegistry.Namespace = "nci"
	tSubjectRegistry.Subject = "turn_debug_on"
	subjectRegistry = append(subjectRegistry, tSubjectRegistry)

	tSubjectRegistry.Namespace = "nci"
	tSubjectRegistry.Subject = "turn_debug_off"
	subjectRegistry = append(subjectRegistry, tSubjectRegistry)

	return
}
