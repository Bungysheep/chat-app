#!/bin/bash

protoc chat-protobuf/chat.proto --go_out=plugins=grpc:.