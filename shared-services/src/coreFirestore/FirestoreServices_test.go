// Package coreFirestore
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
package coreFirestore

import (
	"fmt"
	"runtime"
	"testing"

	"albert/constants"
	"albert/core/coreError"
	"albert/core/coreFirebase"
	"albert/core/coreHelpers"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

var (
	tFireTestNameValue = map[any]interface{}{
		constants.FN_REQUESTOR_ID: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
		constants.FN_EMAIL:        constants.TEST_STRING,
	}
	tFireTestNameValueSubCollection = map[any]interface{}{
		constants.FN_INSTITUTION_NAME:   constants.TEST_INSTITUTION_CHASE,
		constants.FN_PLAID_ACCESS_TOKEN: constants.TEST_STRING,
	}
)

func TestBuildFirestoreUpdate(tPtr *testing.T) {

	var (
		errorInfo          coreError.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
		tNameValues        = make(map[any]interface{})
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		if _, errorInfo = BuildFirestoreUpdate(tNameValues); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		tNameValues["Field_1"] = "Value_1"
		tNameValues["Field_2"] = "Value_2"
		tNameValues["Field_3"] = "Value_3"
		if _, errorInfo = BuildFirestoreUpdate(tNameValues); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}

	})
}

func TestDoesDocumentExist(tPtr *testing.T) {

	var (
		tDocumentReferencePtr *firestore.DocumentRef
		tFirestoreClientPtr   *firestore.Client
		tFunction, _, _, _    = runtime.Caller(0)
		tFunctionName         = runtime.FuncForPC(tFunction).Name()
	)

	tFirestoreClientPtr = getTestFirestoreConnection()
	buildTestDocuments(tFirestoreClientPtr, 1)
	tDocumentReferencePtr, _ = getDocumentRef(tFirestoreClientPtr, constants.TEST_DATASTORE, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 0))

	tPtr.Run(tFunctionName, func(t *testing.T) {
		// Document exists
		if doesDocumentExist(tDocumentReferencePtr) == false {
			tPtr.Errorf("%v Failed: Was expecting true and got false.", tFunctionName)
		}
		// Document doesn't exist
		RemoveDocument(tFirestoreClientPtr, constants.TEST_DATASTORE, NameValueQuery{
			FieldName:  constants.FN_REQUESTOR_ID,
			FieldValue: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
		})
		if doesDocumentExist(tDocumentReferencePtr) {
			tPtr.Errorf("%v Failed: Was expecting an false and got true.", tFunctionName)
		}
	})
}

func TestFindDocument(tPtr *testing.T) {

	type arguments struct {
		createDocument bool
		dataStore      string
		nameValues1    NameValueQuery
		nameValues2    NameValueQuery
	}

	var (
		tFirestoreClientPtr *firestore.Client
		errorInfo           coreError.ErrorInfo
		gotError            bool
		tNameValues         = make(map[string]interface{})
	)

	tNameValues[constants.FN_REQUESTOR_ID] = constants.TEST_USERNAME_SAVUP_REQUESTOR_ID
	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful - single param!",
			arguments: arguments{
				createDocument: true,
				dataStore:      constants.TEST_DATASTORE,
				nameValues1: NameValueQuery{
					FieldName:  constants.FN_REQUESTOR_ID,
					FieldValue: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				},
			},
			wantError: false,
		},
		{
			name: "Positive Case: Successful - double param!",
			arguments: arguments{
				createDocument: true,
				dataStore:      constants.TEST_DATASTORE,
				nameValues1: NameValueQuery{
					FieldName:  constants.FN_REQUESTOR_ID,
					FieldValue: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				},
				nameValues2: NameValueQuery{
					FieldName:  constants.FN_EMAIL,
					FieldValue: constants.TEST_STRING,
				},
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing datastore!",
			arguments: arguments{
				createDocument: true,
				dataStore:      constants.EMPTY,
				nameValues1: NameValueQuery{
					FieldName:  constants.FN_REQUESTOR_ID,
					FieldValue: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing name/values field name value!",
			arguments: arguments{
				createDocument: true,
				dataStore:      constants.TEST_DATASTORE,
				nameValues1: NameValueQuery{
					FieldName:  constants.EMPTY,
					FieldValue: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing name/values!",
			arguments: arguments{
				createDocument: true,
				dataStore:      constants.TEST_DATASTORE,
				nameValues1:    NameValueQuery{},
			},
			wantError: true,
		},
	}

	tFirestoreClientPtr = getTestFirestoreConnection()
	buildTestDocuments(tFirestoreClientPtr, 1)

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if ts.arguments.nameValues2.FieldName == constants.EMPTY {
				_, _, errorInfo = FindDocument(tFirestoreClientPtr, ts.arguments.dataStore, ts.arguments.nameValues1)
			} else {
				_, _, errorInfo = FindDocument(tFirestoreClientPtr, ts.arguments.dataStore, ts.arguments.nameValues1, ts.arguments.nameValues2)
			}
			if errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo.Error.Error())
			}
		})
	}

	removeTestDocument(tFirestoreClientPtr, 1)
}

