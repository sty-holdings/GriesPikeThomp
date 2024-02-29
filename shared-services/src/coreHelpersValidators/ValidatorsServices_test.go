// Package coreHelpersValidators
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
package coreHelpersValidators

import (
	"bytes"
	"os"
	"runtime"
	"strings"
	"testing"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

var (
	// TestValidJson          = []byte("{\"name\": \"Test Name\"}")

	tWorkingDirectory, _ = os.Getwd()
)

// func TestAreMapKeysPopulated(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tMyMap             map[any]interface{}
// 	)
//
// 	tPtr.Run(tFunctionName, func(t *testing.T) {
// 		tMyMap = make(map[any]interface{})
// 		if AreMapKeysPopulated(tMyMap) {
// 			tPtr.Errorf("%v Failed: Expected map keys to fail.", tFunctionName)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[rcv.EMPTY] = "string"
// 		if AreMapKeysPopulated(tMyMap) {
// 			tPtr.Errorf("%v Failed: Expected map keys to fail.", tFunctionName)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap["string"] = "string"
// 		if AreMapKeysPopulated(tMyMap) == false {
// 			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[1] = "string"
// 		if AreMapKeysPopulated(tMyMap) == false {
// 			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[1] = 1
// 		if AreMapKeysPopulated(tMyMap) == false {
// 			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
// 		}
// 	})
// }

// func TestAreMapKeysValuesPopulated(tPtr *testing.T) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tMyMap             map[any]interface{}
// 	)
//
// 	tPtr.Run(tFunctionName, func(t *testing.T) {
// 		tMyMap = make(map[any]interface{})
// 		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != rcv.EMPTY_WORD {
// 			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, rcv.EMPTY_WORD)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[rcv.EMPTY] = "string"
// 		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != rcv.MISSING_KEY {
// 			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, rcv.MISSING_KEY)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[1] = rcv.EMPTY
// 		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != rcv.MISSING_VALUE {
// 			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, rcv.MISSING_VALUE)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap["string"] = "string"
// 		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != rcv.GOOD {
// 			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, rcv.GOOD)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[1] = "string"
// 		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != rcv.GOOD {
// 			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, rcv.GOOD)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[1] = 1
// 		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != rcv.GOOD {
// 			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, rcv.GOOD)
// 		}
// 	})
// }

// func TestAreMapValuesPopulated(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tMyMap             map[any]interface{}
// 	)
//
// 	tPtr.Run(tFunctionName, func(t *testing.T) {
// 		tMyMap = make(map[any]interface{})
// 		tMyMap["string"] = rcv.EMPTY
// 		if AreMapValuesPopulated(tMyMap) {
// 			tPtr.Errorf("%v Failed: Expected map keys to fail.", tFunctionName)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[1] = rcv.EMPTY
// 		if AreMapValuesPopulated(tMyMap) {
// 			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap["string"] = "string"
// 		if AreMapValuesPopulated(tMyMap) == false {
// 			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
// 		}
// 		tMyMap = make(map[any]interface{})
// 		tMyMap[1] = 0
// 		if AreMapValuesPopulated(tMyMap) == false {
// 			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
// 		}
// 	})
// }

func TestBase64Decode(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tValue             []byte
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if tValue, _ = Base64Decode(TEST_BASE64_STRING); bytes.Equal(tValue, TEST_BYTE_ARRAY) {
			} else {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.VAL_EMPTY)
			}
			if tValue, _ = Base64Decode(TEST_STRING); bytes.Equal(tValue, TEST_BYTE_ARRAY) {
				tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
			}
		},
	)
}

func TestBase64Encode(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// Adds working directory to file name
			if Base64Encode(TEST_STRING) != TEST_BASE64_STRING {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.VAL_EMPTY)
			}
		},
	)
}

