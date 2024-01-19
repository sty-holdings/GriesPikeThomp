/*
This is the STY-Holdings NATS-Connect service

NOTES:

	None

COPYRIGHT & WARRANTY:

	Copyright (c) 2022 STY-Holdings, inc
	All rights reserved.

	This software is the confidential and proprietary information of STY-Holdings, Inc.
	Use is subject to license terms.

	Unauthorized copying of this file, via any medium is strictly prohibited.

	Proprietary and confidential

	Written by Scott Yacko / syacko
	STY-Holdings, Inc.
	support@sty-holdings.com
	<Replace with WEBSITE_NAME>

	12-2023

	USA

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package main

import (
	"fmt"
	"log"
	"os"

	"GriesPikeThomp/servers/nats-connect/src"
	cc "GriesPikeThomp/shared-services/src/coreConfiguration"
	cpi "GriesPikeThomp/shared-services/src/coreProgramInfo"
	"github.com/integrii/flaggy"
	rcv "github.com/sty-holdings/resuable-const-vars/src"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Add types to the request_reply_types.go or the data_structure_types.go file

var (
	// Add Variables here for the file (Remember, they are global)
	// Start up values for a service
	configFileFQN      string
	generateConfigFile bool
	serverName         = "nats-connect"
	testingOn          bool
	version            = "9999.9999.9999"
)

func init() {

	appDescription := cases.Title(language.English).String(serverName) + " bridges the gap between NATS users and the wider world, enabling effortless integration" +
		" with third-party platforms and services.\n" +
		"\nVersion: \n" +
		rcv.SPACES_FOUR + "- " + version + "\n" +
		"\nConstraints: \n" +
		rcv.SPACES_FOUR + "- When using -c you must pass the fully qualified configuration file name.\n" +
		rcv.SPACES_FOUR + "- There is no console available at this time and all log messages are output to Log_Directory specified in the config file.\n" +
		"\nNotes:\n" +
		rcv.SPACES_FOUR + "None\n" +
		"\nFor more info, see link below:\n"

	// Set your program's name and description.  These appear in help output.
	flaggy.SetName("\n" + serverName) // "\n" is added to the start of the name to make the output easier to read.
	flaggy.SetDescription(appDescription)

	// You can disable various things by changing bool on the default parser
	// (or your own parser if you have created one).
	flaggy.DefaultParser.ShowHelpOnUnexpected = true

	// You can set a help prepend or append on the default parser.
	flaggy.DefaultParser.AdditionalHelpPrepend = "https://github.com/styh-dev/albert"

	// Add a flag to the main program (this will be available in all subcommands as well).
	flaggy.String(&configFileFQN, "c", "config", "Provides the setup information needed by and is required to start the server.")
	flaggy.Bool(&generateConfigFile, "g", "genconfig", "This will output a skeleton configuration file.\n\t\t\tThis will cause all other options to be ignored.")
	flaggy.Bool(&testingOn, "t", "testingOn", "This puts the server into testing mode.")

	// Set the version and parse all inputs into variables.
	flaggy.SetVersion(version)
	flaggy.Parse()
}

func main() {

	var (
		returnCode = 0
		tServer    *src.Server
	)

	fmt.Println()
	log.Printf("Starting %v server.\n", serverName)

	if serverName == rcv.VAL_EMPTY {
		cpi.PrintError(cpi.ErrMissingServerName, fmt.Sprintf("%v %v", rcv.TXT_SERVER_NAME, serverName), rcv.MODE_OUTPUT_DISPLAY)
		os.Exit(1)
	}

	if generateConfigFile {
		cc.GenerateConfigFileSkeleton(serverName, cc.DEFAULT_SKELETON_CONFIG_DIRECTORY, cc.DEFAULT_SKELETON_CONFIG_FILENAME_NO_SUFFIX)
		os.Exit(0)
	}

	// Has the config file location and name been provided, if not, return help.
	if configFileFQN == "" && testingOn == false {
		flaggy.ShowHelpAndExit("")
	}

	if tServer, returnCode = src.NewServer(configFileFQN, version, testingOn); returnCode > 0 {
		Shutdown(returnCode)
	}

	Shutdown(tServer.Run()) // Start things up. Block here until done.
}

func Shutdown(returnCode int) {
	log.Printf("Shutting down %v server.\n", serverName)
	os.Exit(returnCode)
}
