// Package corePostgres
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
package corePostgres

import (
	"runtime"
	"testing"

	"albert/constants"
	"github.com/nats-io/nats.go"
)

func TestSetConnectionValue(tPtr *testing.T) {

	type arguments struct {
		dbName      string
		user        string
		password    string
		host        string
		sslMode     string
		port        uint
		timeout     uint
		poolMaxConn uint
	}

	var (
		err           error
		gotError      bool
		tDBConnString string
	)

	tests := []struct {
		name      string
		arguments arguments
		wantError bool
	}{
		{
			name: "Negative Case: Missing DBName!",
			arguments: arguments{
				dbName:      rcv.EMPTY,
				user:        "string",
				password:    "string",
				host:        "string",
				sslMode:     "string",
				port:        1,
				timeout:     1,
				poolMaxConn: 1,
			},
			wantError: false,
		},
	}

	for _, ts := range tests {
		tPtr.Run(ts.name, func(t *testing.T) {
			tDBConnString = setConnectionValues(ts.arguments.dbName, ts.arguments.user, ts.arguments.password, ts.arguments.sslMode, ts.arguments.port, ts.arguments.timeout, ts.arguments.poolMaxConn)
			if _, err = executePlaidGetLinkToken(tPlaidClient, tFirebase.FirestoreClientPtr, &nats.Msg{Data: tRequestJSON}); err != nil {
				gotError = true
			} else {
				gotError = false
			}
			if gotError != ts.wantError {
				tPtr.Error(err.Error())
			}
		})
	}

	dbConnString := setConnectionValues("dbName", "User", "Password", "Host", "INVALID", 1, 1, 1)
	if soteErr.ErrCode != 209220 {
		tPtr.Errorf("setConnectionValues Failed: Error code is not for an invalid sslMode.")
		tPtr.Fail()
	}
	_, soteErr = setConnectionValues("dbName", "User", "Password", "Host", "disable", 1, 1)
	if soteErr.ErrCode != nil {
		tPtr.Errorf("setConnectionValues Failed: Expected a nil error code.")
		tPtr.Fail()
	}
}
func TestSetConnectionValues(tPtr *testing.T) {
	_, soteErr := setConnectionValues("dbName", "User", "Password", "Host", "disable", 1, 1)
	if soteErr.ErrCode != nil {
		tPtr.Errorf("setConnectionValues Failed: Expected a nil error code.")
		tPtr.Fail()
	}
}
func TestVerifyConnection(tPtr *testing.T) {
	var tConnInfo ConnInfo
	soteErr := VerifyConnection(tConnInfo)
	if soteErr.ErrCode != 209299 {
		tPtr.Errorf("VerifyConnection Failed: Expected 209299 error code.")
		tPtr.Fail()
	}

	if soteErr = GetAWSParams(); soteErr.ErrCode != nil {
		tPtr.Errorf("GetAWSParams Failed: Expected error code to be nil.")
		tPtr.Fatal()
	}

	tConnInfo, soteErr = GetConnection(DBName, DBUser, DBPassword, DBHost, DBSSLMode, DBPort, 3)
	if soteErr.ErrCode != nil {
		tPtr.Errorf("setConnectionValues Failed: Expected a nil error code.")
		tPtr.Fail()
	}

	soteErr = VerifyConnection(tConnInfo)
	if soteErr.ErrCode != nil {
		tPtr.Errorf("VerifyConnection Failed: Expected a nil error code.")
		tPtr.Fail()
	}

	// This will test the condition that no database is available to connect
	tConnInfo, soteErr = GetConnection(DBName, DBUser, DBPassword, DBHost, DBSSLMode, 65000, 3)
	if soteErr.ErrCode != 209299 {
		tPtr.Errorf("setConnectionValues Failed: Expected 209299 error code.")
		tPtr.Fail()
	}

}
func TestToJSONString(tPtr *testing.T) {
	if soteErr := GetAWSParams(); soteErr.ErrCode != nil {
		tPtr.Errorf("GetAWSParams Failed: Expected error code to be nil.")
		tPtr.Fatal()
	}

	tConnInfo, soteErr := GetConnection(DBName, DBUser, DBPassword, DBHost, DBSSLMode, DBPort, 3)
	if soteErr.ErrCode != nil {
		tPtr.Errorf("GetConnection Failed: Please Investigate")
		tPtr.Fail()
	}

	var dbConnJSONString string
	if dbConnJSONString, soteErr = ToJSONString(tConnInfo.DSConnValues); soteErr.ErrCode != nil {
		tPtr.Errorf("ToJSONString Failed: Please Investigate")
		tPtr.Fail()
	}

	if len(dbConnJSONString) == 0 {
		tPtr.Errorf("ToJSONString Failed: Please Investigate")
		tPtr.Fail()
	}
}
func TestContext(tPtr *testing.T) {
	if soteErr := GetAWSParams(); soteErr.ErrCode != nil {
		tPtr.Errorf("GetAWSParams Failed: Expected error code to be nil.")
		tPtr.Fatal()
	}

	tConnInfo, soteErr := GetConnection(DBName, DBUser, DBPassword, DBHost, DBSSLMode, DBPort, 3)
	if soteErr.ErrCode != nil {
		tPtr.Errorf("setConnectionValues Failed: Expected a nil error code.")
		tPtr.Fail()
	}

	if tConnInfo.DBContext == nil {
		tPtr.Errorf("TestContext testing DBContext Failed: Expected a non-nil error code.")
		tPtr.Fail()
	}
}
func TestSRow(tPtr *testing.T) {
	tRow := SRow(nil)
	if tRow != nil {
		tPtr.Errorf("TestSRow testing creation of SRow variable Failed: Expected error code to be nil.")
		tPtr.Fail()
	}
}
func TestSRows(tPtr *testing.T) {
	tRows := SRows(nil)
	if tRows != nil {
		tPtr.Errorf("TestSRows testing creation of SRows variable Failed: Expected error code to be nil.")
		tPtr.Fail()
	}
}
func getMyDBConn(tPtr *testing.T) (myDBConn ConnInfo, soteErr sError.SoteError) {
	var (
		function, _, _, _ = runtime.Caller(0)
		testName          = runtime.FuncForPC(function).Name()
	)

	if soteErr = GetAWSParams(); soteErr.ErrCode != nil {
		tPtr.Errorf("%v Failed: Expected error code to be nil.", testName)
		tPtr.Fatal()
	}

	myDBConn, soteErr = GetConnection(DBName, DBUser, DBPassword, DBHost, DBSSLMode, DBPort, 3)
	if soteErr.ErrCode != nil {
		tPtr.Errorf("%v Failed: Expected a nil error code.", testName)
		tPtr.Fail()
	}

	return
}
