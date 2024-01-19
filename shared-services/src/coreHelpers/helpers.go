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
	"log"
	"os"
	"strings"
	"time"

	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// BuildJSONReply - return a JSON reply object
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func BuildJSONReply(reply interface{}, requestorId, message string) (jsonReply []byte) {
//
// 	var (
// 		errorInfo          cpi.ErrorInfo
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	if jsonReply, errorInfo.Error = json.Marshal(reply); errorInfo.Error != nil {
// 		errorInfo.AdditionalInfo = errorInfo.Error.Error()
// 		errorInfo.Error = cpi.ErrJSONGenrationFailed
// 		cpi.PrintError(errorInfo)
// 		// 	todo Error Handling & Notification
// 	}
//
// 	if coreValidators.IsJSONValid(jsonReply) == false {
// 		jsonReply = nil // Initializing the field
// 		errorInfo.AdditionalInfo = errorInfo.Error.Error()
// 		errorInfo.Error = cpi.ErrJSONInvalid
// 		cpi.PrintError(errorInfo)
// 		// 	todo Error Handling & Notification
// 	}
//
// 	return
// }

// BuildJSONRequest
// func BuildJSONRequest(request interface{}) (jsonRequest []byte) {
//
// 	var (
// 		err error
// 	)
//
// 	if jsonRequest, err = json.Marshal(request); err != nil {
// 		err = errors.New(fmt.Sprintf("Failed to generate JSON payload. Error: %v", err.Error()))
// 		log.Println(err.Error())
// 		// 	todo Error Handling & Notification
// 	}
//
// 	if coreValidators.IsJSONValid(jsonRequest) == false {
// 		jsonRequest = nil
// 		err = errors.New(fmt.Sprintf("Was not able to generate valid json for request %v", request))
// 		log.Println(err.Error())
// 		// 	todo Error Handling & Notification
// 	}
//
// 	return
// }

// BuildLegalName
// func BuildLegalName(firstName, lastName string) (legalName string) {
//
// 	if firstName != constants.EMPTY && lastName != constants.EMPTY {
// 		legalName = fmt.Sprintf("%v %v", firstName, lastName)
// 	}
//
// 	return
// }

// CapitalizeFirstLetter - will make the first letter of the string to upper case and the other letters to lower
// func CapitalizeFirstLetter(stringIn string) string {
//
// 	if stringIn == constants.EMPTY {
// 		return constants.EMPTY
// 	}
//
// 	x := []byte(stringIn)
// 	y := bytes.ToUpper([]byte{x[0]})
// 	z := bytes.ToLower(x[1:])
//
// 	return string(bytes.Join([][]byte{y, z}, nil))
// }

// ConvertMapAnyToMapString
// func ConvertMapAnyToMapString(mapIn map[any]interface{}) (mapOut map[string]interface{}) {
//
// 	mapOut = make(map[string]interface{})
//
// 	if coreValidators.IsMapPopulated(mapIn) {
// 		for key, value := range mapIn {
// 			mapOut[key.(string)] = value
// 		}
// 	}
//
// 	return
// }

// FloatToPennies
// func FloatToPennies(amount float64) (pennies int64) {
//
// 	return int64(amount * 100)
// }

// formatURL - will return a formatted url with the protocol, domain, and port.
//
//	Validation: none
//	Format: "%v://%v:%v"
//	Example: http://verifyemail.savup.com:2134, https://verifyemail.savup.com:2134, http://localhost:2134, https://localhost:2134
// func formatURL(protocol, domain string, port uint) (url string) {
//
// 	if domain == constants.ENVIRONMENT_LOCAL {
// 		url = fmt.Sprintf("%v://%v:%v", protocol, constants.HTTP_DOMAIN_LOCALHOST, port)
// 	} else {
// 		url = fmt.Sprintf("%v://%v:%v", protocol, domain, port)
// 	}
//
// 	return
// }

// GenerateEndDate - will return a string by taking the startDate and adding months.
// If the startDate is empty the endDate will be empty.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func GenerateEndDate(startDate string, months int) (endDate string) {
//
// 	var (
// 		err    error
// 		tStart time.Time
// 	)
//
// 	if startDate == "" {
// 		endDate = ""
// 	} else {
// 		// Parse the start date string.
// 		tStart, err = time.Parse("2006-01-02", startDate)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// Calculate the end date.
// 		end := tStart.AddDate(0, months, 0)
// 		endDate = end.Format("2006-01-02")
// 	}
//
// 	return
// }

// GenerateURL - will return the protocol, domain, and port. Using HTTP_PROTOCOL_SECURE or HTTP_PROTOCOL_NON_SECURE,
// ENDPOINT_VERIFY_EMAIL and HTTP_SECURE_PORT or HTTP_NON_SECURE_PORT based on the arguments.
//
//	Customer Messages: None
//	Errors: None
//	Verification: none
//	Example: http://localhost:1234, https://localhost:1234, http://api-dev.savup.com:1234, https://api-dev.savup.com:1234
//
// ToDo Change the Environment_local domain to local host once we have resolved the handshake issue happening when savup-http is run locally.
// func GenerateURL(environment string, secure bool) (url string) {
//
// 	switch fmt.Sprintf("%v,%v", strings.ToUpper(environment), secure) {
// 	case fmt.Sprintf("%v,%v", constants.ENVIRONMENT_LOCAL, true):
// 		url = formatURL(constants.HTTP_PROTOCOL_SECURE, constants.HTTP_DOMAIN_API_LOCAL, constants.HTTP_PORT_SECURE)
// 	case fmt.Sprintf("%v,%v", constants.ENVIRONMENT_LOCAL, false):
// 		url = formatURL(constants.HTTP_PROTOCOL_NON_SECURE, constants.HTTP_DOMAIN_API_LOCAL, constants.HTTP_PORT_NON_SECURE)
// 	case fmt.Sprintf("%v,%v", constants.ENVIRONMENT_DEVELOPMENT, true):
// 		url = formatURL(constants.HTTP_PROTOCOL_SECURE, constants.HTTP_DOMAIN_API_DEV, constants.HTTP_PORT_SECURE)
// 	case fmt.Sprintf("%v,%v", constants.ENVIRONMENT_DEVELOPMENT, false):
// 		url = formatURL(constants.HTTP_PROTOCOL_NON_SECURE, constants.HTTP_DOMAIN_API_DEV, constants.HTTP_PORT_NON_SECURE)
// 	case fmt.Sprintf("%v,%v", constants.ENVIRONMENT_PRODUCTION, true):
// 		url = formatURL(constants.HTTP_PROTOCOL_SECURE, constants.HTTP_DOMAIN_API_PROD, constants.HTTP_PORT_SECURE)
// 	case fmt.Sprintf("%v,%v", constants.ENVIRONMENT_PRODUCTION, false):
// 		url = formatURL(constants.HTTP_PROTOCOL_NON_SECURE, constants.HTTP_DOMAIN_API_PROD, constants.HTTP_PORT_NON_SECURE)
// 	}
//
// 	return
// }

// GenerateUUIDType1
// func GenerateUUIDType1(removeDashes bool) (myUUID string) {
//
// 	_UUID, _ := uuid.NewUUID()
// 	myUUID = fmt.Sprintf("%v", _UUID)
//
// 	if removeDashes {
// 		myUUID = strings.Replace(myUUID, "-", "", -1)
// 	}
//
// 	return
// }

// GenerateUUIDType4
// func GenerateUUIDType4(removeDashes bool) (myUUID string) {
//
// 	_UUID, _ := uuid.NewRandom()
// 	myUUID = fmt.Sprintf("%v", _UUID)
//
// 	if removeDashes {
// 		myUUID = strings.Replace(myUUID, "-", "", -1)
// 	}
//
// 	return
// }

// GenerateVerifyEmailURLWithUUID - return the url and uuid for the Verify Email.
// func GenerateVerifyEmailURLWithUUID(environment string, secure bool) (url, uuid string) {
//
// 	uuid = GenerateUUIDType4(false)
// 	url = fmt.Sprintf("%v?uuid=%v", GenerateVerifyEmailURL(environment, secure), uuid)
//
// 	return
// }

// GenerateVerifyEmailURLWithUUIDUsername - return the url, uuid and the username for the Verify Email.
// func GenerateVerifyEmailURLWithUUIDUsername(username, environment string, secure bool) (url, uuid string) {
//
// 	uuid = GenerateUUIDType4(false)
// 	url = fmt.Sprintf("%v?%v=%v&%v=%v", GenerateVerifyEmailURL(environment, secure), constants.FN_UUID, uuid, constants.FN_USERNAME, username)
//
// 	return
// }

// GenerateVerifyEmailURL - return the url.
// func GenerateVerifyEmailURL(environment string, secure bool) (url string) {
//
// 	url = fmt.Sprintf("%v/%v", GenerateURL(environment, secure), constants.ENDPOINT_VERIFY_EMAIL)
//
// 	return
// }

// GetDate - return the current date in YYYY-MM-DD format
//
//	Customer Message: None
//	Errors: None
//	Verification: None
// func GetDate() string {
// 	return time.Now().Format("2006-01-02")
// }

// GetTime - return the current time in HH-mm-ss.00000 format, where hour is in military time.
//
//	Customer Message: None
//	Errors: None
//	Verification: None
// func GetTime() string {
// 	return time.Now().Format("15-04-05.00000")
// }

// getType
// func getType(myVar interface{}) (myType string) {
//
// 	if t := reflect.TypeOf(myVar); t.Kind() == reflect.Ptr {
// 		myType = "*" + t.Elem().Name()
// 	} else {
// 		myType = t.Name()
// 	}
//
// 	return
// }

// PenniesToFloat
// func PenniesToFloat(pennies int64) float64 {
//
// 	return float64(pennies) / 100
// }

// PrintAndDie - is exported for access in other packages. Not going to test
// func PrintAndDie(msg string) {
//
// 	_, _ = fmt.Fprintln(os.Stderr, msg)
// 	os.Exit(1)
//
// }

// PrependWorkingDirectory - will add the working directory and return the result.
// if the filename first character is a /, the passed value will be returned
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func PrependWorkingDirectory(filename string) (fqn string) {

	var (
		tWorkingDirectory, _ = os.Getwd()
	)

	if strings.HasPrefix(filename, rcv.FORWARD_SLASH) {
		fqn = filename
		return
	}

	return fmt.Sprintf("%v/%v", tWorkingDirectory, filename)
}

// printDashLine - will output a given number of dashed lines based on the outputMode.
// The default is to output to the log
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func printDashLines(lines int, outputMode string) {
//
// 	for i := 0; i < lines; i++ {
// 		if strings.ToLower(outputMode) == rcv.MODE_OUTPUT_DISPLAY {
// 			fmt.Println("------------------------------------------")
// 		} else {
// 			log.Println("------------------------------------------")
// 		}
// 	}
// }

// PrintLinesAtStartOfRequest - will output dashed lines when a new request arrives.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func PrintLinesAtStartOfRequest(lines int, outputMode string) {
//
// 	printDashLines(lines, outputMode)
//
// }

// RedirectLogOutput - will send log records to the config file log directory or the default directory.
// The log file name uses this format: 2006-01-02 15:04:05.000 Z0700. All spaces, colons, and periods
// are replaced with underscores.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func RedirectLogOutput(logDirectory string) (logFileHandlerPtr *os.File, logFQN string, errorInfo cpi.ErrorInfo) {

	var (
		tLogFileName string
	)

	tDateTime := time.Now().Format("2006-01-02 15:04:05.000 Z0700")
	tLogFileName = strings.Replace(strings.Replace(strings.Replace(tDateTime, rcv.SPACES_ONE, rcv.UNDERSCORE, -1), rcv.COLON, rcv.UNDERSCORE, -1), rcv.PERIOD, rcv.UNDERSCORE, -1)
	logFQN = logDirectory + "/" + tLogFileName

	// Set log file output
	if logFileHandlerPtr, errorInfo.Error = os.OpenFile(logFQN, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666); errorInfo.Error != nil {
		errorInfo = cpi.NewErrorInfo(errorInfo.Error, fmt.Sprintf("%v%v", rcv.TXT_FILENAME, logFQN))
		return
	}

	log.Printf("The log file has been redirected to %v.", logFQN)
	log.SetOutput(logFileHandlerPtr)

	return
}

// RemovePidFile
// func RemovePidFile(pidFQN string) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if coreValidators.DoesFileExist(pidFQN) {
// 		if errorInfo.Error = os.Remove(pidFQN); errorInfo.Error != nil {
// 			errorInfo.Error = errors.New(fmt.Sprintf("The removal of the pid file (%v) failed: %v", pidFQN, errorInfo.Error.Error()))
// 			log.Println(errorInfo.Error.Error())
// 		}
// 	} else {
// 		errorInfo.Error = errors.New(fmt.Sprintf("The pid file (%v) does not exist.", pidFQN))
// 		log.Println(errorInfo.Error.Error())
// 	}
//
// 	return
// }

// SendReply
// func SendReply(functionName string, jsonReply []byte, msg *nats.Msg) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if errorInfo.Error = msg.Respond(jsonReply); errorInfo.Error != nil {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Failed to create reply for the executed %v request! ERROR: %v", functionName, errorInfo.Error.Error()))
// 		log.Println(errorInfo.Error.Error())
// 		// ToDo Handle Error & Notification
// 	}
//
// 	return
// }

// UnmarshalMessageData - reads the message data into the pointer. The second argument must be a pointer. If you pass something else, the unmarshal will fail.
// func UnmarshalMessageData(msg *nats.Msg, requestPtr any) (errorInfo cpi.ErrorInfo) {
//
// 	if errorInfo.Error = json.Unmarshal(msg.Data, requestPtr); errorInfo.Error != nil {
// 		errorInfo.Error = errors.New(fmt.Sprintf("ERROR: Unable to unmarshal message data: '%v' Request: %v", string(msg.Data), getType(requestPtr)))
// 		log.Println(errorInfo.Error.Error())
// 		// ToDo Handle Error & Notification
// 	}
//
// 	return
// }

// WriteFile
// func WriteFile(fullQualifiedName string, fileData []byte, filePermissions os.FileMode) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if errorInfo.Error = os.WriteFile(fullQualifiedName, fileData, filePermissions); errorInfo.Error != nil {
// 		errorInfo.Error = errors.New(fmt.Sprintf("%vERROR: The creation of the file (%v) failed: %v", constants.COLOR_RED, fullQualifiedName, errorInfo.Error.Error()))
// 		log.Println(errorInfo.Error.Error())
// 	} else {
// 		log.Printf("Wrote %v to the file system.", fullQualifiedName)
// 	}
//
// 	return
// }

// WritePidFile
// func WritePidFile(directory string) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tPIDFileName = directory + constants.PID_FILENAME
// 	)
//
// 	if errorInfo = coreValidators.ValidateDirectory(directory); errorInfo.Error == nil {
// 		pidStr := strconv.Itoa(os.Getpid())
// 		if errorInfo = WriteFile(tPIDFileName, []byte(pidStr), 0766); errorInfo.Error == nil {
// 			log.Printf("Wrote pid file (%v) to the file system. It must be removed if the system exits unexspectedly.", tPIDFileName)
// 		}
// 	}
//
// 	return
// }
