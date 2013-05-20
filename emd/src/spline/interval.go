package spline

import (
)

func Interval(abcissa float64, x []float64) int {

	n := len (x)

	lower := 0
	upper := n-1
	tmp   := 0
	
	for {
		if ((upper-lower) <= 1) {
			break
		}
		tmp = lower + ((upper-lower)/2)
		if (x[tmp] > abcissa) {
			upper = tmp
		} else {
			lower = tmp
		}
	}
	p := lower
	if (x[upper] < abcissa) {
		p = upper
	} 
	return p
}