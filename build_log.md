In order to build this:

First, build healpix_cxx and libsharp.

This will require the GNU build tools.

To do that:
```
svn checkout https://svn.code.sf.net/p/healpix/code/trunk healpix-code
cd healpix-code

cd src/healpix_cxx
autoreconf -i
popd

cd src/common_libraries/libsharp
autoreconf -i
popd

./configure --auto=cxx
make
cp lib/pkgconfig/healpix_cxx.pc /usr/share/pkgconfig/healpix_cxx.pc
cp lib/pkgconfig/libsharp.pc /usr/share/pkgconfig/libsharp.pc
```

Now, you should be able to run SWIG:
```
swig4.0 -c++ -go -verbose -intgosize 64 $(pkg-config --cflags-only-I libsharp) $(pkg-config --cflags-only-I healpix_cxx) healpix.i
```

And you can go build:
```
CGO_CXXFLAGS="$(pkg-config --cflags healpix_cxx libsharp)" CGO_LDFLAGS="$(pkg-config --libs healpix_cxx libsharp)" go build ./cmd/tester
```

Running the code requires explicitly linking the library:
(TODO: this is specific to my machine...)

```
LD_LIBRARY_PATH=$PWD/../../3p/healpix-code/lib ./tester
```