func TestGetAllDocuments(tPtr *testing.T) {

	var (
		tFirestoreClientPtr *firestore.Client
		tFunction, _, _, _  = runtime.Caller(0)
		tFunctionName       = runtime.FuncForPC(tFunction).Name()
		errorInfo           coreError.ErrorInfo
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		if _, errorInfo = GetAllDocuments(tFirestoreClientPtr, constants.TEST_DATASTORE); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		//
		tFirestoreClientPtr = getTestFirestoreConnection()
		if _, errorInfo = GetAllDocuments(tFirestoreClientPtr, constants.TEST_DATASTORE); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		//
		buildTestDocuments(tFirestoreClientPtr, 6)
		if _, errorInfo = GetAllDocuments(tFirestoreClientPtr, constants.TEST_DATASTORE); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
	})

	removeTestDocument(tFirestoreClientPtr, 6)
}

func TestGetAllDocumentsWhere(tPtr *testing.T) {

	var (
		tDocuments          []*firestore.DocumentSnapshot
		tFirestoreClientPtr *firestore.Client
		tFunction, _, _, _  = runtime.Caller(0)
		tFunctionName       = runtime.FuncForPC(tFunction).Name()
		errorInfo           coreError.ErrorInfo
	)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		//  No Pointer
		if _, errorInfo = GetAllDocumentsWhere(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.FN_REQUESTOR_ID, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		//
		// No Data
		tFirestoreClientPtr = getTestFirestoreConnection()
		if tDocuments, errorInfo = GetAllDocumentsWhere(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.FN_REQUESTOR_ID, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID); errorInfo.Error != nil && len(tDocuments) > 0 {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		//
		// Success
		buildTestDocuments(tFirestoreClientPtr, 6)
		if tDocuments, errorInfo = GetAllDocumentsWhere(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.FN_REQUESTOR_ID, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID); errorInfo.Error != nil && len(tDocuments) == 0 {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
	})

	removeTestDocument(tFirestoreClientPtr, 6)
}

func TestGetDocumentById(tPtr *testing.T) {

	var (
		tFirestoreClientPtr *firestore.Client
		tFunction, _, _, _  = runtime.Caller(0)
		tFunctionName       = runtime.FuncForPC(tFunction).Name()
		errorInfo           coreError.ErrorInfo
	)

	tFirestoreClientPtr = getTestFirestoreConnection()
	buildTestDocuments(tFirestoreClientPtr, 1)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		// Successful
		if _, errorInfo = GetDocumentById(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.TEST_DOCUMENT_ID_0); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		_ = RemoveDocumentById(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.TEST_DOCUMENT_ID_0)
		// Not found
		if _, errorInfo = GetDocumentById(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.TEST_DOCUMENT_ID_0); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		// Missing Datastore name
		if _, errorInfo = GetDocumentById(tFirestoreClientPtr, constants.EMPTY, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		// Missing document id
		if _, errorInfo = GetDocumentById(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.EMPTY); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})

	removeTestDocument(tFirestoreClientPtr, 1)
}

