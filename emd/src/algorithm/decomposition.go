package algorithm

import (
	"util"
	"fmt"
	"algebra"
)

func Decompose(f []float64, maxImf int) (int) {

	h := make([]float64, len(f))
	residual := make([]float64, len(f))
	copy (h, f)
	copy (residual, f)
	const eps = 0.05
	imfCount := 0
	loopCount := 0
	rmin := 0.0
	rmax := 0.0
	monotonic := false
	
	n := len(f)
	
	imfs := algebra.NewMat(n, maxImf)
	
	for {
		fmt.Printf("loop %d\n",loopCount)
		loopCount = loopCount + 1
		util.DumpSlice("current h",h)
		fmt.Printf("Sifting...\n")
		minMean, zeroCrossings, minCount, maxCount := Sift(h)
		extrema := minCount + maxCount
		fmt.Printf("found minMean of %f\n",minMean)
		fmt.Printf("min mean %f zc %d ex %d %d\n",minMean, zeroCrossings, minCount, maxCount)
		pointMatch := ((zeroCrossings-extrema < 2) || (extrema-zeroCrossings < 2)) 
		if ((minMean < eps) && pointMatch) {
			fmt.Printf("storing imf at index %d\n",imfCount)
			for i:=0; i<n; i++ {
				imfs[i][imfCount] = h[i]
			}
			imfCount = imfCount + 1
			fmt.Printf("found imf number %d\n",imfCount)
			util.SliceSub(residual, h, residual)
			copy(h, residual)
			if (imfCount >= maxImf) {
				break
			}
			rmin, rmax, monotonic = util.SliceStats(residual)
			if (monotonic) {
				fmt.Printf("terminating with monotonic residual\n")
			}
			if ((rmax-rmin) < eps) {
				fmt.Printf("terminating with small residual %f\n",rmax-rmin)
			}
		}
		
    }
    
    algebra.MatInfo(imfs)
    
    return imfCount
}