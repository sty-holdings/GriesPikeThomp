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
// 		errorInfo  cpi.ErrorInfo
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
// 					Name: rcv.TEST_FIELD_NAME,
// 					Blah: rcv.TEST_STRING,
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
// 				reply: rcv.TEST_STRING,
// 			},
// 			wantError: true,
// 		},
// 	}
//
// 	for _, ts := range tests {
// 		tPtr.Run(ts.name, func(t *testing.T) {
// 			if tJSONReply = BuildJSONReply(ts.arguments.reply, rcv.EMPTY, rcv.EMPTY); len(tJSONReply) == 0 {
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

func TestConvertSliceToSliceOfPtrs(tPtr *testing.T) {

	type arguments struct {
		paymentMethodTypes []string
	}

	var (
		gotError       bool
		sliceOut       []*string
		paymentMethods []string
	)

	// Append the constants to the slice
	paymentMethods = append(paymentMethods, rcv.PAYMENT_METHOD_CARD)
	paymentMethods = append(paymentMethods, rcv.PAYMENT_METHOD_PAYNOW)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: rcv.TEST_POSITIVE_SUCCESS + "Successful!",
			arguments: arguments{
				paymentMethodTypes: paymentMethods,
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(
			ts.name, func(t *testing.T) {
				if sliceOut = ConvertStringSliceToSliceOfPtrs(ts.arguments.paymentMethodTypes); len(sliceOut) == 0 {
					gotError = true
				} else {
					gotError = false
				}
				if gotError != ts.wantError {
					tPtr.Error("TEST failed, investigate.")
				}
			},
		)
	}
}

// This is needed, because GIT must have read access for push,
// and it must be the first test in this file.
// func TestCreateUnreadableFile(tPtr *testing.T) {
// 	_, _ = os.OpenFile(rcv.TEST_UNREADABLE_FQN, os.O_CREATE, 0333)
// }

