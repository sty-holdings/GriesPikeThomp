# Makefile_testing_core.mk
#
# This will run the core tests based on the parameters provided
#
# NOTE: If you have an improvement or a correct, please make the change and check it into the repository.
#
# User Controlled
ROOT_DIRECTORY=/Users/syacko/workspace/styh-dev/src/albert
SERVER_CONFIG_FILE=${ROOT_DIRECTORY}/savup-server/config/development/local-savup-config-with-TLS.json
#
# System Controlled
CORE_ROOT_DIRECTORY=${ROOT_DIRECTORY}/core
export TEST_TOKEN_DIRECTORY=${CORE_ROOT_DIRECTORY}/testTokens

define run_Go_Test
	echo $(1)
	go test -coverprofile ${CORE_ROOT_DIRECTORY}/$(1)/coverage.out ${CORE_ROOT_DIRECTORY}/$(1)/*.go
	echo -------
endef

define run_Firebase_Auth_Test
	echo $(1)
	go test -coverprofile ${CORE_ROOT_DIRECTORY}/$(1)/coverage.out ${CORE_ROOT_DIRECTORY}/$(1) -run TestNewFirebaseApp
	go test -coverprofile ${CORE_ROOT_DIRECTORY}/$(1)/coverage.out ${CORE_ROOT_DIRECTORY}/$(1) -run TestFindFirebaseAuthUser
	echo -------
endef

define run_AWS_No_Token_Test
	echo $(1)
	go test -coverprofile ${CORE_ROOT_DIRECTORY}/$(1)/coverage.out ${CORE_ROOT_DIRECTORY}/$(1) -run TestNewAWSSession
	go test -coverprofile ${CORE_ROOT_DIRECTORY}/$(1)/coverage.out ${CORE_ROOT_DIRECTORY}/$(1) -run TestAWSHelper_ConfirmUser
	go test -coverprofile ${CORE_ROOT_DIRECTORY}/$(1)/coverage.out ${CORE_ROOT_DIRECTORY}/$(1) -run TestGetPublicKeySet
	echo -------
endef

all: build buildTestTokenFiles awsTest errorTest firebaseTest firestoreTest helpersTest jwtTest natsTest sendgridTest validatorsTest
noAccessToken: build awsNoTokenTest errorTest firebaseTest firestoreTest helpersTest jwtTest natsTest sendgridTest validatorsTest

build:
	go build ${CORE_ROOT_DIRECTORY}/coreAWS/AWSServices.go
	go build ${CORE_ROOT_DIRECTORY}/coreError/ErrorService.go
	go build ${CORE_ROOT_DIRECTORY}/coreFirebase/FirebaseServices.go
	go build ${CORE_ROOT_DIRECTORY}/coreFirestore/FirestoreServices.go
	go build ${CORE_ROOT_DIRECTORY}/coreHelpers/helpers.go
	go build ${CORE_ROOT_DIRECTORY}/coreJWT/JWTServices.go
	go build ${CORE_ROOT_DIRECTORY}/coreNATS/NATSServices.go
	go build ${CORE_ROOT_DIRECTORY}/coreSendGrid/SendGridEmailService.go
	go build ${CORE_ROOT_DIRECTORY}/coreValidators/ValidatorServices.go

buildTestTokenFiles:
	-rm ${TEST_TOKEN_DIRECTORY}/token* ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh

	envsubst < ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh.template > ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh
	chmod 777 ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh
	sh ${CORE_ROOT_DIRECTORY}/build_test_token_files.sh

awsTest:
	$(call run_Go_Test,coreAWS)

awsNoTokenTest:
	$(call run_AWS_No_Token_Test,coreAWS)

errorTest:
	$(call run_Go_Test,coreError)

firebaseTest:
	cat ${ROOT_DIRECTORY}/savup-server/config/development/local-savup-config-with-TLS.json | awk ' { if ( /COGNITO/ ) { print > "/tmp/cognito.tmp" } } '
ifeq ($(wildcard "/tmp/cognito.tmp"),)
	$(call run_Firebase_Auth_Test,coreFirebase)
else
	$(call run_Go_Test,coreFirebase)
endif

firestoreTest:
	$(call run_Go_Test,coreFirestore)

helpersTest:
	$(call run_Go_Test,coreHelpers)

jwtTest:
	$(call run_Go_Test,coreJWT)

natsTest:
	$(call run_Go_Test,coreNATS)

sendgridTest:
	$(call run_Go_Test,coreSendGrid)

validatorsTest:
	$(call run_Go_Test,coreValidators)
