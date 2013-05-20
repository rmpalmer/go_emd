package spline

import (
	"solver"
	"fmt"
)

func Points (x, y []float64, c *CoefsType) *CoefsType {

	// make sure size is right
	if (c == nil) {
		c = New()
	} 
	c.setSize(len(x))
	
	// set up to compute coefficients
	
	// h has length n-1
	for i:=0; i<(c.n-1); i++ {
		c.h[i] = x[i+1] - x[i]
	}
	
	// diagonal of size n-2
	for i:=0; i<(c.n-2); i++ {
		c.b[i] = 2.0 * (c.h[i] + c.h[i+1]) 
	}
	
	// rhs of size n-2
	for i:=1; i<(c.n-1); i++ {
		c.d[i-1] = 6.0 * ( ((y[i+1]-y[i])/c.h[i]) - ((y[i]-y[i-1])/c.h[i-1]) )
	}
	
	// set second derivatives for natural spline
	c.S[0] = 0.0
	c.S[c.n-1] = 0.0
	
	fmt.Printf("calling solver\n")
	
	//util.DumpSlice("h",c.h)
	
	// call tridiagonal solver
	solver.Tdma(c.h[:], // lower diagonal
		        c.b[:],  // diagonal
		        c.h[1:], // upper diagonal
		        c.d[:],  // rhs
		        c.S[1:], // unkknown
		        c.cp,    // work_a
		        c.dp)    // work_b
		        
	// save the coefficients
	for i:=0; i<(c.n-1); i++ {
		c.coef_a[i] = (c.S[i+1]-c.S[i])/(6.0 * c.h[i])
		c.coef_b[i] = c.S[i] / 2.0
		c.coef_c[i] = ( (y[i+1]-y[i])/c.h[i]) - ( (2.0*c.h[i]*c.S[i] + c.h[i]*c.S[i+1]) / 6.0)
		c.coef_d[i] = y[i]
	}
	
	//util.DumpSlice("coef_a",c.coef_a)
	//util.DumpSlice("coef_b",c.coef_b)
	//util.DumpSlice("coef_c",c.coef_c)
	//util.DumpSlice("coef_d",c.coef_d)
	
	return c
}
