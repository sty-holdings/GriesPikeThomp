#!/bin/bash
#
# Name: loadTestingTokens.sh
#
# Description: This will display the alias file supplied.
#
# Why: For troubleshooting where the logs (in Papertrail or CloudWatch) do not provide
#      sufficient information
#
# Installation:
#
# Copyright (c) 2022-Present STY-Holdings Inc
# All Rights Reserved
#
#
# User Controlled
ROOT_DIRECTORY=/Users/syacko/workspace/styh-dev/src/albert
#
# System Controlled
CORE_ROOT_DIRECTORY=${ROOT_DIRECTORY}/core
export TEST_TOKEN_DIRECTORY=${CORE_ROOT_DIRECTORY}/testTokens

rm ${TEST_TOKEN_DIRECTORY}/token* ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh 2> /dev/null
envsubst < ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh.template > ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh
chmod 777 ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh
sh ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh
