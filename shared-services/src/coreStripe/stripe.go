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
	"encoding/json"
	"fmt"
	"sync"

	ext "GriesPikeThomp/servers/nats-connect/loadExtensions"
	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cn "GriesPikeThomp/shared-services/src/coreNATS"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/nats-io/nats.go"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
	"github.com/stripe/stripe-go/v76/paymentmethodconfiguration"
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

	handlers[STRIPE_CANCEL_PAYMENT_INTENT] = cn.MessageHandler{
		Handler: stripeInstancePtr.CancelPaymentIntent(),
	}
	handlers[STRIPE_CONFIRM_PAYMENT_INTENT] = cn.MessageHandler{
		Handler: stripeInstancePtr.ConfirmPaymentIntent(),
	}
	handlers[STRIPE_LIST_PAYMENT_METHODS] = cn.MessageHandler{
		Handler: stripeInstancePtr.listPaymentMethods(),
	}
	handlers[STRIPE_LIST_PAYMENT_INTENTS] = cn.MessageHandler{
		Handler: stripeInstancePtr.listPaymentIntents(),
	}
	handlers[STRIPE_CREATE_PAYMENT_INTENT] = cn.MessageHandler{
		Handler: stripeInstancePtr.createPaymentIntent(),
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

func (stripeInstancePtr *stripeInstance) CancelPaymentIntent() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo                      cpi.ErrorInfo
			tCancelPaymentIntentResultsPtr *stripe.PaymentIntent
			tReply                         cn.NATSReply
			tRequest                       CancelPaymentIntentRequest
		)

		if errorInfo = cn.UnmarshalMessageData(cpi.GetFunctionInfo(1).Name, msg, &tRequest); errorInfo.Error == nil {
			tCancelPaymentIntentResultsPtr, errorInfo = processCancelPaymentIntent(tRequest)
		}

		if errorInfo.Error == nil {
			tReply.Response = *tCancelPaymentIntentResultsPtr
		} else {
			cpi.PrintErrorInfo(errorInfo)
			tReply.ErrorInfo = errorInfo
		}

		if errorInfo = cn.SendReply(tReply, msg); errorInfo.Error != nil {
			cpi.PrintErrorInfo(errorInfo)
		}

		return
	}
}

func (stripeInstancePtr *stripeInstance) ConfirmPaymentIntent() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo                       cpi.ErrorInfo
			tConfirmPaymentIntentResultsPtr *stripe.PaymentIntent
			tReply                          cn.NATSReply
			tRequest                        ConfirmPaymentIntentRequest
		)

		if errorInfo = cn.UnmarshalMessageData(cpi.GetFunctionInfo(1).Name, msg, &tRequest); errorInfo.Error == nil {
			tConfirmPaymentIntentResultsPtr, errorInfo = processConfirmPaymentIntent(tRequest)
		}

		if errorInfo.Error == nil {
			tReply.Response = *tConfirmPaymentIntentResultsPtr
		} else {
			cpi.PrintErrorInfo(errorInfo)
			tReply.ErrorInfo = errorInfo
		}

		if errorInfo = cn.SendReply(tReply, msg); errorInfo.Error != nil {
			cpi.PrintErrorInfo(errorInfo)
		}

		return
	}
}

func (stripeInstancePtr *stripeInstance) listPaymentMethods() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo                  cpi.ErrorInfo
			tListPaymentMethodsResults = make(map[string]PaymentMethodDetails)
			tReply                     cn.NATSReply
			tRequest                   ListPaymentMethodRequest
		)

		if errorInfo = cn.UnmarshalMessageData(cpi.GetFunctionInfo(1).Name, msg, &tRequest); errorInfo.Error == nil {
			tListPaymentMethodsResults, errorInfo = processListPaymentMethods(tRequest)
		}

		if errorInfo.Error == nil {
			tReply.Response = tListPaymentMethodsResults
		} else {
			cpi.PrintErrorInfo(errorInfo)
			tReply.ErrorInfo = errorInfo
		}

		if errorInfo = cn.SendReply(tReply, msg); errorInfo.Error != nil {
			cpi.PrintErrorInfo(errorInfo)
		}

		return
	}
}

