// Package src
/*
General description of the purpose of the go file.

RESTRICTIONS:
    AWS functions:
    * Program must have access to a .aws/credentials file in the default location.
    * This will only access system parameters that start with '/sote' (ROOTPATH).
    * {Enter other restrictions here for AWS

    {Other catagories of restrictions}
    * {List of restrictions for the catagory

NOTES:
    {Enter any additional notes that you believe will help the next developer.}

COPYRIGHT:
	Copyright 2022
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.

*/
package src

import (
	"os"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	ext "github.com/sty-holdings/GriesPikeThomp/servers/nats-connect/loadExtensions"
)

//goland:noinspection GoSnakeCaseUsage,GoCommentStart
const (
	// Subjects
	NCI_PING           = "ping"
	NCI_TURN_DEBUG_OFF = "turn_debug_off"
	NCI_TURN_DEBUG_ON  = "turn_debug_on"
)

type Auth struct {
	authenticatorService string
}

// Configuration is a generic config file structure for application servers.
type BaseConfiguration struct {
	ConfigFQN               string
	DebugModeOn             bool                       `json:"debug_mode_on"`
	Environment             string                     `json:"environment"`
	LogDirectory            string                     `json:"log_directory"`
	MaxThreads              int                        `json:"max_threads"`
	PIDDirectory            string                     `json:"pid_directory"`
	Extensions              []ext.BaseConfigExtensions `json:"load_extensions"`
	ServerInstanceNumber    string                     `json:"server_instance_number"`
	SkeletonConfigDirectory string                     `json:"skeleton_config_directory"`
}

// type BaseConfigExtensions struct {
// 	Name           string `json:"name"`
// 	ConfigFilename string `json:"config_filename"`
// }

// ExtInstance
// READ ME: Shorten the name to resolve naming conflict
// *********************
type ExtInstance struct {
	InstanceName      string
	NatsConnectionPtr *nats.Conn
	ProcessChannel    chan string
	SubscriptionPtrs  map[string]*nats.Subscription
	WaitGroup         sync.WaitGroup
}

// Instance - Some of these values can change over the life of the instance.
type Instance struct {
	baseURL           string
	debugModeOn       bool
	extInstances      map[string]ExtInstance
	hostname          string
	logFileHandlerPtr *os.File
	logFQN            string
	messageHandlers   map[string]nats.MsgHandler
	mu                sync.RWMutex
	numberCPUS        int
	outputMode        string
	pid               int
	pidFQN            string
	nciProcessChannel chan string
	running           bool
	runStartTime      time.Time
	serverName        string
	testingOn         bool
	threadsAssigned   uint
	version           string
	waitGroupCreated  bool
	workingDirectory  string
}

type Server struct {
	config           BaseConfiguration
	instance         Instance
	extensionConfigs map[string]ext.ExtensionConfiguration
}
