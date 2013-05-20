package spline

import (
)

func Evaluate (abcissa float64, x []float64, cind int, c *CoefsType) (ind int, result float64){

	ind = cind
	
	// if necessary, find index of current interval
	if ((ind < 0) || 
		(ind >= len(x)) ||
		(abcissa <= x[ind]) ||
		(abcissa > x[ind+1])) {
		ind = Interval(abcissa, x)
	}
	
	
	z := abcissa - x[ind]

	//fmt.Printf("evaluating spline at %f on interval %d\n",abcissa, ind)
	
	result = c.coef_a[ind] * z*z*z +
			 c.coef_b[ind] * z*z   +
			 c.coef_c[ind] * z     +
			 c.coef_d[ind]
	
	//fmt.Printf("evaluating spline at %f on interval %d result %f\n",abcissa, ind, result)
	
	return ind, result
}