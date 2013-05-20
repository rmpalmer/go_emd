package main 

import (
	"fmt"
	"regression"
)

func main() {
	fmt.Printf("emd version 0.1\n")
	
		
	//regression.TestPeaks()
	//regression.TestInterp()
	//regression.TestSolver(8)
	//regression.TestMatMul()
	//regression.TestSpline(4)
	//regression.TestResamp()
	
	regression.TestDecompose()
	//regression.TestSubtract()
	
	fmt.Printf("emd done\n")
}

