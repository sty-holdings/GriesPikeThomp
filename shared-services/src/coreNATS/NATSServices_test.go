// Package sharedServices
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc.
	Use is subject to license terms.

	Unauthorized copying of this file, via any medium is strictly prohibited.

	Proprietary and confidential

	Written by <Replace with FULL_NAME> / syacko
	STY-Holdings, Inc.
	support@sty-holdings.com
	www.sty-holdings.com

	01-2024
	USA

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package sharedServices

import (
	"errors"
	"fmt"
	"runtime"
	"testing"

	ext "GriesPikeThomp/servers/nats-connect/loadExtensions"
	cj "GriesPikeThomp/shared-services/src/coreJWT"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

func TestGetNATSConnection(tPtr *testing.T) {

	type arguments struct {
		instanceName string
		config       ext.ExtensionConfiguration
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
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: ext.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				instanceName: rcv.VAL_EMPTY,
				config: ext.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing Credential filename.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: ext.ExtensionConfiguration{
					NATSCredentialsFilename: rcv.VAL_EMPTY,
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Port is zero.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: ext.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                0,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing certificate FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: ext.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       rcv.VAL_EMPTY,
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing private key FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: ext.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: rcv.VAL_EMPTY,
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing CA bundle FQN.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: ext.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   rcv.VAL_EMPTY,
					},
					NATSURL: "savup-local-0030.savup.com",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing URL.",
			arguments: arguments{
				instanceName: "scott-test-connection",
				config: ext.ExtensionConfiguration{
					NATSCredentialsFilename: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key",
					NATSPort:                4222,
					NATSTLSInfo: cj.TLSInfo{
						TLSCert:       "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt",
						TLSPrivateKey: "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key",
						TLSCABundle:   "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt",
					},
					NATSURL: rcv.VAL_EMPTY,
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if _, errorInfo = GetConnection(ts.arguments.instanceName, ts.arguments.config); errorInfo.Error != nil {
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
			},
		)
	}
}
