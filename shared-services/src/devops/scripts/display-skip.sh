#!/bin/bash
#
# This is reusable code for displaying skip message on the screen
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

function display_skip_message() {
  if [[ ("$1" == "Y") || ("$1" == "y") ]]; then
    echo " Do you want to ${BLACK}${ON_YELLOW}SKIP${COLOR_OFF} this step? (${BLACK}${ON_YELLOW}Y${COLOR_OFF}/n)"
  else
    echo " Do you want to ${BLACK}${ON_YELLOW}SKIP${COLOR_OFF} this step? (y/${BLACK}${ON_YELLOW}N${COLOR_OFF})"
  fi
}
