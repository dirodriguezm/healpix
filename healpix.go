package healpix

import (
	"errors"
	"runtime"

	"github.com/spenczar/healpix/internal/healpix_cxx"
)

type Pointing struct {
	Theta float64
	Phi   float64
}

func (p Pointing) to_c() healpix_cxx.Pointing {
	return healpix_cxx.NewPointing(p.Theta, p.Phi)
}

func ptgFromC(cptg healpix_cxx.Pointing, destroy bool) Pointing {
	ptg := Pointing{
		Theta: cptg.GetTheta(),
		Phi:   cptg.GetPhi(),
	}
	if destroy {
		healpix_cxx.DeletePointing(cptg)
	}
	return ptg
}

type OrderingScheme int

const (
	Ring OrderingScheme = 1
	Nest OrderingScheme = 1
)

func (s OrderingScheme) to_c() healpix_cxx.Healpix_Ordering_Scheme {
	if s == Ring {
		return healpix_cxx.RING
	}
	if s == Nest {
		return healpix_cxx.NEST
	}
	panic("invalid ordering scheme")
}

type HEALPixMapper struct {
	cobj *healpix_cxx.SwigcptrHealpix_Base
}

func NewHEALPixMapper(order int, scheme OrderingScheme) (*HEALPixMapper, error) {
	if order < 0 || order > 13 {
		return nil, errors.New("invalid order, must be between 0 and 14")
	}

	cobj := healpix_cxx.NewHealpix_Base(order, scheme.to_c()).(healpix_cxx.SwigcptrHealpix_Base)

	mapper := HEALPixMapper{
		cobj: &cobj,
	}
	runtime.SetFinalizer(&mapper, func(m *HEALPixMapper) {
		healpix_cxx.DeleteHealpix_Base(m.cobj)
		m.cobj = nil
	})
	return &mapper, nil
}

// PixelAt returns the number of the pixel which contains the given angular
// coordinates indicated by ptg.
func (m *HEALPixMapper) PixelAt(ptg Pointing) int {
	return m.cobj.Ang2pix(ptg.to_c())
}

// PointingToCenter returns a pointing towards the center of the pixel with the
// given number.
func (m *HEALPixMapper) PointingToCenter(pixel int) Pointing {
	return ptgFromC(m.cobj.Pix2ang(pixel), true)
}

// QueryDisc returns the set of all pixels whose centers lie within a disk. The
// disc is centered at pointing, and has a radius of r radians.
func (m *HEALPixMapper) QueryDisc(pointing Pointing, r float64) []PixelRange {
	rangeset := m.cobj.Query_disc__SWIG_1(pointing.to_c(), r)
	defer healpix_cxx.DeleteRangeset(rangeset)
	data := rangeset.Data()

	var ranges []PixelRange
	for i := int64(0); i < data.Size(); i += 2 {
		pr := PixelRange{
			Start: data.Get(int(i)),
			Stop:  data.Get(int(i + 1)),
		}
		ranges = append(ranges, pr)
	}
	return ranges
}

// QueryDiscInclusive returns the set of all pixels which overlap with the disk
// defined by pointing and radius r (measured in radians). resolution should be
// an integer which determines the resolution used for the overlapping test.
//
// For Nested HEALPix, resolution must be a power of 2. For Ring, it can be any
// positive integer. A typical choice would be 4.
//
// Note that this method may return some pixels which don't overlap wiht the
// disk at all. The higher resolution is chosen, the fewer false positives are
// returned, at the cost of increased run time.
//
// This method is more efficient in the Ring scheme.
func (m *HEALPixMapper) QueryDiscInclusive(pointing Pointing, r float64, resolution int) []PixelRange {
	rangeset := m.cobj.Query_disc_inclusive__SWIG_2(pointing.to_c(), r, resolution)
	defer healpix_cxx.DeleteRangeset(rangeset)

	data := rangeset.Data()
	var ranges []PixelRange
	for i := int64(0); i < data.Size(); i += 2 {
		pr := PixelRange{
			Start: data.Get(int(i)),
			Stop:  data.Get(int(i + 1)),
		}
		ranges = append(ranges, pr)
	}
	return ranges
}

// PixelRange represents a contiguous sequence of pixels. Its Stop field
// indicates the first pixel in the range, and its Stop field indicates the
// first pixel which is not in the range (in other words, this represents
// `[Start, Stop)`).
type PixelRange struct {
	Start, Stop int
}
