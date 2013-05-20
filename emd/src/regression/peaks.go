package regression

import (
	"fmt"
	"util"
	"peaks"
)

func TestPeaks() {

	const n = 100
	
	var dArray      [n] float64
	var minValArray [n] float64
	var minLocArray [n] float64
	var maxValArray [n] float64
	var maxLocArray [n] float64
	
	var zeroCrossings int
	
	d := dArray[:]
	
	minVals := minValArray[:]
	minLocs := minLocArray[:]
	maxVals := maxValArray[:]
	maxLocs := maxLocArray[:]
	
	util.Synth(d)
	
	minVals, minLocs, maxVals, maxLocs, zeroCrossings = peaks.Detect(d, 
		minVals, minLocs, 
		maxVals, maxLocs,
		0.1, false)
		
	fmt.Printf("found %d minima and %d maxima and %d zeros\n",
		len(minVals), len(maxVals), zeroCrossings)
	
	fmt.Printf("minima:\n")
	util.DumpSlice("min locs",minLocs);
	util.DumpSlice("min vals",minVals);
	
	fmt.Printf("maxima:\n");
	util.DumpSlice("max locs",maxLocs)
	util.DumpSlice("max vals",maxVals)
		

}

