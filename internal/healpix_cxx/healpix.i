%module healpix_cxx
%{
#include "vec3.h"
#include "geom_utils.h"
#include "healpix_base.h"
#include "healpix_map.h"
%}

// General datatypes
%include "stdint.i"
typedef int8_t int8;
typedef uint8_t uint8;

typedef int16_t int16;
typedef uint16_t uint16;

typedef int32_t int32;
typedef uint32_t uint32;

typedef int64_t int64;
typedef uint64_t uint64;

/*! unsigned integer type which should be used for array sizes */
namespace std {
  typedef size_t tsize;
/*! signed integer type which should be used for relative array indices */
  typedef ptrdiff_t tdiff;
}

// Pointings
%include "pointing.h"

// std::vector
%include "std_vector.i"
namespace std {
  %template(VectorF64) vector<double>;
  %template(VectorF32) vector<float>;
  %template(VectorI64) vector<int64_t>;
  %template(VectorInt) vector<int>;
  %template(VectorPointing) vector<pointing>;
};

// 3-vectors
%rename(subtractInPlace) vec3_t::operator-=;
%rename(add) vec3_t::operator+;
%rename(addInPlace) vec3_t::operator+=;
%rename(multiply) vec3_t::operator*;
%rename(divide) vec3_t::operator/;
%rename(multiplyInPlace) vec3_t::operator*=;
%rename(invert) vec3_t::operator-() const;
%rename(subtract) vec3_t::operator-(const vec3_t &) const;
%include "vec3.h"
%template(Vec3Floats) vec3_t<double>;

// Rangesets
%nodefaultctor rangeset;
%include "rangeset.h"
%template(Rangeset) rangeset<int64_t>;

// Ordering schemes
%include "healpix_tables.h"

// Bases
%include "healpix_base.h"
%template(Healpix_Base) T_Healpix_Base<int64_t>;


%include "geom_utils.h"
