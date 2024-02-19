// package extensions
/*
NOTES:
    None

COPYRIGHT:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package extensions

import (
	"runtime"
	"testing"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
)

func Test(tPtr *testing.T) {

	type arguments struct {
		configFilename string
	}

	var (
		errorInfo          cpi.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		gotError           bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Valid Config.",
			arguments: arguments{
				configFilename: "test-stripe-config.json",
			},
			wantError: false,
		},
		{
			name: "Negative Case: Invalid Config",
			arguments: arguments{
				configFilename: "test-invalid_stripe-config.json",
			},
			wantError: true,
		},
	}

	for _, ts := range tests {
		tPtr.Run(tFunctionName, func(t *testing.T) {
			if _, errorInfo := LoadNCInternal(ts.arguments.configFilename); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(errorInfo.Error.Error())
			}
		})
	}
}
