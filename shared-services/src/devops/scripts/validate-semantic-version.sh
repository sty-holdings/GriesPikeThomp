#!/bin/bash
#
# This is reusable code for validating a semantic number up to 9999.9999.9999
# A return value of zero means good, 1 means the first node is bad, 2 the second,
# and 3 for the third
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

function validate_semantic {
  local count=1
  local nodeCount=0
  re='^[0-9]+$'
  IFS="."
  for node in ${1}; do
    if ! [[ $node =~ $re ]]; then
      break
    fi
    if [[ -z "$node" || $node -lt 0 || $node -gt 9999 ]]; then
      break
    fi
    count=$((count + 1))
    nodeCount=$((nodeCount + 1))
  done

  if [[ $count -eq 4 ]]; then
    validate_semantic_result='ok' # Valid semantic number
  else
    validate_semantic_result='failed'
  fi
}

# Test
#result=$(validate_semantic 2023)
#echo $result
#result=$(validate_semantic 2023.0)
#echo $result
#result=$(validate_semantic 2023..0)
#echo $result
#result=$(validate_semantic 99999.0.0)
#echo $result
#result=$(validate_semantic 2023.99999.0)
#echo $result
#result=$(validate_semantic 2023.9999.99999)
#echo $result
#result=$(validate_semantic 2023.9999.aa)
#echo $result
#result=$(validate_semantic 2023.aa.99)
#echo $result
#result=$(validate_semantic aa.99.99)
#echo $result
#result=$(validate_semantic 99.99.99)
#echo $result
