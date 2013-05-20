package regression

import (
	"algorithm"
	"util"
	"fmt"
)

func TestDecompose() {
	data := make([]float64, 100)
	util.Synth(data)
	cnt := algorithm.Decompose(data, 10)
	fmt.Printf("found %d imfs\n",cnt)
}