# healpix

This is a library providing Go bindings to the [official C++
library](https://healpix.sourceforge.io/) for
[HEALPix](https://healpix.jpl.nasa.gov/), which is a scheme for partitioning a
sphere into numbered regions.

I wanted this to support implementing "cone search" with a B-Tree, so the only
APIs which are exposed are ones that assist with that goal. It's not hard to add
more if you want more of the features of the C++ library though, please open an
issue if you could use something more.
