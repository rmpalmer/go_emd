package resample

import (
	"spline"
	"fmt"
)

func ResampSpline(x, y, result []float64) ([]float64) {

	// is there a way to keep this as global, re-usable data?
	var coefs *spline.CoefsType = nil

	coefs = spline.Points(x, y, coefs)

	outCount := len(result)
	
	fmt.Printf("ResampSpline outCount %d\n",outCount)

	// current index
	curind := -1
	
	for i:=0; i<outCount; i++ {
		curind, result[i] = spline.Evaluate(float64(i),x,curind,coefs)
	}

	return result
}
