package regression

import (
	"resample"
	"util"
)

func TestResamp() {

	const n = 100;

	regular := make([]float64, n)
	
	abcissa  := []float64 {2, 19, 36, 53, 80}
	ordinate := []float64 {1.503, 1.967, 1.810, 1.089, 0.5882}
	//abcissa  := []float64 {2, 19, 36, 53}
	//ordinate := []float64 {1.503, 1.967, 1.810, 1.089}
	
	regular = resample.Resamp(abcissa, ordinate, regular)
	
	util.DumpSlice("regular",regular)
	
}
