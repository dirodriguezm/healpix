#!/bin/bash
cd internal/healpix_cxx
swig \
    -c++ \
    -go \
    -verbose \
    -intgosize 64 \
    $(pkg-config --cflags-only-I libsharp healpix_cxx) \
    healpix.i
