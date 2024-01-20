// Package shared_services
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc.
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
package shared_services

import (
	// Add imports here

	"errors"
)

//goland:noinspection ALL
const (
	ERROR    = "error"
	NO_ERROR = "no error"
	//
	//
	EXPECTED_ERROR_FORMAT     = "%v Failed: Was expecting an err."
	EXPECTING_NO_ERROR_FORMAT = "%v Failed: Wasn't expecting an err. ERROR: %v"
	UNEXPECTED_ERROR_FORMAT   = "%v Failed: Unexpected err. ERROR: %v"
	//
	// Messages
	ACCESS_TOKEN_MISSING           = "No access token was provided."
	ADDRESS_MISSING                = "The profile address is missing information. Please review the street, city, state, and zip code."
	ALREADY_CONFIRMED_EMAIL        = "The SavUp account has already need confirmed by email."
	ALREADY_CONFIRMED_PHONE        = "The SavUp account has already need confirmed by phone."
	AREA_CODE_PHONE_NUMBER_MISSING = "Either the area code or the phone number is missing."
	ATTEMPTS_EXCEEDED              = "LimitExceededException: Attempt limit exceeded, please try after some time."
	BUCKET_NOT_FOUND               = "The bucket was not found."
	BUFFER_EMPTY                   = "The buffer is empty"
	BUNDLE_ALREADY_EXISTS          = "Bundle already exists in the SavUp system."
	BUNDLE_MISSING                 = "Bundle is not in the SavUp system."
	COGNITO_USER_NAME_MISSING      = "Username is not in the Cognito user pool."
	CONFIG_FILE_MISSING            = "Not able to read the supplied config file. "
	DIRECTORY_MISSING              = "The directory does not exist."
	DIRECTORY_NOT_FULLY_QUALIFIED  = "The directory doesn't start and end with slash."
	DOCUMENT_NOT_FOUND             = "The document was not found."
	DOCUMENTS_NONE_FOUND           = "No documents were found."
	EVIRNOMENT_INVALID             = "The environment value is invalid."
	ERROR_MISSING                  = "ERROR MISSING"
	FALSE_SHOULD_BE_TRUE           = "The result should have been true."
	FILE_CREATION_FAILED           = "Create the file failed."
	FILE_MISSING                   = "The file doesn't exist."
	FILE_REMOVAL_FAILED            = "The file was not deleted."
	FILE_UNREADABLE                = "[ERROR} The file is not readable."
	FIREBASE_GC_PROJECT_ID_MISSING = "No Firebase project id was provided."
	FIRST_NAME_MISSING             = "The first name is empty."
	FIRST_LAST_NAME_MISSING        = "Either the first or the last name is missing."
	HTTP_SECURE_SERVER_FAILED      = "The savup-http secure server failed."
	JSON_GENERATION_FAILED         = "Failed to generate JSON payload"
	JSON_INVALID                   = "The JSON provided is invalid"
	LAST_NAME_MISSING              = "The last name is empty."
	MAP_IS_EMPTY                   = "Provided map is not populated."
	MAP_MISSING_KEY                = "Provided map has a nil or empty key."
	MAP_MISSING_VALUE              = "Provided map has a nil or empty value."
	MAX_THREADS_INVALID            = "The config file max threads value is invalid."
	MESSAGE_JSON_INVALID           = "The message body is not valid JSON."
	MESSAGE_PREFIX_INVALID         = "The Message Prefix value is invalid."
	MISSING_SERVER_NAME            = "The server name in main.go is empty."
	NATS_URL_INVALID               = "The NATS URL value is invalid."
	NATS_CONNECTION_FAILED         = "Connecting to NATS server failed"
	PID_FILE_EXISTS                = "A PID file already exists. Delete the 'server.pid' file in '.run' directory and start the server again."
	PLAID_INVALID_PUBLIC_TOKEN     = "INVALID_PUBLIC_TOKEN" // DO NOT change this, it is used to test a condition
	POINTER_MISSING                = "You must pass a pointer. Nil is not valid!"
	POSTGRES_SSL_MODE              = "Only disable, allow, prefer and required are supported."
	POSTGRES_CONN_FALIED           = "No database connection has been established"
	POSTGRES_CONN_EMPTY            = "Database connection is empty"
	REDIRECT_MODE_MISSING          = "The redirect mode is missing."
	REDIRECT_MODE_INVALID          = "The redirect mode is invalid."
	REFRESH_TOO_SOON               = "Too soon to refresh balances."
	REQUESTOR_ID_MISSING           = "The requestor id is missing."
	REQUIRED_ARGUMENT_MISSING      = "A required argument is empty."
	REQUIRED_FILE_MISSING          = "A required file is missing."
	RETRY_LIMIT_HIT                = "You have tried too many times. Please try again in 15 mins or contact support@sty-holdings.com."
	SERVER_CONFIGURATION_INVALID   = "The setting in the configuration file are inconsistant."
	SERVICE_FAILED_AWS             = "AWS service has failed. Investigate right away!"
	SERVICE_FAILED_FIREBASE        = "FIREBASE service has failed. Investigate right away!"
	SERVICE_FAILED_FIRESTORE       = "FIRESTORE service has failed. Investigate right away!"
	SERVICE_FAILED_PLAID           = "PLAID service has failed. Investigate right away!"
	SERVICE_FAILED_POSTGRES        = "POSTGRES service has failed. Investigate right away!"
	SERVICE_FAILED_SENDGRID        = "SENDGRID service has failed. Investigate right away!"
	SERVICE_FAILED_STRIPE          = "STRIPE service has failed. Investigate right away!"
	SHORT_URL_ALREADY_EXISTS       = "Short URL already exists in the SavUp system."
	SHORT_URL_MISSING              = "Short URL is not in the SavUp system."
	SIGNAL_UNKNOWN                 = "Unknown signal was caught and ignored."
	STRIPE_CUSTOMER_FAILED         = "Creating a Stripe customer failed."
	STRIPE_ONE_TIME_CODE_FAILED    = "Generating the Stripe One Time Use Token failed."
	STRIPE_OUT_NOT_SUPPORTED       = "Transfers out of SavUp using Stripe are not supported."
	TLS_FILES_MISSING              = "TLS files are missing."
	TOKEN_CLAIMS_INVALID           = "The token claims are invalid."
	TOKEN_EXPIRED                  = "The token has expired."
	TOKEN_INVALID                  = "The token is invalid."
	TRANSFER_AMOUNT_INVALID        = "The transfer amount is not support for this transfer method!"
	TRANSFER_IN_NOT_ALLOWED        = "Transferring money to SavUp is not allowed for this transfer method."
	TRANSFER_METHOD_INVALID        = "The transfer method is not support! (Transfer Method is case insensitive)"
	TRANSFER_OUT_NOT_ALLOWED       = "Transferring money from SavUp is not allowed for this transfer method."
	TRUE_SHOULD_BE_FALSE           = "The result should have been false."
	UNABLE_READ_FILE               = "Unable to read file."
	UNCONFIRMED_EMAIL              = "Users email has not been confirmed."
	UNEXPECTED_ERROR               = "The system has experienced an unexpected issue. Investigate right away!"
	UNSUPPORTED_TRANSFER_METHOD    = "The transfer method is not supported."
	USER_ACCOUNT_ALREADY_EXISTS    = "User account already exists in the SavUp system."
	USER_ACCOUNT_MISSING           = "User account is not in the SavUp system."
	USER_ALREADY_EXISTS            = "User already exists in the SavUp system."
	UNAUTHORIZED_REQUEST           = "You are not authorized to use this system."
	USER_BUNDLE_ALREADY_EXISTS     = "User bundle already exists in the SavUp system."
	USER_BUNDLE_MISSING            = "User bundle is not in the SavUp system."
	USER_MISSING                   = "User is not in the SavUp system."
	VERSION_INVALID                = "The software version is invalid. Use @env GOOS=linux GOARCH=amd64 go build -ldflags \"-X main.version=$(VERSION)\" -o ${ROOT_DIRECTORY}/servers/${SERVER_NAME}/bin/${SERVER_NAME} ${ROOT_DIRECTORY}/servers/${SERVER_NAME}/main.go"
	//
	// String that are used to determine third party error messages
	USER_DOES_NOT_EXIST = "User does not exist."
	NOT_FOUND           = "not found"
	UNKNOWN             = "UNKNOWN"
	//
	// Testing Strings
	TEST_STRING = "TEST STRING"
)

