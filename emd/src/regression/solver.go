package regression

import (
	"fmt"
	"algebra"
	"math/rand"
	"util"
	"solver"
	"time"
)

func TestSolver(n int) {
	fmt.Printf("Testing TDMA Solver\n")
	
	now := time.Now()
	seed := now.Unix()
	rand.Seed(seed)
	
	A := algebra.NewMat(n, n)
	x := make([]float64, n)
	b := make([]float64, n)
	
	for i:=0; i<n; i++ {
		x[i] = rand.Float64()
		A[i][i] = rand.Float64()
		if (i<(n-1)) {
			A[i+1][i] = rand.Float64()
			A[i][i+1] = rand.Float64()
		}
	}
	
	algebra.MatMul(A,x,b)
	algebra.MatInfo(A)
	util.DumpSlice("x",x)
	util.DumpSlice("b",b)
	
	diag := make([]float64, n)
	unk  := make([]float64, n)
	od_l := make([]float64, n)
	od_u := make([]float64, n)
	
	for i:=0; i<n; i++ {
		diag[i] = A[i][i]
		if (i>0) {
			od_l[i] = A[i][i-1]
		}
		if (i<(n-1)) {
			od_u[i] = A[i][i+1]
		}
	}
	
	cp := make([]float64, n)
	dp := make([]float64, n)
	
	//util.DumpSlice("diag",diag)
	//util.DumpSlice("od_lower",od_l)
	//util.DumpSlice("od_upper",od_u)
	
	solver.Tdma(od_l, diag, od_u, b, unk, cp, dp)
	util.DumpSlice("result",unk)
	
}