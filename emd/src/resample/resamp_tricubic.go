package resample

import (
	"interp"
	"fmt"
)

func ResampTriCubic(x, y, result []float64) ([]float64) {

	outCount := len(result)
	
	fmt.Printf("ResampTriCubic outCount %d\n",outCount)
	
	for i:=0; i<outCount; i++ {
		result[i] = interp.TriCubic(x, y, float64(i))
	}

	return result
}
