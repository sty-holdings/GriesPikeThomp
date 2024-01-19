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
	"fmt"
	"os"
	"runtime"
	"testing"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

var (
// TestMsg       nats.Msg
// TestMsgPtr    = &TestMsg
// TestValidJson = []byte("{\"name\": \"Test Name\"}")
)

// func TestBuildJSONReply(tPtr *testing.T) {
//
// 	type GoodReply struct {
// 		Name string
// 		Blah string
// 	}
//
// 	type arguments struct {
// 		reply interface{}
// 	}
//
// 	var (
// 		errorInfo  coreError.ErrorInfo
// 		gotError   bool
// 		tJSONReply []byte
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
// 				reply: GoodReply{
// 					Name: constants.TEST_FIELD_NAME,
// 					Blah: constants.TEST_STRING,
// 				},
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Empty Reply!",
// 			arguments: arguments{
// 				reply: nil,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Empty Reply!",
// 			arguments: arguments{
// 				reply: constants.TEST_STRING,
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if tJSONReply = BuildJSONReply(ts.arguments.reply, constants.EMPTY, constants.EMPTY); len(tJSONReply) == 0 {
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

// func TestConvertMapAnyToMapString(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tMapIn             = make(map[any]interface{})
// 		tMapOut            = make(map[string]interface{})
// 	)
//
// 	tPtr.Run(tFunctionName, func(t *testing.T) {
// 		if tMapOut = ConvertMapAnyToMapString(tMapIn); len(tMapOut) > 0 {
// 			tPtr.Errorf("%v Failed: Was not expecting a map with any entries.", tFunctionName)
// 		}
// 		tMapIn["string"] = "string"
// 		if tMapOut = ConvertMapAnyToMapString(tMapIn); len(tMapOut) == 0 {
// 			tPtr.Errorf("%v Failed: Was expecting a map to have entries.", tFunctionName)
// 		}
// 	})
//
// }

// This is needed, because GIT must have read access for push,
// and it must be the first test in this file.
// func TestCreateUnreadableFile(tPtr *testing.T) {
// 	_, _ = os.OpenFile(constants.TEST_UNREADABLE_FQN, os.O_CREATE, 0333)
// }

// func TestDoesDirectoryExist(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if coreValidators.DoesDirectoryExist(constants.TEST_GOOD_FQN) == false {
// 			tPtr.Errorf("%v Failed: DoesDirectoryExist returned false for %v which should exist.", tFunctionName, constants.TEST_GOOD_FQN)
// 		}
// 		_ = os.Remove(constants.TEST_NO_SUCH_FILE)
// 		if coreValidators.DoesDirectoryExist(constants.TEST_NO_SUCH_FILE) {
// 			tPtr.Errorf("%v Failed: DoesDirectoryExist returned true for %v afer it was removed.", tFunctionName, constants.TEST_NO_SUCH_FILE)
// 		}
// 	})
// }

// func TestDoesFileExist(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if coreValidators.DoesFileExist(constants.TEST_GOOD_FQN) == false {
// 			tPtr.Errorf("%v Failed: DoesFileExist returned false for %v which should exist.", tFunctionName, constants.TEST_GOOD_FQN)
// 		}
// 		_ = os.Remove(constants.TEST_NO_SUCH_FILE)
// 		if coreValidators.DoesFileExist(constants.TEST_NO_SUCH_FILE) {
// 			tPtr.Errorf("%v Failed: DoesFileExist returned true for %v afer it was removed.", tFunctionName, constants.TEST_NO_SUCH_FILE)
// 		}
// 	})
// }

// func TestFloatToPennies(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(t *testing.T) {
// 		if FloatToPennies(constants.TEST_FLOAT_123_01) != constants.TEST_FLOAT_123_01*100 {
// 			tPtr.Errorf("%v Failed: Expected the numbers to match", tFunctionName)
// 		}
// 	})
//
// }

// func TestFormatURL(tPtr *testing.T) {
//
// 	type arguments struct {
// 		protocol string
// 		domain   string
// 		port     uint
// 	}
//
// 	var (
// 		errorInfo          coreError.ErrorInfo
// 		gotError           bool
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tUrl               string
// 	)
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 		wantError bool
// 	}{
// 		{
// 			name: "Positive Case: Successful Secure, localhost, 1234",
// 			arguments: arguments{
// 				protocol: constants.HTTP_PROTOCOL_SECURE,
// 				domain:   constants.HTTP_DOMAIN_LOCALHOST,
// 				port:     1234,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful Non-Secure, localhost, 1234",
// 			arguments: arguments{
// 				protocol: constants.HTTP_PROTOCOL_NON_SECURE,
// 				domain:   constants.HTTP_DOMAIN_LOCALHOST,
// 				port:     1234,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful Secure, api-dev.savup.com, 1234",
// 			arguments: arguments{
// 				protocol: constants.HTTP_PROTOCOL_SECURE,
// 				domain:   constants.HTTP_DOMAIN_API_DEV,
// 				port:     1234,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful Non-Secure, api-dev.savup.com, 1234",
// 			arguments: arguments{
// 				protocol: constants.HTTP_PROTOCOL_NON_SECURE,
// 				domain:   constants.HTTP_DOMAIN_API_DEV,
// 				port:     1234,
// 			},
// 			wantError: false,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if tUrl = formatURL(ts.arguments.protocol, ts.arguments.domain, ts.arguments.port); tUrl == fmt.Sprintf("%v://%v:%v", ts.arguments.protocol, ts.arguments.domain, ts.arguments.port) {
// 				gotError = false
// 			} else {
// 				gotError = true
// 			}
// 			if gotError != ts.wantError {
// 				tPtr.Error(tFunctionName, ts.name, errorInfo)
// 			}
// 		})
// 	}
//
// }

// func TestGenerateEndDate(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tEnd               string
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
//
// 		if tEnd = GenerateEndDate("2024-01-10", 10); tEnd != "2024-01-20" {
// 			tPtr.Errorf("%v Failed: End date was not 10 days greater than start date.", tFunctionName)
// 		}
// 		if tEnd = GenerateEndDate("2024-01-10", 0); tEnd != "2024-01-10" {
// 			tPtr.Errorf("%v Failed: End date was not equal to start date.", tFunctionName)
// 		}
// 		if tEnd = GenerateEndDate("", 0); tEnd != constants.EMPTY {
// 			tPtr.Errorf("%v Failed: End date was not empty.", tFunctionName)
// 		}
// 	})
// }

// func TestGenerateUUIDType1(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tUUID              string
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
//
// 		if tUUID = GenerateUUIDType1(true); strings.Contains(tUUID, "-") {
// 			tPtr.Errorf("%v Failed: UUID contains dashes when removeDashes was set to true.", tFunctionName)
// 		}
// 		if tUUID = GenerateUUIDType1(false); strings.Contains(tUUID, "-") == false {
// 			tPtr.Errorf("%v Failed: UUID does not contain dashes when 'removeDashes' was set to false.", tFunctionName)
// 		}
// 		if coreValidators.IsUUIDValid(tUUID) == false {
// 			tPtr.Errorf("%v Failed: UUID is not a valid type 4.", tFunctionName)
// 		}
// 	})
// }

// func TestGenerateUUIDType4(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tUUID              string
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
//
// 		if tUUID = GenerateUUIDType4(true); strings.Contains(tUUID, "-") {
// 			tPtr.Errorf("%v Failed: UUID contains dashes when removeDashes was set to true.", tFunctionName)
// 		}
// 		if tUUID = GenerateUUIDType4(false); strings.Contains(tUUID, "-") == false {
// 			tPtr.Errorf("%v Failed: UUID does not contain dashes when 'removeDashes' was set to false.", tFunctionName)
// 		}
// 		if coreValidators.IsUUIDValid(tUUID) == false {
// 			tPtr.Errorf("%v Failed: UUID is not a valid type 4.", tFunctionName)
// 		}
// 	})
// }

// func TestGenerateURL(tPtr *testing.T) {
//
// 	//  This test is only for code coverage.
//
// 	type arguments struct {
// 		environment string
// 		secure      bool
// 	}
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 	}{
// 		{
// 			name: "Positive Case: Successful local and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_LOCAL,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful local and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_LOCAL,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_DEVELOPMENT,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_DEVELOPMENT,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_PRODUCTION,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_PRODUCTION,
// 				secure:      false,
// 			},
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			GenerateURL(ts.arguments.environment, ts.arguments.secure)
// 		})
// 	}
//
// }

// func TestGenerateVerifyEmailURL(tPtr *testing.T) {
//
// 	//  This test is only for code coverage.
//
// 	type arguments struct {
// 		environment string
// 		secure      bool
// 	}
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 	}{
// 		{
// 			name: "Positive Case: Successful local and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_LOCAL,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful local and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_LOCAL,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_DEVELOPMENT,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_DEVELOPMENT,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_PRODUCTION,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_PRODUCTION,
// 				secure:      false,
// 			},
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			fmt.Println(GenerateVerifyEmailURL(ts.arguments.environment, ts.arguments.secure))
// 		})
// 	}
//
// }

// func TestGenerateVerifyEmailURLWithUUID(tPtr *testing.T) {
//
// 	//  This test is only for code coverage.
//
// 	type arguments struct {
// 		environment string
// 		secure      bool
// 	}
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 	}{
// 		{
// 			name: "Positive Case: Successful local and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_LOCAL,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful local and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_LOCAL,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_DEVELOPMENT,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_DEVELOPMENT,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_PRODUCTION,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and non-secure",
// 			arguments: arguments{
// 				environment: constants.ENVIRONMENT_PRODUCTION,
// 				secure:      false,
// 			},
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			fmt.Println(GenerateVerifyEmailURLWithUUID(ts.arguments.environment, ts.arguments.secure))
// 		})
// 	}
//
// }

// func TestGetDate(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		_ = GetDate()
// 	})
// }

// func TestGetLegalName(tPtr *testing.T) {
//
// 	type arguments struct {
// 		firstName string
// 		lastName  string
// 	}
//
// 	var (
// 		gotError bool
// 	)
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 		wantError bool
// 	}{
// 		{
// 			name: "Positive Case: Connect to Firebase.",
// 			arguments: arguments{
// 				firstName: "first",
// 				lastName:  "last",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Missing first name",
// 			arguments: arguments{
// 				firstName: "",
// 				lastName:  "last",
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Missing last name",
// 			arguments: arguments{
// 				firstName: "first",
// 				lastName:  "",
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if tLegalName := BuildLegalName(ts.arguments.firstName, ts.arguments.lastName); tLegalName == constants.EMPTY {
// 				gotError = true
// 			} else {
// 				gotError = false
// 			}
// 			if gotError != ts.wantError {
// 				tPtr.Error(ts.name)
// 			}
// 		})
// 	}
// }

// func TestGetTime(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		_ = GetTime()
// 	})
// }

// func TestGetType(tPtr *testing.T) {
//
// 	type arguments struct {
// 		tVar          any
// 		tExpectedType string
// 	}
//
// 	type testStruct struct {
// 	}
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		err                error
// 		gotError           bool
// 		tTestStruct        testStruct
// 	)
//
// 	tests := []struct {
// 		name      string
// 		arguments arguments
// 		wantError bool
// 	}{
// 		{
// 			name: "Positive Case: Type is string.",
// 			arguments: arguments{
// 				tVar:          "first",
// 				tExpectedType: "string",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Type is Struct.",
// 			arguments: arguments{
// 				tVar:          tTestStruct,
// 				tExpectedType: "testStruct",
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Type is pointer to Struct.",
// 			arguments: arguments{
// 				tVar:          &tTestStruct,
// 				tExpectedType: "*testStruct",
// 			},
// 			wantError: false,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if tTypeGot := getType(ts.arguments.tVar); tTypeGot == ts.arguments.tExpectedType {
// 				gotError = false
// 			} else {
// 				gotError = true
// 				err = errors.New(fmt.Sprintf("%v failed: Was expecting %v and got %v! Error: %v", tFunctionName, ts.arguments.tExpectedType, tTypeGot, err.Error()))
// 			}
// 			if gotError != ts.wantError {
// 				tPtr.Error(ts.name)
// 			}
// 		})
// 	}
// }

// func TestIsFileReadable(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if coreValidators.IsFileReadable(constants.TEST_GOOD_FQN) == false {
// 			tPtr.Errorf("%v Failed: File is not readable.", tFunctionName)
// 		}
// 		_, _ = os.ReadFile(constants.TEST_NO_SUCH_FILE)
// 		if coreValidators.IsFileReadable(constants.TEST_NO_SUCH_FILE) == true {
// 			tPtr.Errorf("%v Failed: File is not readable.", tFunctionName)
// 		}
// 		if coreValidators.IsFileReadable(constants.TEST_UNREADABLE_FQN) == true {
// 			tPtr.Errorf("%v Failed: File is not readable.", tFunctionName)
// 		}
// 	})
// }

// func TestPenniesToFloat(tPtr *testing.T) {
//
// 	var (
// 		tAmount            float64
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if tAmount = PenniesToFloat(constants.TEST_NUMBER_44); tAmount != constants.TEST_NUMBER_44/100 {
// 			tPtr.Errorf("%v Failed: Was expected %v and got error.", tFunctionName, constants.TEST_NUMBER_44/100)
// 		}
// 		if tAmount = PenniesToFloat(0); tAmount != 0 {
// 			tPtr.Errorf("%v Failed: Was expected zero and got %v.", tFunctionName, tAmount)
// 		}
// 	})
// }

// func TestRedirectLogOutput(tPtr *testing.T) {
//
// 	var (
// 		tLogFileHandlerPtr *os.File
// 		tLogFQN            string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if tLogFileHandlerPtr, _ = RedirectLogOutput("/tmp"); tLogFileHandlerPtr == nil {
// 			tPtr.Errorf("%v Failed: Was expecting a pointer to be returned and got nil.", tFunctionName)
// 		}
// 		if _, tLogFQN = RedirectLogOutput("/tmp"); tLogFQN == constants.EMPTY {
// 			tPtr.Errorf("%v Failed: Was expecting the LogFQN to be populated and it was empty.", tFunctionName)
// 		}
// 	})
// }

// func TestRemovePidFile(tPtr *testing.T) {
//
// 	var (
// 		errorInfo          coreError.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		_ = WritePidFile("/tmp")
// 		if errorInfo = RemovePidFile("/tmp/server.pid"); errorInfo.Error != nil {
// 			tPtr.Errorf("%v Failed: Was expected err to be nil and got %v.", tFunctionName, errorInfo.Error.Error())
// 		}
// 		if errorInfo = RemovePidFile("/xxx/server.pid"); errorInfo.Error == nil {
// 			tPtr.Errorf("%v Failed: Was expected err to be nil and got %v.", tFunctionName, errorInfo.Error.Error())
// 		}
// 	})
// }

// func TestUnmarshalRequest(tPtr *testing.T) {
//
// 	type testStruct struct {
// 		TestField1 int `json:"test_field1"`
// 	}
//
// 	var (
// 		errorInfo          coreError.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tTestStruct        = testStruct{
// 			TestField1: 0,
// 		}
// 		tTestStructPtr = &tTestStruct
// 	)
//
// 	TestMsg.Data = []byte("{\"test_field1\": 12345}")
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if errorInfo = UnmarshalMessageData(TestMsgPtr, tTestStructPtr); errorInfo.Error != nil {
// 			tPtr.Errorf("%v Failed: Expected to get error message.", tFunctionName)
// 		}
// 		TestMsg.Data = nil
// 		if errorInfo = UnmarshalMessageData(TestMsgPtr, testStruct{}); errorInfo.Error == nil {
// 			tPtr.Errorf("%v Failed: Expected to get error message.", tFunctionName)
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
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				service: constants.AUTH_COGNITO,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Not Supported!",
// 			arguments: arguments{
// 				service: constants.AUTH_FIREBASE,
// 			},
// 			wantError: true,
// 		},
// 		{
// 			name: "Negative Case: Empty method!",
// 			arguments: arguments{
// 				service: constants.EMPTY,
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if errorInfo = coreValidators.ValidateAuthenticatorService(ts.arguments.service); errorInfo.Error != nil {
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

// func TestValidateDirectory(tPtr *testing.T) {
//
// 	var (
// 		errorInfo          coreError.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if errorInfo = coreValidators.ValidateDirectory(constants.TEST_PID_DIRECTORY); errorInfo.Error != nil {
// 			tPtr.Errorf("%v Failed: Expected err to be 'nil' and got %v.", tFunctionName, errorInfo.Error.Error())
// 		}
// 		if errorInfo = coreValidators.ValidateDirectory(constants.TEST_STRING); errorInfo.Error == nil {
// 			tPtr.Errorf("%v Failed: Expected an error and got nil.", tFunctionName)
// 		}
// 	})
// }

// func TestValidateTransferMethod(tPtr *testing.T) {
//
// 	type arguments struct {
// 		method string
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
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				method: constants.TRANFER_STRIPE,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				method: constants.TRANFER_WIRE,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				method: constants.TRANFER_CHECK,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful!",
// 			arguments: arguments{
// 				method: constants.TRANFER_ZELLE,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Negative Case: Empty method!",
// 			arguments: arguments{
// 				method: constants.EMPTY,
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if errorInfo = coreValidators.ValidateTransferMethod(ts.arguments.method); errorInfo.Error != nil {
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

// func TestWritePidFile(tPtr *testing.T) {
//
// 	var (
// 		errorInfo          coreError.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		// Create PID file
// 		if errorInfo = WritePidFile(constants.TEST_PID_DIRECTORY); errorInfo.Error != nil {
// 			tPtr.Errorf("%v Failed: Expected err to be 'nil'.", tFunctionName)
// 		}
// 		// PID directory is not provided
// 		if errorInfo = WritePidFile(constants.EMPTY); errorInfo.Error == nil {
// 			tPtr.Errorf("%v Failed: Expected err to be 'nil'.", tFunctionName)
// 		}
// 		// PID file exists
// 		if errorInfo = WritePidFile(constants.TEST_PID_DIRECTORY); errorInfo.Error != nil {
// 			tPtr.Errorf("%v Failed: Expected err to be 'nil'.", tFunctionName)
// 		}
// 	})
//
// 	_ = RemovePidFile(constants.TEST_PID_DIRECTORY)
//
// }

func TestPrependWorkingDirectory(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tPrependedFileName string
		tWorkingDirectory  string
		tTestFileName      string
	)

	tWorkingDirectory, _ = os.Getwd()
	tTestFileName = fmt.Sprintf("%v/%v", tWorkingDirectory, TEST_FILE_NAME)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Adds working directory to file name
		if tPrependedFileName = PrependWorkingDirectory(TEST_FILE_NAME); tPrependedFileName != tTestFileName {
			tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_DID_NOT_MATCH)
		}
		// Pass working directory and get back working directory
		if tPrependedFileName = PrependWorkingDirectory(tWorkingDirectory); tPrependedFileName != tWorkingDirectory {
			tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_DID_NOT_MATCH)
		}
	})
}
