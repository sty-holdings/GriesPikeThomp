// Package stripe
/*
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
package stripe

import (
	"encoding/json"
	"runtime"
	"testing"

	ext "GriesPikeThomp/servers/nats-connect/loadExtensions"
	"github.com/nats-io/nats.go"
)

func TestBuildExtension(tPtr *testing.T) {

	type arguments struct {
		extensionKey string
		config       ext.ExtensionConfiguration
	}

	var (
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
			name: "Positive Case: Successful!",
			arguments: arguments{
				field1: TEST_GOOD_FQN,
				field2: "Test Good FQN",
			},
			wantError: false,
		},
	}

	tEmailServerPtr, err = NewSendGridServer(constants.TEST_EMAIL_ADDRESS, constants.TEST_STRING, constants.ENVIRONMENT_DEVELOPMENT, constants.TEST_STRING)

	tFirebase.AppPtr, tFirebase.AuthPtr, tFirebase.FirestoreClientPtr, _ = shared.GetFirebaseFirestoreConnection(constants.TEST_FIREBASE_CREDENTIALS)

	tPlaidInfo := plaidGetKey(constants.TEST_PLAID_KEY_FILE)
	tPlaidClient, _ = plaidGetClientConnection(tPlaidInfo.ClientId, tPlaidInfo.Secret)

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			tRequestJSON, _ = json.Marshal(sut.TestCreateUserRequest)
			_, _ = executeCreateUser(tFirebase.FirestoreClientPtr, &nats.Msg{Data: tRequestJSON})
			tRequest.RequestorId = constants.TEST_USER_BANK_ACCOUNT_ID
			tRequestJSON, _ = json.Marshal(tRequest)
			if _, err = executePlaidGetLinkToken(tPlaidClient, tFirebase.FirestoreClientPtr, &nats.Msg{Data: tRequestJSON}); err != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(err.Error())
			}
		})
	}

	_ = shared.RemoveDocument(tFirebase.FirestoreClientPtr, constants.USERS_DATASTORE, shared.NameValueQuery{
		FieldName:  constants.REQUESTOR_ID_FN,
		FieldValue: constants.TEST_USER_REQUESTOR_ID,
	})
}
