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
package coreStripe

import (
	"runtime"
	"testing"

	"albert/constants"
	"albert/core/coreError"
	"albert/core/coreHelpers"
)

func TestBuildStripeCustomerAccountId(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		_ = buildStripeCustomerAccountId(constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CHASE, constants.TEST_USER_BANK_ACCOUNT_ID_1)
	})
}

func TestIsStripeLockSet(tPtr *testing.T) {

	var (
		myServer           *Server
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	myServer = StartTest(tFunctionName, true, false)
	BuildTestUserInstitutions(myServer)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Returns true
		if isStripeLockSet(myServer.MyFireBase.FirestoreClientPtr, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CHASE_CLONE) == false {
			tPtr.Errorf("%v Failed: Was expecting no err.", tFunctionName)
		}
		// Returns false
		if isStripeLockSet(myServer.MyFireBase.FirestoreClientPtr, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CHASE) {
			tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
		}
	})

	RemoveTestInstitutions(myServer)
	StopTest(myServer)

}

func TestProcessCreateStripeCustomer(tPtr *testing.T) {

	type arguments struct {
		requestorId     string
		institutionName string
	}

	var (
		errorInfo          coreError.ErrorInfo
		myServer           *Server
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	// Since the system pulls the requestor id from the access token, the struct's only have an access token field.
	tests := []struct {
		name                    string
		arguments               arguments
		buildUserNoFederalTaxId bool
		wantError               bool
	}{
		{
			name: "Positive Case: Successful update!",
			arguments: arguments{
				requestorId:     constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				institutionName: constants.TEST_INSTITUTION_CITIZEN_BANK,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing requestor id!",
			arguments: arguments{
				requestorId:     constants.EMPTY,
				institutionName: constants.TEST_INSTITUTION_CHASE,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing institution",
			arguments: arguments{
				requestorId:     constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				institutionName: constants.EMPTY,
			},
			wantError: true,
		},
	}

	myServer = StartTest(tFunctionName, true, false)
	BuildTestUserInstitutions(myServer)

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = processCreateStripeCustomer(myServer.MyPlaid, myServer.MyFireBase, ts.arguments.requestorId, ts.arguments.institutionName); errorInfo.Error == nil {
				gotError = false
			} else {
				gotError = true
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
			}
		})
	}

	RemoveTestInstitutions(myServer)
	StopTest(myServer)

}

func TestProcessStripeCustomerTransfer(tPtr *testing.T) {

	type arguments struct {
		requestorId     string
		institutionName string
		plaidAccountId  string
		transferAmount  float64
		reportedBalance float64
	}

	var (
		myServer           *Server
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tReply             CustomerTransferReply
	)

	// Since the system pulls the requestor id from the access token, the struct's only have an access token field.
	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful transfer!",
			arguments: arguments{
				requestorId:     constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				institutionName: constants.TEST_INSTITUTION_CITIZEN_BANK,
				plaidAccountId:  constants.TEST_CITIZEN_PLAID_ACCOUNT_ID,
				transferAmount:  40,
				reportedBalance: 100,
			},
			wantError: false,
		},
	}

	myServer = StartTest(tFunctionName, true, false)
	BuildTestUserInstitutions(myServer)
	BuildTestStripeCustomer(myServer)

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if tReply = processStripeCustomerTransfer(myServer.MyFireBase, ts.arguments.requestorId, ts.arguments.institutionName, ts.arguments.plaidAccountId, ts.arguments.transferAmount, ts.arguments.reportedBalance); tReply.Error == constants.EMPTY {
				gotError = false
			} else {
				gotError = true
			}
			if gotError != ts.wantError {
				tPtr.Error(tFunctionName, ts.name, tReply.Error)
			}
		})
	}

	deleteStripeCustomer(buildStripeCustomerAccountId(constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CITIZEN_BANK, constants.TEST_CITIZEN_PLAID_ACCOUNT_ID))
	RemoveTestTransferRecords(myServer, constants.TRANFER_STRIPE, coreHelpers.GetDate())
	RemoveTestUserInstitutions(myServer)
	StopTest(myServer)

}

