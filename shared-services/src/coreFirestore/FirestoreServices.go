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
	"context"

	"cloud.google.com/go/firestore"
)

const (
	NOT_FOUND_MAYBE_CORRECT = "Getting the 'The document was found ' error maybe correct. Review code logic."
)

type NameValueQuery struct {
	FieldName  string
	FieldValue interface{}
}

var (
	CTXBackground = context.Background()
)

// BuildFirestoreUpdate - while the nameValues is a map[any], the function using a string assertion on the key.
// func BuildFirestoreUpdate(nameValues map[any]interface{}) (firestoreUpdateFields []firestore.Update, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tUpdate            firestore.Update
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if tFinding = coreValidators.AreMapKeysValuesPopulated(nameValues); tFinding == rcv.GOOD {
// 		for field, value := range nameValues {
// 			tUpdate.Path = field.(string)
// 			tUpdate.Value = value
// 			firestoreUpdateFields = append(firestoreUpdateFields, tUpdate)
// 		}
// 	} else {
// 		errorInfo.Error = cpi.GetMapKeyPopulatedError(tFinding)
// 	}
//
// 	return
// }

// DoesDocumentExist
func doesDocumentExist(documentReferencePtr *firestore.DocumentRef) bool {

	if _, err := documentReferencePtr.Get(CTXBackground); err == nil {
		return true
	}

	return false
}

// FindDocument - Returns an error for documents not found, but it doesn't print the error to the log.
//
//	Customer Messages: None
//	Errors: cpi.ErrRequiredArgumentMissing, cpi.ErrDocumentNotFound, cpi.ErrServiceFailedFIRESTORE
//	Verifications: None
// func FindDocument(firestoreClientPtr *firestore.Client, datastore string, queryParameters ...NameValueQuery) (found bool, documentSnapshotPtr *firestore.DocumentSnapshot, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tQuery             firestore.Query
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || len(queryParameters) < 1 {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 	} else {
// 		tQuery = firestoreClientPtr.Collection(datastore).Query
// 		for _, parameter := range queryParameters {
// 			if parameter.FieldName == rcv.EMPTY || parameter.FieldValue == rcv.EMPTY {
// 				errorInfo.FileName, errorInfo.ErrorLineNumber = cpi.GetFileLineNumber()
// 				errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 				cpi.PrintError(errorInfo)
// 				break
// 			} else {
// 				tQuery = tQuery.Where(parameter.FieldName, rcv.EQUALS, parameter.FieldValue)
// 			}
// 		}
// 	}
//
// 	if errorInfo.Error == nil {
// 		tDocuments := tQuery.Documents(CTXBackground)
// 		for {
// 			documentSnapshotPtr, errorInfo.Error = tDocuments.Next()
// 			if errorInfo.Error != nil {
// 				if errors.Is(errorInfo.Error, iterator.Done) {
// 					errorInfo.FileName, errorInfo.ErrorLineNumber = cpi.GetFileLineNumber()
// 					errorInfo.Error = cpi.ErrDocumentNotFound
// 					errorInfo.AdditionalInfo = NOT_FOUND_MAYBE_CORRECT
// 					cpi.PrintError(errorInfo)
// 					break
// 				} else {
// 					errorInfo.FileName, errorInfo.ErrorLineNumber = cpi.GetFileLineNumber()
// 					errorInfo.Error = cpi.ErrServiceFailedFIRESTORE
// 					cpi.PrintError(errorInfo)
// 					// todo handle error & notification
// 					break
// 				}
// 			}
// 			if len(documentSnapshotPtr.Data()) > 0 {
// 				found = true
// 				break
// 			}
// 		}
// 	}
//
// 	return
// }

// GetAllDocuments will return snapshot pointers to each document in the datastore.
// If no documents are found, the documents will have a count of zero.
//
//	Customer Messages: None
//	Errors: cpi.ErrRequiredArgumentMissing
//	Verifications: None
// func GetAllDocuments(firestoreClientPtr *firestore.Client, datastore string) (documents []*firestore.DocumentSnapshot, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tCollectionReferencePtr *firestore.CollectionRef
// 		tFunction, _, _, _      = runtime.Caller(0)
// 		tFunctionName           = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if firestoreClientPtr == nil || datastore == rcv.EMPTY {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 		errorInfo.AdditionalInfo = fmt.Sprintf("Firestore Client Pointer: %v Datastore: %v", firestoreClientPtr, datastore)
// 		cpi.PrintError(errorInfo)
// 	} else {
// 		tCollectionReferencePtr = firestoreClientPtr.Collection(datastore)
// 		documents, errorInfo.Error = tCollectionReferencePtr.Documents(CTXBackground).GetAll()
// 		if documents == nil && errorInfo.Error == nil {
// 			errorInfo.Error = cpi.ErrDocumentsNoneFound
// 		}
// 	}
//
// 	return
// }

// GetAllDocumentsWhere will return snapshot pointers to each document in the datastore that meet the where condition.
// If no documents are found, the documents will have a count of zero.
//
//	Customer Messages: None
//	Errors: cpi.ErrRequiredArgumentMissing, cpi.ErrDocumentsNoneFound, cpi.ErrServiceFailedFIRESTORE
//	Verifications: None
// func GetAllDocumentsWhere(firestoreClientPtr *firestore.Client, datastore, fieldName string, fieldValue interface{}) (documents []*firestore.DocumentSnapshot, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tQuery             firestore.Query
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if firestoreClientPtr == nil || datastore == rcv.EMPTY || fieldName == rcv.EMPTY || fieldValue == nil {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 		errorInfo.AdditionalInfo = fmt.Sprintf("Firestore Client Pointer: %v Datastore: %v Field Name: %v Field Value: %v", firestoreClientPtr, datastore, fieldName, fieldValue)
// 		cpi.PrintError(errorInfo)
// 	} else {
// 		tQuery = firestoreClientPtr.Collection(datastore).Where(fieldName, "==", fieldValue)
// 		if documents, errorInfo.Error = tQuery.Documents(CTXBackground).GetAll(); len(documents) == 0 {
// 			if errorInfo.Error == nil {
// 				errorInfo.AdditionalInfo = rcv.NOT_FOUND + rcv.IS_OK
// 				errorInfo.Error = cpi.ErrDocumentsNoneFound
// 				cpi.PrintError(errorInfo)
// 			} else {
// 				errorInfo.AdditionalInfo = errorInfo.Error.Error()
// 				errorInfo.Error = cpi.ErrServiceFailedFIRESTORE
// 				cpi.PrintError(errorInfo)
// 			}
// 		}
// 	}
//
// 	return
// }

// GetSomeDocumentsWhere provides snapshot pointers to documents in the datastore that meet the specified 'where' condition, limited by the record count and starting from the offset position.
// If no documents are found, the documents variable will have a zero length.
//
//	Customer Messages: None
//	Errors: cpi.ErrRequiredArgumentMissing
//	Verifications: None
// func GetSomeDocumentsWhere(firestoreClientPtr *firestore.Client, datastore, fieldName string, fieldValue interface{}, offset, recordCount int) (documents []*firestore.DocumentSnapshot, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tQuery             firestore.Query
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if firestoreClientPtr == nil || datastore == rcv.EMPTY || fieldName == rcv.EMPTY || fieldValue == nil {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 		errorInfo.AdditionalInfo = fmt.Sprintf("Firestore Client Pointer: %v Datastore: %v Field Name: %v Field Value: %v", firestoreClientPtr, datastore, fieldName, fieldValue)
// 		cpi.PrintError(errorInfo)
// 	} else {
// 		tQuery = firestoreClientPtr.Collection(datastore).Where(fieldName, rcv.EQUALS, fieldValue).Offset(offset).Limit(recordCount)
// 		documents, errorInfo.Error = tQuery.Documents(CTXBackground).GetAll()
// 	}
//
// 	return
// }

// GetDocumentById - will return a non-nil documentSnapshotPtr if the document is found.
// func GetDocumentById(firestoreClientPtr *firestore.Client, datastore string, documentId string) (documentSnapshotPtr *firestore.DocumentSnapshot, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if firestoreClientPtr == nil || datastore == rcv.EMPTY || documentId == rcv.EMPTY {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Firestore Client Pointer or Datastore: '%v' Document Id: '%v'", datastore, documentId))
// 	} else {
// 		if documentSnapshotPtr, errorInfo.Error = firestoreClientPtr.Doc(datastore + "/" + documentId).Get(CTXBackground); documentSnapshotPtr == nil || errorInfo.Error != nil {
// 			if strings.Contains(errorInfo.Error.Error(), rcv.NOT_FOUND) {
// 				errorInfo.Error = cpi.ErrDocumentNotFound
// 			}
// 			documentSnapshotPtr = nil
// 		}
// 	}
//
// 	return
// }

// getDocumentRef
// func getDocumentRef(firestoreClientPtr *firestore.Client, datastore, documentId string) (documentReferencePtr *firestore.DocumentRef, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || documentId == rcv.EMPTY {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' Document Id: '%v'", datastore, documentId))
// 		log.Println(errorInfo.Error.Error())
// 	} else {
// 		documentReferencePtr = firestoreClientPtr.Collection(datastore).Doc(documentId)
// 		if doesDocumentExist(documentReferencePtr) == false {
// 			errorInfo.Error = errors.New(fmt.Sprintf("The document was not found. %v: '%v'", rcv.FN_DOCUMENT_ID, documentId))
// 			log.Println(errorInfo.Error.Error())
// 			documentReferencePtr = nil
// 		}
// 	}
//
// 	return
// }

// GetDocumentIdsWithSubCollections
// func GetDocumentIdsWithSubCollections(firestoreClientPtr *firestore.Client, datastore, parentDocumentId, subCollectionName string) (documentRefIds []string, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tPath              string
// 		tDocumentPtr       []*firestore.DocumentSnapshot
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || parentDocumentId == rcv.EMPTY || subCollectionName == rcv.EMPTY {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 		log.Println(errorInfo.Error)
// 	} else {
// 		tPath = fmt.Sprintf("%v/%v/%v", datastore, parentDocumentId, subCollectionName)
// 		tDocumentPtr, errorInfo.Error = firestoreClientPtr.Collection(tPath).Documents(CTXBackground).GetAll()
// 		for _, snapshot := range tDocumentPtr {
// 			documentRefIds = append(documentRefIds, snapshot.Ref.ID)
// 		}
// 	}
//
// 	return
// }

// GetDocumentFromSubCollectionByDocumentId
//
//	If the document is not found, an error will be returned.
// func GetDocumentFromSubCollectionByDocumentId(firestoreClientPtr *firestore.Client, datastore, parentDocumentId, subCollectionName, documentId string) (data map[string]interface{}, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tDocumentRefPtr    *firestore.DocumentRef
// 		tDocumentPtr       *firestore.DocumentSnapshot
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tPath              string
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || parentDocumentId == rcv.EMPTY || subCollectionName == rcv.EMPTY || documentId == rcv.EMPTY {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 		log.Println(errorInfo.Error)
// 	} else {
// 		tPath = fmt.Sprintf("%v/%v/%v/%v", datastore, parentDocumentId, subCollectionName, documentId)
// 		if tDocumentRefPtr = firestoreClientPtr.Doc(tPath); errorInfo.Error == nil {
// 			if tDocumentPtr, errorInfo.Error = tDocumentRefPtr.Get(CTXBackground); errorInfo.Error == nil {
// 				data = tDocumentPtr.Data()
// 			}
// 		}
// 	}
//
// 	return
// }

// GetFirestoreClientConnection
// func GetFirestoreClientConnection(appPtr *firebase.App) (firestoreClientPtr *firestore.Client, errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if appPtr == nil {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! %v: '%v'", "appPtr", appPtr))
// 	} else {
// 		// firestoreClientPtr is in the function definition because error is passed up the stack by Firebase/Firestore
// 		if firestoreClientPtr, errorInfo.Error = appPtr.Firestore(context.Background()); errorInfo.Error != nil {
// 			log.Println(errorInfo.Error.Error() + rcv.ENDING_EXECUTION)
// 		} else {
// 			log.Printf("The Firebase app connection has been established along with the Firestore Client.")
// 		}
// 	}
//
// 	return
// }

// RemoveDocument
// func RemoveDocument(firestoreClientPtr *firestore.Client, datastore string, queryParameters ...NameValueQuery) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tDocument          *firestore.DocumentSnapshot
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tQuery             firestore.Query
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || len(queryParameters) < 1 {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' nameValueQuery argument is '%v'", datastore, rcv.EMPTY))
// 	} else {
// 		tQuery = firestoreClientPtr.Collection(datastore).Query
// 		for _, parameter := range queryParameters {
// 			if parameter.FieldName == rcv.EMPTY || parameter.FieldValue == rcv.EMPTY {
// 				errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' nameValueQuery parameter is '%v' Field Name: %v, Field Value: %v", datastore, rcv.EMPTY,
// 					parameter.FieldName, parameter.FieldValue))
// 				break
// 			} else {
// 				tQuery = tQuery.Where(parameter.FieldName, rcv.EQUALS, parameter.FieldValue)
// 			}
// 		}
// 	}
//
// 	if errorInfo.Error == nil {
// 		tDocuments := tQuery.Documents(CTXBackground)
// 		for {
// 			tDocument, errorInfo.Error = tDocuments.Next()
// 			if errors.Is(errorInfo.Error, iterator.Done) {
// 				errorInfo.Error = nil
// 				break
// 			}
// 			if errorInfo.Error != nil {
// 				errorInfo.AdditionalInfo = fmt.Sprintf("An error occurred trying to remove a document. Error: %v", errorInfo.Error)
// 				errorInfo.Error = cpi.ErrServiceFailedFIRESTORE
// 				cpi.PrintError(errorInfo)
// 				// todo handle error & notification
// 			}
// 			if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(tDocument.Ref.ID).Delete(CTXBackground); errorInfo.Error != nil {
// 				errorInfo.AdditionalInfo = fmt.Sprintf("%v Failed: Investigate, there is something wrong! Error: %v", tFunctionName, errorInfo.Error.Error())
// 				errorInfo.Error = cpi.ErrServiceFailedFIRESTORE
// 				cpi.PrintError(errorInfo)
// 				// todo Handle error and Notification
// 			}
// 		}
// 	}
//
// 	return
// }

// RemoveDocumentById
// func RemoveDocumentById(firestoreClientPtr *firestore.Client, datastore, documentId string) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || documentId == rcv.EMPTY {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' Document Id: '%v'", datastore, documentId))
// 	} else {
// 		_, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(documentId).Delete(CTXBackground)
// 	}
//
// 	return
// }

// RemoveDocumentFromSubCollectionByDocumentId
// func RemoveDocumentFromSubCollectionByDocumentId(firestoreClientPtr *firestore.Client, datastore, parentDocumentId, subCollectionName, documentId string) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || parentDocumentId == rcv.EMPTY || subCollectionName == rcv.EMPTY || documentId == rcv.EMPTY {
// 		errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' Parent Document Id: '%v' Sub-Collection Name: '%v' Document Id: '%v'", datastore, parentDocumentId,
// 			subCollectionName, documentId))
// 	} else {
// 		if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(parentDocumentId).Collection(subCollectionName).Doc(documentId).Delete(CTXBackground); errorInfo.Error != nil {
// 			errorInfo.Error = errors.New(fmt.Sprintf("%v Failed: Investigate, there is something wrong! Error: %v", "removeDocument", errorInfo.Error.Error()))
// 			log.Println(errorInfo.Error.Error())
// 			// todo Handle error and Notification
// 		}
// 	}
//
// 	return
// }

// RemoveDocumentFromSubCollection
//
//	Customer Messages: None
//	Errors: cpi.ErrRequiredArgumentMissing
//	Verification: Check datastore, parentDocumentId, and subCollectionName are populated
// func RemoveDocumentFromSubCollection(firestoreClientPtr *firestore.Client, datastore, parentDocumentId, subCollectionName string) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tDocumentRefIterPtr *firestore.DocumentRefIterator
// 		tDocumentRefPtr     *firestore.DocumentRef
// 		tFunction, _, _, _  = runtime.Caller(0)
// 		tFunctionName       = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || parentDocumentId == rcv.EMPTY || subCollectionName == rcv.EMPTY {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 	} else {
// 		tDocumentRefIterPtr = firestoreClientPtr.Collection(datastore).Doc(parentDocumentId).Collection(subCollectionName).DocumentRefs(CTXBackground)
// 		for {
// 			tDocumentRefPtr, errorInfo.Error = tDocumentRefIterPtr.Next()
// 			if errors.Is(errorInfo.Error, iterator.Done) {
// 				errorInfo.Error = nil
// 				break
// 			}
// 			if errorInfo.Error != nil {
// 				break
// 			}
// 			_, _ = tDocumentRefPtr.Delete(CTXBackground)
// 		}
// 	}
//
// 	return
// }

// SetDocument - This will create or overwrite the record. While nameValues is a map[any], this function will apply a string assertion on the key.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
// func SetDocument(firestoreClientPtr *firestore.Client, datastore, documentId string, nameValues map[any]interface{}) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if coreValidators.AreMapKeysPopulated(nameValues) == false {
// 		errorInfo.Error = cpi.GetMapKeyPopulatedError(tFinding)
// 	} else {
// 		if firestoreClientPtr == nil || datastore == rcv.EMPTY || documentId == rcv.EMPTY {
// 			errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 			cpi.PrintError(errorInfo)
// 			// todo Handle errors and Notifications
// 		} else {
// 			if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(documentId).Set(CTXBackground, coreHelpers.ConvertMapAnyToMapString(nameValues)); errorInfo.Error != nil {
// 				errorInfo.Error = cpi.ErrServiceFailedFIRESTORE
// 				cpi.PrintError(errorInfo)
// 				// todo Handle errors and Notifications
// 			}
// 		}
// 	}
//
// 	return
// }

// SetDocumentWithSubCollection - This will create or overwrite the existing record that is in a sub-collection. While nameValues is a map[any], this function will apply a string assertion on the key.
// func SetDocumentWithSubCollection(firestoreClientPtr *firestore.Client, datastore, parentDocumentId, subCollectionName, documentId string, nameValues map[any]interface{}) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if tFinding = coreValidators.AreMapKeysValuesPopulated(nameValues); tFinding != rcv.GOOD {
// 		errorInfo.Error = cpi.GetMapKeyPopulatedError(tFinding)
// 	} else {
// 		// if datastore == rcv.EMPTY || parentDocumentId == rcv.EMPTY || subCollectionName == rcv.EMPTY || documentId == rcv.EMPTY {
// 		if datastore == rcv.EMPTY || parentDocumentId == rcv.EMPTY || subCollectionName == rcv.EMPTY || documentId == rcv.EMPTY {
// 			errorInfo.Error = errors.New(fmt.Sprintf("Require information is missing! Datastore: '%v' Parent Document Id: '%v' Sub-collection Name: '%v' Document Id: '%v' Function Name: %v", datastore, parentDocumentId, subCollectionName, documentId, tFunctionName))
// 			log.Println(errorInfo.Error.Error())
// 			// todo Handle errors and Notifications
// 		} else {
// 			if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(parentDocumentId).Collection(subCollectionName).Doc(documentId).Set(CTXBackground, coreHelpers.ConvertMapAnyToMapString(nameValues)); errorInfo.Error != nil {
// 				errorInfo.Error = errors.New(fmt.Sprintf("An error has occurred creating Document Id: %v for Datastore: %v Parent Document Id: '%v' Subcollection Name: '%v' Error: %v", documentId, datastore,
// 					parentDocumentId, subCollectionName, errorInfo.Error.Error()))
// 				log.Println(errorInfo.Error.Error())
// 				// todo Handle errors and Notifications
// 			}
// 		}
// 	}
//
// 	return
// }

// UpdateDocument- will return an error of nil when successful. If the document is not found, shared_services.ErrDocumentNotFound will be returned, otherwise the error from Firestore will be returned.
// func UpdateDocument(firestoreClientPtr *firestore.Client, datastore, documentId string, nameValues map[any]interface{}) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFinding           string
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tUpdateFields      []firestore.Update
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	errorInfo.AdditionalInfo = fmt.Sprintf("Datastore: %v Document Id: %v", datastore, documentId)
//
// 	if tFinding = coreValidators.AreMapKeysValuesPopulated(nameValues); tFinding != rcv.GOOD {
// 		errorInfo.Error = cpi.GetMapKeyPopulatedError(tFinding)
// 		cpi.PrintError(errorInfo)
// 	} else {
// 		if datastore == rcv.EMPTY || documentId == rcv.EMPTY {
// 			errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 			cpi.PrintError(errorInfo)
// 			// todo Handle errors and Notifications
// 		} else {
// 			if tUpdateFields, errorInfo = BuildFirestoreUpdate(nameValues); errorInfo.Error == nil {
// 				if _, errorInfo.Error = firestoreClientPtr.Collection(datastore).Doc(documentId).Update(CTXBackground, tUpdateFields); errorInfo.Error != nil {
// 					cpi.PrintError(errorInfo)
// 				}
// 			}
// 		}
// 	}
//
// 	return
// }

// UpdateDocumentFromSubCollectionByDocumentId
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing, Any error from Firestore
//	Verifications: None
// func UpdateDocumentFromSubCollectionByDocumentId(firestoreClientPtr *firestore.Client, datastore, parentDocumentId, subCollectionName, documentId string, updateFields []firestore.Update) (errorInfo cpi.ErrorInfo) {
//
// 	var (
// 		tFunction, _, _, _ = runtime.Caller(0)
// 		tFunctionName      = runtime.FuncForPC(tFunction).Name()
// 		tPath              string
// 	)
//
// 	cpi.PrintDebugTrail(tFunctionName)
//
// 	if datastore == rcv.EMPTY || parentDocumentId == rcv.EMPTY || subCollectionName == rcv.EMPTY || documentId == rcv.EMPTY {
// 		errorInfo.Error = cpi.ErrRequiredArgumentMissing
// 		log.Println(errorInfo.Error)
// 	} else {
// 		tPath = fmt.Sprintf("%v/%v/%v/%v", datastore, parentDocumentId, subCollectionName, documentId)
// 		_, errorInfo.Error = firestoreClientPtr.Doc(tPath).Update(CTXBackground, updateFields)
// 	}
//
// 	return
// }
