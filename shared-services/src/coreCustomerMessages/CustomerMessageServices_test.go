// Package coreCustomerMessages
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
package coreCustomerMessages

import (
	"runtime"
	"testing"

	"albert/constants"
)

var (
	tFunction, _, _, _ = runtime.Caller(0)
	tFunctionName      = runtime.FuncForPC(tFunction).Name()
)

func TestUsernameAlreadyExists(tPtr *testing.T) {

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		_ = UsernameAlreadyExists(constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE)
	})
}

func TestUserCreatedVerifyEmailNext(tPtr *testing.T) {

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		_ = UserCreatedVerifyEmailNextStep()
	})
}

func TestUnexpectedError(tPtr *testing.T) {

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		_ = UnexpectedError()
	})
}

func TestUsernameNotFound(tPtr *testing.T) {

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		_ = UsernameNotFound(constants.TEST_USERNAME_SAVUP_TEST_DO_NOT_DELETE)
	})
}

func TestUserRetryLimitHit(tPtr *testing.T) {

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		_ = UserRetryLimitHit()
	})
}
