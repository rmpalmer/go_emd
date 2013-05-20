package resample

import (
	"interp"
	"fmt"
)

func ResampBiCubic(x, y, result []float64) ([]float64) {

	outCount := len(result)
	
	fmt.Printf("ResampBiCubic outCount %d\n",outCount)
	
	for i:=0; i<outCount; i++ {
		result[i] = interp.BiCubic(x, y, float64(i))
	}

	return result
}


