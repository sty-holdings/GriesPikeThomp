// Package sharedServices
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
package sharedServices

import (
	// Add imports here

	"log"
	"runtime"

	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type ErrorInfo struct {
	AdditionalInfo string `json:"error_additional_info"`
	Error          error  `json:"-"`
	FileName       string `json:"error_filename"`
	FunctionName   string `json:"error_function_name"`
	LineNumber     int    `json:"error_line_number"`
	Message        string `json:"error_message"`
}

// NewErrorInfo - will return an ErrorInfo object.
//
//	Customer Messages: None
//	Errors: Missing values will be filled with 'MISSING'.
//	Verifications: None
func NewErrorInfo(
	myError error,
	additionalInfo string,
) (errorInfo ErrorInfo) {

	if myError == nil {
		return
	} else {
		errorInfo = newError(myError)
	}
	if additionalInfo == rcv.VAL_EMPTY {
		errorInfo.AdditionalInfo = rcv.TXT_EMPTY
	} else {
		errorInfo.AdditionalInfo = additionalInfo
	}
	errorInfo.Message = myError.Error()

	return
}

// PrintError - will output error information using this format:
// "[ERROR] {Error Message} Additional Info: '{Additional Info}' File: {Filename} Near Line Number: {Line Number}\n"
// If the outputMode is display, the color will be red. The default is to output to the log.
//
//	Customer Messages: None
//	Errors: Missing values will be filled with 'MISSING'.
//	Verifications: None
func PrintError(
	myError error,
	additionalInfo string,
) {

	var (
		errorInfo ErrorInfo
	)

	if myError == nil {
		errorInfo = newError(ErrErrorMissing)
	} else {
		errorInfo = newError(myError)
	}
	if additionalInfo == rcv.VAL_EMPTY {
		errorInfo.AdditionalInfo = rcv.TXT_EMPTY
	} else {
		errorInfo.AdditionalInfo = additionalInfo
	}

	outputError(errorInfo)
}

// PrintErrorInfo - will output error information using this format:
// "[ERROR] {Error Message} Additional Info: '{Additional Info}' File: {Filename} Near Line Number: {Line Number}\n"
// If the outputMode is display, the color will be red. The default is to output to the log.
//
//	Customer Messages: None
//	Errors: ErrErrorMissing
//	Verifications: None
func PrintErrorInfo(errorInfo ErrorInfo) {

	if errorInfo.Error == nil {
		errorInfo = newError(ErrErrorMissing)
	}

	outputError(errorInfo)
}

// Private Functions
func outputError(errorInfo ErrorInfo) {

	log.Printf(
		"[ERROR] %v Additional Info: '%v' File: %v Near Line Number: %v\n",
		errorInfo.Error.Error(),
		errorInfo.AdditionalInfo,
		errorInfo.FileName,
		errorInfo.LineNumber,
	)
}

func newError(myError error) (errorInfo ErrorInfo) {

	errorInfo = getErrorFunctionFileNameLineNumber(3)
	errorInfo.Error = myError

	return
}

func getErrorFunctionFileNameLineNumber(level int) (errorInfo ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(level)
	)

	errorInfo.FunctionName = runtime.FuncForPC(tFunction).Name()
	_, errorInfo.FileName, errorInfo.LineNumber, _ = runtime.Caller(level)

	return
}

// DumpErrorInfos - outputs multiple error messages
//
//	func DumpErrorInfos(ErrorInfos []ErrorInfo) {
//		for _, info := range ErrorInfos {
//			PrintError(info)
//		}
//	}

// ToDo move to validation which should return the errors
// func GetMapKeyPopulatedError(finding string) (errorInfo ErrorInfo) {
//
// 	GetFunctionInfo(1)
//
// 	switch strings.ToLower(finding) {
// 	case rcv.TXT_EMPTY:
// 		errorInfo = ErrorInfo{
// 			Error:   ErrMapIsEmpty,
// 			Message: ErrMapIsEmpty.Error(),
// 		}
// 	case rcv.TXT_MISSING_KEY:
// 		errorInfo = ErrorInfo{
// 			Error:   ErrMapIsMissingKey,
// 			Message: ErrMapIsMissingKey.Error(),
// 		}
// 	case rcv.TXT_MISSING_VALUE:
// 		errorInfo = ErrorInfo{
// 			Error:   ErrMapIsMissingValue,
// 			Message: ErrMapIsMissingValue.Error(),
// 		}
// 	case rcv.VAL_EMPTY:
// 		fallthrough
// 	default:
// 		errorInfo.Error = ErrRequiredArgumentMissing
// 		errorInfo.Message = ErrRequiredArgumentMissing.Error()
// 		errorInfo.AdditionalInfo = "The 'finding' argument is empty."
// 	}
//
// 	return
// }
