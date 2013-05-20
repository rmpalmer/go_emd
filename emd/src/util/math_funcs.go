package util

import (
	"math"
	"fmt"
)

func SliceStats (data[]float64) (low, high float64, monotonic bool) {
	n := len(data)
	low = data[0]
	high = data[0]
	lastVal := data[0]
	increasing := false
	decreasing := false
	for i:=1; i<n; i++ {
		low = math.Min(low,data[i])
		high = math.Max(high, data[i])
		if (data[i] > lastVal) {
			increasing = true
		} else if (data[i] < lastVal) {
			decreasing = true
		}
		lastVal = data[i]
	}
	if (increasing && decreasing) {
		monotonic = false
	} else {
		monotonic = true
	}
	return low, high, monotonic
}

func Energy(data []float64) (float64) {
	n := len(data)
	sum := 0.0
	for i:=0; i<n; i++ {
		sum = sum + (data[i] * data[i])
	}
	return sum
}

func SliceSub(a, b, c []float64) {
	n := len(a)
	fmt.Printf("subtracting two slices of length %d\n", n)
	for i:=0; i<n; i++ {
		c[i] = a[i] - b[i]
	}
	return
}

func SliceMin(data []float64) (val float64) {
	val = data[0]
	for i:=1; i<len(data); i++ {
		val = math.Min(val, data[i]) 
	}
	return val
}

func SliceMax(data []float64) (val float64) {
	val = data[0]
	for i:=1; i<len(data); i++ {
		val = math.Max(val, data[i]) 
	}
	return val
}