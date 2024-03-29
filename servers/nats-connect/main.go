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

	"github.com/integrii/flaggy"
	"github.com/sty-holdings/GriesPikeThomp/servers/nats-connect/src"
	ctv "github.com/sty-holdings/constant-type-vars-go/v2024"
	config "github.com/sty-holdings/sty-shared/v2024/configuration"
	pi "github.com/sty-holdings/sty-shared/v2024/programInfo"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// Add types to the request_reply_types.go or the data_structure_types.go file

var (
	// Add Variables here for the file (Remember, they are global)
	// Start up values for a service
	configFileFQN     string
	generateConfigFQD string
	serverName        = "nats-connect"
	testingOn         bool
	version           = "9999.9999.9999"
)

func init() {

	appDescription := cases.Title(language.English).String(serverName) + " bridges the gap between NATS users and the wider world, enabling effortless integration" +
		" with third-party platforms and services.\n" +
		"\nVersion: \n" +
		ctv.SPACES_FOUR + "- " + version + "\n" +
		"\nConstraints: \n" +
		ctv.SPACES_FOUR + "- When using -c you must pass the fully qualified configuration file name.\n" +
		ctv.SPACES_FOUR + "- There is no console available at this time and all log messages are output to Log_Directory specified in the config file.\n" +
		"\nNotes:\n" +
		ctv.SPACES_FOUR + "None\n" +
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
	flaggy.String(
		&generateConfigFQD,
		"g",
		"genconfig",
		"This will output a skeleton configuration and note files.\n\t\t\tThis will cause all other options to be ignored.",
	)
	flaggy.Bool(&testingOn, "t", "testingOn", "This puts the server into testing mode.")

	// Set the version and parse all inputs into variables.
	flaggy.SetVersion(version)
	flaggy.Parse()
}

func main() {

	var (
		returnCode = 0
	)

	fmt.Println()
	log.Printf("Starting %v server.\n", serverName)

	if len(generateConfigFQD) > 0 {
		config.GenerateConfigFileSkeleton(serverName, generateConfigFQD)
		os.Exit(0)
	}

	// This is to prevent the serverName from being empty.
	if serverName == ctv.VAL_EMPTY {
		pi.PrintError(pi.ErrServerNameMissing, fmt.Sprintf("%v %v", ctv.TXT_SERVER_NAME, serverName))
		os.Exit(1)
	}

	// This is to prevent the version from being empty or not being set during the build process.
	if (version == ctv.VAL_EMPTY || version == "9999.9999.9999") && testingOn == false {
		pi.PrintError(pi.ErrVersionInvalid, fmt.Sprintf("%v %v", ctv.TXT_SERVER_VERSION, version))
		os.Exit(1)
	}

	// Has the config file location and name been provided, if not, return help.
	if (configFileFQN == "" || configFileFQN == "-t") && testingOn == false {
		flaggy.ShowHelpAndExit("")
	}

	returnCode = src.RunServer(configFileFQN, serverName, version, testingOn)
	log.Printf("%v server is stopped. REMIDER: Remove the pid file in the .run directory if it exists before running again.\n", serverName)
	os.Exit(returnCode)
}
