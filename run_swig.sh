#!/bin/bash
cd internal/healpix_cxx
ifile="healpix_amd64.i"
if [[ $(uname -a) == *"Darwin"* ]]; then
    ifile="healpix_darwin.i" 
fi
swig \
    -c++ \
    -go \
    -verbose \
    -intgosize 64 \
    -o healpix_wrap.cxx \
    $(pkg-config --cflags-only-I libsharp healpix_cxx) \
    $ifile
