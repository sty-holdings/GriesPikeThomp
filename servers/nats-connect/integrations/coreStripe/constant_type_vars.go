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
	"sync"

	"github.com/nats-io/nats.go"
	sc "github.com/stripe/stripe-go/v76"
	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
)

//goland:noinspection GoSnakeCaseUsage,GoCommentStart
const (
	// Subjects
	STRIPE_CANCEL_PAYMENT_INTENT  = "stripe.cancel-payment-intent"
	STRIPE_CONFIRM_PAYMENT_INTENT = "stripe.confirm-payment-intent"
	STRIPE_LIST_PAYMENT_METHODS   = "stripe.list-payment-methods"
	STRIPE_LIST_PAYMENT_INTENTS   = "stripe.list-payment-intents"
	STRIPE_CREATE_PAYMENT_INTENT  = "stripe.create-payment-intent"
)

type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}

type Customer struct {
	Address              Address     `json:"address"`
	Balance              int         `json:"balance"`
	CashBalance          interface{} `json:"cash_balance"`
	Created              int         `json:"created"`
	Currency             string      `json:"currency"`
	DefaultSource        string      `json:"default_source"`
	Deleted              bool        `json:"deleted"`
	Delinquent           bool        `json:"delinquent"`
	Description          string      `json:"description"`
	Discount             interface{} `json:"discount"`
	Email                string      `json:"email"`
	Id                   string      `json:"id"`
	InvoiceCreditBalance interface{} `json:"invoice_credit_balance"`
	InvoicePrefix        string      `json:"invoice_prefix"`
	InvoiceSettings      struct {
		CustomFields         interface{} `json:"custom_fields"`
		DefaultPaymentMethod interface{} `json:"default_payment_method"`
		Footer               string      `json:"footer"`
		RenderingOptions     interface{} `json:"rendering_options"`
	} `json:"invoice_settings"`
	Livemode bool `json:"livemode"`
	Metadata struct {
	} `json:"metadata"`
	Name                string        `json:"name"`
	NextInvoiceSequence int           `json:"next_invoice_sequence"`
	Object              string        `json:"object"`
	Phone               string        `json:"phone"`
	PreferredLocales    []interface{} `json:"preferred_locales"`
	Shipping            interface{}   `json:"shipping"`
	Sources             interface{}   `json:"sources"`
	Subscriptions       interface{}   `json:"subscriptions"`
	Tax                 interface{}   `json:"tax"`
	TaxExempt           string        `json:"tax_exempt"`
	TaxIds              interface{}   `json:"tax_ids"`
	TestClock           interface{}   `json:"test_clock"`
}

type PaymentMethodDetails struct {
	Available         bool `json:"available"`
	DisplayPreference struct {
		Overridable bool   `json:"overridable"`
		Preference  string `json:"preference"`
		Value       string `json:"value"`
	} `json:"display_preference"`
}

// NATS Request go below here

type ListPaymentMethodRequest struct {
	Key string `json:"key"`
}

type ListPaymentIntentRequest struct {
	Key           string `json:"key"`
	CustomerId    string `json:"customer_id,omitempty"`
	Limit         int64  `json:"limit,omitempty"`
	StartingAfter string `json:"starting_after,omitempty"`
}

type PaymentIntentRequest struct {
	Amount                  float64 `json:"amount"`
	AutomaticPaymentMethods bool    `json:"automatic_payment_methods,omitempty"`
	// Confirm            bool     `json:"confirm,omitempty"`
	Currency    string `json:"currency"`
	Description string `json:"description,omitempty"`
	Key         string `json:"key"`
	// PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
	ReceiptEmail string `json:"receipt_email"`
	ReturnURL    string `json:"return_url,omitempty"`
}

type CancelPaymentIntentRequest struct {
	Key                string `json:"key"`
	PaymentIntentId    string `json:"id"`
	CancellationReason string `json:"cancellation_reason"`
}

type ConfirmPaymentIntentRequest struct {
	CaptureMethod   string `json:"capture_method,omitempty"`
	Key             string `json:"key"`
	PaymentIntentId string `json:"id"`
	PaymentMethod   string `json:"payment_method,omitempty"`
	ReceiptEmail    string `json:"receipt_email,omitempty"`
	ReturnURL       string `json:"return_url,omitempty,omitempty"`
}

