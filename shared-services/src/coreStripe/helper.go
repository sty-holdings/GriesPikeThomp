// Package coreStripe
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

	"github.com/stripe/stripe-go/v76"
	chv "github.com/sty-holdings/GriesPikeThomp/shared-services/src/coreHelpersValidators"
	cpi "github.com/sty-holdings/GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// checkSetAmount - will see if the amount is positive. If not, an error is returned, otherwise, the amount.
//
//	Customer Messages: None
//	Errors: cpi.ErrStripeAmountInvalid
//	Verifications: None
func checkSetAmount(amount float64) (
	stripeAmountPtr *int64,
	errorInfo cpi.ErrorInfo,
) {

	if amount > 0 {
		x := chv.FloatToPennies(amount)
		stripeAmountPtr = &x
		return
	}
	errorInfo = cpi.NewErrorInfo(cpi.ErrStripeAmountInvalid, fmt.Sprintf("%v%v", rcv.TXT_AMOUNT, amount))

	return
}

// checkSetCurrency - will see if the currency is valid. If not, then an error is returned, otherwise,
// the stripe currency is returned.
//
//	Customer Messages: None
//	Errors: cpi.ErrStripeCurrencyInvalid
//	Verifications: None
func checkSetCurrency(currency string) (
	stripeCurrencyPtr *string,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tCurrency stripe.Currency
	)

	if currency == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripeCurrencyInvalid, rcv.VAL_EMPTY)
		return
	}
	for _, tCurrency = range currencyValidValues {
		if stripe.Currency(strings.ToLower(strings.Trim(currency, rcv.SPACES_ONE))) == tCurrency {
			x := string(tCurrency)
			stripeCurrencyPtr = &x
			return
		}
	}

	errorInfo = cpi.NewErrorInfo(cpi.ErrStripeCurrencyInvalid, fmt.Sprintf("%v%v", rcv.TXT_CURRENCY, currency))

	return
}

// checkSetId - will see if the payment intent id is empty. If it is, then an error is return
// otherwise it returns the id.
//
//	Customer Messages: None
//	Errors: cpi.ErrStripeIdInvalid
//	Verifications: None
func checkSetId(paymentIntentId string) (
	stripePaymentIntentId string,
	errorInfo cpi.ErrorInfo,
) {

	if paymentIntentId == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripePaymentIntentIdEmpty, rcv.VAL_EMPTY)
		return
	}
	stripePaymentIntentId = paymentIntentId

	return
}

// checkSetKey - will see if the key is empty. If it is, then an error is return
// otherwise it returns the key.
//
//	Customer Messages: None
//	Errors: cpi.ErrStripeKeyInvalid
//	Verifications: None
func checkSetKey(key string) (
	stripeKey string,
	errorInfo cpi.ErrorInfo,
) {

	if key == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripeKeyInvalid, rcv.VAL_EMPTY)
		return
	}
	stripeKey = key

	return
}

// checkSetPaymentMethod - will see if the payment method is empty. If it is, then an error is return
// otherwise it returns the payment method as a string pointer.
//
//	Customer Messages: None
//	Errors: cpi.ErrStripePaymentMethodInvalid
//	Verifications: None
func checkSetPaymentMethod(paymentMethod string) (
	stripePaymentMethod *string,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tPaymentMethod string
	)

	if paymentMethod == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripePaymentMethodEmpty, rcv.VAL_EMPTY)
		return
	}
	for _, tPaymentMethod = range paymentMethodValidValues {
		if tPaymentMethod == strings.ToLower(strings.Trim(paymentMethod, rcv.SPACES_ONE)) {
			stripePaymentMethod = &paymentMethod
			return
		}
	}

	errorInfo = cpi.NewErrorInfo(cpi.ErrStripePaymentMethodInvalid, fmt.Sprintf("%v%v", rcv.TXT_PAYMENT_METHOD, paymentMethod))

	return
}

// checkSetPaymentMethodType - will see if the payment method type is empty. If it is, then an error is return
// otherwise it returns the payment method type as a string pointer.
//
//	Customer Messages: None
//	Errors: cpi.ErrStripePaymentMethodInvalid
//	Verifications: None
func checkSetPaymentMethodType(paymentMethodType string) (
	stripePaymentMethodType *string,
	errorInfo cpi.ErrorInfo,
) {

	var (
		tPaymentMethodType string
	)

	if paymentMethodType == rcv.VAL_EMPTY {
		errorInfo = cpi.NewErrorInfo(cpi.ErrStripePaymentMethodTypeEmpty, rcv.VAL_EMPTY)
		return
	}
	for _, tPaymentMethodType = range paymentMethodValidValues {
		if tPaymentMethodType == strings.ToLower(strings.Trim(paymentMethodType, rcv.SPACES_ONE)) {
			stripePaymentMethodType = &paymentMethodType
			return
		}
	}

	errorInfo = cpi.NewErrorInfo(cpi.ErrStripePaymentMethodTypeInvalid, fmt.Sprintf("%v%v", rcv.TXT_PAYMENT_METHOD_TYPE, paymentMethodType))

	return
}
