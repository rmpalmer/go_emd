package interp

import (

)

func Linear(x, y []float64, abcissa float64) (ordinate float64) {
	
	ordinate = y[0] + ((abcissa-x[0]) * ((y[1]-y[0])/(x[1]-x[0])))
	
	return
}