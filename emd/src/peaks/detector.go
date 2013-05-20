package peaks

import (
	"util"
	"math"
)

func Detect(data []float64, 
	minVals []float64, minLocs []float64, 
	maxVals []float64, maxLocs []float64, 
	sens float64, endsArePeaks bool) ([]float64, []float64, []float64, []float64, int) {	

	// reset slices to their full capacities
	minVals = minVals[0:cap(minVals)]
	minLocs = minLocs[0:cap(minLocs)]
	maxVals = maxVals[0:cap(maxVals)]
	maxLocs = maxLocs[0:cap(maxLocs)]
	
	// initialize local variables
	n  := len(data)
	minCnt := 0
	maxCnt := 0
	maxPos := 0
	minPos := 0
	lookForMax := true
	var val float64
	zeroCrossings := 0
	lastSign := math.Signbit(data[0])
	thisSign := lastSign
	
	// range of samples to loop over
	loopMin := 0
	loopMax := n
	
	// get range of input values
	mn := util.SliceMin(data)
	mx := util.SliceMax(data)
	delta := sens * math.Abs(mx - mn)
	
	if endsArePeaks {
		// force first data to be used as the first local max 
		maxLocs[0] = float64(0)
		maxVals[0] = data[0]
		mx = data[0]
		maxPos = 0
		maxCnt = maxCnt + 1
		
		// also as the first local min
		minLocs[0] = float64(0)
		minVals[0] = data[0]
		mn = data[0]
		minPos = 0
		minCnt = minCnt + 1
		
		// and shorten the main loop
		loopMin = loopMin + 1
		loopMax = loopMax - 1
	}
	
	for i:=loopMin; i<loopMax; i++ {
	
		thisSign = math.Signbit(data[i])
		if (thisSign != lastSign) {
			zeroCrossings = zeroCrossings + 1
		}
		lastSign = thisSign
	
		// current sample value
		val = data[i]
		
		// update position and value of local max
		if (val > mx) {
			mx = val
			maxPos = i
		}
		
		// update position and value of local min
		if (val < mn) {
			mn = val
			minPos = i
		}
		if lookForMax {
		
			// passed a local max; record it and start looking for min
			if (val < (mx-delta)) {
				mn = val
				minPos = i
				lookForMax = false
				if (math.Abs(maxLocs[maxCnt]-float64(maxPos)) < 0.1) {
					continue
				}
				maxCnt = maxCnt + 1
				maxLocs[maxCnt-1] = float64(maxPos)
				maxVals[maxCnt-1] = data[maxPos]
			}
		} else {
			if (val > (mn+delta)) {
				mx = val
				maxPos = i
				lookForMax = true
				if (math.Abs(minLocs[minCnt]-float64(minPos)) < 0.1) {
					continue
				}
				minCnt = minCnt + 1
				minLocs[minCnt-1] = float64(minPos) 
				minVals[minCnt-1] = data[minPos]
			}
		}
	} // end of loop over samples of interest
	
	if (endsArePeaks) {
		// if last point is not already a local max, force it to be 
		if (math.Abs(maxLocs[maxCnt] - float64(n-1)) > 0.1) {
			maxCnt = maxCnt + 1
			maxLocs[maxCnt-1] = float64(n-1)
			maxVals[maxCnt-1] = data[n-1]
		}
		
		// if last point is not already a local min, force it to be
		if (math.Abs(minLocs[minCnt] - float64(n-1)) > 0.1) {
			minCnt = minCnt + 1
			minLocs[minCnt-1] = float64(n-1)
			minVals[minCnt-1] = data[n-1]
		}
		
	}
	
	// check for a final zero crossing
	thisSign = math.Signbit(data[n-1])
	if (thisSign != lastSign) {
		zeroCrossings = zeroCrossings + 1
	}
	
	// set lengths of output slices
	minVals = minVals[0:minCnt]
	minLocs = minLocs[0:minCnt]
	maxVals = maxVals[0:maxCnt]
	maxLocs = maxLocs[0:maxCnt]
		
	return minVals, minLocs, maxVals, maxLocs, zeroCrossings
}