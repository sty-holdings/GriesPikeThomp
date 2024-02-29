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
	"encoding/json"
	"fmt"
	"log"
	"reflect"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	ext "github.com/sty-holdings/GriesPikeThomp/servers/nats-connect/loadExtensions"
	cpi "github.com/sty-holdings/GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type NATSService struct {
	ConnPtr        *nats.Conn
	CredentialsFQN string
	Secure         bool
	URL            string
}

// BuildInstanceName - will create the NATS connection name with dashes, underscores between nodes or as provided.
// The method can be cn.METHOD_DASHES, cn.METHOD_UNDERSCORES, rcv.VAL_EMPTY, "dashes", "underscores" or ""
//
//	Customer Messages: None
//	Errors: error returned by nats.Connect
//	Verifications: None
func BuildInstanceName(
	method string,
	nodes ...string,
) (
	instanceName string,
	errorInfo cpi.ErrorInfo,
) {

	if len(nodes) == 1 {
		method = METHOD_BLANK
	}
	switch strings.Trim(strings.ToLower(method), rcv.SPACES_ONE) {
	case METHOD_DASHES:
		instanceName, errorInfo = buildInstanceName(rcv.DASH, nodes...)
	case METHOD_UNDERSCORES:
		instanceName, errorInfo = buildInstanceName(rcv.UNDERSCORE, nodes...)
	default:
		instanceName, errorInfo = buildInstanceName(rcv.VAL_EMPTY, nodes...)
	}

	return
}

// GetConnection - will connect to a NATS leaf server with either a ssl or non-ssl connection.
//
//	Customer Messages: None
//	Errors: error returned by nats.Connect
//	Verifications: None
func GetConnection(
	instanceName string,
	config ext.ExtensionConfiguration,
) (
	connPtr *nats.Conn,
	errorInfo cpi.ErrorInfo,
) {

	var (
		opts []nats.Option
		tURL string
	)

	opts = []nats.Option{
		nats.Name(instanceName),             // Set a client name
		nats.MaxReconnects(5),               // Set maximum reconnection attempts
		nats.ReconnectWait(5 * time.Second), // Set reconnection wait time
		nats.UserCredentials(config.NATSCredentialsFilename),
		nats.RootCAs(config.NATSTLSInfo.TLSCABundle),
		nats.ClientCert(config.NATSTLSInfo.TLSCert, config.NATSTLSInfo.TLSPrivateKey),
	}

	if tURL, errorInfo = buildURLPort(config.NATSURL, config.NATSPort); errorInfo.Error != nil {
		return
	}
	if connPtr, errorInfo.Error = nats.Connect(tURL, opts...); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v: %v", instanceName, rcv.TXT_SECURE_CONNECTION_FAILED))
		return
	}

	log.Printf("%v: A connection has been established with the NATS server at %v.", instanceName, config.NATSURL)
	log.Printf(
		"%v: URL: %v Server Name: %v Server Id: %v Address: %v",
		instanceName,
		connPtr.ConnectedUrl(),
		connPtr.ConnectedClusterName(),
		connPtr.ConnectedServerId(),
		connPtr.ConnectedAddr(),
	)

	return
}

// SendReply - will take in an object, build a json object out of it, and send out the reply
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func SendReply(
	reply interface{},
	msg *nats.Msg,
) (errorInfo cpi.ErrorInfo) {

	var (
		tJSONReply []byte
	)

	if tJSONReply, errorInfo = buildJSONReply(reply); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v%v%v", rcv.TXT_SUBJECT, msg.Subject, rcv.TXT_MESSAGE_HEADER, msg.Header))
		return
	}

	if errorInfo.Error = msg.Respond(tJSONReply); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v%v%v", rcv.TXT_SUBJECT, msg.Subject, rcv.TXT_MESSAGE_HEADER, msg.Header))
	}

	return
}

// Subscribe - will create a NATS subscription
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func Subscribe(
	connectionPtr *nats.Conn,
	instanceName, subject string,
	handler nats.MsgHandler,
) (
	subscriptionPtr *nats.Subscription,
	errorInfo cpi.ErrorInfo,
) {

	if subscriptionPtr, errorInfo.Error = connectionPtr.Subscribe(subject, handler); errorInfo.Error != nil {
		log.Printf("%v: Subscribe failed on subject: %v", instanceName, subject)
		return
	}
	log.Printf("%v Subscribed to subject: %v", instanceName, subject)

	return
}

// UnmarshalMessageData - reads the message data into the pointer. The second argument must be a pointer. If you pass something else, the unmarshal will fail.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func UnmarshalMessageData(
	functionName string,
	msg *nats.Msg,
	requestPtr any,
) (errorInfo cpi.ErrorInfo) {

	if string(msg.Data) == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrRequiredArgumentMissing, fmt.Sprintf("%v%v", rcv.TXT_FUNCTION_NAME, functionName))
		return
	}

	if errorInfo.Error = json.Unmarshal(msg.Data, requestPtr); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_FUNCTION_NAME, functionName))
	}

	return
}

//  Private Functions

// buildInstanceName - will create the NATS connection name with the delimiter between nodes.
//
//	Customer Messages: None
//	Errors: error returned by nats.Connect
//	Verifications: None
func buildInstanceName(
	delimiter string,
	nodes ...string,
) (
	instanceName string,
	errorInfo cpi.ErrorInfo,
) {

	if len(nodes) == rcv.VAL_ZERO {
		errorInfo = cpi.NewErrorInfo(cpi.ErrRequiredArgumentMissing, fmt.Sprint(rcv.TXT_AT_LEAST_ONE))
		return
	}
	for index, node := range nodes {
		if index == 0 {
			instanceName = strings.Trim(node, rcv.SPACES_ONE)
		} else {
			instanceName = fmt.Sprintf("%v%v%v", instanceName, delimiter, strings.Trim(node, rcv.SPACES_ONE))
		}
	}

	return
}

// buildJSONReply - return a JSON reply object
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func buildJSONReply(reply interface{}) (
	jsonReply []byte,
	errorInfo cpi.ErrorInfo,
) {

	if jsonReply, errorInfo.Error = json.Marshal(reply); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_REPLY_TYPE, reflect.ValueOf(reply).Type().String()))
		return
	}

	return
}

// buildURLPort - will create the NATS URL with the port.
//
//	Customer Messages: None
//	Errors: error returned by nats.Connect
//	Verifications: None
func buildURLPort(
	url string,
	port int,
) (
	natsURL string,
	errorInfo cpi.ErrorInfo,
) {

	if url == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrRequiredArgumentMissing, fmt.Sprint(rcv.FN_URL))
		return
	}
	if port == rcv.VAL_ZERO {
		errorInfo = cpi.NewErrorInfo(cpi.ErrGreatThanZero, fmt.Sprint(rcv.FN_PORT))
		return
	}

	return fmt.Sprintf("%v:%d", url, port), cpi.ErrorInfo{}
}
