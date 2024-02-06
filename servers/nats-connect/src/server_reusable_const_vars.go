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

//goland:noinspection GoSnakeCaseUsage,GoCommentStart
const (
	// Services
	//
<<<<<<< HEAD
	NATS_INTERNAL = "nats_internal"
=======
	NATS_INTERNAL = "internal_nats"
>>>>>>> parent of bc61635 (Working HTTP ListenAndServe)
	//
	// Subject
	//
	TURN_DEBUG_ON  = "server_cmd.turn_debug_on"
	TURN_DEBUG_OFF = "server_cmd.turn_debug_off"
)
