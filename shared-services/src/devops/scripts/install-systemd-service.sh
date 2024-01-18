#!/bin/bash
#
# This will create a systemd service file and reload the daemon
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
# shellcheck disable=SC2016
# shellcheck disable=SC2029
function install_systemd_service() {
  local identity=$1
  local ssh_user=$2
  local dns_ipv4=$3
  local server_name=$4
  export install_directory=$5 # This is a local variable, but envsubst only works with environment variable, so it must be exported.
  export service_user_name=$6 # This is a local variable, but envsubst only works with environment variable, so it must be exported.
  local template_directory=$7
  local template_filename=$8
  local service_filename=$9

#  echo
#  echo identity=$identity
#  echo ssh_user=$ssh_user
#  echo dns_ipv4=$dns_ipv4
#  echo server_name=$server_name
#  echo install_directory=$install_directory
#  echo service_user_name=$service_user_name
#  echo template_directory=$template_directory
#  echo template_filename=$template_filename
#  echo service_filename=$service_filename
#  echo

  echo "Installing systemd service file for $server_name."
  envsubst <$template_directory/$template_filename >/tmp/$service_filename
  scp $identity /tmp/$service_filename $ssh_user@$dns_ipv4:$install_directory/$service_filename
  ssh $identity $ssh_user@$dns_ipv4 "sudo mv $install_directory/$service_filename /etc/systemd/system/$service_filename; sudo chmod 755 /etc/systemd/system/$service_filename; sudo systemctl daemon-reload;"
}

# Test
#export IDENTITY="-i /Users/syacko/.ssh/savup-local-0030"
#export WORKING_AS=savup
#export INSTANCE_DNS_IPV4=154.12.225.56
#export NATS_SERVER_USER=savup-nats
#export NATS_INSTALL_DIRECTORY=/home/NATS
#export ROOT_DIRECTORY=/Users/syacko/workspace/styh-dev/src/albert
#export TEMPLATE_DIRECTORY=$ROOT_DIRECTORY/savup-nats/build-deploy/templates
#install_systemd_service $IDENTITY $WORKING_AS $INSTANCE_DNS_IPV4 $NATS_SERVER_USER $NATS_INSTALL_DIRECTORY $WORKING_AS $TEMPLATE_DIRECTORY 'nats-server-servicefile.template' 'nats-server.service'
