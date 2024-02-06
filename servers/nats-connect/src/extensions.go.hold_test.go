// Package coreOptions
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, Inc
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
package src

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

var (
// Global variables here
)

func TestGenerateConfigFileSkeleton(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		tReader, tWriter, _ := os.Pipe()
		old := os.Stdout
		os.Stdout = tWriter

		GenerateConfigFileSkeleton("NATS Connect Test",
			"GriesPikeThomp/shared-services/src/coreConfiguration",
			"/skeleton-config-file.")

		tBuffer := make([]byte, 3072)
		_, _ = tReader.Read(tBuffer)

		_ = tWriter.Close()
		os.Stdout = old

		if len(tBuffer) == 0 && errorInfo.Error != nil {
			tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, cpi.ErrBufferEmpty)
		}
	})
}

func TestReadAndParseConfigFile(tPtr *testing.T) {

	type arguments struct {
		configFileName string
	}

	var (
		gotError           bool
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Valid config file",
			arguments: arguments{
				configFileName: "skeleton-config-file.json",
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Invalid config file",
			arguments: arguments{
				configFileName: "invalid-skeleton-config-file.json",
			},
			wantError: true,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Unreadable config file",
			arguments: arguments{
				configFileName: "unreadable-skeleton-config-file.json",
			},
			wantError: true,
		},
	}

	fmt.Println(os.Getwd())

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _, errorInfo = ReadAndParseConfigFile(ts.arguments.configFileName); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Errorf(cpi.UNEXPECTED_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
		})
	}
}

