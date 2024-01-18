#!/bin/bash
#
# This is reusable code for generating a private/public key
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

# shellcheck disable=SC2016
function build_private_public_key() {
  if [ -f $IDENTITY_FILENAME ]; then
    echo "Backing up $IDENTITY_FILENAME to $IDENTITY_FILENAME.bkup"
    cp $IDENTITY_FILENAME $IDENTITY_FILENAME.bkup
  fi
  echo "Generating private and public key at $IDENTITY_FILENAME"
  ssh-keygen -b 2048 -t rsa -f $1 -q -N "" -C "$SERVER_INSTANCE_IPV4 key pair for $WORKING_AS" <<<y >/dev/null 2>&1
}

# Test
#export IDENTITY_FILENAME=/Users/syacko/.ssh/root-lcl-0030
#export SERVER_INSTANCE_IPV4=savup-local-0030.savup.com
#export WORKING_AS=savup
#build_private_public_key /tmp/test_Key_$WORKING_AS
#cat /tmp/test_Key
## shellcheck disable=SC2028
#echo "\n"
#cat /tmp/test_Key_$WORKING_AS.pub
# shellcheck disable=SC2028
#echo "\n"
