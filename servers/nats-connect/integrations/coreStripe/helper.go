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
	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	hv "github.com/sty-holdings/sty-shared/v2024/helpersValidators"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
)

// checkSetAmount - will see if the amount is positive. If not, an error is returned, otherwise, the amount.
//
//	Customer Messages: None
//	Errors: pi.ErrStripeAmountInvalid
//	Verifications: None
func checkSetAmount(amount float64) (
	stripeAmountPtr *int64,
	errorInfo pi.ErrorInfo,
) {

	if amount > 0 {
		x := hv.FloatToPennies(amount)
		stripeAmountPtr = &x
		return
	}
	errorInfo = pi.NewErrorInfo(pi.ErrStripeAmountInvalid, fmt.Sprintf("%v%v", ctv.TXT_AMOUNT, amount))

	return
}

// checkSetCurrency - will see if the currency is valid. If not, then an error is returned, otherwise,
// the stripe currency is returned.
//
//	Customer Messages: None
//	Errors: pi.ErrStripeCurrencyInvalid
//	Verifications: None
func checkSetCurrency(currency string) (
	stripeCurrencyPtr *string,
	errorInfo pi.ErrorInfo,
) {

	var (
		tCurrency stripe.Currency
	)

	if currency == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrStripeCurrencyInvalid, ctv.VAL_EMPTY)
		return
	}
	for _, tCurrency = range currencyValidValues {
		if stripe.Currency(strings.ToLower(strings.Trim(currency, ctv.SPACES_ONE))) == tCurrency {
			x := string(tCurrency)
			stripeCurrencyPtr = &x
			return
		}
	}

	errorInfo = pi.NewErrorInfo(pi.ErrStripeCurrencyInvalid, fmt.Sprintf("%v%v", ctv.TXT_CURRENCY, currency))

	return
}

// checkSetId - will see if the payment intent id is empty. If it is, then an error is return
// otherwise it returns the id.
//
//	Customer Messages: None
//	Errors: pi.ErrStripeIdInvalid
//	Verifications: None
func checkSetId(paymentIntentId string) (
	stripePaymentIntentId string,
	errorInfo pi.ErrorInfo,
) {

	if paymentIntentId == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrStripePaymentIntentIdEmpty, ctv.VAL_EMPTY)
		return
	}
	stripePaymentIntentId = paymentIntentId

	return
}

// checkSetKey - will see if the key is empty. If it is, then an error is return
// otherwise it returns the key.
//
//	Customer Messages: None
//	Errors: pi.ErrStripeKeyInvalid
//	Verifications: None
func checkSetKey(key string) (
	stripeKey string,
	errorInfo pi.ErrorInfo,
) {

	if key == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrStripeKeyInvalid, ctv.VAL_EMPTY)
		return
	}
	stripeKey = key

	return
}

// checkSetPaymentMethod - will see if the payment method is empty. If it is, then an error is return
// otherwise it returns the payment method as a string pointer.
//
//	Customer Messages: None
//	Errors: pi.ErrStripePaymentMethodInvalid
//	Verifications: None
func checkSetPaymentMethod(paymentMethod string) (
	stripePaymentMethod *string,
	errorInfo pi.ErrorInfo,
) {

	var (
		tPaymentMethod string
	)

	if paymentMethod == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrStripePaymentMethodEmpty, ctv.VAL_EMPTY)
		return
	}
	for _, tPaymentMethod = range paymentMethodValidValues {
		if tPaymentMethod == strings.ToLower(strings.Trim(paymentMethod, ctv.SPACES_ONE)) {
			stripePaymentMethod = &paymentMethod
			return
		}
	}

	errorInfo = pi.NewErrorInfo(pi.ErrStripePaymentMethodInvalid, fmt.Sprintf("%v%v", ctv.TXT_PAYMENT_METHOD, paymentMethod))

	return
}

// checkSetPaymentMethodType - will see if the payment method type is empty. If it is, then an error is return
// otherwise it returns the payment method type as a string pointer.
//
//	Customer Messages: None
//	Errors: pi.ErrStripePaymentMethodInvalid
//	Verifications: None
func checkSetPaymentMethodType(paymentMethodType string) (
	stripePaymentMethodType *string,
	errorInfo pi.ErrorInfo,
) {

	var (
		tPaymentMethodType string
	)

	if paymentMethodType == ctv.VAL_EMPTY {
		errorInfo = pi.NewErrorInfo(pi.ErrStripePaymentMethodTypeEmpty, ctv.VAL_EMPTY)
		return
	}
	for _, tPaymentMethodType = range paymentMethodValidValues {
		if tPaymentMethodType == strings.ToLower(strings.Trim(paymentMethodType, ctv.SPACES_ONE)) {
			stripePaymentMethodType = &paymentMethodType
			return
		}
	}

	errorInfo = pi.NewErrorInfo(pi.ErrStripePaymentMethodTypeInvalid, fmt.Sprintf("%v%v", ctv.TXT_PAYMENT_METHOD_TYPE, paymentMethodType))

	return
}
