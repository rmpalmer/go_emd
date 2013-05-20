package regression

import (
	"util"
)

func TestSubtract() {
	a := []float64 {3, 5, 7, 9}
	b := make([]float64, len(a))
	copy(b, a)
	util.DumpSlice("a",a)
	util.DumpSlice("b",b)
	//c:= make([]float64, len(a)
	util.SliceSub(a,b,a)
	util.DumpSlice("result",a)
	util.DumpSlice("b",b)
}
