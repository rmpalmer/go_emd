package interp

import (

)

func BiCubic(x, y []float64, abcissa float64) (ordinate float64) {
	
	switch {
		case abcissa < x[1] :
			ordinate = Cubic(x[0:4], y[0:4], abcissa)
		case abcissa > x[3] :
			ordinate = Cubic(x[1:5], y[1:5],abcissa) 
		default:
			ordinate = (Cubic(x[0:4], y[0:4], abcissa) + 
					    Cubic(x[1:5], y[1:5], abcissa)) / 2.0 
	}
	
	return
}