// func TestDoesDirectoryExist(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if coreValidators.DoesDirectoryExist(rcv.TEST_GOOD_FQN) == false {
// 			tPtr.Errorf("%v Failed: DoesDirectoryExist returned false for %v which should exist.", tFunctionName, rcv.TEST_GOOD_FQN)
// 		}
// 		_ = os.Remove(rcv.TEST_NO_SUCH_FILE)
// 		if coreValidators.DoesDirectoryExist(rcv.TEST_NO_SUCH_FILE) {
// 			tPtr.Errorf("%v Failed: DoesDirectoryExist returned true for %v afer it was removed.", tFunctionName, rcv.TEST_NO_SUCH_FILE)
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
// 		if coreValidators.DoesFileExist(rcv.TEST_GOOD_FQN) == false {
// 			tPtr.Errorf("%v Failed: DoesFileExist returned false for %v which should exist.", tFunctionName, rcv.TEST_GOOD_FQN)
// 		}
// 		_ = os.Remove(rcv.TEST_NO_SUCH_FILE)
// 		if coreValidators.DoesFileExist(rcv.TEST_NO_SUCH_FILE) {
// 			tPtr.Errorf("%v Failed: DoesFileExist returned true for %v afer it was removed.", tFunctionName, rcv.TEST_NO_SUCH_FILE)
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
// 		if FloatToPennies(rcv.TEST_FLOAT_123_01) != rcv.TEST_FLOAT_123_01*100 {
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
// 		errorInfo          cpi.ErrorInfo
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
// 				protocol: rcv.HTTP_PROTOCOL_SECURE,
// 				domain:   rcv.HTTP_DOMAIN_LOCALHOST,
// 				port:     1234,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful Non-Secure, localhost, 1234",
// 			arguments: arguments{
// 				protocol: rcv.HTTP_PROTOCOL_NON_SECURE,
// 				domain:   rcv.HTTP_DOMAIN_LOCALHOST,
// 				port:     1234,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful Secure, api-dev.savup.com, 1234",
// 			arguments: arguments{
// 				protocol: rcv.HTTP_PROTOCOL_SECURE,
// 				domain:   rcv.HTTP_DOMAIN_API_DEV,
// 				port:     1234,
// 			},
// 			wantError: false,
// 		},
// 		{
// 			name: "Positive Case: Successful Non-Secure, api-dev.savup.com, 1234",
// 			arguments: arguments{
// 				protocol: rcv.HTTP_PROTOCOL_NON_SECURE,
// 				domain:   rcv.HTTP_DOMAIN_API_DEV,
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
// 		if tEnd = GenerateEndDate("", 0); tEnd != rcv.EMPTY {
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
// 				environment: rcv.ENVIRONMENT_LOCAL,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful local and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_LOCAL,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_DEVELOPMENT,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_DEVELOPMENT,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_PRODUCTION,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_PRODUCTION,
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
// 				environment: rcv.ENVIRONMENT_LOCAL,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful local and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_LOCAL,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_DEVELOPMENT,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_DEVELOPMENT,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_PRODUCTION,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_PRODUCTION,
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
// 				environment: rcv.ENVIRONMENT_LOCAL,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful local and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_LOCAL,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_DEVELOPMENT,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful development and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_DEVELOPMENT,
// 				secure:      false,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_PRODUCTION,
// 				secure:      true,
// 			},
// 		},
// 		{
// 			name: "Positive Case: Successful production and non-secure",
// 			arguments: arguments{
// 				environment: rcv.ENVIRONMENT_PRODUCTION,
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
// 			if tLegalName := BuildLegalName(ts.arguments.firstName, ts.arguments.lastName); tLegalName == rcv.EMPTY {
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
//		type arguments struct {
//			tVar          any
//			tExpectedType string
//		}
//
//		type testStruct struct {
//		}
//
//		var (
//			tFunction, _, _, _ = runtime.Caller(0)
//			tFunctionName      = runtime.FuncForPC(tFunction).Name()
//			err                error
//			gotError           bool
//			tTestStruct        testStruct
//		)
//
//		tests := []struct {
//			name      string
//			arguments arguments
//			wantError bool
//		}{
//			{
//				name: "Positive Case: Type is string.",
//				arguments: arguments{
//					tVar:          "first",
//					tExpectedType: "string",
//				},
//				wantError: false,
//			},
//			{
//				name: "Positive Case: Type is Struct.",
//				arguments: arguments{
//					tVar:          tTestStruct,
//					tExpectedType: "testStruct",
//				},
//				wantError: false,
//			},
//			{
//				name: "Positive Case: Type is pointer to Struct.",
//				arguments: arguments{
//					tVar:          &tTestStruct,
//					tExpectedType: "*testStruct",
//				},
//				wantError: false,
//			},
//		}
//
//		for _, ts := range tests {
//			tPtr.Run(ts.name, func(t *testing.T) {
//				if tTypeGot := getType(ts.arguments.tVar); tTypeGot == ts.arguments.tExpectedType {
//					gotError = false
//				} else {
//					gotError = true
//					err = errors.New(fmt.Sprintf("%v failed: Was expecting %v and got %v! Error: %v", tFunctionName, ts.arguments.tExpectedType, tTypeGot, err.Error()))
//				}
//				if gotError != ts.wantError {
//					tPtr.Error(ts.name)
//				}
//			})
//		}
//	}

