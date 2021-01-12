package main

import (
	"fmt"

	"github.com/spenczar/healpix"
)

func main() {
	v := healpix.NewVec3Floats(0.0, 0.2, 0.3)
	defer healpix.DeleteVec3Floats(v)
	fmt.Println(v.Length())

}