type PaymentIntentInfo struct {
	Amount           int `json:"amount"`
	AmountCapturable int `json:"amount_capturable"`
	AmountDetails    struct {
		Tip struct {
			Amount int `json:"amount"`
		} `json:"tip"`
	} `json:"amount_details"`
	AmountReceived          int         `json:"amount_received"`
	Application             interface{} `json:"application"`
	ApplicationFeeAmount    int         `json:"application_fee_amount"`
	AutomaticPaymentMethods struct {
		AllowRedirects string `json:"allow_redirects"`
		Enabled        bool   `json:"enabled"`
	} `json:"automatic_payment_methods"`
	CanceledAt         int         `json:"canceled_at"`
	CancellationReason string      `json:"cancellation_reason"`
	CaptureMethod      string      `json:"capture_method"`
	ClientSecret       string      `json:"client_secret"`
	ConfirmationMethod string      `json:"confirmation_method"`
	Created            int         `json:"created"`
	Currency           string      `json:"currency"`
	Customer           interface{} `json:"customer"`
	Description        string      `json:"description"`
	Id                 string      `json:"id"`
	Invoice            interface{} `json:"invoice"`
	LastPaymentError   interface{} `json:"last_payment_error"`
	LatestCharge       interface{} `json:"latest_charge"`
	Livemode           bool        `json:"livemode"`
	Metadata           struct {
	} `json:"metadata"`
	NextAction                        interface{} `json:"next_action"`
	Object                            string      `json:"object"`
	OnBehalfOf                        interface{} `json:"on_behalf_of"`
	PaymentMethod                     interface{} `json:"payment_method"`
	PaymentMethodConfigurationDetails struct {
		Id     string `json:"id"`
		Parent string `json:"parent"`
	} `json:"payment_method_configuration_details"`
	PaymentMethodOptions struct {
		AcssDebit        interface{} `json:"acss_debit"`
		Affirm           interface{} `json:"affirm"`
		AfterpayClearpay interface{} `json:"afterpay_clearpay"`
		Alipay           interface{} `json:"alipay"`
		AuBecsDebit      interface{} `json:"au_becs_debit"`
		BacsDebit        interface{} `json:"bacs_debit"`
		Bancontact       interface{} `json:"bancontact"`
		Blik             interface{} `json:"blik"`
		Boleto           interface{} `json:"boleto"`
		Card             struct {
			CaptureMethod                   string      `json:"capture_method"`
			Installments                    interface{} `json:"installments"`
			MandateOptions                  interface{} `json:"mandate_options"`
			Network                         string      `json:"network"`
			RequestExtendedAuthorization    string      `json:"request_extended_authorization"`
			RequestIncrementalAuthorization string      `json:"request_incremental_authorization"`
			RequestMulticapture             string      `json:"request_multicapture"`
			RequestOvercapture              string      `json:"request_overcapture"`
			RequestThreeDSecure             string      `json:"request_three_d_secure"`
			RequireCvcRecollection          bool        `json:"require_cvc_recollection"`
			SetupFutureUsage                string      `json:"setup_future_usage"`
			StatementDescriptorSuffixKana   string      `json:"statement_descriptor_suffix_kana"`
			StatementDescriptorSuffixKanji  string      `json:"statement_descriptor_suffix_kanji"`
		} `json:"card"`
		CardPresent     interface{} `json:"card_present"`
		Cashapp         interface{} `json:"cashapp"`
		CustomerBalance interface{} `json:"customer_balance"`
		Eps             interface{} `json:"eps"`
		Fpx             interface{} `json:"fpx"`
		Giropay         interface{} `json:"giropay"`
		Grabpay         interface{} `json:"grabpay"`
		Ideal           interface{} `json:"ideal"`
		InteracPresent  interface{} `json:"interac_present"`
		Klarna          interface{} `json:"klarna"`
		Konbini         interface{} `json:"konbini"`
		Link            interface{} `json:"link"`
		Oxxo            interface{} `json:"oxxo"`
		P24             interface{} `json:"p24"`
		Paynow          interface{} `json:"paynow"`
		Paypal          interface{} `json:"paypal"`
		Pix             interface{} `json:"pix"`
		Promptpay       interface{} `json:"promptpay"`
		RevolutPay      interface{} `json:"revolut_pay"`
		SepaDebit       interface{} `json:"sepa_debit"`
		Sofort          interface{} `json:"sofort"`
		Swish           interface{} `json:"swish"`
		UsBankAccount   struct {
			FinancialConnections struct {
				Permissions []string      `json:"permissions"`
				Prefetch    []interface{} `json:"prefetch"`
				ReturnUrl   string        `json:"return_url"`
			} `json:"financial_connections"`
			MandateOptions struct {
				CollectionMethod string `json:"collection_method"`
			} `json:"mandate_options"`
			PreferredSettlementSpeed string `json:"preferred_settlement_speed"`
			SetupFutureUsage         string `json:"setup_future_usage"`
			VerificationMethod       string `json:"verification_method"`
		} `json:"us_bank_account"`
		WechatPay interface{} `json:"wechat_pay"`
		Zip       interface{} `json:"zip"`
	} `json:"payment_method_options"`
	PaymentMethodTypes        []string    `json:"payment_method_types"`
	Processing                interface{} `json:"processing"`
	ReceiptEmail              string      `json:"receipt_email"`
	Review                    interface{} `json:"review"`
	SetupFutureUsage          string      `json:"setup_future_usage"`
	Shipping                  interface{} `json:"shipping"`
	Source                    interface{} `json:"source"`
	StatementDescriptor       string      `json:"statement_descriptor"`
	StatementDescriptorSuffix string      `json:"statement_descriptor_suffix"`
	Status                    string      `json:"status"`
	TransferData              interface{} `json:"transfer_data"`
	TransferGroup             string      `json:"transfer_group"`
}

