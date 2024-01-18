// Package corePostgres
/*
This is the STY-Holdings shared services

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc..
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
	"context"
	"fmt"
	"strings"

	"albert/constants"
	"albert/core/coreError"
	"albert/core/coreValidators"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ConnValues struct {
	DBName   string `json:"dbName"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Timeout  int    `json:"timeout"`
	SSLMode  string `json:"sslMode"`
}

var (
	CTXBackground = context.Background()
)

// Row and Rows are so pgx package doesn't need to be imported everywhere there are queries to the database.
type Transaction pgx.Tx
type Rows pgx.Rows
type Row pgx.Row

// GetPostgresConnection will create a connection to a database.
// dbName   Name of the Postgres database
// user     Account that GetPostgresConnection will use to authenticate
// password Users password for authentication
// host     Internet DNS or IP address of the server running the instance of Postgres
// sslMode  Type of encryption used for the connection (https://www.postgresql.org/docs/12/libpq-ssl.html for version 12)
// port     Interface the connection communicates with Postgres
// timeout  Number of seconds a request must complete (3 seconds is normal setting)
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetPostgresConnection(dbName, user, password, host, sslMode string, port, timeout, poolMaxConns int) (DBPoolPtr *pgxpool.Pool, errorInfo coreError.ErrorInfo) {

	if dbName == constants.EMPTY || user == constants.EMPTY || password == constants.EMPTY || host == constants.EMPTY || coreValidators.IsPostgresSSLModeValid(sslMode) {
		errorInfo.Error = coreError.ErrRequiredArgumentMissing
		coreError.PrintError(errorInfo)
		return
	}

	if port == 0 || timeout == 0 || poolMaxConns == 0 {
		errorInfo.Error = coreError.ErrRequiredArgumentMissing
		coreError.PrintError(errorInfo)
		return
	}

	if coreValidators.IsPostgresSSLModeValid(sslMode) == false {
		errorInfo.Error = coreError.ErrPostgresSSLMode
		coreError.PrintError(errorInfo)
		return
	}

	if DBPoolPtr, errorInfo.Error = pgxpool.New(context.Background(), setConnectionValues(dbName, user, password, host, sslMode, port, timeout, poolMaxConns)); errorInfo.Error != nil {
		if strings.Contains(errorInfo.Error.Error(), "dial") {
			errorInfo.Error = coreError.ErrPostgresConnFailed
			coreError.PrintError(errorInfo)
		} else {
			errorInfo.Error = coreError.ErrServiceFailedPOSTGRES
			coreError.PrintError(errorInfo)
		}
	}

	return
}

func setConnectionValues(dbName, user, password, host, sslMode string, port, timeout, poolMaxConns int) (dbConnString string) {

	return fmt.Sprintf(constants.POSTGRES__CONN_STRING, dbName, user, password, host, port, timeout, sslMode, poolMaxConns)

}

// Verify that the pointer to the database connection is active.
func VerifyConnection(dbPoolPtr *pgxpool.Pool, dbName string) (errorInfo coreError.ErrorInfo) {

	var (
		tRows Rows
	)

	if dbPoolPtr == nil {
		errorInfo.Error = coreError.ErrPostgresConnEmpty
		coreError.PrintError(errorInfo)
	} else {
		qStmt := "SELECT * FROM pg_stat_activity WHERE datname = $1 and state = 'active';"
		tRows, errorInfo.Error = dbPoolPtr.Query(CTXBackground, qStmt, dbName)
		if errorInfo.Error != nil {
			errorInfo.Error = coreError.ErrPostgresConnFailed
			coreError.PrintError(errorInfo)
		}
		defer tRows.Close()
	}

	return
}