func (stripeInstancePtr *stripeInstance) listPaymentIntents() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo                 cpi.ErrorInfo
			tListPaymentIntentResults []stripe.PaymentIntent
			tReply                    cn.NATSReply
			tRequest                  ListPaymentIntentRequest
		)

		if errorInfo = cn.UnmarshalMessageData(cpi.GetFunctionInfo(1).Name, msg, &tRequest); errorInfo.Error == nil {
			tListPaymentIntentResults, errorInfo = processListPaymentIntents(tRequest)
		}

		if errorInfo.Error == nil {
			tReply.Response = tListPaymentIntentResults
		} else {
			cpi.PrintErrorInfo(errorInfo)
			tReply.ErrorInfo = errorInfo
		}

		if errorInfo = cn.SendReply(tReply, msg); errorInfo.Error != nil {
			cpi.PrintErrorInfo(errorInfo)
		}

		return
	}
}

func (stripeInstancePtr *stripeInstance) createPaymentIntent() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo             cpi.ErrorInfo
			tPaymentIntentResults *stripe.PaymentIntent
			tReply                cn.NATSReply
			tRequest              PaymentIntentRequest
		)

		if errorInfo = cn.UnmarshalMessageData(cpi.GetFunctionInfo(1).Name, msg, &tRequest); errorInfo.Error == nil {
			tPaymentIntentResults, errorInfo = processCreatePaymentIntent(tRequest)
		}

		if errorInfo.Error == nil {
			tReply.Response = tPaymentIntentResults
		} else {
			cpi.PrintErrorInfo(errorInfo)
			tReply.ErrorInfo = errorInfo
		}

		if errorInfo = cn.SendReply(tReply, msg); errorInfo.Error != nil {
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

// processCancelPaymentIntent - will cancel a payment intent request
//
//	Customer Message: none
//	Errors: checkSetId returns, checkSetPaymentMethod returns
//	Verifications: none
func processCancelPaymentIntent(
	request CancelPaymentIntentRequest,
) (
	cancelPaymentIntentResultsPtr *stripe.PaymentIntent,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tPaymentIntentId string
	)

	if stripe.Key, errorInfo = checkSetKey(request.Key); errorInfo.Error != nil {
		return
	}

	tPaymentIntentConfirmParamsPtr := &stripe.PaymentIntentCancelParams{}

	if tPaymentIntentId, errorInfo = checkSetId(request.PaymentIntentId); errorInfo.Error != nil {
		return
	}

	cancelPaymentIntentResultsPtr, errorInfo.Error = paymentintent.Cancel(tPaymentIntentId, tPaymentIntentConfirmParamsPtr)

	return
}

// processConfirmPaymentIntent - will handle a confirm payment intent request
//
//	Customer Message: none
//	Errors: checkSetId returns, checkSetPaymentMethod returns
//	Verifications: none
func processConfirmPaymentIntent(
	request ConfirmPaymentIntentRequest,
) (
	paymentIntentResultsPtr *stripe.PaymentIntent,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tPaymentIntentId string
	)

	if stripe.Key, errorInfo = checkSetKey(request.Key); errorInfo.Error != nil {
		return
	}

	paymentParams := &stripe.PaymentIntentConfirmParams{}

	if tPaymentIntentId, errorInfo = checkSetId(request.PaymentIntentId); errorInfo.Error != nil {
		return
	}
	if paymentParams.PaymentMethod, errorInfo = checkSetPaymentMethod(request.PaymentMethod); errorInfo.Error != nil {
		return
	}

	if chv.IsPopulated(request.ReturnURL) {
		paymentParams.ReturnURL = &request.ReturnURL
	}

	paymentIntentResultsPtr, errorInfo.Error = paymentintent.Confirm(tPaymentIntentId, paymentParams)

	return
}

