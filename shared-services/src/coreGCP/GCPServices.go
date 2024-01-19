// Package coreGCP
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
package coreGCP

import (
	"context"
	"io/ioutil"
	"log"
	"os"
	"runtime"

	"albert/constants"
	"albert/core/coreError"
	"cloud.google.com/go/storage"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var (
	ctx = context.Background()
)

// CreateStorageClient - connect to Google Cloud Platform services
//
//	Customer Messages: None
//	Errors: Return an errors generated by GCP with Ending Execution appended
//	Verifications: None
func CreateStorageClient(credentialsFile string, test bool) (client *storage.Client, errorInfo coreError.ErrorInfo) {

	var (
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	if client, errorInfo.Error = storage.NewClient(ctx, option.WithCredentialsJSON(getGCPKey(credentialsFile, test))); errorInfo.Error != nil {
		log.Println(errorInfo.Error.Error(), constants.ENDING_EXECUTION)
	}

	return
}

// getBucket - return a pointer to a storage bucket
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing
//	Verifications: storageClientPtr is not nil
func getBucket(storageClientPtr *storage.Client, bucketName string) (bucketPtr *storage.BucketHandle, errorInfo coreError.ErrorInfo) {

	if storageClientPtr == nil || bucketName == constants.EMPTY {
		errorInfo.Error = coreError.ErrRequiredArgumentMissing
	} else {
		// Create a bucket object for the specified bucket.
		bucketPtr = storageClientPtr.Bucket(bucketName)
	}

	return
}

// getGCPKey - will read the JSON key file. If either fail, exit is called.
//
//	Customer Messages: None
//	Errors: ErrUnableReadFile
//	Validations: File readable
func getGCPKey(GCPCredentialsFQN string, test bool) (GCPCredentials []byte) {

	var (
		errorInfo          coreError.ErrorInfo
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	coreError.PrintDebugTrail(tFunctionName)

	if GCPCredentials, errorInfo.Error = os.ReadFile(GCPCredentialsFQN); errorInfo.Error != nil {
		errorInfo.Error = coreError.ErrUnableReadFile
		log.Println(errorInfo.Error.Error())
	}

	if errorInfo.Error != nil && test == false {
		os.Exit(1)
	}

	return
}

// ListFilesInBucket - return all the object names in a bucket, folders and files
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing
//	Verifications: storageClientPtr is not nil
func ListObjectsInBucket(storageClientPtr *storage.Client, bucketName string) (bucketList []string, errorInfo coreError.ErrorInfo) {

	var (
		tBucketPtr        *storage.BucketHandle
		tObjectAttributes *storage.ObjectAttrs
		tObjectIterator   *storage.ObjectIterator
	)

	if storageClientPtr == nil || bucketName == constants.EMPTY {
		errorInfo.Error = coreError.ErrRequiredArgumentMissing
	} else {
		tBucketPtr, errorInfo = getBucket(storageClientPtr, bucketName)

		// Create a list object for the bucket.
		tObjectIterator = tBucketPtr.Objects(ctx, nil)

		for {
			tObjectAttributes, errorInfo.Error = tObjectIterator.Next()
			if errorInfo.Error == iterator.Done {
				errorInfo.Error = nil
				break
			}
			if errorInfo.Error != nil {
				break
			}
			bucketList = append(bucketList, tObjectAttributes.Name)
		}
	}

	return
}

// ReadBucketObject - returns the contains of the bucket's named file.
//
//	Customer Messages: None
//	Errors: ErrRequiredArgumentMissing
//	Verifications: bucketPtr is not nil
func ReadBucketObject(storageClientPtr *storage.Client, bucketName string, fileName string) (contents []byte, errorInfo coreError.ErrorInfo) {

	var (
		tBucketPtr *storage.BucketHandle
		tReader    *storage.Reader
	)

	if storageClientPtr == nil || bucketName == constants.EMPTY || fileName == constants.EMPTY {
		errorInfo.Error = coreError.ErrRequiredArgumentMissing
	} else {
		if tBucketPtr, errorInfo = getBucket(storageClientPtr, bucketName); errorInfo.Error == nil {
			// Create an object for the specified file.
			if tReader, errorInfo.Error = tBucketPtr.Object(fileName).NewReader(context.Background()); errorInfo.Error == nil {
				defer func(tReader *storage.Reader) {
					_ = tReader.Close()
				}(tReader)
				// Read the contents of the reader into a byte slice.
				contents, errorInfo.Error = ioutil.ReadAll(tReader)
			}
		}
	}

	return
}
