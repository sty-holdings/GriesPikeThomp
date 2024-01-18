#!/bin/bash
#
# This will check if a process is running on the system
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

set -eo pipefail

# shellcheck disable=SC2034
# shellcheck disable=SC2154
function process_running() {
  local identity=$1
  local user=$2
  local dns_ip=$3
  local process_name=$4
  local exclude_string=$5

#  echo "identity=$identity"
#  echo "user=$user"
#  echo "dns_ip=$dns_ip"
#  echo "process_name=$process_name"
#  echo "exclude_string=$exclude_string"

  ssh $identity $user@$dns_ip "ps aux > /tmp/processes.tmp"
  find_string_exclude_string_in_remote_file "$identity" $user $dns_ip $process_name $exclude_string '/tmp/processes.tmp'
  process_running_result=$find_string_in_remote_file_result
}

# Test
#. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/find-string-in-file.sh
#process_running "-i /Users/syacko/.ssh/savup-local-0030" savup savup-local-0030.savup.com 'nats-server'
#echo "\$process_running_result=$process_running_result"
#process_running "-i /Users/syacko/.ssh/savup-local-0030" savup savup-local-0030.savup.com 'sshd:'
#echo "\$process_running_result=$process_running_result"
