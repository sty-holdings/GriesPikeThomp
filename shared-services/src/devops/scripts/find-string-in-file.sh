#!/bin/bash
#
# This is reusable code for logging and execute NATS NSC commands
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

# shellcheck disable=SC2129
function find_string_in_file() {
  local search_string=$1
  local file=$2

  #  echo
  #  echo search_string=$1
  #  echo file=$2
  #  echo

  if grep -q "$search_string" "$file"; then
    find_string_in_file_result="found"
  else
    find_string_in_file_result="missing"
  fi
}

# Test
#find_string_in_file  'alias root-local-0030' '/Users/syacko/.oh-my-zsh/custom/cloud-aliases.zsh'
#echo "\$find_string_in_file_result=$find_string_in_file_result"

# shellcheck disable=SC2029
function find_string_in_remote_file() {
  local identity=$1
  local user=$2
  local dns_ip=$3
  local search_string=$4
  local file=$5

#  echo "find_string_in_remote_file()"
#  echo "identity=$identity"
#  echo "user=$user"
#  echo "dns_ip=$dns_ip"
#  echo "search_string=$search_string"
#  echo "file=$file"

  find_string_in_remote_file_result=$(ssh $identity $user@$dns_ip "grep -q '$search_string' $file && echo 'found' || echo 'missing'")
}

# Test
#find_string_in_remote_file '-i /Users/syacko/.ssh/savup-local-0030' savup 154.12.225.56 'nats-server' '/tmp/processes.tmp'
#echo "find nats-server:  $find_string_in_remote_file_result"
#find_string_in_remote_file '-i /Users/syacko/.ssh/savup-local-0030' savup 154.12.225.56 'sshd:' '/tmp/processes.tmp'
#echo "find sshd: $find_string_in_remote_file_result"
#find_string_in_remote_file '-i /Users/syacko/.ssh/savup-local-0030' savup 154.12.225.56 $NATS_SYSTEM_USER '/etc/passwd'
#echo "find sshd: $find_string_in_remote_file_result"

# shellcheck disable=SC2029
# shellcheck disable=SC2034
function find_string_exclude_string_in_remote_file() {
  local identity=$1
  local user=$2
  local dns_ip=$3
  local search_string=$4
  local exclude_string=$5
  local file=$6

  #  echo "identity=$identity"
  #  echo "user=$user"
  #  echo "dns_ip=$dns_ip"
  #  echo "search_string=$search_string"
  #  echo "exclude_string=$exclude_string"
  #  echo "file=$file"

  find_string_in_remote_file_result=$(ssh $identity $user@$dns_ip "grep -q -E '$search_string | ^$exclude_string' $file && echo 'found' || echo 'missing'")
}

# Test
#find_string_in_remote_file '-i /Users/syacko/.ssh/savup-local-0030' savup 154.12.225.56 'nats-server' '/tmp/processes.tmp'
#echo "find nats-server:  $find_string_in_remote_file_result"
#find_string_in_remote_file '-i /Users/syacko/.ssh/savup-local-0030' savup 154.12.225.56 'sshd:' '/tmp/processes.tmp'
#echo "find sshd: $find_string_in_remote_file_result"
#find_string_in_remote_file '-i /Users/syacko/.ssh/savup-local-0030' savup 154.12.225.56 $NATS_SYSTEM_USER '/etc/passwd'
#echo "find sshd: $find_string_in_remote_file_result"
