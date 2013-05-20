package algorithm

import (
	"algebra"
	"resample"
	"peaks"
	"math"
	"fmt"
	"util"
)

func Sift(h []float64) (float64, int, int, int) {
	
	inEnergy := util.Energy(h)
	n := len(h)
	
	minVals := make([]float64, n)
	minLocs := make([]float64, n)
	maxVals := make([]float64, n)
	maxLocs := make([]float64, n)
	upperEnvelope := make([]float64, n)
	lowerEnvelope := make([]float64, n)
	var mean float64
	var smallestMean float64 
	var zeroCrossings int
	
	detail := algebra.NewMat(n, 5)
	
	minVals, minLocs, maxVals, maxLocs, zeroCrossings = peaks.Detect(h, 
		minVals, minLocs, 
		maxVals, maxLocs,
		0.1, true)
		
	fmt.Printf("found %d minima and %d maxima and %d zero crossings\n",
		len(minVals), len(maxVals), zeroCrossings)
		
	util.DumpSlice("minima",minVals)
	util.DumpSlice("minlocs",minLocs)
	util.DumpSlice("maxima",maxVals)
	util.DumpSlice("maxlocs",maxLocs)
		
	fmt.Printf("getting upper envelope\n")
	upperEnvelope = resample.Resamp(maxLocs, maxVals, upperEnvelope)
	//util.DumpSlice("upper",upperEnvelope)
	fmt.Printf("getting lower envelope\n")
	lowerEnvelope = resample.Resamp(minLocs, minVals, lowerEnvelope)
	//util.DumpSlice("lower",lowerEnvelope)

	for i:=0; i<n; i++ {
		detail[i][0] = h[i]
		detail[i][1] = upperEnvelope[i]
		detail[i][2] = lowerEnvelope[i]
	}
	
			
	fmt.Printf("subtracting the mean\n")
	for i:=0; i<n; i++ {
		mean = 0.5 * (upperEnvelope[i] + lowerEnvelope[i])
		h[i] = h[i] - mean
		detail[i][3] = mean
		detail[i][4] = h[i]
		if (i == 0) {
			smallestMean = math.Abs(mean)
		} else {
			smallestMean = math.Min(math.Abs(mean), smallestMean)
		}
	}
	
	algebra.MatInfo(detail)
	
	outEnergy := util.Energy(h)
	
	fmt.Printf("... reduced energy from %f to %f\n",inEnergy,outEnergy)
	
	return smallestMean, zeroCrossings, len(minVals), len(maxVals)
}