func TestGetDocumentRef(tPtr *testing.T) {

	var (
		tFirestoreClientPtr *firestore.Client
		tFunction, _, _, _  = runtime.Caller(0)
		tFunctionName       = runtime.FuncForPC(tFunction).Name()
		errorInfo           coreError.ErrorInfo
	)

	tFirestoreClientPtr = getTestFirestoreConnection()
	buildTestDocuments(tFirestoreClientPtr, 1)

	tPtr.Run(tFunctionName, func(t *testing.T) {
		//  Found
		if _, errorInfo = getDocumentRef(tFirestoreClientPtr, constants.TEST_DATASTORE, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 0)); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, "nil", errorInfo.Error.Error())
		}
		//  Not Found
		_ = RemoveDocument(tFirestoreClientPtr, constants.TEST_DATASTORE, NameValueQuery{
			FieldName:  constants.FN_REQUESTOR_ID,
			FieldValue: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
		})
		if _, errorInfo = getDocumentRef(tFirestoreClientPtr, constants.TEST_DATASTORE, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 0)); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})

	removeTestDocument(tFirestoreClientPtr, 1)
}

func TestGetDocumentIdsWithSubCollections(tPtr *testing.T) {

	type arguments struct {
		datastoreName     string
		requestorId       string
		subCollectionName string
	}

	var (
		errorInfo          coreError.ErrorInfo
		gotError           bool
		tFirebase          coreHelpers.FirebaseFirestoreHelper
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
				datastoreName:     constants.DATASTORE_USER_INSTITUTIONS,
				requestorId:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollectionName: constants.COLLECTION_INSTITUTIONS,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing datastore!",
			arguments: arguments{
				datastoreName:     constants.EMPTY,
				requestorId:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollectionName: constants.COLLECTION_INSTITUTIONS,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing requestor id!",
			arguments: arguments{
				datastoreName:     constants.DATASTORE_USER_INSTITUTIONS,
				requestorId:       constants.EMPTY,
				subCollectionName: constants.COLLECTION_INSTITUTIONS,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing sub collection!",
			arguments: arguments{
				datastoreName:     constants.DATASTORE_USER_INSTITUTIONS,
				requestorId:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollectionName: constants.EMPTY,
			},
			wantError: true,
		},
	}

	tFirebase.AppPtr, _, _ = coreFirebase.GetFirebaseAppAuthConnection(constants.TEST_FIREBASE_CREDENTIALS)
	tFirebase.FirestoreClientPtr, _ = GetFirestoreClientConnection(tFirebase.AppPtr)

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _, errorInfo = GetDocumentIdsWithSubCollections(tFirebase.FirestoreClientPtr, ts.arguments.datastoreName, ts.arguments.requestorId, ts.arguments.subCollectionName); errorInfo.Error == nil {
				gotError = false
			} else {
				gotError = true
			}
			if gotError != ts.wantError {
				tPtr.Error(tFunctionName, ts.name, errorInfo.Error.Error())
			}
		})
	}
}

