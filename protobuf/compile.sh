#!/bin/bash -e
protoc -I=. --go_out=./generated/ ./protobuf/measurement.proto