func TestCheckFileExistsAndReadable(tPtr *testing.T) {

	type arguments struct {
		FQN       string
		fileLabel string
	}

	var (
		errorInfo cpi.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "File exists and is readable.",
			arguments: arguments{
				FQN:       TEST_FILE_EXISTS_FILENAME,
				fileLabel: "Test Good filename",
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "File exists and is readable - No Label.",
			arguments: arguments{
				FQN: TEST_FILE_EXISTS_FILENAME,
			},
			wantError: false,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "File doesn't exist.",
			arguments: arguments{
				FQN:       rcv.VAL_EMPTY,
				fileLabel: "Test No Such filename",
			},
			wantError: true,
		},
		{
			name: rcv.TEST_NEGATIVE_SUCCESS + "File is not readable",
			arguments: arguments{
				FQN:       TEST_FILE_UNREADABLE,
				fileLabel: "Test Unreadable FQN",
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if errorInfo = DoesFileExistsAndReadable(ts.arguments.FQN, ts.arguments.fileLabel); errorInfo.Error != nil {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error(ts.name)
					tPtr.Error(errorInfo)
				}
			},
		)
	}
}

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
// 			if errorInfo = CheckFileValidJSON(ts.arguments.FQN, ts.arguments.fileLabel); errorInfo.Error != nil {
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

func TestDoesDirectoryExist(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if DoesDirectoryExist(tWorkingDirectory) == false {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_DIRECTORY_DOES_NOT_EXIST)
			}
			if DoesDirectoryExist(rcv.VAL_EMPTY) {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_DIRECTORY_EXISTS)
			}
		},
	)
}

func TestDoesFileExist(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if DoesFileExist(TEST_FILE_EXISTS_FILENAME) == false {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_FILENAME_DOES_NOT_EXISTS)
			}
			if DoesFileExist(rcv.VAL_EMPTY) {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_FILENAME_EXISTS)
			}
		},
	)
}

func TestIsBase64Encode(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if IsBase64Encode(TEST_STRING) {
				tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
			}
			if IsBase64Encode(TEST_BASE64_STRING) == false {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, cpi.FALSE_SHOULD_BE_TRUE)
			}
			if IsBase64Encode(rcv.VAL_EMPTY) == false {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, cpi.FALSE_SHOULD_BE_TRUE)
			}
		},
	)

}

func TestIsDomainValid(tPtr *testing.T) {

	type arguments struct {
		domain string
	}

	var (
		gotError bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "No domain",
			arguments: arguments{
				domain: rcv.VAL_EMPTY,
			},
			wantError: true,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "invalid domain",
			arguments: arguments{
				domain: TEST_INVALID_DOMAIN,
			},
			wantError: true,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "good domain",
			arguments: arguments{
				domain: TEST_DOMAIN,
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(tPtr *testing.T) {
				if IsDomainValid(ts.arguments.domain) {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, ts.name, rcv.TXT_GOT_WRONG_BOOLEAN)
				}
			},
		)
	}
}

func TestIsGinModeValid(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if IsGinModeValid(rcv.MODE_DEBUG) == false {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, cpi.FALSE_SHOULD_BE_TRUE)
			}
			if IsGinModeValid(rcv.MODE_RELEASE) == false {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, cpi.FALSE_SHOULD_BE_TRUE)
			}
			if IsGinModeValid(rcv.TXT_EMPTY) {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, cpi.TRUE_SHOULD_BE_FALSE)
			}
			if IsGinModeValid(rcv.VAL_EMPTY) {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, cpi.TRUE_SHOULD_BE_FALSE)
			}
		},
	)
}

