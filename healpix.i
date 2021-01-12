%module healpix
%{
#include "vec3.h"
#include "geom_utils.h"
#include "healpix_base.h"
#include "healpix_map.h"
%}
// 3-vectors
%rename(subtractInPlace) vec3_t::operator-=;
%rename(add) operator+;
%rename(addInPlace) operator+=;
%rename(multiply) operator*;
%rename(divide) operator/;
%rename(multiplyInPlace) operator*=;

%rename(invert) vec3_t::operator-() const;
%rename(subtract) vec3_t::operator-(const vec3_t &) const;
%include "vec3.h"
%template(Vec3Floats) vec3_t<double>;


%include "healpix_base.h"
%include "healpix_map.h"

enum Healpix_Ordering_Scheme {
  RING, /*!< RING scheme */
  NEST  /*!< NESTED scheme */
};
%rename(HEALPixOrderingScheme) Healpix_Ordering_Scheme;

%rename(HEALPixMap) Healpix_Map<double>;
class Healpix_Map<double> {
public:
  Healpix_Map();
  Healpix_Map(int order, Healpix_Ordering_Scheme scheme);
  long pix2ring(long);
};

%include "geom_utils.h"
%include "std_vector.i"
