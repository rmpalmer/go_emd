package interp

import (

)

func Quadratic(x, y []float64, abcissa float64) (ordinate float64) {
	
	ordinate = 
		y[0] * ( (abcissa-x[1]) * (abcissa-x[2]) ) / ( (x[0]-x[1]) * (x[0]-x[2]) )  +
		y[1] * ( (abcissa-x[0]) * (abcissa-x[2]) ) / ( (x[1]-x[0]) * (x[1]-x[2]) )  +
		y[2] * ( (abcissa-x[0]) * (abcissa-x[1]) ) / ( (x[2]-x[0]) * (x[2]-x[1]) ) 
	
	return
}