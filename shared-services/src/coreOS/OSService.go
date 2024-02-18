// Package coreOS
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
package coreOS

import (
	"net"
	"regexp"

	"GriesPikeThomp/shared-services/src/coreValidators"
)

var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

// GetIPAddresses - returns a list of all IP addresses
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetIPAddresses() (ipList []string) {

	var (
		err            error
		tNetAddr       []net.Addr
		tNetInterfaces []net.Interface
	)

	if tNetInterfaces, err = net.Interfaces(); err == nil {
		// handle err
		for _, i := range tNetInterfaces {
			tNetAddr, err = i.Addrs()
			// handle err
			for _, addr := range tNetAddr {
				var ip net.IP
				switch v := addr.(type) {
				case *net.IPNet:
					ip = v.IP
				case *net.IPAddr:
					ip = v.IP
				}
				ipList = append(ipList, ip.String())
			}
		}
	}

	return
}

// GetIPv4Addresses - returns a list of all IPv4 addresses
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetIPv4Addresses() (ipV4List []string) {

	for _, ipAddress := range GetIPAddresses() {
		if coreValidators.IsIPv4Valid(ipAddress) {
			ipV4List = append(ipV4List, ipAddress)
		}
	}

	return
}

// GetIPv6Addresses - returns a list of all IPv6 addresses
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None
func GetIPv6Addresses() (ipV6List []string) {

	for _, ipAddress := range GetIPAddresses() {
		if coreValidators.IsIPv6Valid(ipAddress) {
			ipV6List = append(ipV6List, ipAddress)
		}
	}

	return
}
