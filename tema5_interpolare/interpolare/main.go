package r

import (
	"fmt"
	"os"
)

func main(){

	var x[] float64
	var y[] float64

	var n int
	file,_ := os.Open("fisier.txt")
	fmt.Fscan(file, &n)

	for i:=0; i<n; i++{
		fmt.Fscan(file,&x[i],&y[i])
	}

	var s = NewSpline(x,y, CubicFirstDeriv,0,0)
	fmt.Println(s.At(7))

}

