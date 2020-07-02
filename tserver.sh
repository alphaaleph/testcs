#!/bin/bash

echo ">>  create bin directory"
mkdir -p ./bin

echo ">>  build the test server"
go build -o ./bin/testcs .
chmod a+x ./bin/testcs

#execute entries
printf "\n\n>>  execute the server\n\n"
./bin/testcs

