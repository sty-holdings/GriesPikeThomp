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
	"fmt"
	"runtime"
	"testing"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/stripe/stripe-go/v76"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

//goland:noinspection GoSnakeCaseUsage,GoCommentStart
const (
	TEST_KEY = "sk_test_51LalVGK3aJ31D0ASERSRRZ5bxTaMBMm7v5CYgCtLkJ8QCzyd3TecGD4Kv3Wk6NkCWL3LOplumLK30cA3RqOnNtK400cDqiATbp"
)

// func TestBuildStripeCustomerAccountId(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(
// 		tFunctionName, func(t *testing.T) {
// 			_ = buildStripeCustomerAccountId(
// 				constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 				constants.TEST_INSTITUTION_CHASE,
// 				constants.TEST_USER_BANK_ACCOUNT_ID_1,
// 			)
// 		},
// 	)
// }

// func TestIsStripeLockSet(tPtr *testing.T) {
//
// 	var (
// 		myServer           *Server
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	myServer = StartTest(tFunctionName, true, false)
// 	BuildTestUserInstitutions(myServer)
//
// 	tPtr.Run(
// 		tFunctionName, func(tPtr *testing.T) {
// 			// Returns true
// 			if isStripeLockSet(
// 				myServer.MyFireBase.FirestoreClientPtr,
// 				constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 				constants.TEST_INSTITUTION_CHASE_CLONE,
// 			) == false {
// 				tPtr.Errorf("%v Failed: Was expecting no err.", tFunctionName)
// 			}
// 			// Returns false
// 			if isStripeLockSet(myServer.MyFireBase.FirestoreClientPtr, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.TEST_INSTITUTION_CHASE) {
// 				tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
// 			}
// 		},
// 	)
//
// 	RemoveTestInstitutions(myServer)
// 	StopTest(myServer)
//
// }

// func TestProcessCreateStripeCustomer(tPtr *testing.T) {
//
// 	type arguments struct {
// 		requestorId     string
// 		institutionName string
// 	}
//
// 	var (
// 		errorInfo          coreError.ErrorInfo
// 		myServer           *Server
// 		gotError           bool
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	// Since the system pulls the requestor id from the access token, the struct's only have an access token field.
// 	tests := []struct {
// 		name                    string
// 		arguments               arguments
// 		buildUserNoFederalTaxId bool
// 		wantError               bool
// 	}{
// 		{
// 			name: "Positive Case: Successful update!",
// 			arguments: arguments{
// 				requestorId:     constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 				institutionName: constants.TEST_INSTITUTION_CITIZEN_BANK,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Missing requestor id!",
// 			arguments: arguments{
// 				requestorId:     constants.EMPTY,
// 				institutionName: constants.TEST_INSTITUTION_CHASE,
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Missing institution",
// 			arguments: arguments{
// 				requestorId:     constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 				institutionName: constants.EMPTY,
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	myServer = StartTest(tFunctionName, true, false)
// 	BuildTestUserInstitutions(myServer)
//
// 	for _, ts := range tests {
// 		tPtr.Run(
// 			ts.name, func(t *testing.T) {
// 				if errorInfo = processCreateStripeCustomer(
// 					myServer.MyPlaid,
// 					myServer.MyFireBase,
// 					ts.arguments.requestorId,
// 					ts.arguments.institutionName,
// 				); errorInfo.Error == nil {
// 					gotError = false
// 				} else {
// 					gotError = true
// 				}
// 				if gotError != ts.wantError {
// 					tPtr.Error(ts.name)
// 				}
// 			},
// 		)
// 	}
//
// 	RemoveTestInstitutions(myServer)
// 	StopTest(myServer)
//
// }

