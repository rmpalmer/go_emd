package regression

import (
	"algebra"
	"util"
)

func TestMatMul() {

	a := algebra.NewMat(3,4)
	x := make([]float64, 4)
	b := make([]float64, 3)
	
	for i:=0; i<len(x); i++ {
		x[i] = float64(i+1)
	}
	
	v := 1.0
	
	for r,row := range a {
		for c,_ := range row {
			a[r][c] = v
			v = v + 1.0
		}
	}
	
	util.DumpSlice("x",x)
	algebra.MatInfo(a)
	algebra.MatMul(a,x,b)
	util.DumpSlice("b",b)
}