func TestGetDocumentFromSubCollectionByDocumentId(tPtr *testing.T) {

	type arguments struct {
		datastoreName     string
		requestorId       string
		subCollectionName string
		documentId        string
	}

	var (
		errorInfo          coreError.ErrorInfo
		gotError           bool
		tFirebase          coreHelpers.FirebaseFirestoreHelper
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
				datastoreName:     constants.DATASTORE_USER_INSTITUTIONS,
				requestorId:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollectionName: constants.COLLECTION_INSTITUTIONS,
				documentId:        constants.TEST_INSTITUTION_CHASE,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing datastore!",
			arguments: arguments{
				datastoreName:     constants.EMPTY,
				requestorId:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollectionName: constants.COLLECTION_INSTITUTIONS,
				documentId:        constants.TEST_INSTITUTION_CHASE_CLONE,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing requestor id!",
			arguments: arguments{
				datastoreName:     constants.DATASTORE_USER_INSTITUTIONS,
				requestorId:       constants.EMPTY,
				subCollectionName: constants.COLLECTION_INSTITUTIONS,
				documentId:        constants.TEST_INSTITUTION_CHASE_CLONE,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing sub collection!",
			arguments: arguments{
				datastoreName:     constants.DATASTORE_USER_INSTITUTIONS,
				requestorId:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollectionName: constants.EMPTY,
				documentId:        constants.TEST_INSTITUTION_CHASE_CLONE,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing document id!",
			arguments: arguments{
				datastoreName:     constants.DATASTORE_USER_INSTITUTIONS,
				requestorId:       constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollectionName: constants.COLLECTION_INSTITUTIONS,
				documentId:        constants.EMPTY,
			},
			wantError: true,
		},
	}

	tFirebase.AppPtr, _, _ = coreFirebase.GetFirebaseAppAuthConnection(constants.TEST_FIREBASE_CREDENTIALS)
	tFirebase.FirestoreClientPtr, _ = GetFirestoreClientConnection(tFirebase.AppPtr)
	_ = SetDocumentWithSubCollection(tFirebase.FirestoreClientPtr, constants.DATASTORE_USER_INSTITUTIONS, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.COLLECTION_INSTITUTIONS, constants.TEST_INSTITUTION_CHASE, tFireTestNameValueSubCollection)

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if _, errorInfo = GetDocumentFromSubCollectionByDocumentId(tFirebase.FirestoreClientPtr, ts.arguments.datastoreName, ts.arguments.requestorId, ts.arguments.subCollectionName, ts.arguments.documentId); errorInfo.Error == nil {
				gotError = false
			} else {
				gotError = true
			}
			if gotError != ts.wantError {
				tPtr.Error(tFunctionName, ts.name, errorInfo.Error.Error())
			}
		})
	}

	_ = RemoveDocumentFromSubCollectionByDocumentId(tFirebase.FirestoreClientPtr, constants.DATASTORE_USER_INSTITUTIONS, constants.TEST_USERNAME_SAVUP_REQUESTOR_ID, constants.COLLECTION_INSTITUTIONS, constants.TEST_INSTITUTION_CHASE)
}

func TestGetFirestoreClientConnection(tPtr *testing.T) {

	var (
		errorInfo          coreError.ErrorInfo
		tAppPtr            *firebase.App
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tAppPtr, _ = coreFirebase.NewFirebaseApp(constants.TEST_FIREBASE_CREDENTIALS)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Test connection with good Firebase app pointer
		if _, errorInfo = GetFirestoreClientConnection(tAppPtr); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Firebase app was not created.", tFunctionName)
		}
		// Test connection with nil firebase app pointer
		if _, errorInfo = GetFirestoreClientConnection(nil); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Firestore connection was established.", tFunctionName)
		}
	})
}