func TestProcessCancelPaymentIntent(tPtr *testing.T) {

	type arguments struct {
		request CancelPaymentIntentRequest
	}

	var (
		errorInfo          cpi.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	// Since the system pulls the requestor id from the access token, the struct's only have an access token field.
	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "missing key",
			arguments: arguments{
				request: CancelPaymentIntentRequest{
					Key:             rcv.VAL_EMPTY,
					PaymentIntentId: "pi_3OotQbK3aJ31D0AS0QeNlREa",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "missing payment intent id",
			arguments: arguments{
				request: CancelPaymentIntentRequest{
					Key: TEST_KEY,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Successful!",
			arguments: arguments{
				request: CancelPaymentIntentRequest{
					Key:             TEST_KEY,
					PaymentIntentId: "pi_3OotPzK3aJ31D0AS0kHvX5Fs",
				},
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				fmt.Println(rcv.LINE_LONG)
				if _, errorInfo = processCancelPaymentIntent(ts.arguments.request); errorInfo.Error == nil {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Error(tFunctionName, ts.name, errorInfo.Error.Error())
				}
			},
		)
	}
}

func TestProcessConfirmPaymentIntent(tPtr *testing.T) {

	type arguments struct {
		request ConfirmPaymentIntentRequest
	}

	var (
		errorInfo          cpi.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	// Since the system pulls the requestor id from the access token, the struct's only have an access token field.
	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "missing key",
			arguments: arguments{
				request: ConfirmPaymentIntentRequest{
					Key: rcv.VAL_EMPTY,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "missing payment intent id",
			arguments: arguments{
				request: ConfirmPaymentIntentRequest{
					Key: TEST_KEY,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "missing return URL",
			arguments: arguments{
				request: ConfirmPaymentIntentRequest{
					Key:             TEST_KEY,
					PaymentIntentId: "pi_3OotA4K3aJ31D0AS0BPWECZn",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "missing payment method",
			arguments: arguments{
				request: ConfirmPaymentIntentRequest{
					Key:             TEST_KEY,
					PaymentIntentId: "pi_3OotA4K3aJ31D0AS0BPWECZn",
					ReturnURL:       "https://stripe.natsconnect.com",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Successful!",
			arguments: arguments{
				request: ConfirmPaymentIntentRequest{
					Key:             TEST_KEY,
					PaymentIntentId: "pi_3OotA4K3aJ31D0AS0BPWECZn",
					PaymentMethod:   rcv.CARD_BRAND_VISA,
					ReturnURL:       "https://stripe.natsconnect.com",
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				fmt.Println(rcv.LINE_LONG)
				if _, errorInfo = processConfirmPaymentIntent(ts.arguments.request); errorInfo.Error == nil {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Error(tFunctionName, ts.name, errorInfo.Error.Error())
				}
			},
		)
	}
}

func TestProcessListPaymentIntents(tPtr *testing.T) {

	type arguments struct {
		request ListPaymentIntentRequest
	}

	var (
		errorInfo          cpi.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	// Since the system pulls the requestor id from the access token, the struct's only have an access token field.
	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Successful List!",
			arguments: arguments{
				request: ListPaymentIntentRequest{
					Key: TEST_KEY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "key is missing",
			arguments: arguments{
				request: ListPaymentIntentRequest{
					Key: rcv.VAL_EMPTY,
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if _, errorInfo = processListPaymentIntents(ts.arguments.request); errorInfo.Error == nil {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Error(tFunctionName, ts.name, errorInfo.Error.Error())
				}
			},
		)
	}
}

func TestProcessListPaymentMethods(tPtr *testing.T) {

	type arguments struct {
		request ListPaymentMethodRequest
	}

	var (
		errorInfo          cpi.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	// Since the system pulls the requestor id from the access token, the struct's only have an access token field.
	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Successful List!",
			arguments: arguments{
				request: ListPaymentMethodRequest{
					Key: TEST_KEY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Successful key is missing",
			arguments: arguments{
				request: ListPaymentMethodRequest{
					Key: rcv.VAL_EMPTY,
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if _, errorInfo = processListPaymentMethods(ts.arguments.request); errorInfo.Error == nil {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Error(tFunctionName, ts.name, errorInfo.Error.Error())
				}
			},
		)
	}
}

func TestProcessCreatePaymentIntent(tPtr *testing.T) {

	type arguments struct {
		request PaymentIntentRequest
	}

	var (
		errorInfo          cpi.ErrorInfo
		gotError           bool
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	// Since the system pulls the requestor id from the access token, the struct's only have an access token field.
	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Successful payment no automatic payment method or description!",
			arguments: arguments{
				request: PaymentIntentRequest{
					Amount:   1.01,
					Currency: string(stripe.CurrencyUSD),
					Key:      TEST_KEY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Successful payment no automatic payment method!",
			arguments: arguments{
				request: PaymentIntentRequest{
					Amount:      1.01,
					Currency:    string(stripe.CurrencyUSD),
					Description: rcv.TXT_EMPTY,
					Key:         TEST_KEY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Successful payment with automatic payment method no description!",
			arguments: arguments{
				request: PaymentIntentRequest{
					Amount:                  1.01,
					AutomaticPaymentMethods: true,
					Currency:                string(stripe.CurrencyUSD),
					Key:                     TEST_KEY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Successful payment with automatic payment method and description!",
			arguments: arguments{
				request: PaymentIntentRequest{
					Amount:                  1.01,
					AutomaticPaymentMethods: true,
					Currency:                string(stripe.CurrencyUSD),
					Description:             rcv.TXT_EMPTY,
					Key:                     TEST_KEY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Zero amount",
			arguments: arguments{
				request: PaymentIntentRequest{
					Amount:      0,
					Currency:    string(stripe.CurrencyUSD),
					Description: rcv.TXT_EMPTY,
					Key:         TEST_KEY,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Empty currency",
			arguments: arguments{
				request: PaymentIntentRequest{
					Amount:      1.01,
					Currency:    rcv.VAL_EMPTY,
					Description: rcv.TXT_EMPTY,
					Key:         TEST_KEY,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Uppercase currency",
			arguments: arguments{
				request: PaymentIntentRequest{
					Amount:      1.01,
					Currency:    "USD",
					Description: rcv.TXT_EMPTY,
					Key:         rcv.VAL_EMPTY,
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "Empty key",
			arguments: arguments{
				request: PaymentIntentRequest{
					Amount:      1.01,
					Currency:    string(stripe.CurrencyUSD),
					Description: rcv.TXT_EMPTY,
					Key:         rcv.VAL_EMPTY,
				},
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				fmt.Println(rcv.LINE_LONG)
				if _, errorInfo = processCreatePaymentIntent(ts.arguments.request); errorInfo.Error == nil {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Error(tFunctionName, ts.name, errorInfo.Error.Error())
				}
			},
		)
	}
}

// func TestStripeCreateCustomer(tPtr *testing.T) {
//
// 	var (
// 		errorInfo               coreError.ErrorInfo
// 		myServer                *Server
// 		tFunction, _, _, _      = runtime.Caller(0)
// 		tFunctionName           = runtime.FuncForPC(tFunction).Name()
// 		tStripeBankAccountToken string
// 	)
//
// 	myServer = StartTest(tFunctionName, true, false)
// 	tStripeCustomerAccountId := buildStripeCustomerAccountId(
// 		constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 		constants.TEST_INSTITUTION_CITIZEN_BANK,
// 		constants.TEST_CITIZEN_PLAID_ACCOUNT_ID,
// 	)
// 	tStripeBankAccountToken, _ = getPlaidStripeBankToken(
// 		myServer.MyPlaid.PlaidClient,
// 		constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 		constants.TEST_CITIZEN_PLAID_ACCESS_TOKEN,
// 		constants.TEST_CITIZEN_PLAID_ACCOUNT_ID,
// 	)
//
// 	tPtr.Run(
// 		tFunctionName, func(tPtr *testing.T) {
// 			// Successful is not testable, at this time, because of the many steps to create the needed test data.
// 			// If this test fails, must likely it is because the Stripe Customer database already has this record from the last test run.
// 			if errorInfo = createStripeCustomer(
// 				myServer.MyFireBase,
// 				constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 				constants.TEST_INSTITUTION_CITIZEN_BANK,
// 				tStripeCustomerAccountId,
// 				tStripeBankAccountToken,
// 			); errorInfo.Error != nil {
// 				tPtr.Errorf("%v Failed: Was expecting no err.", tFunctionName)
// 			}
// 			// Duplicate create request
// 			if errorInfo = createStripeCustomer(
// 				myServer.MyFireBase,
// 				constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 				constants.TEST_INSTITUTION_CITIZEN_BANK,
// 				tStripeCustomerAccountId,
// 				tStripeBankAccountToken,
// 			); errorInfo.Error == nil {
// 				tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
// 			}
// 			// Missing Stripe customer accountId
// 			if errorInfo = createStripeCustomer(
// 				myServer.MyFireBase,
// 				constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 				constants.TEST_INSTITUTION_CHASE,
// 				constants.EMPTY,
// 				constants.EMPTY,
// 			); errorInfo.Error == nil {
// 				tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
// 			}
// 			// Missing requestor id
// 			if errorInfo = createStripeCustomer(
// 				myServer.MyFireBase,
// 				constants.EMPTY,
// 				constants.TEST_INSTITUTION_CHASE,
// 				tStripeCustomerAccountId,
// 				constants.EMPTY,
// 			); errorInfo.Error == nil {
// 				tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
// 			}
// 			// Missing institution name
// 			if errorInfo = createStripeCustomer(
// 				myServer.MyFireBase,
// 				constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 				constants.EMPTY,
// 				tStripeCustomerAccountId,
// 				constants.EMPTY,
// 			); errorInfo.Error == nil {
// 				tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
// 			}
// 		},
// 	)
//
// 	deleteStripeCustomer(tStripeCustomerAccountId)
// 	StopTest(myServer)
//
// }

// func TestSearchStripeCustomer(tPtr *testing.T) {
//
// 	var (
// 		myServer           *Server
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tCustomer          Customer
// 	)
//
// 	myServer = StartTest(tFunctionName, true, false)
// 	BuildTestStripeCustomer(myServer)
//
// 	tPtr.Run(
// 		tFunctionName, func(tPtr *testing.T) {
// 			// Not Found
// 			if tCustomer = searchStripeCustomer(constants.EMPTY); tCustomer.Name != constants.EMPTY {
// 				tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
// 			}
// 			//  Found
// 			if tCustomer = searchStripeCustomer(
// 				buildStripeCustomerAccountId(
// 					constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 					constants.TEST_INSTITUTION_CITIZEN_BANK,
// 					constants.TEST_CITIZEN_PLAID_ACCOUNT_ID,
// 				),
// 			); tCustomer.Name == constants.EMPTY {
// 				tPtr.Errorf("%v Failed: Was expecting no err.", tFunctionName)
// 			}
// 		},
// 	)
//
// 	StopTest(myServer)
//
// }

// func TestDeleteStripeCustomer(tPtr *testing.T) {
//
// 	var (
// 		myServer           *Server
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tCustomer          Customer
// 	)
//
// 	myServer = StartTest(tFunctionName, true, false)
// 	BuildTestStripeCustomer(myServer)
//
// 	tPtr.Run(
// 		tFunctionName, func(tPtr *testing.T) {
// 			// Not Found
// 			if tCustomer = deleteStripeCustomer(constants.EMPTY); tCustomer.Name != constants.EMPTY {
// 				tPtr.Errorf("%v Failed: Was expecting an err.", tFunctionName)
// 			}
// 			//  Found
// 			if tCustomer = deleteStripeCustomer(
// 				buildStripeCustomerAccountId(
// 					constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
// 					constants.TEST_INSTITUTION_CITIZEN_BANK,
// 					constants.TEST_CITIZEN_PLAID_ACCOUNT_ID,
// 				),
// 			); tCustomer.Name == constants.EMPTY {
// 				tPtr.Errorf("%v Failed: Was expecting no err.", tFunctionName)
// 			}
// 		},
// 	)
//
// 	StopTest(myServer)
//
// }
