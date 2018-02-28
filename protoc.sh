#!/bin/bash
PATH=$PATH:$(pwd)/bin protoc -Iproto --go_out=plugins=grpc:src/ proto/payments.proto
