#!/bin/bash

PROTO_DIR=api
OUTPUT_DIR=pkg/grpc

CompileProto()
{
    protoc --go_out=$OUTPUT_DIR \
        --go_opt=paths=source_relative \
        --go-grpc_out=$OUTPUT_DIR \
        --go-grpc_opt=paths=source_relative \
        $1
}

for file in `\find $PROTO_DIR -name '*.proto'`; do
    echo $file
    CompileProto $file
done
