package util

import (
	"fmt"
)

func DumpSlice(name string, d []float64) {
	fmt.Printf("%s\n",name)
	for i:=0; i<len(d); i++ {
	fmt.Printf("%d [%f]\n",i,d[i])
	}
}

func DumpVal(val float64, rtn string, msg string) {
	fmt.Printf("[%s] %s : %f\n",rtn,msg,val)
}