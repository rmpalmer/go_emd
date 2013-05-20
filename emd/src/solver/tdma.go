package solver

import (
)

func Tdma(od_lower, diag, od_upper, rhs, unk, cp, dp []float64) {

	// from http://en.wikipedia.org/wiki/Tridiagonal_matrix_algorithm

	n := len (diag)
	var m float64
	
	//util.DumpSlice("diag",diag)
	//util.DumpSlice("lower",od_lower)
	//util.DumpSlice("upper",od_upper)
	//util.DumpSlice("rhs",rhs)
	
	// initialize cp, dp
	cp[0] = od_upper[0] / diag[0]
	dp[0] = rhs[0] / diag[0]

	// solver for vectors cp and dp
	for i:=1; i<n; i++ {
		m = diag[i] - cp[i-1] * od_lower[i]
		cp[i] = od_upper[i] / m
		dp[i] = (rhs[i] - dp[i-1]*od_lower[i]) / m
	}
	
	// initialize solution
	unk[n-1] = dp[n-1]
	
	// back substitution to get the rest of solution
	for i:=n-2; i>=0; i-- {
		unk[i] = dp[i] - cp[i] * unk[i+1]
	}

}
