// Package sharedServices
/*
General description of the purpose of the go file.

<<<<<<< HEAD
=======
RESTRICTIONS:
    AWS functions:
    * Program must have access to a .aws/credentials file in the default location.
    * This will only access system parameters that start with '/sote' (ROOTPATH).
    * {Enter other restrictions here for AWS

    {Other catagories of restrictions}
    * {List of restrictions for the catagory

>>>>>>> fbf9762 (Fixed the label)
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
package sharedServices

//goland:noinspection GoSnakeCaseUsage
const (
	DEFAULT_LOG_DIRECTORY                    = "/var/log/nats-connect"
	DEFAULT_MAX_THREADS                      = 1
	DEFAULT_PID_DIRECTORY                    = "/var/run/nats-connect"
	DEFAULT_SKELETON_CONFIG_FQD              = "/Users/syacko/workspace/sty-holdings/GriesPikeThomp/shared-services/src/coreConfiguration/"
	DEFAULT_SKELETON_CONFIG_FILENAME         = "skeleton-config-file.json"
	DEFAULT_SKELETON_CONFIG_NOTE_FILENAME    = "skeleton-config-file.txt"
	DEFAULT_INVALID_SKELETON_CONFIG_FILENAME = "invalid-skeleton-config-file.json"
	DEFAULT_UNREADABLE_CONFIG_FILENAME       = "unreadable-skeleton-config-file.json"
	THREAD_CAP                               = 25 // This is set so system performance can be controlled. Update this as more is learned about the system.
)
