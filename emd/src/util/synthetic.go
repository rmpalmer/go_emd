package util

import (
	"math"
)

func Synth(data []float64) {

	f := float64(len(data)) / 100.0
	c1 := 2.5 * f
	c2 := 3.0 * f

	for i:=0; i<len(data); i++ {
		data[i] = math.Sin(float64(i)/c1) + math.Cos(float64(i)/c2)
	}
}