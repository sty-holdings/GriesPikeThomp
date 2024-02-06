// Package coreOS
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
package coreOS

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGetIPAddresses(tPtr *testing.T) {
	var (
		ipAddresses        []string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Success
		if ipAddresses = GetIPAddresses(); len(ipAddresses) == 0 {
			tPtr.Errorf("%v FAILED - THIS IS NOT WORKING. THERE SHOULD ALWAYS BE AN IP ADDRESS.", tFunctionName)
		} else {
			for _, address := range ipAddresses {
				fmt.Println("IP: ", address)
			}
		}
	})
}

func TestGetIPv4Addresses(tPtr *testing.T) {
	var (
		ipAddresses        []string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Success
		if ipAddresses = GetIPv4Addresses(); len(ipAddresses) == 0 {
			tPtr.Errorf("%v FAILED - THIS IS NOT WORKING. THERE SHOULD ALWAYS BE AN IPv4 ADDRESS.", tFunctionName)
		} else {
			for _, address := range ipAddresses {
				fmt.Println("IP: ", address)
			}
		}
	})
}

func TestGetIPv6Addresses(tPtr *testing.T) {
	var (
		ipAddresses        []string
		tFunction, _, _, _ = runtime.Caller(0)
		tFunctionName      = runtime.FuncForPC(tFunction).Name()
	)

	tPtr.Run(tFunctionName, func(tPtr *testing.T) {
		// Success
		if ipAddresses = GetIPv6Addresses(); len(ipAddresses) == 0 {
			fmt.Printf("%v WARNING - Your system may not have an IPv6 assigned.", tFunctionName)
		} else {
			for _, address := range ipAddresses {
				fmt.Println("IP: ", address)
			}
		}
	})
}
