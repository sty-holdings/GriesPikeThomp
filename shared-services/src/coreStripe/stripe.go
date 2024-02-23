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
	"log"
	"sync"

	ext "GriesPikeThomp/servers/nats-connect/loadExtensions"
	chv "GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cns "GriesPikeThomp/shared-services/src/coreNATS"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/nats-io/nats.go"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
	// "github.com/stripe/stripe-go/v76"
	// "github.com/stripe/stripe-go/v76/charge"
	// "github.com/stripe/stripe-go/v76/customer"
)

// NewExtension - creates an instance by setting the values for the extension struct
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func NewExtension(
	config ext.ExtensionConfiguration,
	hostname string,
	workingDirectory string,
	testingOn bool,
) (
	errorInfo cpi.ErrorInfo,
) {

	var (
		stripeInstancePtr *stripeInstance
	)

	stripeInstancePtr = &stripeInstance{
		subscriptionPtrs: make(map[string]*nats.Subscription),
		waitGroup:        sync.WaitGroup{},
	}
	if stripeInstancePtr.instanceName, errorInfo = cns.BuildInstanceName(cns.METHOD_BLANK, rcv.STRIPE_EXTENSION); errorInfo.Error != nil {
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

// messageHandles - builds a map of messages handlers
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (stripeInstancePtr *stripeInstance) messageHandles() (
	handlers map[string]cns.MessageHandler,
) {

	handlers = make(map[string]cns.MessageHandler)

	handlers[PRINT_INSTANCE_NAME] = cns.MessageHandler{
		Handler: stripeInstancePtr.printInstanceName(),
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
		tHandlers        = make(map[string]cns.MessageHandler)
		tSubscriptionPtr *nats.Subscription
	)

	if stripeInstancePtr.natsConnectionPtr, errorInfo = cns.GetConnection(stripeInstancePtr.instanceName, config); errorInfo.Error != nil {
		return
	}

	tHandlers = stripeInstancePtr.messageHandles()

	// Use a WaitGroup to wait for a message to arrive
	stripeInstancePtr.waitGroup.Add(1)
	for _, loadSubject := range config.SubjectRegistry {
		if _, found := tHandlers[loadSubject.Subject]; found == false {
			errorInfo = cpi.NewErrorInfo(cpi.ErrSubjectInvalid, fmt.Sprintf("%v%v", rcv.TXT_SUBJECT, loadSubject.Subject))
			return
		}
		if tSubscriptionPtr, errorInfo.Error = stripeInstancePtr.natsConnectionPtr.Subscribe(
			loadSubject.Subject,
			tHandlers[loadSubject.Subject].Handler,
		); errorInfo.Error != nil {
			log.Printf("Subscribe failed on subject: %v", loadSubject.Subject)
			return
		}
		log.Printf("Subscribed to subject: %v", loadSubject.Subject)
		stripeInstancePtr.subscriptionPtrs[loadSubject.Subject] = tSubscriptionPtr
	}

	if stripeInstancePtr.testingOn {
		stripeInstancePtr.waitGroup.Done()
	} else {
		stripeInstancePtr.waitGroup.Wait()
	}

	return

}

// NATS Message Handlers go below this line.

// printInstanceName - sets serverPtr.instance.debugModeOn to false
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func (stripeInstancePtr *stripeInstance) printInstanceName() nats.MsgHandler {

	return func(msg *nats.Msg) {

		var (
			errorInfo cpi.ErrorInfo
			tReply    cns.NATSReply
		)

		tReply.Response = stripeInstancePtr.instanceName
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
// 	myFirebase coreHelpers.FirebaseFirestoreHelper,
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

// stripeGetKey - will read and parse the JSON key file. If either fail, exit is called.
//
//	Validations: File readable and JSON valid
// func getStripeKey(
// 	stripeFQN string,
// 	test bool,
// ) (stripeKey StripeHelper) {
//
// 	var (
// 		errorInfo          cpi.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tStripe            []byte
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if tStripe, errorInfo.Error = os.ReadFile(stripeFQN); errorInfo.Error != nil {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Required Stripe key file %v has issue.%v", stripeFQN, constants.ENDING_EXECUTION))
// 		log.Println(errorInfo.Error.Error())
// 	} else {
// 		if errorInfo.Error = json.Unmarshal(tStripe, &stripeKey); errorInfo.Error != nil {
// 			errorInfo.Error = errors.New(fmt.Sprintf("Stripe JSON file %v is corrupt.%v", stripeFQN, constants.ENDING_EXECUTION))
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
// func isStripeLockSet(
// 	firestoreClientPtr *firestore.Client,
// 	requestorId, institutionName string,
// ) bool {
//
// 	var (
// 		data      = make(map[string]interface{})
// 		errorInfo cpi.ErrorInfo
// 	)
//
// 	if data, errorInfo = coreFirestore.GetDocumentFromSubCollectionByDocumentId(firestoreClientPtr, constants.DATASTORE_USER_INSTITUTIONS, requestorId, constants.COLLECTION_INSTITUTIONS, institutionName); errorInfo.Error == nil {
// 		if _, exists := data[constants.FN_STRIPE_LOCK]; exists {
// 			return data[constants.FN_STRIPE_LOCK].(bool)
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
// func processCreateStripeCustomer(
// 	myPlaid PlaidHelper,
// 	myFirebase coreHelpers.FirebaseFirestoreHelper,
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

// processStripeCustomerTransfer - will handle a Stripe transfer into SavUp.
//
//	Errors: cpi.ErrTransferOutNotAllowed, cpi.ErrUserMissing
//	Customer Message: coreCustomerMessages.StripeTransferOut, coreCustomerMessages.UserMissing
//	Verifications: requestorId must exist, institutionName and transfer amount are required, transferAmount must be greater than zero.
// func processStripeCustomerTransfer(
// 	myFirebase coreHelpers.FirebaseFirestoreHelper,
// 	requestorId, institutionName, plaidAccountId string,
// 	transferAmount, reportedBalance float64,
// ) (reply CustomerTransferReply) {
//
// 	var (
// 		errorInfo     cpi.ErrorInfo
// 		tChargeResult *stripe.Charge
// 		tCustomer     Customer
// 	)
//
// 	if doesRequestorIdExist(myFirebase.FirestoreClientPtr, requestorId) {
// 		//  Stripe is a payment gateway, so only payments to SavUp are valid
// 		if transferAmount > 0 {
// 			if tCustomer = searchStripeCustomer(buildStripeCustomerAccountId(requestorId, institutionName, plaidAccountId)); tCustomer.Name == constants.EMPTY {
// 				reply.Error = cpi.ErrUserMissing.Error()
// 			} else {
// 				chargeParams := &stripe.ChargeParams{
// 					Amount:      stripe.Int64(coreHelpers.FloatToPennies(transferAmount)),
// 					Currency:    stripe.String(string(stripe.CurrencyUSD)),
// 					Customer:    stripe.String(tCustomer.Id),
// 					Description: stripe.String(fmt.Sprintf("ACH transfer to SavUp - powered by STY-Holdings")),
// 				}
// 				if tChargeResult, errorInfo.Error = charge.New(chargeParams); errorInfo.Error == nil {
// 					_tChargeResult, _ := json.Marshal(tChargeResult)
// 					if errorInfo = recordTransfer(myFirebase.FirestoreClientPtr, requestorId, institutionName, plaidAccountId, constants.TRANFER_STRIPE, string(_tChargeResult), constants.FN_PLAID_ACCOUNTS, constants.TRANSFER_IN, constants.EMPTY, transferAmount, reportedBalance); errorInfo.Error == nil {
// 						reply.Message = coreCustomerMessages.PaymentSuccessful
// 					} else {
// 						reply.Error = errorInfo.Error.Error()
// 					}
// 				}
// 			}
// 		} else {
// 			reply.Error = cpi.ErrTransferOutNotAllowed.Error()
// 			reply.Message = coreCustomerMessages.StripeTransferOut
// 		}
// 	} else {
// 		reply.Error = cpi.ErrUserMissing.Error()
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
