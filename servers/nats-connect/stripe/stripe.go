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
	"fmt"
	"log"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/customer"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// CreateStripeCustomer - will create a customer in the Stripe database using the stripeCustomerAccountId retrieved from buildStripeCustomerAccountId function, if the Stripe customer
// doesn't already exist. The Stripe customer account id is used by Stripe customer.Search API, which returns the Stripe customer object if found.
//
//	Customer Messages: None
//	Errors: coreError.ErrRequiredArgumentMissing, coreError.ErrStripeCreateCustomerFailed
//	Verifications: requestorId, institutionName must be populated.
//			The requestorId must exist is the SavUp database.
//			Checks if the Stripe Customer already exists in the Stripe database.
func CreateStripeCustomer(
	institutionName, customerAccountId, bankAccountToken, email string,
	address Address,
) (
	result string,
	errorInfo cpi.ErrorInfo,
) {

	// var (
	// 	tStripeCustomer Customer
	// )

	if institutionName == rcv.VAL_EMPTY || customerAccountId == rcv.VAL_EMPTY || bankAccountToken == rcv.VAL_EMPTY {
		errorInfo.Error = cpi.ErrRequiredArgumentMissing
		log.Println(errorInfo.Error.Error())
	} else {

		// if tStripeCustomer = searchStripeCustomer(customerAccountId); tStripeCustomer.Id == rcv.VAL_EMPTY {
		// 	params := &stripe.CustomerParams{
		// 		Address: buildCustomerAddress(address),
		// 		Email:   stripe.String(email),
		// 		Phone:   stripe.String(fmt.Sprintf("(%v) %v", tUserInfo.AreaCode, tUserInfo.PhoneNumber)),
		// 	}
		// 	// Creating Stripe customer
		// 	params.Name = stripe.String(stripeCustomerAccountId)
		// 	params.Source = stripe.String(stripeBankAccountToken)
		// 	params.TaxExempt = stripe.String("exempt")
		// 	if _, errorInfo.Error = customer.New(params); errorInfo.Error != nil {
		// 		errorInfo.Error = coreError.ErrStripeCreateCustomerFailed
		// 		log.Println(errorInfo.Error.Error())
		// 	}
		// } else {
		// 	errorInfo.Error = coreError.ErrUserAlreadyExists
		// }
	}

	return
}

// Private functions

// buildCustomerAddress - returns a Stripe formatted address.
//
//		Errors: None
//		Customer Messages: None
//	 Validations: None
func buildCustomerAddress(address Address) (customerAddressPtr *stripe.AddressParams) {

	customerAddressPtr = &stripe.AddressParams{
		City:       stripe.String(address.City),
		Country:    stripe.String(address.Country),
		Line1:      stripe.String(address.Line1),
		Line2:      stripe.String(address.Line2),
		PostalCode: stripe.String(address.PostalCode),
		State:      stripe.String(address.State),
	}

	return
}

// buildPlaidInstitutionCustomerAccountId
//
//	Errors: None
//	Customer Messages: None
func buildPlaidInstitutionCustomerAccountId(plaidAccountId, institutionName string) string {

	return fmt.Sprintf("%v-%v",
		institutionName,
		plaidAccountId,
	)
}

