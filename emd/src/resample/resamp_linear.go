package resample

import (
	"interp"
)

func ResampLinear(x, y, result []float64) ([]float64) {


	outCount := len(result)

	for i:=0; i<outCount; i++ {
		result[i] = interp.Linear(x, y, float64(i))
	}

	return result
}
