package main

import (
	"fmt"
	"math"
	"os"

	"github.com/spenczar/healpix"
)

func main() {
	mapper, err := healpix.NewHEALPixMapper(18, healpix.Nest)
	if err != nil {
		fmt.Printf("error: %v\n", err)
		os.Exit(1)
	}
	pointing := healpix.Pointing{
		Phi:   0.1,
		Theta: 0.5,
	}
	pixels := mapper.QueryDisc(pointing, 2*math.Pi)
	for _, pr := range pixels {
		fmt.Printf("[%d, %d)\n", pr.Start, pr.Stop)
	}
}
