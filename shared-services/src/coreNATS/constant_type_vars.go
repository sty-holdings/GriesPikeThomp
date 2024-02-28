// Package sharedServices
/*
General description of the purpose of the go file.

RESTRICTIONS:
    AWS functions:
    * Program must have access to a .aws/credentials file in the default location.
    * This will only access system parameters that start with '/sote' (ROOTPATH).
    * {Enter other restrictions here for AWS

    {Other catagories of restrictions}
    * {List of restrictions for the catagory

NOTES:
    {Enter any additional notes that you believe will help the next developer.}

COPYRIGHT:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the License);
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an AS IS BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package sharedServices

import (
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/nats-io/nats.go"
)

//goland:noinspection GoSnakeCaseUsage,GoCommentStart
const (
	METHOD_DASHES      = "dashes"
	METHOD_UNDERSCORES = "underscores"
	METHOD_BLANK       = ""

	// Test constants
	TEST_CREDENTIALS_FILENAME = "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/nats-savup-backend.key"
	TEST_MESSAGE_ENVIRONMENT  = "local"
	TEST_MESSAGE_NAMESPACE    = "nci"
	TEST_URL                  = "savup-local-0030.savup.com"
	TEST_PORT                 = 4222
	TEST_PORT_EMPTY           = ""
	TEST_TLS_CERT             = "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/STAR_savup_com.crt"
	TEST_TLS_PRIVATE_KEY      = "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/savup.com.key"
	TEST_TLS_CA_BUNDLE        = "/Users/syacko/workspace/styh-dev/src/albert/keys/local/.keys/savup/STAR_savup_com/CAbundle.crt"
	//
	TEST_INVALID_URL = "invalid URL"
)

type MessageHandler struct {
	Handler nats.MsgHandler
}

type NATSReply struct {
	Response  interface{}   `json:"response,omitempty"`
	ErrorInfo cpi.ErrorInfo `json:"error,omitempty"`
}