func TestRemoveDocument(tPtr *testing.T) {

	type arguments struct {
		tDataStore       string
		tQueryParameters NameValueQuery
	}

	var (
		errorInfo           coreError.ErrorInfo
		gotError            bool
		tFirestoreClientPtr *firestore.Client
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Success",
			arguments: arguments{
				tDataStore: constants.TEST_DATASTORE,
				tQueryParameters: NameValueQuery{
					FieldName:  constants.TEST_FIELD_NAME,
					FieldValue: constants.TEST_FIELD_VALUE,
				},
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing datastore",
			arguments: arguments{
				tDataStore: constants.EMPTY,
				tQueryParameters: NameValueQuery{
					FieldName:  constants.TEST_FIELD_NAME,
					FieldValue: constants.TEST_FIELD_VALUE,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing Field Nane",
			arguments: arguments{
				tDataStore: constants.TEST_DATASTORE,
				tQueryParameters: NameValueQuery{
					FieldName:  constants.EMPTY,
					FieldValue: constants.TEST_FIELD_VALUE,
				},
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing Field Value",
			arguments: arguments{
				tDataStore: constants.TEST_DATASTORE,
				tQueryParameters: NameValueQuery{
					FieldName:  constants.TEST_FIELD_NAME,
					FieldValue: constants.EMPTY,
				},
			},
			wantError: true,
		},
	}

	tFirestoreClientPtr = getTestFirestoreConnection()

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = RemoveDocument(tFirestoreClientPtr, ts.arguments.tDataStore, ts.arguments.tQueryParameters); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}

}

func TestRemoveDocumentById(tPtr *testing.T) {

	type arguments struct {
		tDataStore  string
		tDocumentId string
	}

	var (
		tFirestoreClientPtr *firestore.Client
		errorInfo           coreError.ErrorInfo
		gotError            bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Success",
			arguments: arguments{
				tDataStore:  constants.TEST_DATASTORE,
				tDocumentId: constants.TEST_DOCUMENT_ID_0,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Document not found",
			arguments: arguments{
				tDataStore:  constants.TEST_DATASTORE,
				tDocumentId: constants.TEST_DOCUMENT_ID_0,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing datastore",
			arguments: arguments{
				tDataStore:  constants.EMPTY,
				tDocumentId: constants.TEST_DOCUMENT_ID_0,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing Document Id",
			arguments: arguments{
				tDataStore:  constants.TEST_DATASTORE,
				tDocumentId: constants.EMPTY,
			},
			wantError: true,
		},
	}

	tFirestoreClientPtr = getTestFirestoreConnection()
	buildTestDocuments(tFirestoreClientPtr, 1)

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = RemoveDocumentById(tFirestoreClientPtr, ts.arguments.tDataStore, ts.arguments.tDocumentId); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}

	removeTestDocument(tFirestoreClientPtr, 1)
}

func TestRemoveDocumentFromSubCollection(tPtr *testing.T) {

	type arguments struct {
		dataStore        string
		parentDocumentId string
		subCollection    string
		documentId       string
	}

	var (
		tFirestoreClientPtr *firestore.Client
		errorInfo           coreError.ErrorInfo
		gotError            bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Success",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 0),
				subCollection:    constants.TEST_DATASTORE_SUBCOLLECTION,
				documentId:       fmt.Sprintf(constants.TEST_USER_REQUESTOR_ID_F, 0),
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing datastore",
			arguments: arguments{
				dataStore:        constants.EMPTY,
				parentDocumentId: fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 1),
				subCollection:    constants.TEST_DATASTORE_SUBCOLLECTION,
				documentId:       fmt.Sprintf(constants.TEST_USER_REQUESTOR_ID_F, 1),
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing parent document id",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: constants.EMPTY,
				subCollection:    constants.TEST_DATASTORE_SUBCOLLECTION,
				documentId:       fmt.Sprintf(constants.TEST_USER_REQUESTOR_ID_F, 1),
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing sub-collection",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 2),
				subCollection:    constants.EMPTY,
				documentId:       fmt.Sprintf(constants.TEST_USER_REQUESTOR_ID_F, 1),
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing sub-collection",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 3),
				subCollection:    constants.TEST_DATASTORE_SUBCOLLECTION,
				documentId:       constants.EMPTY,
			},
			wantError: true,
		},
	}

	tFirestoreClientPtr = getTestFirestoreConnection()
	buildTestDocumentsWithSubCollection(tFirestoreClientPtr, 1)

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = RemoveDocumentFromSubCollectionByDocumentId(tFirestoreClientPtr, ts.arguments.dataStore, ts.arguments.parentDocumentId, ts.arguments.subCollection, ts.arguments.documentId); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}
}

func TestSetDocument(tPtr *testing.T) {

	type arguments struct {
		dataStore  string
		documentId string
		nameValues map[any]interface{}
	}

	var (
		errorInfo           coreError.ErrorInfo
		gotError            bool
		tFirestoreClientPtr *firestore.Client
		tNameValues         = make(map[any]interface{})
	)

	tNameValues[constants.FN_REQUESTOR_ID] = constants.TEST_USERNAME_SAVUP_REQUESTOR_ID
	tNameValues[constants.FN_EMAIL] = constants.TEST_STRING

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				dataStore:  constants.TEST_DATASTORE,
				documentId: constants.TEST_DOCUMENT_ID_0,
				nameValues: tNameValues,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing datastore!",
			arguments: arguments{
				dataStore:  constants.EMPTY,
				documentId: constants.TEST_DOCUMENT_ID_0,
				nameValues: tNameValues,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing document id!",
			arguments: arguments{
				dataStore:  constants.TEST_DATASTORE,
				documentId: constants.EMPTY,
				nameValues: tNameValues,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing name/values!",
			arguments: arguments{
				dataStore:  constants.TEST_DATASTORE,
				documentId: constants.EMPTY,
				nameValues: nil,
			},
			wantError: true,
		},
	}

	tFirestoreClientPtr = getTestFirestoreConnection()

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = SetDocument(tFirestoreClientPtr, ts.arguments.dataStore, ts.arguments.documentId, ts.arguments.nameValues); errorInfo.Error != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
		})
	}

	_ = RemoveDocumentById(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.TEST_DOCUMENT_ID_0)
}

