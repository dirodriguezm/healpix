#!/bin/bash

export CGO_CXXFLAGS="$(pkg-config --cflags healpix_cxx libsharp)"
export CGO_LDFLAGS="$(pkg-config --libs healpix_cxx libsharp)"
go build -o ./build/tester -ldflags "-linkmode external -extldflags -static" ./cmd/tester
