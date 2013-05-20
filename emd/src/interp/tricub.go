package interp

import (

)

func TriCubic(x, y []float64, abcissa float64) (ordinate float64) {

	switch {
		case abcissa < x[1] :
			ordinate = Cubic(x[0:4], y[0:4], abcissa)
		case abcissa < x[2] :
			ordinate = (Cubic(x[0:4], y[0:4],abcissa) + 
						Cubic(x[1:5], y[1:5],abcissa)) / 2.0
		case abcissa > x[4] :
			ordinate = Cubic(x[2:6], y[2:6], abcissa)
		case abcissa > x[3] :
			ordinate = (Cubic(x[1:5], y[1:5],abcissa) + 
						Cubic(x[2:6], y[2:6],abcissa)) / 2.0
		default:
			ordinate = (Cubic(x[0:4], y[0:4], abcissa) + 
						Cubic(x[0:4], y[0:4], abcissa) + 
						Cubic(x[2:6], y[2:6], abcissa)) / 3.0
	}
	
	return
}