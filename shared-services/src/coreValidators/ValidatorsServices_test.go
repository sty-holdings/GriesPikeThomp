// Package coreValidators
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
package coreValidators

import (
	"fmt"
	"os"
	"runtime"
	"testing"

	"albert/constants"
	"albert/core/coreError"
)

var (
	TestValidJson = []byte("{\"name\": \"Test Name\"}")
)

func TestAreMapKeysPopulated(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tMyMap             map[any]interface{}
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		tMyMap = make(map[any]interface{})
		if AreMapKeysPopulated(tMyMap) {
			tPtr.Errorf("%v Failed: Expected map keys to fail.", tFunctionName)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[constants.EMPTY] = "string"
		if AreMapKeysPopulated(tMyMap) {
			tPtr.Errorf("%v Failed: Expected map keys to fail.", tFunctionName)
		}
		tMyMap = make(map[any]interface{})
		tMyMap["string"] = "string"
		if AreMapKeysPopulated(tMyMap) == false {
			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[1] = "string"
		if AreMapKeysPopulated(tMyMap) == false {
			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[1] = 1
		if AreMapKeysPopulated(tMyMap) == false {
			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
		}
	})
}

func TestAreMapKeysValuesPopulated(tPtr *testing.T) {

	var (
		tFinding           string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tMyMap             map[any]interface{}
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		tMyMap = make(map[any]interface{})
		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != constants.EMPTY_WORD {
			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, constants.EMPTY_WORD)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[constants.EMPTY] = "string"
		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != constants.MISSING_KEY {
			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, constants.MISSING_KEY)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[1] = constants.EMPTY
		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != constants.MISSING_VALUE {
			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, constants.MISSING_VALUE)
		}
		tMyMap = make(map[any]interface{})
		tMyMap["string"] = "string"
		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != constants.GOOD {
			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, constants.GOOD)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[1] = "string"
		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != constants.GOOD {
			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, constants.GOOD)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[1] = 1
		if tFinding = AreMapKeysValuesPopulated(tMyMap); tFinding != constants.GOOD {
			tPtr.Errorf("%v Failed: Expected a finding of %v.", tFunctionName, constants.GOOD)
		}
	})
}

func TestAreMapValuesPopulated(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tMyMap             map[any]interface{}
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		tMyMap = make(map[any]interface{})
		tMyMap["string"] = constants.EMPTY
		if AreMapValuesPopulated(tMyMap) {
			tPtr.Errorf("%v Failed: Expected map keys to fail.", tFunctionName)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[1] = constants.EMPTY
		if AreMapValuesPopulated(tMyMap) {
			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
		}
		tMyMap = make(map[any]interface{})
		tMyMap["string"] = "string"
		if AreMapValuesPopulated(tMyMap) == false {
			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
		}
		tMyMap = make(map[any]interface{})
		tMyMap[1] = 0
		if AreMapValuesPopulated(tMyMap) == false {
			tPtr.Errorf("%v Failed: Expected map keys to pass.", tFunctionName)
		}
	})
}

func TestCheckFileExistsAndReadable(tPtr *testing.T) {

	type arguments struct {
		FQN       string
		fileLabel string
	}

	var (
		errorInfo coreError.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: File exists and is readable.",
			arguments: arguments{
				FQN:       constants.TEST_GOOD_FQN,
				fileLabel: "Test Good FQN",
			},
			wantError: false,
		},
		{
			name: "Negative Case: File doesn't exist.",
			arguments: arguments{
				FQN:       constants.TEST_NO_SUCH_FILE,
				fileLabel: "Test Bad - No Such FQN",
			},
			wantError: true,
		},
		{
			name: "Negative Case: File is not readable",
			arguments: arguments{
				FQN:       constants.TEST_UNREADABLE_FQN,
				fileLabel: "Test Bad - Unreadable FQN",
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = CheckFileExistsAndReadable(ts.arguments.FQN, ts.arguments.fileLabel); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}
}

func TestCheckFileValidJSON(tPtr *testing.T) {

	type arguments struct {
		FQN       string
		fileLabel string
	}

	var (
		errorInfo coreError.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: File contains valid JSON.",
			arguments: arguments{
				FQN:       constants.TEST_GOOD_FQN,
				fileLabel: "Test Good JSON",
			},
			wantError: false,
		},
		{
			name: "Negative Case: File is not readable.",
			arguments: arguments{
				FQN:       constants.TEST_UNREADABLE_FQN,
				fileLabel: "Test Unreadable File",
			},
			wantError: true,
		},
		{
			name: "Negative Case: File contains INVALID JSON.",
			arguments: arguments{
				FQN:       constants.TEST_MALFORMED_JSON_FILE,
				fileLabel: "Test Bad JSON",
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = CheckFileValidJSON(ts.arguments.FQN, ts.arguments.fileLabel); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			fmt.Println(gotError)
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}
}

func TestDoesDirectoryExist(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if DoesDirectoryExist(constants.TEST_GOOD_FQN) == false {
			tPtr.Errorf("%v Failed: DoesDirectoryExist returned false for %v which should exist.", tFunctionName, constants.TEST_GOOD_FQN)
		}
		_ = os.Remove(constants.TEST_NO_SUCH_FILE)
		if DoesDirectoryExist(constants.TEST_NO_SUCH_FILE) {
			tPtr.Errorf("%v Failed: DoesDirectoryExist returned true for %v afer it was removed.", tFunctionName, constants.TEST_NO_SUCH_FILE)
		}
	})
}

