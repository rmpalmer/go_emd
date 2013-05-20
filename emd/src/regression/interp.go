package regression

import (
	"fmt"
	"util"
	"resample"
)

func TestInterp() {
	fmt.Printf ("Testing (polynomial) interpolation\n")
	
	x := []float64{0.0, 1.0, 2.0, 3.0}
	y := []float64{0.0, 1.0, 8.0, 27.0}
	 
	r := make([]float64, 10)
	
	r = resample.Resamp(x, y, r)
	
	util.DumpSlice("r",r)

	
}