type stripeInstance struct {
	instanceName      string
	natsConnectionPtr *nats.Conn
	processChannel    chan string
	subscriptionPtrs  map[string]*nats.Subscription
	testingOn         bool
	waitGroup         sync.WaitGroup
}

var (
	currencyValidValues = []sc.Currency{
		sc.CurrencyAED,
		sc.CurrencyAFN,
		sc.CurrencyALL,
		sc.CurrencyAMD,
		sc.CurrencyANG,
		sc.CurrencyAOA,
		sc.CurrencyARS,
		sc.CurrencyAUD,
		sc.CurrencyAWG,
		sc.CurrencyAZN,
		sc.CurrencyBAM,
		sc.CurrencyBBD,
		sc.CurrencyBDT,
		sc.CurrencyBGN,
		sc.CurrencyBIF,
		sc.CurrencyBMD,
		sc.CurrencyBND,
		sc.CurrencyBOB,
		sc.CurrencyBRL,
		sc.CurrencyBSD,
		sc.CurrencyBWP,
		sc.CurrencyBZD,
		sc.CurrencyCAD,
		sc.CurrencyCDF,
		sc.CurrencyCHF,
		sc.CurrencyCLP,
		sc.CurrencyCNY,
		sc.CurrencyCOP,
		sc.CurrencyCRC,
		sc.CurrencyCVE,
		sc.CurrencyCZK,
		sc.CurrencyDJF,
		sc.CurrencyDKK,
		sc.CurrencyDOP,
		sc.CurrencyDZD,
		sc.CurrencyEEK,
		sc.CurrencyEGP,
		sc.CurrencyETB,
		sc.CurrencyEUR,
		sc.CurrencyFJD,
		sc.CurrencyFKP,
		sc.CurrencyGBP,
		sc.CurrencyGEL,
		sc.CurrencyGIP,
		sc.CurrencyGMD,
		sc.CurrencyGNF,
		sc.CurrencyGTQ,
		sc.CurrencyGYD,
		sc.CurrencyHKD,
		sc.CurrencyHNL,
		sc.CurrencyHRK,
		sc.CurrencyHTG,
		sc.CurrencyHUF,
		sc.CurrencyIDR,
		sc.CurrencyILS,
		sc.CurrencyINR,
		sc.CurrencyISK,
		sc.CurrencyJMD,
		sc.CurrencyJPY,
		sc.CurrencyKES,
		sc.CurrencyKGS,
		sc.CurrencyKHR,
		sc.CurrencyKMF,
		sc.CurrencyKRW,
		sc.CurrencyKYD,
		sc.CurrencyKZT,
		sc.CurrencyLAK,
		sc.CurrencyLBP,
		sc.CurrencyLKR,
		sc.CurrencyLRD,
		sc.CurrencyLSL,
		sc.CurrencyLTL,
		sc.CurrencyLVL,
		sc.CurrencyMAD,
		sc.CurrencyMDL,
		sc.CurrencyMGA,
		sc.CurrencyMKD,
		sc.CurrencyMNT,
		sc.CurrencyMOP,
		sc.CurrencyMRO,
		sc.CurrencyMUR,
		sc.CurrencyMVR,
		sc.CurrencyMWK,
		sc.CurrencyMXN,
		sc.CurrencyMYR,
		sc.CurrencyMZN,
		sc.CurrencyNAD,
		sc.CurrencyNGN,
		sc.CurrencyNIO,
		sc.CurrencyNOK,
		sc.CurrencyNPR,
		sc.CurrencyNZD,
		sc.CurrencyPAB,
		sc.CurrencyPEN,
		sc.CurrencyPGK,
		sc.CurrencyPHP,
		sc.CurrencyPKR,
		sc.CurrencyPLN,
		sc.CurrencyPYG,
		sc.CurrencyQAR,
		sc.CurrencyRON,
		sc.CurrencyRSD,
		sc.CurrencyRUB,
		sc.CurrencyRWF,
		sc.CurrencySAR,
		sc.CurrencySBD,
		sc.CurrencySCR,
		sc.CurrencySEK,
		sc.CurrencySGD,
		sc.CurrencySHP,
		sc.CurrencySLL,
		sc.CurrencySOS,
		sc.CurrencySRD,
		sc.CurrencySTD,
		sc.CurrencySVC,
		sc.CurrencySZL,
		sc.CurrencyTHB,
		sc.CurrencyTJS,
		sc.CurrencyTOP,
		sc.CurrencyTRY,
		sc.CurrencyTTD,
		sc.CurrencyTWD,
		sc.CurrencyTZS,
		sc.CurrencyUAH,
		sc.CurrencyUGX,
		sc.CurrencyUSD,
		sc.CurrencyUYU,
		sc.CurrencyUZS,
		sc.CurrencyVEF,
		sc.CurrencyVND,
		sc.CurrencyVUV,
		sc.CurrencyWST,
		sc.CurrencyXAF,
		sc.CurrencyXCD,
		sc.CurrencyXOF,
		sc.CurrencyXPF,
		sc.CurrencyYER,
		sc.CurrencyZAR,
		sc.CurrencyZMW,
	}
	paymentMethodValidValues = []string{
		ctv.CARD_BRAND_VISA,
		ctv.CARD_BRAND_VISA_DEBIT,
		ctv.CARD_BRAND_MASTERCARD,
		ctv.CARD_BRAND_MASTERCARD_DEBIT,
		ctv.CARD_BRAND_MASTERCARD_PREPAID,
		ctv.CARD_BRAND_AMEX,
		ctv.CARD_BRAND_DISCOVER,
		ctv.CARD_BRAND_DINERS_CLUB,
		ctv.CARD_BRAND_JCB,
		ctv.CARD_BRAND_UNION_PAY,
		ctv.CARD_BRAND_CARTES_BANCAIRES_VISA,
		ctv.CARD_BRAND_CARTES_BANCAIRES_MASTERCARD,
		ctv.CARD_BRAND_EFTPOS_AUSTRALIA_VISA,
		ctv.CARD_BRAND_EFTPOS_AUSTRALIA_MASTERCARD,
		ctv.CARD_USA_VISA,
		ctv.CARD_ARGENTINA_VISA,
		ctv.CARD_BRAZIL_VISA,
		ctv.CARD_CANADA_VISA,
		ctv.CARD_MEXICO_VISA,
		ctv.CARD_UNITED_ARAB_EMIRATES_VISA,
		ctv.CARD_UNITED_ARAB_EMIRATES_MASTERCARD,
		ctv.CARD_AUSTRIA_VISA,
		ctv.CARD_BELGIUM_VISA,
		ctv.CARD_BULGARIA_VISA,
		ctv.CARD_BELARUS_VISA,
		ctv.CARD_CROATIA_VISA,
		ctv.CARD_CYPRUS_VISA,
		ctv.CARD_CZECH_REPUBLIC_VISA,
		ctv.CARD_DENMARK_VISA,
		ctv.CARD_ESTONIA_VISA,
		ctv.CARD_FINLAND_VISA,
		ctv.CARD_FRANCE_VISA,
		ctv.CARD_GERMANY_VISA,
		ctv.CARD_GIBRALTAR_VISA,
		ctv.CARD_GREECE_VISA,
		ctv.CARD_HUNGARY_VISA,
		ctv.CARD_IRELAND_VISA,
		ctv.CARD_ITALY_VISA,
		ctv.CARD_LATVIA_VISA,
		ctv.CARD_LIECHTENSTEIN_VISA,
		ctv.CARD_LITHUANIA_VISA,
		ctv.CARD_LUXEMBOURG_VISA,
		ctv.CARD_MALTA_VISA,
		ctv.CARD_NETHERLANDS_VISA,
		ctv.CARD_NORWAY_VISA,
		ctv.CARD_POLAND_VISA,
		ctv.CARD_PORTUGAL_VISA,
		ctv.CARD_ROMANIA_VISA,
		ctv.CARD_SLOVENIA_VISA,
		ctv.CARD_SLOVAKIA_VISA,
		ctv.CARD_SPAIN_VISA,
		ctv.CARD_SWEDEN_VISA,
		ctv.CARD_SWITZERLAND_VISA,
		ctv.CARD_UNITED_KINGDOM_VISA,
		ctv.CARD_UNITED_KINGDOM_DEBIT,
		ctv.CARD_UNITED_KINGDOM_MASTERCARD,
		ctv.CARD_AUSTRALIA_VISA,
		ctv.CARD_CHINA_VISA,
		ctv.CARD_HONG_KONG_VISA,
		ctv.CARD_INDIA_VISA,
		ctv.CARD_JAPAN_VISA,
		ctv.CARD_JAPAN_JCB,
		ctv.CARD_MALAYSIA_VISA,
		ctv.CARD_NEW_ZEALAND_VISA,
		ctv.CARD_SINGAPORE_VISA,
		ctv.CARD_THAILAND_VISA_CREDIT,
		ctv.CARD_THAILAND_DEBIT_VISA_DEBIT,
	}
	paymentMethodTypeValidValues = []string{
		ctv.PAYMENT_METHOD_TYPE_ACSSDEBIT,
		ctv.PAYMENT_METHOD_TYPE_AFFIRM,
		ctv.PAYMENT_METHOD_TYPE_AFTERPAYCLEARPAY,
		ctv.PAYMENT_METHOD_TYPE_ALIPAY,
		ctv.PAYMENT_METHOD_TYPE_AUBECSDEBIT,
		ctv.PAYMENT_METHOD_TYPE_BACSDEBIT,
		ctv.PAYMENT_METHOD_TYPE_BANCONTACT,
		ctv.PAYMENT_METHOD_TYPE_BLIK,
		ctv.PAYMENT_METHOD_TYPE_BOLETO,
		ctv.PAYMENT_METHOD_TYPE_CARD,
		ctv.PAYMENT_METHOD_TYPE_CARD_PRESENT,
		ctv.PAYMENT_METHOD_TYPE_CASHAPP,
		ctv.PAYMENT_METHOD_TYPE_CUSTOMER_BALANCE,
		ctv.PAYMENT_METHOD_TYPE_EPS,
		ctv.PAYMENT_METHOD_TYPE_FPX,
		ctv.PAYMENT_METHOD_TYPE_GIROPAY,
		ctv.PAYMENT_METHOD_TYPE_GRABPAY,
		ctv.PAYMENT_METHOD_TYPE_IDEAL,
		ctv.PAYMENT_METHOD_TYPE_INTERAC_PRESENT,
		ctv.PAYMENT_METHOD_TYPE_KLARNA,
		ctv.PAYMENT_METHOD_TYPE_KONBINI,
		ctv.PAYMENT_METHOD_TYPE_LINK,
		ctv.PAYMENT_METHOD_TYPE_OXXO,
		ctv.PAYMENT_METHOD_TYPE_P24,
		ctv.PAYMENT_METHOD_TYPE_PAYNOW,
		ctv.PAYMENT_METHOD_TYPE_PAYPAL,
		ctv.PAYMENT_METHOD_TYPE_PIX,
		ctv.PAYMENT_METHOD_TYPE_PROMPTPAY,
		ctv.PAYMENT_METHOD_TYPE_REVOLUTPAY,
		ctv.PAYMENT_METHOD_TYPE_SEPADEBIT,
		ctv.PAYMENT_METHOD_TYPE_SOFORT,
		ctv.PAYMENT_METHOD_TYPE_SWISH,
		ctv.PAYMENT_METHOD_TYPE_USBANKACCOUNT,
		ctv.PAYMENT_METHOD_TYPE_WECHATPAY,
		ctv.PAYMENT_METHOD_TYPE_ZIP,
	}
)
