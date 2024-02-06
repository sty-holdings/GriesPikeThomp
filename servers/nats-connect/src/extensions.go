// Package src
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
package src

import (
	"fmt"
	"log"
	"strings"

	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
<<<<<<< HEAD
	cn "GriesPikeThomp/shared-services/src/coreNATS"
=======
	ns "GriesPikeThomp/shared-services/src/coreNATS"
>>>>>>> parent of bc61635 (Working HTTP ListenAndServe)
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
)

type Extensions[T any] struct {
	ExtensionsData map[string]T
}

// HandleExtension - will route the extension configuration file.
//
//	Customer Messages: None
//	Errors: None
//	Verifications: None/
func (serverPtr *Server) HandleExtension(hostname string, configExtensions []cc.BaseConfigExtensions) (errorInfo cpi.ErrorInfo) {

	var (
<<<<<<< HEAD
		tNATSService cn.NATSService
=======
		tNATSService ns.NATSService
>>>>>>> parent of bc61635 (Working HTTP ListenAndServe)
	)

	for _, values := range configExtensions {
		switch strings.ToLower(values.Name) {
		case NATS_INTERNAL:
			tNATSService, errorInfo = cn.NewNATS(hostname, values.ConfigFilename)
			serverPtr.extensions[NATS_INTERNAL] = tNATSService
		default:
			errorInfo = cpi.NewErrorInfo(cpi.ErrExtensionInvalid, fmt.Sprintf("%v%v", rcv.TXT_EXTENSION_NAME, values.Name))
			log.Printf("%v failed to load. Removing all extensions.", values.Name)
			serverPtr.extensions = make(map[string]interface{})
			return
		}
		log.Printf("%v is loaded.", values.Name)
	}

	return
}