func TestDoesFileExist(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if DoesFileExist(constants.TEST_GOOD_FQN) == false {
			tPtr.Errorf("%v Failed: DoesFileExist returned false for %v which should exist.", tFunctionName, constants.TEST_GOOD_FQN)
		}
		_ = os.Remove(constants.TEST_NO_SUCH_FILE)
		if DoesFileExist(constants.TEST_NO_SUCH_FILE) {
			tPtr.Errorf("%v Failed: DoesFileExist returned true for %v afer it was removed.", tFunctionName, constants.TEST_NO_SUCH_FILE)
		}
	})
}

func TestIsFileReadable(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if IsFileReadable(constants.TEST_GOOD_FQN) == false {
			tPtr.Errorf("%v Failed: File is not readable.", tFunctionName)
		}
		_, _ = os.ReadFile(constants.TEST_NO_SUCH_FILE)
		if IsFileReadable(constants.TEST_NO_SUCH_FILE) == true {
			tPtr.Errorf("%v Failed: File is not readable.", tFunctionName)
		}
		if IsFileReadable(constants.TEST_UNREADABLE_FQN) == true {
			tPtr.Errorf("%v Failed: File is not readable.", tFunctionName)
		}
	})
}

func TestIsJSONValid(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if IsJSONValid(TestValidJson) == false {
			tPtr.Errorf("%v Failed: Expected JSON string to be valid.", tFunctionName)
		}
		if IsJSONValid([]byte(constants.EMPTY)) == true {
			tPtr.Errorf("%v Failed: Expected JSON string to be invalid.", tFunctionName)
		}
		if IsJSONValid([]byte(constants.TEST_STRING)) == true {
			tPtr.Errorf("%v Failed: Expected JSON string to be invalid.", tFunctionName)
		}
	})
}

func TestIsURLValid(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if IsURLValid(constants.TEST_URL_VALID) == false {
			tPtr.Errorf("%v Failed: Expected JSON string to be valid.", tFunctionName)
		}
		if IsURLValid(constants.TEST_URL_INVALID) == true {
			tPtr.Errorf("%v Failed: Expected JSON string to be invalid.", tFunctionName)
		}
	})
}

func TestIsUUIDValid(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if IsUUIDValid(constants.TEST_UUID_VALID) == false {
			tPtr.Errorf("%v Failed: Expected JSON string to be valid.", tFunctionName)
		}
		if IsUUIDValid(constants.TEST_UUID_INVALID) == true {
			tPtr.Errorf("%v Failed: Expected JSON string to be invalid.", tFunctionName)
		}
	})
}

func TestValidateAuthenticatorService(tPtr *testing.T) {

	type arguments struct {
		service string
	}

	var (
		errorInfo coreError.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				service: constants.AUTH_COGNITO,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Not Supported!",
			arguments: arguments{
				service: constants.AUTH_FIREBASE,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Empty method!",
			arguments: arguments{
				service: constants.EMPTY,
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = ValidateAuthenticatorService(ts.arguments.service); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}

}

func TestValidateDirectory(tPtr *testing.T) {

	var (
		errorInfo          coreError.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if errorInfo = ValidateDirectory(constants.TEST_PID_DIRECTORY); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Expected err to be 'nil' and got %v.", tFunctionName, errorInfo.Error.Error())
		}
		if errorInfo = ValidateDirectory(constants.TEST_STRING); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Expected an error and got nil.", tFunctionName)
		}
	})
}

func TestValidateTransferMethod(tPtr *testing.T) {

	type arguments struct {
		method string
	}

	var (
		errorInfo coreError.ErrorInfo
		gotError  bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				method: constants.TRANFER_STRIPE,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				method: constants.TRANFER_WIRE,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				method: constants.TRANFER_CHECK,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				method: constants.TRANFER_ZELLE,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Empty method!",
			arguments: arguments{
				method: constants.EMPTY,
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = ValidateTransferMethod(ts.arguments.method); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}

}
