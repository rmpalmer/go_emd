package spline

import (

)

type CoefsType struct {
	n      int       // what is the size of the problem
	coef_a []float64 // coefficients length n-1
	coef_b []float64 // coefficients length n-1
	coef_c []float64 // coefficients length n-1
	coef_d []float64 // coefficients length n-1
	b      []float64 // diagonal of matrix length n-2
	d      []float64 // rhs of equation length n-2
	h      []float64 // interval lengths length n-1
	S      []float64 // spline second derivative length n
	cp     []float64 // work array of size n-2
	dp     []float64 // work array of size n-2
}

func New () *CoefsType {
	c := CoefsType{0,nil,nil,nil,nil,nil,nil,nil,nil,nil,nil}
	return &c
} 

func (c *CoefsType) setSize(n int) {

	// n is the number of samples being fit
	// n-1 intervals
	// n-2 equations to be solved

	if (n > c.n){
		c.coef_a = make([]float64, n-1)
		c.coef_b = make([]float64, n-1)
		c.coef_c = make([]float64, n-1)
		c.coef_d = make([]float64, n-1)
		c.b      = make([]float64, n-2)
		c.d      = make([]float64, n-2)
		c.h      = make([]float64, n-1)
		c.S      = make([]float64, n)
		c.cp     = make([]float64, n-2)
		c.dp     = make([]float64, n-2)
	} else {
		c.coef_a = c.coef_a[0:n-1]
		c.coef_b = c.coef_b[0:n-1]
		c.coef_c = c.coef_c[0:n-1]
		c.coef_d = c.coef_d[0:n-1]
		c.b      = c.b[0:n-2]
		c.d      = c.d[0:n-2]
		c.h      = c.h[0:n-1]
		c.S      = c.S[0:n]
		c.cp     = c.cp[0:n-2]
		c.dp     = c.dp[0:n-2]
	}
	c.n = n
}