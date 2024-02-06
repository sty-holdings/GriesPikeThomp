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

	Written by Scott Yacko / syacko
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
	// Add imports here

	"os"
	"runtime"
)

type ProgramInfo struct {
	ErrorInfo        ErrorInfo
	FileName         string       `json:"program_filename"`
	FunctionInfo     FunctionInfo `json:"function_info"`
	GoVersion        string       `json:"go_version"`
	NumberCPUs       int          `json:"number_cpus"`
	DebugModeOn      bool         `json:"debug_mode_on"`
	WorkingDirectory string       `json:"working_directory"`
}

type FunctionInfo struct {
	FileName   string `json:"function_filename"`
	Name       string `json:"function_name"`
	LineNumber int    `json:"function_line_number"`
}

// GetFunctionInfo - returns information about the function based on the level provided.
//
// 0 will always return information about GetFunctionInfo
//
// 1 will return the caller of GetFunctionInfo
//
// 2+ will return the corresponding caller back up the chain.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetFunctionInfo(level int) (functionInfo FunctionInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(level)
	)

	functionInfo.Name = runtime.FuncForPC(tFunction).Name()
	_, functionInfo.FileName, functionInfo.LineNumber, _ = runtime.Caller(level)

	return
}

// GetWorkingDirectory - return the working directory for the program
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetWorkingDirectory() (path string) {
	path, _ = os.Getwd()

	return
}

// GetProgramInfo - returns information about the program and the system where it is executing.
// The level is set to 1, so it will always return information about the caller. It is recommended that
// you use this when you initialize your program.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetProgramInfo() (programInfo ProgramInfo) {

	_, programInfo.FileName, _, _ = runtime.Caller(1)
	programInfo.GoVersion = runtime.Version()
	programInfo.NumberCPUs = runtime.NumCPU()

	return
}
