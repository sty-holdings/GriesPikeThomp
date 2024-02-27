// Package coreStripe
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
	"strings"
	"sync"

	ext "GriesPikeThomp/servers/nats-connect/loadExtensions"
	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cn "GriesPikeThomp/shared-services/src/coreNATS"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/nats-io/nats.go"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// NewExtension - creates an instance by setting the values for the extension struct
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewExtension(
	hostname string,
	config ext.ExtensionConfiguration,
	testingOn bool,
) (
	errorInfo cpi.ErrorInfo,
) {

	var (
		stripeInstancePtr *stripeInstance
	)

	stripeInstancePtr = &stripeInstance{
		subscriptionPtrs: make(map[string]*nats.Subscription),
		testingOn:        testingOn,
		waitGroup:        sync.WaitGroup{},
	}
	if stripeInstancePtr.instanceName, errorInfo = cn.BuildInstanceName(cn.METHOD_DASHES, hostname, rcv.STRIPE_EXTENSION); errorInfo.Error != nil {
		return
	}

	stripeInstancePtr.messageHandles()

	stripeInstancePtr.processChannel = make(chan string) // This is for NC_INTERNAL only
	go func() {
		stripeInstancePtr.buildExtension(config)
	}()
	select {
	case <-stripeInstancePtr.processChannel:
	}

	fmt.Println(stripeInstancePtr.instanceName)

	return
}

// Private methods below here

// messageHandles - builds a map of messages handlers
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (stripeInstancePtr *stripeInstance) messageHandles() (
	handlers map[string]cn.MessageHandler,
) {

	handlers = make(map[string]cn.MessageHandler)

	handlers[STRIPE_PAYMENT_INTENT] = cn.MessageHandler{
		Handler: stripeInstancePtr.paymentIntent(),
	}

	return
}

// extensionHandler - starts extensions in goroutine.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (stripeInstancePtr *stripeInstance) buildExtension(config ext.ExtensionConfiguration) (
	errorInfo cpi.ErrorInfo,
) {
	var (
		tHandlers = make(map[string]cn.MessageHandler)
	)

	if stripeInstancePtr.natsConnectionPtr, errorInfo = cn.GetConnection(stripeInstancePtr.instanceName, config); errorInfo.Error != nil {
		return
	}

	tHandlers = stripeInstancePtr.messageHandles()

	// Use a WaitGroup to wait for a message to arrive
	stripeInstancePtr.waitGroup.Add(1)
	for _, loadSubject := range config.SubjectRegistry {
		if _, found := tHandlers[loadSubject.Subject]; found == false {
			errorInfo = cpi.NewErrorInfo(
				cpi.ErrSubjectInvalid,
				fmt.Sprintf("%v: %v%v", stripeInstancePtr.instanceName, rcv.TXT_SUBJECT, loadSubject.Subject),
			)
			cpi.PrintErrorInfo(errorInfo)
			return
		}
		stripeInstancePtr.subscriptionPtrs[loadSubject.Subject], errorInfo = cn.Subscribe(
			stripeInstancePtr.natsConnectionPtr,
			stripeInstancePtr.instanceName, loadSubject.Subject, tHandlers[loadSubject.Subject].Handler,
		)
	}

	if stripeInstancePtr.testingOn {
		stripeInstancePtr.waitGroup.Done()
	} else {
		stripeInstancePtr.waitGroup.Wait()
	}

	return

}

// NATS Message Handlers go below this line.

func (stripeInstancePtr *stripeInstance) paymentIntent() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo             cpi.ErrorInfo
			tPaymentIntentResults *stripe.PaymentIntent
			tReply                cn.NATSReply
			tRequest              PaymentIntentRequest
		)

		if errorInfo = chv.UnmarshalMessageData(cpi.GetFunctionInfo(1).Name, msg, &tRequest); errorInfo.Error == nil {
			tPaymentIntentResults, errorInfo = processPaymentIntent(tRequest.Amount, tRequest.Currency, tRequest.Description, tRequest.Key)
		}

		tReply.Response = tPaymentIntentResults

		if errorInfo = chv.SendReply(tReply, msg); errorInfo.Error != nil {
			cpi.PrintErrorInfo(errorInfo)
		}

		return
	}
}

// buildStripeCustomerAddress - returns a Stripe formatted address with the country set too USA.
//
//	Errors: None
//	Customer Messages: None
// func buildStripeCustomerAddress(streetAddress, city, state, zipCode string) (stripeCustomerAddressPtr *stripe.AddressParams) {
//
// 	stripeCustomerAddressPtr = &stripe.AddressParams{
// 		City:       stripe.String(city),
// 		Country:    stripe.String("USA"),
// 		Line1:      stripe.String(streetAddress),
// 		PostalCode: stripe.String(zipCode),
// 		State:      stripe.String(state),
// 	}
//
// 	return
// }

// buildStripeCustomerAccountId
//
//	Errors: None
//	Customer Messages: None
// func buildStripeCustomerAccountId(requestorId, institutionName, plaidAccountId string) string {
//
// 	return fmt.Sprintf("%v-%v-%v", requestorId, institutionName, plaidAccountId)
// }

