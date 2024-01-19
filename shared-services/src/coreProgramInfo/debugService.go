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
	// Add imports here

	"fmt"
	"log"

	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

//goland:noinspection GoSnakeCaseUsage
const (
	DEBUG_FUNCTION_FORMAT = "[DEBUG_FUNCTION] File: %v Function Name: '%v' Near Line Number: %v\n"
	DEBUG_MESSAGE_FORMAT  = "[DEBUG_MESSAGE] File: %v Function Name: '%v' Message: %v\n"
)

// PrintDebugFunctionInfo - if debugMode is true the function info of the caller will be output.
// The format of the messages is
// "[DEBUG] File: {Filename} Function: {Function Name} Near Line Number: {Line Number}\n"
// Set shared_services.ProgramInfo.DebugModeOn to true to turn on debug mode. The default is false.
// This function uses the shared_services.ProgramInfo.OutputModeOn. The default is log.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PrintDebugFunctionInfo(debugModeOn bool, outputMode string) {

	var (
		tFunctionInfo = GetFunctionInfo(1)
	)

	if debugModeOn {
		if outputMode == rcv.MODE_OUTPUT_DISPLAY {
			fmt.Printf(DEBUG_FUNCTION_FORMAT, tFunctionInfo.FileName, tFunctionInfo.Name, tFunctionInfo.LineNumber)
		} else {
			log.Printf(DEBUG_FUNCTION_FORMAT, tFunctionInfo.FileName, tFunctionInfo.Name, tFunctionInfo.LineNumber)
		}
	}
}

// PrintDebugLine - if debugMode is true the function info of the caller will be output.
// The format of the messages is
// "[DEBUG] File: {Filename} Function: {Function Name} Near Line Number: {Line Number}\n"
// Set shared_services.ProgramInfo.DebugModeOn to true to turn on debug mode. The default is false.
// This function uses the shared_services.ProgramInfo.OutputModeOn. The default is log.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PrintDebugLine(message string, debugModeOn bool, outputMode string) {

	var (
		tFunctionInfo = GetFunctionInfo(1)
	)

	if message == rcv.VAL_EMPTY {
		message = rcv.TXT_MISSING
	}

	if debugModeOn {
		if outputMode == rcv.MODE_OUTPUT_DISPLAY {
			fmt.Printf(DEBUG_MESSAGE_FORMAT, tFunctionInfo.FileName, tFunctionInfo.Name, message)
		} else {
			log.Printf(DEBUG_MESSAGE_FORMAT, tFunctionInfo.FileName, tFunctionInfo.Name, message)
		}
	}
}