// func TestIsFileReadable(tPtr *testing.T) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if coreValidators.IsFileReadable(rcv.TEST_GOOD_FQN) == false {
// 			tPtr.Errorf("%v Failed: File is not readable.", tFunctionName)
// 		}
// 		_, _ = os.ReadFile(rcv.TEST_NO_SUCH_FILE)
// 		if coreValidators.IsFileReadable(rcv.TEST_NO_SUCH_FILE) == true {
// 			tPtr.Errorf("%v Failed: File is not readable.", tFunctionName)
// 		}
// 		if coreValidators.IsFileReadable(rcv.TEST_UNREADABLE_FQN) == true {
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
// 		if tAmount = PenniesToFloat(rcv.TEST_NUMBER_44); tAmount != rcv.TEST_NUMBER_44/100 {
// 			tPtr.Errorf("%v Failed: Was expected %v and got error.", tFunctionName, rcv.TEST_NUMBER_44/100)
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
// 		if _, tLogFQN = RedirectLogOutput("/tmp"); tLogFQN == rcv.EMPTY {
// 			tPtr.Errorf("%v Failed: Was expecting the LogFQN to be populated and it was empty.", tFunctionName)
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
// 		errorInfo          cpi.ErrorInfo
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
// 		errorInfo          cpi.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		if errorInfo = coreValidators.ValidateDirectory(rcv.TEST_PID_DIRECTORY); errorInfo.Error != nil {
// 			tPtr.Errorf("%v Failed: Expected err to be 'nil' and got %v.", tFunctionName, errorInfo.Error.Error())
// 		}
// 		if errorInfo = coreValidators.ValidateDirectory(rcv.TEST_STRING); errorInfo.Error == nil {
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
// 		errorInfo          cpi.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
// 		// Create PID file
// 		if errorInfo = WritePidFile(rcv.TEST_PID_DIRECTORY); errorInfo.Error != nil {
// 			tPtr.Errorf("%v Failed: Expected err to be 'nil'.", tFunctionName)
// 		}
// 		// PID directory is not provided
// 		if errorInfo = WritePidFile(rcv.EMPTY); errorInfo.Error == nil {
// 			tPtr.Errorf("%v Failed: Expected err to be 'nil'.", tFunctionName)
// 		}
// 		// PID file exists
// 		if errorInfo = WritePidFile(rcv.TEST_PID_DIRECTORY); errorInfo.Error != nil {
// 			tPtr.Errorf("%v Failed: Expected err to be 'nil'.", tFunctionName)
// 		}
// 	})
//
// 	_ = RemovePidFile(rcv.TEST_PID_DIRECTORY)
//
// }

func TestIsDirectoryFullyQualified(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// Adds working directory to file name
			if IsDirectoryFullyQualified(TEST_DIRECTORY_ENDING_SLASH) == false {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_GOT_WRONG_BOOLEAN)
			}
			// Pass working directory and get back working directory
			if IsDirectoryFullyQualified(TEST_DIRECTORY) {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_GOT_WRONG_BOOLEAN)
			}
		},
	)
}

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

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			// Adds working directory to file name
			if tPrependedFileName = PrependWorkingDirectory(TEST_FILE_NAME); tPrependedFileName != tTestFileName {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_DID_NOT_MATCH)
			}
			// Pass working directory and get back working directory
			if tPrependedFileName = PrependWorkingDirectory(tWorkingDirectory); tPrependedFileName != tWorkingDirectory {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, rcv.TXT_DID_NOT_MATCH)
			}
		},
	)
}

func TestPrependWorkingDirectoryWithEndingSlash(tPtr *testing.T) {

	var (
		tFunction, _, _, _   = runtime.Caller(0)
		tFunctionName        = runtime.FuncForPC(tFunction).Name()
		tWorkingDirectory, _ = os.Getwd()
	)

	var tests = []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "TestFileName",
			input:    TEST_FILE_NAME,
			expected: fmt.Sprintf("%v/%v/", tWorkingDirectory, TEST_FILE_NAME),
		},
		{
			name:     "TestDirectory",
			input:    TEST_DIRECTORY,
			expected: TEST_DIRECTORY,
		},
		{
			name:     "TestNonRootDirectory",
			input:    TEST_DIRECTORY_NON_ROOT,
			expected: fmt.Sprintf("%v/%v/", tWorkingDirectory, TEST_DIRECTORY_NON_ROOT),
		},
		{
			name:     "WorkingDirectory",
			input:    tWorkingDirectory,
			expected: tWorkingDirectory,
		},
	}

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			for _, tt := range tests {
				tPtr.Run(
					tt.name, func(t *testing.T) {
						if output := PrependWorkingDirectoryWithEndingSlash(tt.input); output != tt.expected {
							t.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tt.name, rcv.TXT_DID_NOT_MATCH)
						}
					},
				)
			}
		},
	)
}

