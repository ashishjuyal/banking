#!/bin/bash
printf "\n"
echo "**** PLEASE ENSURE THE MOCKS ARE GENERATED BEFORE RUNNING THE UNIT TESTS ****"
echo "     To generate the mocks run the ./generate-mocks.sh file "
printf "\n"
go test -v ./...