//goland:noinspection GoErrorStringFormat
var (
	ErrAccessTokenMissing         = errors.New(ACCESS_TOKEN_MISSING)
	ErrAddressMissing             = errors.New(ADDRESS_MISSING)
	ErrAlreadyConfirmedEmail      = errors.New(ALREADY_CONFIRMED_EMAIL)
	ErrAlreadyConfirmedPhone      = errors.New(ALREADY_CONFIRMED_PHONE)
	ErrAreaCodePhoneNumberMissing = errors.New(AREA_CODE_PHONE_NUMBER_MISSING)
	ErrAttemptsExceeded           = errors.New(ATTEMPTS_EXCEEDED)
	ErrBucketNotFound             = errors.New(BUCKET_NOT_FOUND)
	ErrBufferEmpty                = errors.New(BUFFER_EMPTY)
	ErrBundleAlreadyExists        = errors.New(BUNDLE_ALREADY_EXISTS)
	ErrBundleMissing              = errors.New(BUNDLE_MISSING)
	ErrCognitoUsernameMissing     = errors.New(COGNITO_USER_NAME_MISSING)
	ErrConfigFileMissing          = errors.New(CONFIG_FILE_MISSING)
	ErrDirectoryMissing           = errors.New(DIRECTORY_MISSING)
	ErrDirectoryNotFullyQualified = errors.New(DIRECTORY_NOT_FULLY_QUALIFIED)
	ErrDocumentNotFound           = errors.New(DOCUMENT_NOT_FOUND)
	ErrDocumentsNoneFound         = errors.New(DOCUMENTS_NONE_FOUND)
	ErrEnvironmentInvalid         = errors.New(EVIRNOMENT_INVALID)
	ErrErrorMissing               = errors.New(ERROR_MISSING)
	ErrFalseShouldBeTrue          = errors.New(FALSE_SHOULD_BE_TRUE)
	ErrFileCreationFailed         = errors.New(FILE_CREATION_FAILED)
	ErrFileMissing                = errors.New(FILE_MISSING)
	ErrFileRemovalFailed          = errors.New(FILE_REMOVAL_FAILED)
	ErrFileUnreadable             = errors.New(FILE_UNREADABLE)
	ErrFirstNameMissing           = errors.New(FIRST_NAME_MISSING)
	ErrFirstLastNameMissing       = errors.New(FIRST_LAST_NAME_MISSING)
	ErrFirebaseProjectMissing     = errors.New(FIREBASE_GC_PROJECT_ID_MISSING)
	ErrHTTPSecureServerFailed     = errors.New(HTTP_SECURE_SERVER_FAILED)
	ErrJSONGenerationFailed       = errors.New(JSON_GENERATION_FAILED)
	ErrJSONInvalid                = errors.New(JSON_INVALID)
	ErrLastNameMissing            = errors.New(LAST_NAME_MISSING)
	ErrMapIsEmpty                 = errors.New(MAP_IS_EMPTY)
	ErrMapIsMissingKey            = errors.New(MAP_MISSING_KEY)
	ErrMapIsMissingValue          = errors.New(MAP_MISSING_VALUE)
	ErrMaxThreadsInvalid          = errors.New(MAX_THREADS_INVALID)
	ErrMessageJSONInvalid         = errors.New(MESSAGE_JSON_INVALID)
	ErrMessagePrefixInvalid       = errors.New(MESSAGE_PREFIX_INVALID)
	ErrMissingServerName          = errors.New(MISSING_SERVER_NAME)
	ErrNATSURLInvalid             = errors.New(NATS_URL_INVALID)
	ErrNATSConnectionFailed       = errors.New(NATS_CONNECTION_FAILED)
	ErrPIDFileExists              = errors.New(PID_FILE_EXISTS)
	ErrPlaidInvalidPublicToken    = errors.New(PLAID_INVALID_PUBLIC_TOKEN)
	ErrPointerMissing             = errors.New(POINTER_MISSING)
	ErrPostgresSSLMode            = errors.New(POSTGRES_SSL_MODE)
	ErrPostgresConnFailed         = errors.New(POSTGRES_CONN_FALIED)
	ErrPostgresConnEmpty          = errors.New(POSTGRES_CONN_EMPTY)
	ErrRedirectModeMissing        = errors.New(REDIRECT_MODE_MISSING)
	ErrRedirectModeInvalid        = errors.New(REDIRECT_MODE_INVALID)
	ErrRefreshTooSoon             = errors.New(REFRESH_TOO_SOON)
	ErrRequestorIdMissing         = errors.New(REQUESTOR_ID_MISSING)
	ErrRequiredArgumentMissing    = errors.New(REQUIRED_ARGUMENT_MISSING)
	ErrRequiredFileMissing        = errors.New(REQUIRED_FILE_MISSING)
	ErrRetryLimitHit              = errors.New(RETRY_LIMIT_HIT)
	ErrServerConfigurationInvalid = errors.New(SERVER_CONFIGURATION_INVALID)
	ErrServiceFailedAWS           = errors.New(SERVICE_FAILED_AWS)
	ErrServiceFailedFIREBASE      = errors.New(SERVICE_FAILED_FIREBASE)
	ErrServiceFailedFIRESTORE     = errors.New(SERVICE_FAILED_FIRESTORE)
	ErrServiceFailedPLAID         = errors.New(SERVICE_FAILED_PLAID)
	ErrServiceFailedPOSTGRES      = errors.New(SERVICE_FAILED_POSTGRES)
	ErrServiceFailedSendGrid      = errors.New(SERVICE_FAILED_SENDGRID)
	ErrServiceFailedSTRIPE        = errors.New(SERVICE_FAILED_STRIPE)
	ErrShortURLMissing            = errors.New(SHORT_URL_MISSING)
	ErrSignalUnknown              = errors.New(SIGNAL_UNKNOWN)
	ErrStripeCreateCustomerFailed = errors.New(STRIPE_CUSTOMER_FAILED)
	ErrStripeOneTimeCodeFailed    = errors.New(STRIPE_ONE_TIME_CODE_FAILED)
	ErrStripeOutNotSupported      = errors.New(STRIPE_OUT_NOT_SUPPORTED)
	ErrTLSFilesMissing            = errors.New(TLS_FILES_MISSING)
	ErrTokenClaimsInvalid         = errors.New(TOKEN_CLAIMS_INVALID)
	ErrTokenExpired               = errors.New(TOKEN_EXPIRED)
	ErrTokenInvalid               = errors.New(TOKEN_INVALID)
	ErrTransferAmountInvalid      = errors.New(TRANSFER_AMOUNT_INVALID)
	ErrTransferInNotAllowed       = errors.New(TRANSFER_IN_NOT_ALLOWED)
	ErrTransferMethodInvalid      = errors.New(TRANSFER_METHOD_INVALID)
	ErrTransferOutNotAllowed      = errors.New(TRANSFER_OUT_NOT_ALLOWED)
	ErrTrueShouldBeFalse          = errors.New(TRUE_SHOULD_BE_FALSE)
	ErrUnableReadFile             = errors.New(UNABLE_READ_FILE)
	ErrUnauthorizedRequest        = errors.New(UNAUTHORIZED_REQUEST)
	ErrUnconfirmedEmail           = errors.New(UNCONFIRMED_EMAIL)
	ErrUnexpectedError            = errors.New(UNEXPECTED_ERROR)
	ErrUnsupportedTransferMethod  = errors.New(UNSUPPORTED_TRANSFER_METHOD)
	ErrUserAccountAlreadyExists   = errors.New(USER_ACCOUNT_ALREADY_EXISTS)
	ErrUserAccountMissing         = errors.New(USER_ACCOUNT_MISSING)
	ErrUserAlreadyExists          = errors.New(USER_ALREADY_EXISTS)
	ErrUserBundleAlreadyExists    = errors.New(BUNDLE_ALREADY_EXISTS)
	ErrUserBundleMissing          = errors.New(BUNDLE_MISSING)
	ErrUserMissing                = errors.New(USER_MISSING)
	ErrVersionInvalid             = errors.New(VERSION_INVALID)
)
