package algebra

import (
	"fmt"
)

func NewMat (r, c int) [][]float64 {
	mat := make([][]float64, r)
	for i:= range mat {
		mat[i] = make([]float64, c)
	}
	return mat
}

func MatMul (A [][]float64, x, b []float64) {
	for i := range A {
		b[i] = 0.0
		for j:= range A[i] {
		  	b[i] = b[i] + A[i][j] * x[j] 
		}
	}
	return
}

func MatInfo (A [][]float64) {
	fmt.Printf("printing an array\n")
	for _,r := range A {
		for _,c := range r {
			fmt.Printf("%f\t",c)
		}
	fmt.Printf("\n")
	}
}