func TestIsEnvironmentValid(tPtr *testing.T) {

	type arguments struct {
		environment string
	}

	var (
		gotError bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "No environment",
			arguments: arguments{
				environment: "",
			},
			wantError: true,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "LOCAL environment",
			arguments: arguments{
				environment: strings.ToUpper(rcv.ENVIRONMENT_LOCAL),
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "local environment",
			arguments: arguments{
				environment: strings.ToLower(rcv.ENVIRONMENT_LOCAL),
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "DEVELOPMENT environment",
			arguments: arguments{
				environment: strings.ToUpper(rcv.ENVIRONMENT_DEVELOPMENT),
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "development environment",
			arguments: arguments{
				environment: strings.ToLower(rcv.ENVIRONMENT_DEVELOPMENT),
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "PRODUCTION environment",
			arguments: arguments{
				environment: strings.ToUpper(rcv.ENVIRONMENT_PRODUCTION),
			},
			wantError: false,
		},
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "production environment",
			arguments: arguments{
				environment: strings.ToLower(rcv.ENVIRONMENT_PRODUCTION),
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(tPtr *testing.T) {
				if IsEnvironmentValid(ts.arguments.environment) {
					gotError = false
				} else {
					gotError = true
				}
				if gotError != ts.wantError {
					tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, ts.name, rcv.TXT_GOT_WRONG_BOOLEAN)
				}
			},
		)
	}
}

func TestIsFileReadable(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if IsFileReadable(TEST_FILE_UNREADABLE) == true {
				tPtr.Error(cpi.TRUE_SHOULD_BE_FALSE)
			}
			if IsFileReadable(TEST_FILE_EXISTS_FILENAME) == false {
				tPtr.Error(cpi.FALSE_SHOULD_BE_TRUE)
			}
			if IsFileReadable(rcv.VAL_EMPTY) == true {
				tPtr.Error(cpi.TRUE_SHOULD_BE_FALSE)
			}
		},
	)
}

func TestIsJSONValid(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if IsJSONValid(testValidJson) == false {
				tPtr.Error(cpi.FALSE_SHOULD_BE_TRUE)
			}
			if IsJSONValid([]byte(rcv.VAL_EMPTY)) == true {
				tPtr.Error(cpi.TRUE_SHOULD_BE_FALSE)
			}
			if IsJSONValid([]byte(rcv.TXT_EMPTY)) == true {
				tPtr.Error(cpi.TRUE_SHOULD_BE_FALSE)
			}
		},
	)
}

// func TestIsURLValid(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if IsURLValid(rcv.TEST_URL_VALID) == false {
// 			tPtr.Errorf("%v Failed: Expected JSON string to be valid.", tFunctionName)
// 		}
// 		if IsURLValid(rcv.TEST_URL_INVALID) == true {
// 			tPtr.Errorf("%v Failed: Expected JSON string to be invalid.", tFunctionName)
// 		}
// 	})
// }

// func TestIsUUIDValid(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if IsUUIDValid(rcv.TEST_UUID_VALID) == false {
// 			tPtr.Errorf("%v Failed: Expected JSON string to be valid.", tFunctionName)
// 		}
// 		if IsUUIDValid(rcv.TEST_UUID_INVALID) == true {
// 			tPtr.Errorf("%v Failed: Expected JSON string to be invalid.", tFunctionName)
// 		}
// 	})
// }

// func TestValidateAuthenticatorService(tPtr *testing.T) {
//
// 	type arguments struct {
// 		service string
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
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				service: rcv.AUTH_COGNITO,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Not Supported!",
// 			arguments: arguments{
// 				service: rcv.AUTH_FIREBASE,
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Empty method!",
// 			arguments: arguments{
// 				service: rcv.EMPTY,
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if errorInfo = ValidateAuthenticatorService(ts.arguments.service); errorInfo.Error != nil {
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
//
// }

func TestValidateDirectory(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if errorInfo = ValidateDirectory(tWorkingDirectory); errorInfo.Error != nil {
				tPtr.Errorf("%v Failed: Expected err to be 'nil' and got %v.", tFunctionName, errorInfo.Error.Error())
			}
			if errorInfo = ValidateDirectory(rcv.VAL_EMPTY); errorInfo.Error == nil {
				tPtr.Errorf("%v Failed: Expected an error and got nil.", tFunctionName)
			}
			if errorInfo = ValidateDirectory(rcv.TXT_EMPTY); errorInfo.Error == nil {
				tPtr.Errorf("%v Failed: Expected an error and got nil.", tFunctionName)
			}
		},
	)
}

// func TestValidateTransferMethod(tPtr *testing.T) {
//
// 	type arguments struct {
// 		method string
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
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				method: rcv.TRANFER_STRIPE,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				method: rcv.TRANFER_WIRE,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				method: rcv.TRANFER_CHECK,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				method: rcv.TRANFER_ZELLE,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Empty method!",
// 			arguments: arguments{
// 				method: rcv.EMPTY,
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if errorInfo = ValidateTransferMethod(ts.arguments.method); errorInfo.Error != nil {
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
//
// }
