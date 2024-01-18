// Package coreCustomerMessages
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc..
	Use is subject to license terms.

	Unauthorized copying of this file, via any medium is strictly prohibited.

	Proprietary and confidential

	Written by <Replace with FULL_NAME> / syacko
	STY-Holdings, Inc.
	support@sty-holdings.com
	www.sty-holdings.com

	01-2024
	USA

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package coreCustomerMessages

import (
	"fmt"
)

var (
	AccountBalanceInsufficient     = "In order to try an protect SavUp lenders, your available funds have to be higher than the requested transfer by at least $5."
	AddressMissing                 = "The address is missing information. Please review the street, city, state, and zip code."
	AreaCodePhoneNumberMissing     = "Either the area code or the phone number is missing."
	BankUnavailable                = "SavUp is unable to communication with you bank at this time. Please try again later."
	ConfirmationSuccessfulEmail    = "Your email has been confirmed."
	ConfirmationSuccessfulPhone    = "Your phone has been confirmed."
	EmailDelivery                  = "SavUp email will arrive usually within 5 minutes but could take up to 1 day. \n \nPlease check your inbox, spam, or junk folders for an email from support@sty-holdings.com."
	FederalTaxIdAlreadySet         = "The federal tax identifier has already been provided for this account. For security reasons, we do not display the value. If you need to reset your federal tax identifier, please contact support@sty-holdings.com. \n \nThank you for your understanding about the need for security."
	FirstLastNameMissing           = "Either the first or the last name is missing."
	NoBundles                      = "There are no bundles in the portfolio."
	NoLinkedBanks                  = "There are no linked bank accounts. Please link a bank."
	NoToDos                        = "There are no To Do's for you right now."
	NoUserBundles                  = "There are no bundles for you right now."
	NoUserRegister                 = "There are no history for you right now."
	PaymentSuccessful              = "The transfer to SavUp has been successfully requested. It will take up to 6 days for funds to arrive at SavUp. We will send update when the status changes."
	RefreshTooSoon                 = "SavUp limits balance refreshes to once a day."
	ShortURLMissing                = "Please email support@sty-holdings.com with the following subject line: Account is not confirmed and missing email."
	StripeTransferOut              = "Stripe only allows transfer to SavUp, please select a different transfer method to transfer money out."
	SystemExperiencedIssue         = "The SavUp system has experienced an issue. Support has been notified and is investigating. "
	UnauthorizedTooAccessSystem    = "SavUp has determined that you are not authorized to use the software. If you believe this to be a mistake, please contact support@sty-holdings.com."
	UnconfirmedAccount             = "SavUp has sent you an email to confirm your account. \n \nPlease click on the link in the email or click on \"Need Help\" and \"Resend Confirmation\"."
	UnsupportedTransferMethod      = "The transfer method is not supported."
	UserCreatedVerifyEmailNextStep = "Thank you for creating a SavUp account. \n \nTo complete the process, please look for an email from support@sty-holdings.com in your inbox, spam, or junk folders. \nClick on the link in the email and then you will be able to log into your account.\n \nWelcome to SavUp!"
	UserMissing                    = "We are unable to locate you in our system. Please contact support@sty-holding.com. I'm sure there is a simply way to resolve this issue."
	UserPasswordResetNextStep      = "Your request to reset your password is complete. Next, you will get an email from support@sty-holdings.com in your inbox, spam, or junk folders. The email will have a reset code you will use at login on the mobile app. Please do not share this code with anyone.\n \nHave a good day!"
	UserUpdatedSuccessful          = "The changes have been saved to your profile. Thank you for keep your information current!"
	UserRetryLimitHit              = "There is a limit to the number of retries and it has been reached. Please try again in 30 minutes or contact support@sty-holdings.com."
)

func MissingParameterName(parameterName string) string {
	return fmt.Sprintf("No %v was not provide and is needed.", parameterName)
}

func UsernameAlreadyExists(username string) string {
	return fmt.Sprintf("Username (%v) already exists in the SavUp system.", username)
}

func UsernameNotFound(username string) string {
	return fmt.Sprintf("Username (%v) is not in the SavUp system.", username)
}
