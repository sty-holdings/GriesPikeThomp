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

// ToDo Review where this belongs
type Helper struct {
	CredentialsLocation string
	Key                 string `json:"stripe_key"`
}

type Address struct {
	City       string `json:"city"`
	Country    string `json:"country"`
	Line1      string `json:"line1"`
	Line2      string `json:"line2"`
	PostalCode string `json:"postal_code"`
	State      string `json:"state"`
}
