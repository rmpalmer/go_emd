package resample

import (
	"interp"
	"fmt"
)

func ResampCubic(x, y, result []float64) ([]float64) {

	outCount := len(result)
	
	fmt.Printf("ResampCubic outCount %d\n",outCount)
	
	for i:=0; i<outCount; i++ {
		result[i] = interp.Cubic(x, y, float64(i))
	}

	return result
}