func TestSetDocumentWithSubCollection(tPtr *testing.T) {

	type arguments struct {
		dataStore        string
		parentDocumentId string
		subCollection    string
		documentId       string
	}

	var (
		tFirestoreClientPtr *firestore.Client
		errorInfo           coreError.ErrorInfo
		gotError            bool
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Positive Case: Successful!",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollection:    constants.COLLECTION_INSTITUTIONS,
				documentId:       constants.TEST_DOCUMENT_ID_0,
			},
			wantError: false,
		},
		{
			name: "Negative Case: Missing datastore!",
			arguments: arguments{
				dataStore:        constants.EMPTY,
				parentDocumentId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollection:    constants.COLLECTION_INSTITUTIONS,
				documentId:       constants.TEST_DOCUMENT_ID_0,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing document id!",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollection:    constants.COLLECTION_INSTITUTIONS,
				documentId:       constants.EMPTY,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing parent document id!",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: constants.EMPTY,
				subCollection:    constants.COLLECTION_INSTITUTIONS,
				documentId:       constants.EMPTY,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing sub-collection name!",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollection:    constants.EMPTY,
				documentId:       constants.EMPTY,
			},
			wantError: true,
		},
		{
			name: "Negative Case: Missing name/values!",
			arguments: arguments{
				dataStore:        constants.TEST_DATASTORE,
				parentDocumentId: constants.TEST_USERNAME_SAVUP_REQUESTOR_ID,
				subCollection:    constants.COLLECTION_INSTITUTIONS,
				documentId:       constants.TEST_DOCUMENT_ID_0,
			},
			wantError: true,
		},
	}

	tFirestoreClientPtr = getTestFirestoreConnection()

	tNameValues := make(map[any]interface{})
	tNameValues[constants.FN_REQUESTOR_ID] = constants.TEST_USERNAME_SAVUP_REQUESTOR_ID
	tNameValues[constants.FN_EMAIL] = constants.TEST_STRING

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			if errorInfo = SetDocumentWithSubCollection(tFirestoreClientPtr, ts.arguments.dataStore, ts.arguments.parentDocumentId, ts.arguments.subCollection, ts.arguments.documentId, tNameValues); errorInfo.Error == nil {
				gotError = false
			} else {
				gotError = true
			}
			if gotError != ts.wantError {
				tPtr.Error(ts.name)
				tPtr.Error(errorInfo)
			}
			_ = RemoveDocumentFromSubCollectionByDocumentId(tFirestoreClientPtr, ts.arguments.dataStore, ts.arguments.parentDocumentId, ts.arguments.subCollection, ts.arguments.documentId)
		})
	}
}

func TestUpdateDocument(tPtr *testing.T) {

	var (
		errorInfo           coreError.ErrorInfo
		tFirestoreClientPtr *firestore.Client
		tFunction, _, _, _  = runtime.Caller(0)
		tFunctionName       = runtime.FuncForPC(tFunction).Name()
		tNameValues         = make(map[any]interface{})
	)

	tFirestoreClientPtr = getTestFirestoreConnection()
	buildTestDocuments(tFirestoreClientPtr, 1)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		if errorInfo = UpdateDocument(tFirestoreClientPtr, constants.EMPTY, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 0), tNameValues); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		if errorInfo = UpdateDocument(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.EMPTY, tNameValues); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		if errorInfo = UpdateDocument(tFirestoreClientPtr, constants.TEST_DATASTORE, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 0), tNameValues); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		tNameValues["Field_1"] = "Value_1"
		tNameValues["Field_2"] = "Value_2"
		tNameValues["Field_3"] = "Value_3"
		if errorInfo = UpdateDocument(tFirestoreClientPtr, constants.TEST_DATASTORE, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, 0), tNameValues); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: The Update was not successful! Error: '%v'", tFunctionName, errorInfo.Error.Error())
		}
	})

	removeTestDocument(tFirestoreClientPtr, 1)
}

