#!/bin/bash
#
# This is reusable code for finding a directory
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
function find_directory() {
  local search_directory=$1
  local directory=$2

  if find $search_directory -type d -name "$directory" -print -quit; then
    find_directory_result="found"
  else
    find_directory_result="missing"
  fi
}

# Test
#find_string_in_file  'alias root-local-0030' '/Users/syacko/.oh-my-zsh/custom/cloud-aliases.zsh'
#echo "\$find_string_in_file_result=$find_string_in_file_result"

function find_remote_directory() {
  local identity=$1
  local user=$2
  local dns_ip=$3
  local search_directory=$4
  local directory=$5
  set find_remote_directory_result

#  echo
#  echo "identity=$identity"
#  echo "user=$user"
#  echo "dns_ip=$dns_ip"
#  echo "search_directory=$search_directory"
#  echo "file=$directory"
#  echo

  find_remote_directory_result=$(ssh $identity $user@$dns_ip "if [ \$(find $search_directory -type d -name $directory -print -quit) ]; then echo 'found'; else echo 'missing'; fi")
}

# Test
#export SERVER_NAME=savup-http
#find_remote_directory '-i /Users/syacko/.ssh/savup-local-0030' savup 154.12.225.56 /home/$SERVER_NAME bin
#echo $find_remote_directory_result
#echo '========'
#find_remote_directory '-i /Users/syacko/.ssh/savup-local-0030' savup 154.12.225.56 '/home' 'BILL'
#echo $find_remote_directory_result