// createStripeCustomer - will create a customer in the Stripe database using the stripeCustomerAccountId retrieved from buildStripeCustomerAccountId function, if the Stripe customer
// doesn't already exist. The Stripe customer account id is used by Stripe customer.Search API, which returns the Stripe customer object if found.
//
//	Customer Messages: None
//	Errors: cpi.ErrRequiredArgumentMissing, cpi.ErrStripeCreateCustomerFailed
//	Verifications: requestorId, institutionName must be populated.
//			The requestorId must exist is the SavUp database.
//			Checks if the Stripe Customer already exists in the Stripe database.
// func createStripeCustomer(
// 	myFirebase chv.FirebaseFirestoreHelper,
// 	requestorId, institutionName, stripeCustomerAccountId, stripeBankAccountToken string,
// ) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tStripeCustomer    Customer
// 		tUserInfo          UserInfo
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if requestorId == constants.EMPTY || institutionName == constants.EMPTY {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 		log.Println(errorInfo.Error.Error())
// 	} else {
// 		if tUserInfo, errorInfo = FindSavUpUserByRequestorId(myFirebase.FirestoreClientPtr, requestorId); errorInfo.Error == nil {
// 			if tStripeCustomer = searchStripeCustomer(stripeCustomerAccountId); tStripeCustomer.Id == constants.EMPTY {
// 				params := &stripe.CustomerParams{
// 					Address: buildStripeCustomerAddress(tUserInfo.StreetAddress, tUserInfo.City, tUserInfo.State, tUserInfo.ZipCode),
// 					Email:   stripe.String(tUserInfo.Email),
// 					Phone:   stripe.String(fmt.Sprintf("(%v) %v", tUserInfo.AreaCode, tUserInfo.PhoneNumber)),
// 				}
// 				// Creating Stripe customer
// 				params.Name = stripe.String(stripeCustomerAccountId)
// 				params.Source = stripe.String(stripeBankAccountToken)
// 				params.TaxExempt = stripe.String("exempt")
// 				if _, errorInfo.Error = customer.New(params); errorInfo.Error != nil {
// 					errorInfo.Error = cpi.ErrStripeCreateCustomerFailed
// 					log.Println(errorInfo.Error.Error())
// 				}
// 			} else {
// 				errorInfo.Error = cpi.ErrUserAlreadyExists
// 			}
// 		}
// 	}
//
// 	return
// }

// processCreateStripeCustomer - will create the customer on the Stripe database if the stripe_lock is not set to true for the institution. The Stripe account can be found using the
// customer.Search and buildStripeCustomerAccountId function. All accounts for the institution will get a Stripe customer entry, so payments can be processed. Stripe will create
// multiple records for the same customer, so there is a Stripe Locked field for the institution.
//
//	Customer Messages: None
//	Errors: Any error that is returned by createStripeCustomer or the updateInstitutionStripeLock function.
//	Verification: None
// func processCreateStripeCustomer(
// 	myPlaid PlaidHelper,
// 	myFirebase chv.FirebaseFirestoreHelper,
// 	requestorId, institutionName string,
// ) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _      = runtime.Caller(0)
// 		tFunctionName           = runtime.FuncForPC(tFunction).Name()
// 		tInstitutionAccounts    []string
// 		tPlaidAccessToken       string
// 		tStripeBankAccountToken string
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
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

// Private functions go below here.

// processStripePayment - will handle a Stripe payment request
//
//	Errors: ErrStripeAmountInvalid, cpi.ErrStripeCurrencyInvalid, cpi.ErrStripeKeyInvalid
//	Customer Message: none
//	Verifications: none
func processPaymentIntent(
	amount float64,
	currency, description, key string,
) (
	paymentIntentResults *stripe.PaymentIntent,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tMatchCurrency = false
	)

	if amount <= 0 {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripeAmountInvalid, fmt.Sprintf("%v%v", rcv.TXT_AMOUNT, currency))
		return
	}
	if currency == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripeCurrencyInvalid, rcv.VAL_EMPTY)
		return
	}
	for _, tCurrency := range currencyList {
		if stripe.Currency(strings.ToLower(strings.Trim(currency, rcv.SPACES_ONE))) == tCurrency {
			tMatchCurrency = true
			break
		}
	}
	if tMatchCurrency == false {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripeCurrencyInvalid, fmt.Sprintf("%v%v", rcv.TXT_CURRENCY, currency))
		return
	}
	if key == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripeKeyInvalid, rcv.VAL_EMPTY)
		return
	}
	stripe.Key = key

	paymentParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(chv.FloatToPennies(amount)),
		Currency: stripe.String(currency),
		// Customer:    stripe.String(tCustomer.Id),
		Description: stripe.String(fmt.Sprintf("%v", description)),
	}
	paymentIntentResults, errorInfo.Error = paymentintent.New(paymentParams)

	return
}

// searchStripeCustomer - uses the name field to return the Stripe customer object, if found.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func searchStripeCustomer(stripeCustomerAccountId string) (stripeCustomer Customer) {
//
// 	var (
// 		errorInfo cpi.ErrorInfo
// 	)
//
// 	params := &stripe.CustomerSearchParams{}
// 	params.Query = *stripe.String(fmt.Sprintf("name:'%v'", stripeCustomerAccountId))
// 	iter := customer.Search(params)
// 	for iter.Next() {
// 		result := iter.Current()
// 		jsonResult, _ := json.Marshal(result)
// 		errorInfo.Error = json.Unmarshal(jsonResult, &stripeCustomer)
// 	}
//
// 	return
// }

// deleteStripeCustomer - uses the id field to delete the Stripe customer object.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func deleteStripeCustomer(stripeCustomerAccountId string) (stripeCustomer Customer) {
//
// 	var (
// 		errorInfo      cpi.ErrorInfo
// 		customerParams stripe.CustomerParams
// 	)
//
// 	params := &stripe.CustomerSearchParams{}
// 	params.Query = *stripe.String(fmt.Sprintf("name:'%v'", stripeCustomerAccountId))
// 	iter := customer.Search(params)
// 	for iter.Next() {
// 		result := iter.Current()
// 		jsonResult, _ := json.Marshal(result)
// 		errorInfo.Error = json.Unmarshal(jsonResult, &stripeCustomer)
// 		_, errorInfo.Error = customer.Del(stripeCustomer.Id, &customerParams)
// 	}
//
// 	return
// }
