#!/bin/sh
 
#FIRST_ARGUMENT="$1"
#echo "Hello, world $FIRST_ARGUMENT!"
export PATH=$PATH:$GOPATH/bin
protoc  proto/address.proto  -I. --go_out=plugins=grpc:$GOPATH/src
protoc  proto/person.proto  -I. --go_out=plugins=grpc:$GOPATH/src
protoc  proto/ack.proto  -I. --go_out=plugins=grpc:$GOPATH/src