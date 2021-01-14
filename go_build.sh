#!/bin/bash

go build -o ./build/tester -ldflags "-linkmode external -extldflags -static" ./cmd/tester
