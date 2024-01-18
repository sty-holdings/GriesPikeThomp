#!/bin/bash
#
# This will configure ssh to use an identity key.
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

# Main function of this script
function install_server_user() {
  local identity=$1
  local ssh_user=$2
  local dns_ip=$3
  local create_user=$4
  local parent_group=$5
  local login_allowed=$6

  #  echo "install_server_user()"
  #  echo "identity=$identity"
  #  echo "ssh_user=$ssh_user"
  #  echo "dns_ip=$dns_ip"
  #  echo "create_user=$create_user"
  #  echo "parent_group=$parent_group"
  #  echo "login_allowed=$login_allowed"

  find_string_in_remote_file "$identity" $ssh_user $dns_ip $create_user '/etc/passwd'
  # shellcheck disable=SC2154
  if [ "$find_string_in_remote_file_result" == "found" ]; then
    install_server_user_result="Server user ($create_user) already exist. No action taken."
  else
    if [ "$login_allowed" == true ]; then
      # shellcheck disable=SC2029
      install_server_user_result=$(ssh $identity $ssh_user@$dns_ip "sudo useradd --create-home --shell /bin/bash --gid $parent_group $create_user; sudo chmod 775 /home/$create_user; sudo chmod ug+s /home/$create_user; sudo passwd $create_user")
    else
      #      You can not use -gid for a user who is not able to log into the system.
      # shellcheck disable=SC2029
      install_server_user_result=$(ssh $identity $ssh_user@$dns_ip "sudo useradd --create-home --shell /usr/sbin/nologin $create_user; sudo chmod -R 775 /home/$create_user; sudo chmod ug+s /home/$create_user; sudo chgrp -R $parent_group /home/$create_user")
    fi
    if [ "$?" ]; then
      install_server_user_result="Server user ($create_user) created."
    else
      echo "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
      echo "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
      echo "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
      echo "ERROR:  $install_server_user_result"
      echo "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
      echo "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
      echo "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
      install_server_user_result="failed"
    fi
  fi
}

# Test
#. /Users/syacko/workspace/styh-dev/src/albert/core/devops/scripts/find-string-in-file.sh
#install_server_user "-i /Users/syacko/.ssh/savup-lcl-0030" savup savup-local-0030.savup.com savup-nats 'sty-holdings' false
#echo "\$install_server_user_result=$install_server_user_result"
