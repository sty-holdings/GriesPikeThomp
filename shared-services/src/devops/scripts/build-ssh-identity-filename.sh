#!/bin/bash
#
# This is reusable code for building the ssh identity filename
#
# Copyright (c) 2022 STY-Holdings Inc
# MIT License
# Permission is hereby granted, free of charge, to any person obtaining a copy of this software and
# associated documentation files (the “Software”), to deal in the Software without restriction,
# including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
# and/or sell copies of the Software, and to permit persons to whom the Software is furnished to
# do so, subject to the following conditions:
#
# The above copyright notice and this permission notice shall be included in all copies or
# substantial portions of the Software.
#
# THE SOFTWARE IS PROVIDED “AS IS”, WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING
# BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
# NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,
# DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
# OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

# shellcheck disable=SC2034
function build_ssh_identity_filename() {
  local ssh_user=$1
  local local_user_home_directory=$2
  local environment_alias=$3
  local instance_number=$4

  export IDENTITY_FILENAME="$local_user_home_directory/.ssh/$ssh_user-$environment_alias-$instance_number"
}

# Test
#export WORKING_AS=savup
#export LOCAL_USER_HOME_DIRECTORY=/Users/syacko
#export ENVIRONMENT_ALIAS='lcl'
#export INSTANCE_NUMBER='0030'
#build_ssh_identity_filename $WORKING_AS $LOCAL_USER_HOME_DIRECTORY $ENVIRONMENT_ALIAS $INSTANCE_NUMBER
#echo "\$IDENTITY_FILENAME=$IDENTITY_FILENAME"
