package peaks

import (
	"math"
)

func Find (data *[]float64, 
			minCnt *int, maxCnt *int, 
			frac float64, endsArePeaks bool) {
			
    // size of input data      	
    n := len (*data)
    
	// local variables to avoid awkward dereferencing
	myMinCnt := 0
	myMaxCnt := 0
	
	// will start off looking for a local maxima
	//lookForMax := true
	
	// a sample value
	//val := 0.0
	
	// positions of current local minimum and maximum
	//minPos := 0
	//maxPos := 0
	
	// get the min and max of the entire array
	mn := (*data)[0]
	mx := mn
	for i:=1; i<n; i++ {
		mn = math.Min(mn, (*data)[i])
		mx = math.Max(mx, (*data)[i])
	}
	
	// samples should change by at least this much 
	// going from a minimum to a maximum
	//delta := frac * (mx - mn)
	
	// bounds to search over
	//loopMin := 0
	//loopMax := n
	
	
	// return counts
	*minCnt = myMinCnt
	*maxCnt = myMaxCnt

}