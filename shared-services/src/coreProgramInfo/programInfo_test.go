// Package sharedServices
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
package sharedServices

import (
	"fmt"
	"runtime"
	"testing"

	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

func TestGetFunctionInfo(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tFunctionInfo      FunctionInfo
	)

	type arguments struct {
		level         int
		filenameValue string
	}

	var (
		gotError bool
	)

	tests := []struct {
		name               string
		arguments          arguments
		wantError          bool
		errorMessageFormat string
	}{
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Level 0 Filename Empty",
			arguments: arguments{
				level:         0,
				filenameValue: rcv.VAL_EMPTY,
			},
			errorMessageFormat: EXPECTING_NO_ERROR_FORMAT,
			wantError:          false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Level 1 Filename Empty",
			arguments: arguments{
				level:         1,
				filenameValue: rcv.VAL_EMPTY,
			},
			errorMessageFormat: EXPECTING_NO_ERROR_FORMAT,
			wantError:          false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Level 2 Filename Empty",
			arguments: arguments{
				level:         2,
				filenameValue: rcv.VAL_EMPTY,
			},
			errorMessageFormat: EXPECTING_NO_ERROR_FORMAT,
			wantError:          false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Level 3 Filename Empty",
			arguments: arguments{
				level:         3,
				filenameValue: rcv.VAL_EMPTY,
			},
			errorMessageFormat: EXPECTING_NO_ERROR_FORMAT,
			wantError:          false,
		},
		{
			name: rcv.TEST_POSITVE_SUCCESS + "Level 4 Filename Empty",
			arguments: arguments{
				level:         4,
				filenameValue: rcv.VAL_EMPTY,
			},
			errorMessageFormat: EXPECTING_NO_ERROR_FORMAT,
			wantError:          true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			tFunctionInfo = GetFunctionInfo(ts.arguments.level)
			if tFunctionInfo.FileName == ts.arguments.filenameValue ||
				tFunctionInfo.Name == rcv.VAL_EMPTY {
				gotError = true
			} else {
				fmt.Println("FileName: ", tFunctionInfo.FileName)
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Errorf(EXPECTING_NO_ERROR_FORMAT, tFunctionName, UNKNOWN)
			}
		})
	}
}

func TestGetProgramInfo(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tProgramInfo       ProgramInfo
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		tProgramInfo = GetProgramInfo()
		if tProgramInfo.FileName == rcv.VAL_EMPTY ||
			tProgramInfo.NumberCPUs == rcv.VAL_ZERO ||
			tProgramInfo.GoVersion == rcv.VAL_EMPTY {
			tPtr.Errorf(EXPECTING_NO_ERROR_FORMAT, GetFunctionInfo(1).Name, UNKNOWN)
		}
	})
}

func TestGetWorkingDirectory(tPtr *testing.T) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tProgramInfo       ProgramInfo
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		tProgramInfo.WorkingDirectory = GetWorkingDirectory()
		if tProgramInfo.WorkingDirectory == rcv.VAL_EMPTY {
			tPtr.Errorf(EXPECTING_NO_ERROR_FORMAT, GetFunctionInfo(1).Name, UNKNOWN)
		}
	})
}
