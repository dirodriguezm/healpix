# healpix

This is a library providing Go bindings to the [official C++
library](https://healpix.sourceforge.io/) for
[HEALPix](https://healpix.jpl.nasa.gov/), which is a scheme for partitioning a
sphere into numbered regions.

I wanted this to support implementing "cone search" with a B-Tree, so the only
APIs which are exposed are ones that assist with that goal. It's not hard to add
more if you want more of the features of the C++ library though, please open an
issue if you could use something more.

## Installation and use

You'll need to have the `healpix_cxx` library installed. That is, you'll need
`libhealpix_cxx.so`, and you'll need the `healpix_cxx` development headers.
These need to be on the machine that actually builds your Go code, as well as
where it executes (unless you build statically, see below).

This library uses [`pkg-config`](http://pkg-config.freedesktop.org/) to find the
healpix libraries.

The simplest way to do this may be to use your package manager, like with `sudo
apt install libhealpix-cxx-dev`. But...

### Building healpix_cxx from source
If you want to build from source, here's how I do it on Ubuntu:
```
# Install all required tools
sudo apt-get install autoconf automake libtool pkg-config

# Install cfitsio library
sudo apt install libcfitsio-dev libcfitsio8

# Download source
git checkout git@github.com:spenczar/healpixmirror.git
cd healpixmirror

# Prepare libsharp
cd src/common_libraries/libsharp
autoreconf -i
cd

# Prepare cxx
cd src/cxx
autoreconf -i
cd

# Configure system- you'll be prompted to say where libcfitsio.so is found
./configure

# Actually build
make -j

# Install pkg-config files
cp lib/pkgconfig/libsharp.pc /usr/share/pkgconfig/libsharp.pc
cp lib/pkgconfig/healpix_cxx.pc /usr/share/pkgconfig/healpix_cxx.pc
```

### Static builds

Fully-encapsulated static builds that link against C libraries are possible in
Go, but they require a special invocation of the `go build` command.
Specifically (at least on Linux, with Go 1.14):

```
go build -ldflags "-linkmode external -extldflags -static"
```

Your mileage may vary. This might be different for your system.


## Development

This library uses [SWIG](http://www.swig.org/) to generate wrappers around the
C++ `healpix_cxx` and `libsharp` libraries. Use the `run_swig.sh` script to
execute this. You may get some warnings.
