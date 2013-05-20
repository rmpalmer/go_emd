package resample

import (

)

func Resamp(x, y, result []float64) ([]float64) {
	
	// find out how big the output slice should be and make it so
	n := cap(result) 
	result = result[0:n]
	
	// how many points in input
	samcnt := len(x)
	
	// how many input samples determines which means of resampling
	switch samcnt {
		case 1 :
			for i:=0; i<samcnt; i++ {
				result[i] = y[0]
			}
		case 2 :
			result = ResampLinear(x, y, result)
		case 3 :
			result = ResampQuadratic(x, y, result)
		case 4:
			result = ResampCubic(x, y, result)
		case 5: 
			result = ResampBiCubic(x, y, result)
		case 6:
			result = ResampTriCubic(x, y, result)
		default :
			result = ResampSpline(x, y, result)
	}
	
	return result
}