// processListPaymentMethods - will return the payment methods in the
// Stripe dashboard > Settings > Payments > Payment Methods. Only the first
// configuration is supported.
//
//	Customer Messages: None
//	Errors: checkSetKey returns, GetFieldsNames returns
//	Verifications: None
func processListPaymentMethods(
	request ListPaymentMethodRequest,
) (
	paymentMethodList map[string]PaymentMethodDetails,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tPaymentMethodConfiguration stripe.PaymentMethodConfiguration
		tPaymentMethodListPtr       *paymentmethodconfiguration.Iter
		tPaymentMethodList          = make(map[string]interface{})
		tPaymentMethodParamsPtr     = &stripe.PaymentMethodConfigurationListParams{}
		tJSONPaymentMethod          []byte
		tPaymentMethodDetails       PaymentMethodDetails
	)

	if stripe.Key, errorInfo = checkSetKey(request.Key); errorInfo.Error != nil {
		return
	}

	tPaymentMethodListPtr = paymentmethodconfiguration.List(tPaymentMethodParamsPtr)

	tPaymentMethodConfiguration = *tPaymentMethodListPtr.PaymentMethodConfigurationList().Data[0]
	if tPaymentMethodList, errorInfo = chv.GetFieldsNames(tPaymentMethodConfiguration); errorInfo.Error != nil {
		return
	}

	paymentMethodList = make(map[string]PaymentMethodDetails)
	for key, value := range tPaymentMethodList {
		tJSONPaymentMethod, errorInfo.Error = json.Marshal(&value)
		errorInfo.Error = json.Unmarshal(tJSONPaymentMethod, &tPaymentMethodDetails)
		paymentMethodList[key] = tPaymentMethodDetails
	}

	return
}

// processListPaymentIntents - will return the payment intents based on parameters provide.
//
//	Customer Messages: None
//	Errors: checkSetKey returns,
//	Verifications: None
func processListPaymentIntents(
	request ListPaymentIntentRequest,
) (
	paymentIntentList []stripe.PaymentIntent,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tPaymentIntentListParamsPtr *stripe.PaymentIntentListParams
		tPaymentIntentListPtr       *paymentintent.Iter
		tPaymentIntentList          []stripe.PaymentIntent
	)

	if stripe.Key, errorInfo = checkSetKey(request.Key); errorInfo.Error != nil {
		return
	}
	if chv.IsPopulated(request.CustomerId) {
		tPaymentIntentListParamsPtr.Customer = &request.CustomerId
	}
	if request.Limit > 0 {
		tPaymentIntentListParamsPtr.Limit = stripe.Int64(request.Limit)
	}
	if chv.IsPopulated(request.StartingAfter) {
		tPaymentIntentListParamsPtr.StartingAfter = &request.StartingAfter
	}

	tPaymentIntentListPtr = paymentintent.List(tPaymentIntentListParamsPtr)
	for _, tPaymentIntent := range tPaymentIntentListPtr.PaymentIntentList().Data {
		tPaymentIntentList = append(tPaymentIntentList, *tPaymentIntent)
	}

	return
}

// processCreatePaymentIntent - will handle a create payment intent request
//
//	Errors: ErrStripeAmountInvalid, cpi.ErrStripeCurrencyInvalid, cpi.ErrStripeKeyInvalid
//	Customer Message: none
//	Verifications: none
func processCreatePaymentIntent(
	request PaymentIntentRequest,
) (
	paymentIntentResultsPtr *stripe.PaymentIntent,
	errorInfo cpi.ErrorInfo,
) {

	var (
		paymentParams = &stripe.PaymentIntentParams{}
	)

	if paymentParams.Amount, errorInfo = checkSetAmount(request.Amount); errorInfo.Error != nil {
		return
	}
	if paymentParams.Currency, errorInfo = checkSetCurrency(request.Currency); errorInfo.Error != nil {
		return
	}
	if stripe.Key, errorInfo = checkSetKey(request.Key); errorInfo.Error != nil {
		return
	}

	if chv.IsPopulated(request.AutomaticPaymentMethods) && request.AutomaticPaymentMethods == true {
		paymentParams.AutomaticPaymentMethods = &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(request.AutomaticPaymentMethods),
		}
	}
	if chv.IsPopulated(request.Description) {
		paymentParams.Description = &request.Description
	}
	if chv.IsPopulated(request.ReceiptEmail) {
		paymentParams.ReceiptEmail = &request.ReceiptEmail
	}

	paymentIntentResultsPtr, errorInfo.Error = paymentintent.New(paymentParams)

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
