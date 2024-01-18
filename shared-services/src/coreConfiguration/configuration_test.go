// Package coreOptions
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
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		tReader, tWriter, _ := os.Pipe()
		old := os.Stdout
		os.Stdout = tWriter

		GenerateConfigFileSkeleton("NATS Connect Test",
			"/Users/syacko/workspace/sty-holdings/gpt2/shared-services/src/coreConfiguration",
			"/skeleton-config-file.")

		tBuffer := make([]byte, 3072)
		_, _ = tReader.Read(tBuffer)

		_ = tWriter.Close()
		os.Stdout = old

		if len(tBuffer) == 0 {
			tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, cpi.ErrBufferEmpty)
		}
	})
}

func TestReadAndParseConfigFile(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tConfigFilename    = fmt.Sprintf("%v.json", DEFAULT_SKELETON_CONFIG_FILENAME_NO_SUFFIX)
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {

		fmt.Println(os.Getwd())
		if _, errorInfo = ReadAndParseConfigFile(tConfigFilename); errorInfo.Error != nil {
			tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
		}
		if _, errorInfo = ReadAndParseConfigFile(rcv.VAL_EMPTY); errorInfo.Error == nil {
			tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
		}
	})
}

// func TestValidateOptions(tPtr *testing.T) {
//
// 	type arguments struct {
// 		opts Configuration
// 	}
//
// 	//goland:noinspection ALL
// 	var (
// 		errorInfos         []coreError.ErrorInfo
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
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except for TLS and environment is development.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_DEVELOPMENT,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except for TLS and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     constants.TEST_CERTIFICATE_FQN,
// 						TLSKey:      constants.TEST_SAVUP_PRIVATE_KEY_FQN,
// 						TLSCABundle: constants.TEST_CA_BUNDLE_FQN,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except TLSCert and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     constants.EMPTY,
// 						TLSKey:      constants.TEST_SAVUP_PRIVATE_KEY_FQN,
// 						TLSCABundle: constants.TEST_CA_BUNDLE_FQN,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except TLSKey and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     constants.TEST_CERTIFICATE_FQN,
// 						TLSKey:      constants.EMPTY,
// 						TLSCABundle: constants.TEST_CA_BUNDLE_FQN,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except TLSCABundle and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     constants.TEST_CERTIFICATE_FQN,
// 						TLSKey:      constants.TEST_SAVUP_PRIVATE_KEY_FQN,
// 						TLSCABundle: constants.EMPTY,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: All options values are populated except TLSCert and TLSKey and environment is production.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_PRODUCTION,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 					TLS: coreJWT.TLSInfo{
// 						TLSCert:     constants.EMPTY,
// 						TLSKey:      constants.EMPTY,
// 						TLSCABundle: constants.EMPTY,
// 					},
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Authenticator Service is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.EMPTY,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: AWS credentials FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.EMPTY,
// 					AWSInfoFQN:             constants.TEST_NO_SUCH_FILE,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: All options values are populated and environment is missing.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.EMPTY,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Missing Firebase Project Id",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.EMPTY,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Firebase Credentials FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_NO_SUCH_FILE,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Firebase Credentials FQN has malformed JSON.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_MALFORMED_JSON_FILE,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Log Directory is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_NO_SUCH_DIRECTORY,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Log Directory is missing.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.EMPTY,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: PID Directory is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_NO_SUCH_DIRECTORY,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: PID Directory is missing",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.EMPTY,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Positive Case: Message Prefix is SAVUPLOCAL.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Message Prefix is SAVUPDEV.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPDEV,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Message Prefix is SAVUP.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPPROD,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Message Prefix is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.TEST_STRING,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: NATS Creds FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_NO_SUCH_FILE,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: NATS URL is missing.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.EMPTY,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Private Key FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_MALFORMED_JSON_FILE,
// 					SendGridKeyFQN:         constants.TEST_GOOD_FQN,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: SendGrid Key FQN is invalid.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_NO_SUCH_FILE,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
// 				},
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: SendGrid Key FQN has malformed JSON.",
// 			arguments: arguments{
// 				opts: Options{
// 					AuthenticatorService:   constants.AUTH_COGNITO,
// 					AWSInfoFQN:             constants.TEST_GOOD_FQN,
// 					Environment:            constants.ENVIRONMENT_LOCAL,
// 					FirebaseProjectId:      constants.TEST_STRING,
// 					FirebaseCredentialsFQN: constants.TEST_GOOD_FQN,
// 					LogDirectory:           constants.TEST_GOOD_FQN,
// 					PIDDirectory:           constants.TEST_GOOD_FQN,
// 					MessagePrefix:          constants.MESSAGE_PREFIX_SAVUPLOCAL,
// 					NATSCredsFQN:           constants.TEST_GOOD_FQN,
// 					NATSURL:                constants.TEST_STRING,
// 					PlaidKeyFQN:            constants.TEST_GOOD_FQN,
// 					SendGridKeyFQN:         constants.TEST_MALFORMED_JSON_FILE,
// 					StripeKeyFQN:           constants.TEST_GOOD_FQN,
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
// 		errorInfo coreError.ErrorInfo
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
// 				FQN:       constants.TEST_GOOD_FQN,
// 				fileLabel: "Test Good FQN",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: File doesn't exist.",
// 			arguments: arguments{
// 				FQN:       constants.TEST_NO_SUCH_FILE,
// 				fileLabel: "Test Bad - No Such FQN",
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: File is not readable",
// 			arguments: arguments{
// 				FQN:       constants.TEST_UNREADABLE_FQN,
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
// 		errorInfo coreError.ErrorInfo
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
// 				FQN:       constants.TEST_GOOD_FQN,
// 				fileLabel: "Test Good JSON",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: File is not readable.",
// 			arguments: arguments{
// 				FQN:       constants.TEST_UNREADABLE_FQN,
// 				fileLabel: "Test Unreadable File",
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: File contains INVALID JSON.",
// 			arguments: arguments{
// 				FQN:       constants.TEST_MALFORMED_JSON_FILE,
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
// 		errorInfo coreError.ErrorInfo
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
// 				FQN:       constants.TEST_GOOD_FQN,
// 				fileLabel: "Test Good JSON",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: File is not readable.",
// 			arguments: arguments{
// 				FQN:       constants.TEST_UNREADABLE_FQN,
// 				fileLabel: "Test Unreadable File",
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: File contains INVALID JSON.",
// 			arguments: arguments{
// 				FQN:       constants.TEST_MALFORMED_JSON_FILE,
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
