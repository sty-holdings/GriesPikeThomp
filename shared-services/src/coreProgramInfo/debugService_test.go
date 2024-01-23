// package sharedServices
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
package sharedServices

import (
	"runtime"
	"testing"

	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

func TestPrintDebugFunctionInfo(tPtr *testing.T) {

	type arguments struct {
		debugModeOn bool
		outputMode  string
	}

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Debug False - No output.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  "",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug true - No output.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  "",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug False - Output Display.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  rcv.MODE_OUTPUT_DISPLAY,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug True - Output Display.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  rcv.MODE_OUTPUT_DISPLAY,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug False - Output Log.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  rcv.MODE_OUTPUT_LOG,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug True - Output Log.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  rcv.MODE_OUTPUT_LOG,
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(tFunctionName, func(t *testing.T) {
			PrintDebugFunctionInfo(ts.arguments.debugModeOn, ts.arguments.outputMode)
		})
	}
}

func TestPrintDebugLine(tPtr *testing.T) {

	type arguments struct {
		message     string
		debugModeOn bool
		outputMode  string
	}

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Debug False - No output.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  "",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug true - No output.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  "",
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug False - Output Display.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  rcv.MODE_OUTPUT_DISPLAY,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug True - Output Display.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  rcv.MODE_OUTPUT_DISPLAY,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug False - Output Log.",
			arguments: arguments{
				debugModeOn: false,
				outputMode:  rcv.MODE_OUTPUT_LOG,
			},
			wantError: false,
		},
		{
			name: "Positive Case: Debug True - Output Log.",
			arguments: arguments{
				debugModeOn: true,
				outputMode:  rcv.MODE_OUTPUT_LOG,
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(tFunctionName, func(t *testing.T) {
			PrintDebugLine(ts.arguments.message, ts.arguments.debugModeOn, ts.arguments.outputMode)
		})
	}
}