func TestRedirectLogOutput(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tLogFileHandlerPtr *os.File
		tLogFQN            string
	)

	tLogFileHandlerPtr, tLogFQN, _ = createLogFile(TEST_DIRECTORY_ENDING_SLASH)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if errorInfo = RedirectLogOutput(tLogFileHandlerPtr, rcv.MODE_OUTPUT_LOG); errorInfo.Error != nil {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
			if errorInfo = RedirectLogOutput(tLogFileHandlerPtr, rcv.MODE_OUTPUT_LOG_DISPLAY); errorInfo.Error != nil {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
			if errorInfo = RedirectLogOutput(tLogFileHandlerPtr, rcv.VAL_EMPTY); errorInfo.Error == nil {
				tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
			}
		},
	)

	_ = os.Remove(tLogFQN)
}

func TestRemovePidFile(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tTestFQN           = TEST_DIRECTORY_ENDING_SLASH + TEST_FILE_NAME
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			_ = WritePidFile(tTestFQN, 777)
			if errorInfo = RemovePidFile(tTestFQN); errorInfo.Error != nil {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
			if errorInfo = RemovePidFile(rcv.VAL_EMPTY); errorInfo.Error == nil {
				tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
			}
		},
	)
}

func TestWriteFile(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tTestFQN           = TEST_DIRECTORY_ENDING_SLASH + TEST_FILE_NAME
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if errorInfo = WriteFile(tTestFQN, []byte(rcv.TXT_EMPTY), 0777); errorInfo.Error != nil {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
			_ = os.Remove(tTestFQN)
			if errorInfo = WriteFile(rcv.VAL_EMPTY, []byte(rcv.TXT_EMPTY), 0777); errorInfo.Error == nil {
				tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
			}
		},
	)
}

func TestWritePidFile(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tTestFQN           = TEST_DIRECTORY_ENDING_SLASH + TEST_FILE_NAME
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if errorInfo = WritePidFile(tTestFQN, 777); errorInfo.Error != nil {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
			_ = os.Remove(tTestFQN)
			if errorInfo = WritePidFile(rcv.VAL_EMPTY, 777); errorInfo.Error == nil {
				tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
			}
		},
	)
}

// Private Functions
func TestCreateAndRedirectLogOutput(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tLogFQN            string
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if _, tLogFQN, errorInfo = CreateAndRedirectLogOutput(TEST_DIRECTORY_ENDING_SLASH, rcv.MODE_OUTPUT_LOG); errorInfo.Error != nil {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
			fmt.Println(os.Remove(tLogFQN))
			if _, tLogFQN, errorInfo = CreateAndRedirectLogOutput(TEST_DIRECTORY_ENDING_SLASH, rcv.MODE_OUTPUT_LOG_DISPLAY); errorInfo.Error != nil {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
			fmt.Println(os.Remove(tLogFQN))
			if _, tLogFQN, errorInfo = CreateAndRedirectLogOutput(TEST_DIRECTORY_ENDING_SLASH, rcv.VAL_EMPTY); errorInfo.Error == nil {
				tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
			}
		},
	)

}

func TestCreateLogFile(tPtr *testing.T) {

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tLogFQN            string
	)

	tPtr.Run(
		tFunctionName, func(tPtr *testing.T) {
			if _, tLogFQN, errorInfo = createLogFile(TEST_DIRECTORY_ENDING_SLASH); errorInfo.Error != nil {
				tPtr.Errorf(cpi.EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error.Error())
			}
			_ = os.Remove(tLogFQN)
			if _, _, errorInfo = createLogFile(TEST_DIRECTORY); errorInfo.Error == nil {
				tPtr.Errorf(cpi.EXPECTED_ERROR_FORMAT, tFunctionName)
			}
		},
	)

}