func TestUpdateDocumentFromSubCollectionByDocumentId(tPtr *testing.T) {

	var (
		errorInfo           coreError.ErrorInfo
		tFirestoreClientPtr *firestore.Client
		tFunction, _, _, _  = runtime.Caller(0)
		tFunctionName       = runtime.FuncForPC(tFunction).Name()
		tFieldPath          []string
		tUpdates            []firestore.Update
	)

	tFirestoreClientPtr = getTestFirestoreConnection()
	buildTestDocumentsWithSubCollection(tFirestoreClientPtr, 1)
	tFieldPath = append(tFieldPath, constants.FN_PLAID_ACCOUNTS)
	tFieldPath = append(tFieldPath, "PxEENJbqvGFZRyj6b6MXugrDjgevQaHQRQ9oa")
	tFieldPath = append(tFieldPath, constants.FN_BALANCE)
	tUpdates = []firestore.Update{{
		FieldPath: tFieldPath,
		Value:     123456,
	}}

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Successful
		if errorInfo = UpdateDocumentFromSubCollectionByDocumentId(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.TEST_DOCUMENT_ID_0, constants.TEST_DATASTORE_SUBCOLLECTION, fmt.Sprintf(constants.TEST_USER_REQUESTOR_ID_F, 0), tUpdates); errorInfo.Error != nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
		// Record not found
		//
		RemoveDocumentFromSubCollectionByDocumentId(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.TEST_DOCUMENT_ID_0, constants.TEST_DATASTORE_SUBCOLLECTION, fmt.Sprintf(constants.TEST_USER_REQUESTOR_ID_F, 0))
		//
		if errorInfo = UpdateDocumentFromSubCollectionByDocumentId(tFirestoreClientPtr, constants.TEST_DATASTORE, constants.TEST_DOCUMENT_ID_0, constants.TEST_DATASTORE_SUBCOLLECTION, fmt.Sprintf(constants.TEST_USER_REQUESTOR_ID_F, 0), tUpdates); errorInfo.Error == nil {
			tPtr.Errorf("%v Failed: Was expecting an err of %v but got %v.", tFunctionName, errorInfo.Error.Error(), "nil")
		}
	})
}

func getTestFirestoreConnection() (firestoreClientPtr *firestore.Client) {

	var (
		tAppPtr *firebase.App
	)

	tAppPtr, _ = coreFirebase.NewFirebaseApp(constants.TEST_FIREBASE_CREDENTIALS)
	firestoreClientPtr, _ = GetFirestoreClientConnection(tAppPtr)

	return
}

// buildTestDocuments - creates 1 to 10 test documents with the number starting at 0 and going to count - 1. The document is will be TEST_DOCUMENT_ID_F where is 0 to 9.
func buildTestDocuments(firestoreClientPtr *firestore.Client, count int) {

	if count > 10 {
		count = 10
	} else if count < 1 {
		count = 1
	}
	for i := 0; i < count; i++ {
		_ = SetDocument(firestoreClientPtr, constants.TEST_DATASTORE, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, i), tFireTestNameValue)
	}

}

func buildTestDocumentsWithSubCollection(firestoreClientPtr *firestore.Client, count int) {

	for i := 0; i < count; i++ {
		_ = SetDocumentWithSubCollection(firestoreClientPtr, constants.TEST_DATASTORE, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, i), constants.TEST_DATASTORE_SUBCOLLECTION, fmt.Sprintf(constants.TEST_USER_REQUESTOR_ID_F, i), tFireTestNameValue)
	}

}

func removeTestDocument(firestoreClientPtr *firestore.Client, count int) {

	for i := 0; i < count+1; i++ {
		RemoveDocumentById(firestoreClientPtr, constants.TEST_DATASTORE, fmt.Sprintf(constants.TEST_DOCUMENT_ID_F, i))
	}

}
