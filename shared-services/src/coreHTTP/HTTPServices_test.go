// Package sharedServices
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, Inc
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
	"testing"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

func TestNewHTTP(tPtr *testing.T) {

	type arguments struct {
		hostname       string
		configFilename string
	}

	var (
		errorInfo cpi.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Secure connection.",
			arguments: arguments{
				hostname:       "localhost",
				configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/nats-connect/config/local/http-inbound-config.json",
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Bad URL.",
			arguments: arguments{
				hostname:       "localhost",
				configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/nats-connect/config/local/http-inbound-config.json",
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing credentials location.",
			arguments: arguments{
				hostname:       "localhost",
				configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/nats-connect/config/local/http-inbound-config.json",
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing certificate FQN.",
			arguments: arguments{
				hostname:       "localhost",
				configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/nats-connect/config/local/http-inbound-config.json",
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing private key FQN.",
			arguments: arguments{
				hostname:       "localhost",
				configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/nats-connect/config/local/http-inbound-config.json",
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Missing CA bundle FQN.",
			arguments: arguments{
				hostname:       "localhost",
				configFilename: "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/servers/nats-connect/config/local/http-inbound-config.json",
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _, errorInfo = NewHTTP(ts.arguments.hostname, ts.arguments.configFilename); errorInfo.Error != nil {
				gotError = true
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
