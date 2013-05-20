package regression

import (
	"spline"
)

func TestSpline(n int) {

	var coefs *spline.CoefsType
	coefs = nil

	x := make([]float64, n)
	y := make([]float64, n)
	
	cind := -1
	
	x[0] = 0.0
	x[1] = 1.0
	x[2] = 1.5
	x[3] = 2.25
	
	y[0] = 2.0
	y[1] = 4.4366
	y[2] = 6.7134
	y[3] = 13.9130
	
	coefs = spline.Points(x, y, coefs)
	
	cind, _ = spline.Evaluate(0.66, x, cind, coefs)
	cind, _ = spline.Evaluate(1.75, x, cind, coefs)
	
}