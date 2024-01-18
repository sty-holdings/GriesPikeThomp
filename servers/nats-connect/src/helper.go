// Package src
/*
General description of the purpose of the go file.

RESTRICTIONS:
    AWS functions:
    * Program must have access to a .aws/credentials file in the default location.
    * This will only access system parameters that start with '/sote' (ROOTPATH).
    * {Enter other restrictions here for AWS

    {Other catagories of restrictions}
    * {List of restrictions for the catagory

NOTES:
    {Enter any additional notes that you believe will help the next developer.}

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
package src

import (
	"encoding/json"
	"fmt"
)

const (
// Add Constants to the constants.go file
)

// Add types to the types.go file

var (
// Add Variables here for the file (Remember, they are global)
)

// GenerateConfigFileSkeleton will output to the console a skeleton file with notes.
func GenerateConfigFileSkeleton() {

	var (
		opts   Options
		output []byte
	)

	output, _ = json.Marshal(opts)
	fmt.Println("\nWhen '-g' is used all other program arguments are ignored.")
	fmt.Printf("\nSavUp Config file Skeleton: \n\t%v\n", string(output))
	fmt.Println()
	fmt.Println("NOTES: \n\t==> DEVELOPERS REMEMBER: DO NOT PUT SENSITIVE INFORMATION IN THE LOG!\n\t==> If you are not sure, " +
		"ask the CTO if information is classified as sensitive.")
	//
	fmt.Println("\tAuthenticator_Service is the name of the service provider being used for user login. \n\t\tThe default is FIREBASE and the other option is COGNITO. The value is case-insensitive.")
	//
	fmt.Println("\tAWS_Info_FQN is the fully qualified filename for the AWS project info. \n\t\tThere is no default and a value must be provided. The value is case-sensitive.")
	//
	fmt.Println("\tDebug_Mode is either 'true' or 'false'. \n\t\tThe default is false and a value must be provided. The value is case-sensitive.")
	//
	fmt.Println("\tEnvironment is one of the following, LOCAL, DEVELOPMENT, PRODUCTION. The default is LOCAL and the value is case-insensitive. " +
		"\n\t\tDEVELOPMENT should only be used for small teams of developers doing joint development or integration. \n\t\tPRODUCTION is where we make money.")
	//
	fmt.Println("\tFirebase_Project_Id is the project that contains the FireBase Firestore data structure. \n\t\tThere is no default and a value must be provided. The value is case-sensitive.")
	//
	fmt.Println("\tFirebase_Credentials_FQN is the fully qualified file name of the file used by FireBase to verify access and permissions. " +
		"\n\t\tThere is no default and the value must be provided. It is case-sensitive.")
	//
	fmt.Println("\tGCP_Credentials_FQN is the fully qualified file name of the file used by Google Cloud Platform to verify access and permissions. " +
		"\n\t\tThere is no default and the value must be provided. It is case-sensitive.")
	//
	fmt.Println("\tLog_File_Location is the root directory where logs will be written. The default is ~/styh/server/log and the value is case-sensitive." +
		"The log file name is the date/time the server was started.")
	//
	fmt.Println("\tMessage_Prefix is one of the following, SAVUPLOCAL, SAVUPDEV, or SAVUP " +
		"The default is SAVUPLOCAL and the value is case-insensitive. \n\t\tSAVUPDEV should only be used for small teams of developers doing joint development or integration. " +
		"\n\t\tSAVUP is where we make money.")
	//
	fmt.Println("\tNATS_Creds_FQN is the fully qualified file name of the file used by NATS to verify access and permissions to the NATS server. " +
		"\n\t\tThere is no default and the value must be provided. It is case-sensitive.")
	//
	fmt.Println("\tNATS_URL is the DNS name of the NATS server that this SavUp server will communicate with for messages. " +
		"\n\t\tThere is no default and the value must be provided. It is case-insensitive.")
	//
	fmt.Println("\tPID_File_Location is the root directory where server pid file will be written. The default is ~/styh/server/.run and the value is case-sensitive." +
		"The pid file name is server.pid.")
	//
	fmt.Println("\tPlaid_Key_FQN is the fully qualified file name of the file used by Plaid to verify access and permissions to the Plaid server. \n\t\tThere is no default and the valueis optional. " +
		"The value is case-sensitive.")
	//
	fmt.Println("\tPrivate_Key_FQN is the fully qualified filename for the private key used for JWT. \n\t\tThere is no default and a value must be provided. The value is case-sensitive.")
	//
	fmt.Println("\tSendGrid_Key_FQN is the fully qualified file name of the file used by SendGrid to verify access and permissions to the SendGrid email server. " +
		"\n\t\tThere is no default and the value must be provided. The value is case-sensitive.")
	//
	fmt.Println("\tSendGrid_Template_Ids is the template id found under Email API > Dynamic Templates on https://sendgrid.com. " +
		"\n\t\tThere is no default and the value must be provided. The value is case-sensitive. Valid keys are bank, transferIn, transferOut, and verify.")
	//
	fmt.Println("\tStripe_Key_FQN is the fully qualified file name of the file used by Stripe to verify access and permissions to the Stripe server. \n\t\tThere is no default and the value is optional. " +
		"The value is case-sensitive.")
	//
	fmt.Println("\tTLS is optional.")
	//
	fmt.Println("\t\tTLS_Certificate_FQN is the fully qualified file name of the certificate file used to connect to the NATS server. \n\t\tThere is no default. The value is case-sensitive.")
	//
	fmt.Println("\t\tTLS_Private_Key_FQN is the fully qualified file name of the private key file used to connect to the NATS server. \n\t\tThere is no default. The value is case-sensitive.")
	//
	fmt.Println("\t\tTLS_CABundle_FQN is the fully qualified file name of the CA Bundle file used to connect to the NATS server. \n\t\tThere is no default. The value is case-sensitive.")
	//
	fmt.Println("\tWeb_Root_Assets_URL is the root URL for web assets. \n\t\tThere is no default and a value must be provided. The value is case-insensitive.")
}
