package peakdet

import (
	"math"
	)

func Peaks (data *[]float64, 
			minVals *[]float64, minLocs *[]float64, 
			maxVals *[]float64, maxLocs *[]float64, 
			minCnt *int, maxCnt *int, 
			frac float64, endsArePeaks bool) {
    
    // size of input data      	
    n := len (*data)
    
	// local variables to avoid awkward dereferencing
	myMinCnt := 0
	myMaxCnt := 0
	
	// will start off looking for a local maxima
	lookForMax := true
	
	// a sample value
	val := 0.0
	
	// positions of current local minimum and maximum
	minPos := 0
	maxPos := 0
	
	// get the min and max of the entire array
	mn := (*data)[0]
	mx := mn
	for i:=1; i<n; i++ {
		mn = math.Min(mn, (*data)[i])
		mx = math.Max(mx, (*data)[i])
	}
	
	// samples should change by at least this much 
	// going from a minimum to a maximum
	delta := frac * (mx - mn)
	
	// bounds to search over
	loopMin := 0
	loopMax := n
	
	// force first point to be used as local min and max
	// this block captures the start point;
	// a later block will capture the end point
	if endsArePeaks {
	
		// record first point as the first local maximum
		(*maxLocs)[0] = 0.0
		(*maxVals)[0] = (*data)[0]
		*maxCnt = 1
		mx = (*data)[0]
		maxPos = 0
		
		/// record the first point as the first local minimum
		(*minLocs)[0] = 0.0
		(*minVals)[0] = (*data)[0]
		myMinCnt = 1
		mn = (*data)[0]
		minPos = 0
		
		// subsequent loop can shrink a bit
		loopMin = loopMin + 1
		loopMax = loopMax - 1
	}
	
	// loop over all points
	for i:=loopMin; i<loopMax; i++ {
		 
		 val = (*data)[i]
		 
		 // track (update position of) local minimum and maximum
		 // if mononotically changing.
		 if (val > mx) {
		 	mx = val
		 	maxPos = i
		 }
		 
		 if (val < mn) {
		 	mn = val
		 	minPos = i
		 }
		 
		 // if hunting for a local maximum..
		 if lookForMax {
		 
		 	// if there is a reversal, then switch to looking for min
		 	if (val < (mx-delta)) {
		 		mn = val
		 		minPos = i
		 		lookForMax = false
		 		if (math.Abs((*maxLocs)[myMaxCnt]-float64(maxPos)) < 0.1) {
		 			continue
		 		}
		 		// add the local maximum to the list
		 		myMaxCnt++
		 		(*maxLocs)[myMaxCnt] = float64(maxPos)
		 		(*maxVals)[myMaxCnt] = (*data)[maxPos]
			 }
					 
		// if hunting for a local minumum
		} else {
		
			// if there is a reversal, then switch to looking for max
			if (val > (mn+delta)) {
				mx = val
				maxPos = i
				lookForMax = true
				if (math.Abs((*minLocs)[myMinCnt]-float64(minPos)) < 0.1) {
					continue
				}
				// add the local minimum to the list
				myMinCnt++
				(*minLocs)[myMinCnt] = float64(minPos)
				(*minVals)[myMinCnt] = (*data)[minPos]
			}
		}
			
	}  // end of loop over points
	
	// optionally force last point to be used both as local minimum and maximum
	// ..but make sure not to add to it the lists if it is already there. 
	if endsArePeaks {
		
		if (math.Abs((*maxLocs)[myMaxCnt] - float64(n)) > 0.1) {
			myMaxCnt++
			(*maxLocs)[myMaxCnt] = float64(n)
			(*maxVals)[myMaxCnt] = (*data)[n]
		}
		
		if (math.Abs((*minLocs)[myMinCnt] - float64(n)) > 0.1) {
			myMinCnt++
			(*minLocs)[myMinCnt] = float64(n)
			(*minVals)[myMinCnt] = (*data)[n]
		}
		
	}
	
	// return counts
	*minCnt = myMinCnt
	*maxCnt = myMaxCnt
}	