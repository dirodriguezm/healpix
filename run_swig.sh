#!/bin/bash
swig4.0 \
    -c++ \
    -go \
    -verbose \
    -intgosize 64 \
    $(pkg-config --cflags-only-I libsharp healpix_cxx) \
    healpix.i