func TestValidateConfiguration(tPtr *testing.T) {

	type arguments struct {
		config Configuration
	}

	var (
		gotError           bool
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "All missing",
			arguments: arguments{
				config: Configuration{
					Environment:  "",
					LogDirectory: "",
					MaxThreads:   0,
					PIDDirectory: "",
				},
			},
			wantError: true,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Valid settings Local Environment",
			arguments: arguments{
				config: Configuration{
					Environment:  rcv.ENVIRONMENT_LOCAL,
					LogDirectory: DEFAULT_LOG_DIRECTORY,
					MaxThreads:   DEFAULT_MAX_THREADS,
					PIDDirectory: DEFAULT_PID_DIRECTORY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Valid settings Development Environment",
			arguments: arguments{
				config: Configuration{
					Environment:  rcv.ENVIRONMENT_DEVELOPMENT,
					LogDirectory: DEFAULT_LOG_DIRECTORY,
					MaxThreads:   DEFAULT_MAX_THREADS,
					PIDDirectory: DEFAULT_PID_DIRECTORY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Valid settings Production Environment",
			arguments: arguments{
				config: Configuration{
					Environment:  rcv.ENVIRONMENT_PRODUCTION,
					LogDirectory: DEFAULT_LOG_DIRECTORY,
					MaxThreads:   DEFAULT_MAX_THREADS,
					PIDDirectory: DEFAULT_PID_DIRECTORY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Valid settings Directories missing",
			arguments: arguments{
				config: Configuration{
					Environment:  rcv.ENVIRONMENT_PRODUCTION,
					LogDirectory: rcv.VAL_EMPTY,
					MaxThreads:   1,
					PIDDirectory: rcv.VAL_EMPTY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Valid settings Negative threads",
			arguments: arguments{
				config: Configuration{
					Environment:  rcv.ENVIRONMENT_PRODUCTION,
					LogDirectory: rcv.VAL_EMPTY,
					MaxThreads:   -1,
					PIDDirectory: rcv.VAL_EMPTY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Valid settings Zero threads",
			arguments: arguments{
				config: Configuration{
					Environment:  rcv.ENVIRONMENT_PRODUCTION,
					LogDirectory: rcv.VAL_EMPTY,
					MaxThreads:   0,
					PIDDirectory: rcv.VAL_EMPTY,
				},
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Valid settings Greater than threads cap",
			arguments: arguments{
				config: Configuration{
					Environment:  rcv.ENVIRONMENT_PRODUCTION,
					LogDirectory: rcv.VAL_EMPTY,
					MaxThreads:   THREAD_CAP + 1,
					PIDDirectory: rcv.VAL_EMPTY,
				},
			},
			wantError: false,
		},
	}

	fmt.Println(os.Getwd())

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = ValidateConfiguration(ts.arguments.config); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Errorf(cpi.UNEXPECTED_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
		})
	}
}

// func TestValidateOptions(tPtr *testing.T) {
//
// 	type arguments struct {
// 		opts Configuration
// 	}
//
// 	//goland:noinspection ALL
// 	var (
// 		errorInfos         []cpi.ErrorInfo
// 		gotError           bool
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 		wantError bool
// 	}{
// 		{
// 			name: "Positive Case: All options values are populated except for TLS and AuthenticatorService is Cognito.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except for TLS and environment is local.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except for TLS and environment is development.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_DEVELOPMENT,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except for TLS and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     rcv.TEST_CERTIFICATE_FQN,
// 						TLSKey:      rcv.TEST_SAVUP_PRIVATE_KEY_FQN,
// 						TLSCABundle: rcv.TEST_CA_BUNDLE_FQN,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except TLSCert and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     rcv.EMPTY,
// 						TLSKey:      rcv.TEST_SAVUP_PRIVATE_KEY_FQN,
// 						TLSCABundle: rcv.TEST_CA_BUNDLE_FQN,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except TLSKey and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     rcv.TEST_CERTIFICATE_FQN,
// 						TLSKey:      rcv.EMPTY,
// 						TLSCABundle: rcv.TEST_CA_BUNDLE_FQN,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except TLSCABundle and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     rcv.TEST_CERTIFICATE_FQN,
// 						TLSKey:      rcv.TEST_SAVUP_PRIVATE_KEY_FQN,
// 						TLSCABundle: rcv.EMPTY,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except TLSCert and TLSKey and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     rcv.EMPTY,
// 						TLSKey:      rcv.EMPTY,
// 						TLSCABundle: rcv.EMPTY,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Authenticator Service is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.EMPTY,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: AWS credentials FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.EMPTY,
// 					AWSInfoFQN:             rcv.TEST_NO_SUCH_FILE,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: All options values are populated and environment is missing.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.EMPTY,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Missing Firebase Project Id",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.EMPTY,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Firebase Credentials FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_NO_SUCH_FILE,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Firebase Credentials FQN has malformed JSON.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_MALFORMED_JSON_FILE,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Log Directory is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_NO_SUCH_DIRECTORY,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Log Directory is missing.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.EMPTY,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: PID Directory is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_NO_SUCH_DIRECTORY,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: PID Directory is missing",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.EMPTY,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Positive Case: Message Prefix is SAVUPLOCAL.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Message Prefix is SAVUPDEV.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPDEV,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Message Prefix is SAVUP.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPPROD,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Message Prefix is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.TEST_STRING,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: NATS Creds FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_NO_SUCH_FILE,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: NATS URL is missing.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.EMPTY,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Private Key FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_MALFORMED_JSON_FILE,
// 					SendGridKeyFQN:         rcv.TEST_GOOD_FQN,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: SendGrid Key FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_NO_SUCH_FILE,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: SendGrid Key FQN has malformed JSON.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   rcv.AUTH_COGNITO,
// 					AWSInfoFQN:             rcv.TEST_GOOD_FQN,
// 					Environment:            rcv.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      rcv.TEST_STRING,
// 					FirebaseCredentialsFQN: rcv.TEST_GOOD_FQN,
// 					LogDirectory:           rcv.TEST_GOOD_FQN,
// 					PIDDirectory:           rcv.TEST_GOOD_FQN,
// 					MessagePrefix:          rcv.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           rcv.TEST_GOOD_FQN,
// 					NATSURL:                rcv.TEST_STRING,
// 					PlaidKeyFQN:            rcv.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         rcv.TEST_MALFORMED_JSON_FILE,
// 					StripeKeyFQN:           rcv.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if errorInfos = validateOptions(ts.arguments.opts); len(errorInfos) > 0 {
// 				gotError = true
// 			} else {
// 				gotError = false
// 			}
// 			if gotError != ts.wantError {
// 				tPtr.Error(tFunctionName, ts.name, errorInfos)
// 			}
// 		})
// 	}
// }

// func TestCheckFileExistsAndReadable(tPtr *testing.T) {
//
// 	type arguments struct {
// 		FQN       string
// 		fileLabel string
// 	}
//
// 	var (
// 		errorInfo cpi.ErrorInfo
// 		gotError  bool
// 	)
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 		wantError bool
// 	}{
// 		{
// 			name: "Positive Case: File exists and is readable.",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_GOOD_FQN,
// 				fileLabel: "Test Good FQN",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: File doesn't exist.",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_NO_SUCH_FILE,
// 				fileLabel: "Test Bad - No Such FQN",
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: File is not readable",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_UNREADABLE_FQN,
// 				fileLabel: "Test Bad - Unreadable FQN",
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if errorInfo = coreValidators.CheckFileExistsAndReadable(ts.arguments.FQN, ts.arguments.fileLabel); errorInfo.Error != nil {
// 				gotError = true
// 			} else {
// 				gotError = false
// 			}
// 			if gotError != ts.wantError {
// 				tPtr.Error(ts.name)
// 				tPtr.Error(errorInfo)
// 			}
// 		})
// 	}
// }

// func TestCheckFileValidJSON(tPtr *testing.T) {
//
// 	type arguments struct {
// 		FQN       string
// 		fileLabel string
// 	}
//
// 	var (
// 		errorInfo cpi.ErrorInfo
// 		gotError  bool
// 	)
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 		wantError bool
// 	}{
// 		{
// 			name: "Positive Case: File contains valid JSON.",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_GOOD_FQN,
// 				fileLabel: "Test Good JSON",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: File is not readable.",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_UNREADABLE_FQN,
// 				fileLabel: "Test Unreadable File",
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: File contains INVALID JSON.",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_MALFORMED_JSON_FILE,
// 				fileLabel: "Test Bad JSON",
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if errorInfo = coreValidators.CheckFileValidJSON(ts.arguments.FQN, ts.arguments.fileLabel); errorInfo.Error != nil {
// 				gotError = true
// 			} else {
// 				gotError = false
// 			}
// 			fmt.Println(gotError)
// 			if gotError != ts.wantError {
// 				tPtr.Error(ts.name)
// 				tPtr.Error(errorInfo)
// 			}
// 		})
// 	}
// }

// func TestReadAndParseConfigFile(tPtr *testing.T) {
//
// 	type arguments struct {
// 		FQN       string
// 		fileLabel string
// 	}
//
// 	var (
// 		errorInfo cpi.ErrorInfo
// 		gotError  bool
// 	)
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 		wantError bool
// 	}{
// 		{
// 			name: "Positive Case: File contains valid JSON.",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_GOOD_FQN,
// 				fileLabel: "Test Good JSON",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: File is not readable.",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_UNREADABLE_FQN,
// 				fileLabel: "Test Unreadable File",
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: File contains INVALID JSON.",
// 			arguments: arguments{
// 				FQN:       rcv.TEST_MALFORMED_JSON_FILE,
// 				fileLabel: "Test Bad JSON",
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if _, errorInfo = readAndParseConfigFile(ts.arguments.FQN); errorInfo.Error != nil {
// 				gotError = true
// 			} else {
// 				gotError = false
// 			}
// 			fmt.Println(gotError)
// 			if gotError != ts.wantError {
// 				tPtr.Error(ts.name)
// 				tPtr.Error(errorInfo)
// 			}
// 		})
// 	}
// }