// stripeGetKey - will read and parse the JSON key file. If either fail, exit is called.
//
//	Validations: File readable and JSON valid
// func getStripeKey(stripeFQN string, test bool) (stripeKey Helper) {
//
// 	var (
// 		errorInfo          coreError.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tStripe            []byte
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	if tStripe, errorInfo.Error = os.ReadFile(stripeFQN); errorInfo.Error != nil {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Required Stripe key file %v has issue.%v", stripeFQN, rcv.ENDING_EXECUTION))
// 		log.Println(errorInfo.Error.Error())
// 	} else {
// 		if errorInfo.Error = json.Unmarshal(tStripe, &stripeKey); errorInfo.Error != nil {
// 			errorInfo.Error = errors.New(fmt.Sprintf("Stripe JSON file %v is corrupt.%v", stripeFQN, rcv.ENDING_EXECUTION))
// 			log.Println(errorInfo.Error.Error())
// 		}
// 	}
//
// 	if errorInfo.Error != nil {
// 		os.Exit(1)
// 	}
//
// 	return
// }

// isStripeLockSet - will return the value of stripe_lock or false is the field doesn't exist.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func isStripeLockSet(firestoreClientPtr *firestore.Client, requestorId, institutionName string) bool {
//
// 	var (
// 		data      = make(map[string]interface{})
// 		errorInfo coreError.ErrorInfo
// 	)
//
// 	if data, errorInfo = coreFirestore.GetDocumentFromSubCollectionByDocumentId(firestoreClientPtr, rcv.DATASTORE_USER_INSTITUTIONS, requestorId, rcv.COLLECTION_INSTITUTIONS, institutionName); errorInfo.Error == nil {
// 		if _, exists := data[rcv.FN_STRIPE_LOCK]; exists {
// 			return data[rcv.FN_STRIPE_LOCK].(bool)
// 		}
// 	}
//
// 	return false
// }

// processCreateStripeCustomer - will create the customer on the Stripe database if the stripe_lock is not set to true for the institution. The Stripe account can be found using the
// customer.Search and buildStripeCustomerAccountId function. All accounts for the institution will get a Stripe customer entry, so payments can be processed. Stripe will create
// multiple records for the same customer, so there is a Stripe Locked field for the institution.
//
//	Customer Messages: None
//	Errors: Any error that is returned by createStripeCustomer or the updateInstitutionStripeLock function.
//	Verification: None
// func processCreateStripeCustomer(myPlaid PlaidHelper, myFirebase coreHelpers.FirebaseFirestoreHelper, requestorId, institutionName string) (errorInfo coreError.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _      = runtime.Caller(0)
// 		tFunctionName           = runtime.FuncForPC(tFunction).Name()
// 		tInstitutionAccounts    []string
// 		tPlaidAccessToken       string
// 		tStripeBankAccountToken string
// 	)
//
// 	coreError.PrintDebugTrail(tFunctionName)
//
// 	if isStripeLockSet(myFirebase.FirestoreClientPtr, requestorId, institutionName) == false {
// 		if tPlaidAccessToken, errorInfo = getInstitutionAccessToken(myFirebase.FirestoreClientPtr, requestorId, institutionName); errorInfo.Error == nil {
// 			if tInstitutionAccounts = getInstitutionAccountIds(myFirebase.FirestoreClientPtr, requestorId, institutionName); len(tInstitutionAccounts) > 0 {
// 				for i := 0; i < len(tInstitutionAccounts); i++ {
// 					tStripeBankAccountToken, errorInfo = getPlaidStripeBankToken(myPlaid.PlaidClient, requestorId, tPlaidAccessToken, tInstitutionAccounts[i])
// 					errorInfo = createStripeCustomer(myFirebase, requestorId, institutionName, buildStripeCustomerAccountId(requestorId, institutionName, tInstitutionAccounts[i]), tStripeBankAccountToken)
// 				}
// 				errorInfo = updateInstitutionStripeLock(myFirebase, requestorId, institutionName)
// 			}
// 		}
// 	}
//
// 	return
// }

// processStripeCustomerTransfer - will handle a Stripe transfer into SavUp.
//
//	Errors: coreError.ErrTransferOutNotAllowed, coreError.ErrUserMissing
//	Customer Message: coreCustomerMessages.StripeTransferOut, coreCustomerMessages.UserMissing
//	Verifications: requestorId must exist, institutionName and transfer amount are required, transferAmount must be greater than zero.
// func processStripeCustomerTransfer(myFirebase coreHelpers.FirebaseFirestoreHelper, requestorId, institutionName, plaidAccountId string, transferAmount, reportedBalance float64) (reply CustomerTransferReply) {
//
// 	var (
// 		errorInfo     coreError.ErrorInfo
// 		tChargeResult *stripe.Charge
// 		tCustomer     Customer
// 	)
//
// 	if doesRequestorIdExist(myFirebase.FirestoreClientPtr, requestorId) {
// 		//  Stripe is a payment gateway, so only payments to SavUp are valid
// 		if transferAmount > 0 {
// 			if tCustomer = searchStripeCustomer(buildStripeCustomerAccountId(requestorId, institutionName, plaidAccountId)); tCustomer.Name == rcv.VAL_EMPTY {
// 				reply.Error = coreError.ErrUserMissing.Error()
// 			} else {
// 				chargeParams := &stripe.ChargeParams{
// 					Amount:      stripe.Int64(coreHelpers.FloatToPennies(transferAmount)),
// 					Currency:    stripe.String(string(stripe.CurrencyUSD)),
// 					Customer:    stripe.String(tCustomer.Id),
// 					Description: stripe.String(fmt.Sprintf("ACH transfer to SavUp - powered by STY-Holdings")),
// 				}
// 				if tChargeResult, errorInfo.Error = charge.New(chargeParams); errorInfo.Error == nil {
// 					_tChargeResult, _ := json.Marshal(tChargeResult)
// 					if errorInfo = recordTransfer(myFirebase.FirestoreClientPtr, requestorId, institutionName, plaidAccountId, rcv.TRANFER_STRIPE, string(_tChargeResult), rcv.FN_PLAID_ACCOUNTS, rcv.TRANSFER_IN, rcv.VAL_EMPTY, transferAmount, reportedBalance); errorInfo.Error == nil {
// 						reply.Message = coreCustomerMessages.PaymentSuccessful
// 					} else {
// 						reply.Error = errorInfo.Error.Error()
// 					}
// 				}
// 			}
// 		} else {
// 			reply.Error = coreError.ErrTransferOutNotAllowed.Error()
// 			reply.Message = coreCustomerMessages.StripeTransferOut
// 		}
// 	} else {
// 		reply.Error = coreError.ErrUserMissing.Error()
// 		reply.Message = coreCustomerMessages.UserMissing
// 	}
//
// 	return
// }

// searchStripeCustomer - uses the name field to return the Stripe customer object, if found.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func searchStripeCustomer(stripeCustomerAccountId string) (stripeCustomer Customer) {

	var (
		errorInfo cpi.ErrorInfo
	)

	params := &stripe.CustomerSearchParams{}
	params.Query = *stripe.String(fmt.Sprintf("name:'%v'",
		stripeCustomerAccountId,
	),
	)
	iter := customer.Search(params)
	for iter.Next() {
		result := iter.Current()
		jsonResult, _ := json.Marshal(result)
		errorInfo.Error = json.Unmarshal(jsonResult,
			&stripeCustomer,
		)
	}

	return
}

// deleteStripeCustomer - uses the id field to delete the Stripe customer object.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func deleteStripeCustomer(stripeCustomerAccountId string) (stripeCustomer Customer) {

	var (
		errorInfo      cpi.ErrorInfo
		customerParams stripe.CustomerParams
	)

	params := &stripe.CustomerSearchParams{}
	params.Query = *stripe.String(fmt.Sprintf("name:'%v'",
		stripeCustomerAccountId,
	),
	)
	iter := customer.Search(params)
	for iter.Next() {
		result := iter.Current()
		jsonResult, _ := json.Marshal(result)
		errorInfo.Error = json.Unmarshal(jsonResult,
			&stripeCustomer,
		)
		_, errorInfo.Error = customer.Del(stripeCustomer.Id,
			&customerParams,
		)
	}

	return
}
