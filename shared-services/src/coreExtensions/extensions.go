// Package sharedServices
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
package sharedServices

import (
	"fmt"
	"strings"

	ns "GriesPikeThomp/shared-services/src/coreNATS"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

func HandleExtension(extensions map[string]map[string]interface{}) (extensionPtrs map[string]interface{}, errorInfo cpi.ErrorInfo) {

	var (
		tExtension = make(map[string]interface{})
	)

	extensionPtrs = make(map[string]interface{})
	for extensionName, extensionValues := range extensions {
		switch strings.ToLower(extensionName) {
		case NATS_INTERNAL:
			tExtension = extensionValues
			extensionPtrs[NATS_INTERNAL], errorInfo = ns.NewNATS(tExtension)
		default:
			fmt.Println(rcv.TXT_BAD)
		}
	}

	return
}