func TestStripeCreateCustomer(tPtr *testing.T) {

	var (
		errorInfo               coreError.ErrorInfo
		myServer                *Server
		tFunction, _, _, _      = runtime.Caller(0)
		tFunctionName           = runtime.FuncForPC(tFunction).Name()
		tStripeBankAccountToken string
	)

	myServer = StartTest(tFunctionName, true, false)
	tStripeCustomerAccountId := buildStripeCustomerAccountId(constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CITIZEN_BANK, constants.TEST_CITIZEN_PLAID_ACCOUNT_ID)
	tStripeBankAccountToken, _ = getPlaidStripeBankToken(myServer.MyPlaid.PlaidClient, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_CITIZEN_PLAID_ACCESS_TOKEN, constants.TEST_CITIZEN_PLAID_ACCOUNT_ID)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Successful is not testable, at this time, because of the many steps to create the needed test data.
		// If this test fails, must likely it is because the Stripe Customer database already has this record from the last test run.
		if errorInfo = createStripeCustomer(myServer.MyFireBase, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CITIZEN_BANK, tStripeCustomerAccountId, tStripeBankAccountToken); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting no err.", tFunctionName)
		}
		// Duplicate create request
		if errorInfo = createStripeCustomer(myServer.MyFireBase, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CITIZEN_BANK, tStripeCustomerAccountId, tStripeBankAccountToken); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
		}
		// Missing Stripe customer accountId
		if errorInfo = createStripeCustomer(myServer.MyFireBase, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CHASE, constants.EMPTY, constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
		}
		// Missing requestor id
		if errorInfo = createStripeCustomer(myServer.MyFireBase, constants.EMPTY, constants.TEST_INSTITUTION_CHASE, tStripeCustomerAccountId, constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
		}
		// Missing institution name
		if errorInfo = createStripeCustomer(myServer.MyFireBase, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.EMPTY, tStripeCustomerAccountId, constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
		}
	})

	deleteStripeCustomer(tStripeCustomerAccountId)
	StopTest(myServer)

}

func TestSearchStripeCustomer(tPtr *testing.T) {

	var (
		myServer           *Server
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tCustomer          Customer
	)

	myServer = StartTest(tFunctionName, true, false)
	BuildTestStripeCustomer(myServer)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Not Found
		if tCustomer = searchStripeCustomer(constants.EMPTY); tCustomer.Name != constants.EMPTY {
			tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
		}
		//  Found
		if tCustomer = searchStripeCustomer(buildStripeCustomerAccountId(constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CITIZEN_BANK, constants.TEST_CITIZEN_PLAID_ACCOUNT_ID)); tCustomer.Name == constants.EMPTY {
			tPtr.Errorf("%v Failed: Was expecting no err.", tFunctionName)
		}
	})

	StopTest(myServer)

}

func TestDeleteStripeCustomer(tPtr *testing.T) {

	var (
		myServer           *Server
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tCustomer          Customer
	)

	myServer = StartTest(tFunctionName, true, false)
	BuildTestStripeCustomer(myServer)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Not Found
		if tCustomer = deleteStripeCustomer(constants.EMPTY); tCustomer.Name != constants.EMPTY {
			tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
		}
		//  Found
		if tCustomer = deleteStripeCustomer(buildStripeCustomerAccountId(constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CITIZEN_BANK, constants.TEST_CITIZEN_PLAID_ACCOUNT_ID)); tCustomer.Name == constants.EMPTY {
			tPtr.Errorf("%v Failed: Was expecting no err.", tFunctionName)
		}
	})

	StopTest(myServer)

}

func TestStripeGetKey(tPtr *testing.T) {

	type arguments struct {
		stripeFQN string
	}

	var (
		errorInfo coreError.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Negative Case: File doesn't exist!",
			arguments: arguments{
				stripeFQN: constants.TEST_NO_SUCH_FILE,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Bad JSON!",
			arguments: arguments{
				stripeFQN: constants.TEST_JSON_INVALID,
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _ = getStripeKey(ts.arguments.stripeFQN, true); errorInfo.Error != nil {
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
