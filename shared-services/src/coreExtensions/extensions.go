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

	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	ns "GriesPikeThomp/shared-services/src/coreNATS"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

// HandleExtension - will route the extension configuration file.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None/
func HandleExtension(configExtensions []cc.BaseConfigExtensions) (extensionPtrs map[string]any, errorInfo cpi.ErrorInfo) {

	extensionPtrs = make(map[string]interface{})
	for _, values := range configExtensions {
		switch strings.ToLower(values.Name) {
		case NATS_INTERNAL:
			extensionPtrs[NATS_INTERNAL], errorInfo = ns.NewNATS(values.ConfigFilename)
		default:
			errorInfo = cpi.NewErrorInfo(cpi.ErrExtensionInvalid, fmt.Sprintf("%v%v", rcv.TXT_EXTENSION_NAME, values.Name))
		}
	}

	return
}
