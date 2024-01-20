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
	"runtime"
	"testing"

	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

func TestNewErrorInfo(tPtr *testing.T) {

	type arguments struct {
		additionalInfo string
		myError        error
	}

	tests := []struct {
		name      string
		arguments arguments
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "No Data Supplied",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			NewErrorInfo(ts.arguments.myError, ts.arguments.additionalInfo)
		})
	}
}

func TestPrintError(tPtr *testing.T) {

	type arguments struct {
		additionalInfo string
		myError        error
	}

	tests := []struct {
		name       string
		arguments  arguments
		outputMode string
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "No Data Supplied - No Output Mode",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: "",
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied - No Output Mode",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: "",
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "No Data Supplied - Display Output Mode",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: rcv.MODE_OUTPUT_DISPLAY,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied - Display Output Mode",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: rcv.MODE_OUTPUT_DISPLAY,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "No Data Supplied - Log Output Mode",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: rcv.MODE_OUTPUT_LOG,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied - Log Output Mode",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: rcv.MODE_OUTPUT_LOG,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			PrintError(ts.arguments.myError, ts.arguments.additionalInfo, ts.outputMode)
		})
	}
}

func TestPrintErrorInfo(tPtr *testing.T) {

	type arguments struct {
		additionalInfo string
		myError        error
	}

	tests := []struct {
		name       string
		arguments  arguments
		outputMode string
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "No Data Supplied",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: "",
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: "",
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "No Data Supplied",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: rcv.MODE_OUTPUT_LOG,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: rcv.MODE_OUTPUT_LOG,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "No Data Supplied",
			arguments: arguments{
				additionalInfo: "",
				myError:        nil,
			},
			outputMode: rcv.MODE_OUTPUT_DISPLAY,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: rcv.MODE_OUTPUT_DISPLAY,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			PrintErrorInfo(NewErrorInfo(ts.arguments.myError, ts.arguments.additionalInfo), ts.outputMode)
		})
	}
}

func TestOutputError(tPtr *testing.T) {

	type arguments struct {
		additionalInfo string
		myError        error
	}

	tests := []struct {
		name       string
		arguments  arguments
		outputMode string
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: "",
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: rcv.MODE_OUTPUT_LOG,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Data Supplied",
			arguments: arguments{
				additionalInfo: rcv.TXT_EMPTY,
				myError:        ErrErrorMissing,
			},
			outputMode: rcv.MODE_OUTPUT_DISPLAY,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			outputError(NewErrorInfo(ts.arguments.myError, ts.arguments.additionalInfo), ts.outputMode)
		})
	}
}

func TestNewError(tPtr *testing.T) {

	var (
		errorInfo          ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if errorInfo = newError(ErrErrorMissing); errorInfo.Error == nil {
			tPtr.Errorf(EXPECTED_ERROR_FORMAT, tFunctionName)
		}
		if errorInfo = newError(nil); errorInfo.Error != nil {
			tPtr.Errorf(EXPECTING_NO_ERROR_FORMAT, tFunctionName, errorInfo.Error)
		}
	})
}

// Private Functions
func TestGetErrorFunctionFileNameLineNumber(tPtr *testing.T) {

	type arguments struct {
		level int
	}

	var (
		gotError           bool
		tErrorInfo         ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				level: -1,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				level: 0,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				level: 1,
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(tFunctionName, func(t *testing.T) {
			if tErrorInfo = getErrorFunctionFileNameLineNumber(ts.arguments.level); tErrorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(tErrorInfo.Error.Error())
			}
		})
	}
}
