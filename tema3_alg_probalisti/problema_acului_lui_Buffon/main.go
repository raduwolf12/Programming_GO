package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	var L float64
	fmt.Println("Lungimea dintre linii:")
	fmt.Scan(&L)



	var n int
	fmt.Println("Numarul de aruncari:")
	fmt.Scan(&n)

	min := 0.0
	max := L


	unghiMin := 0.0
	unghiMax := math.Pi / 2.0

	m := 0
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		y := (rand.Float64() * (max - min)) + min
		alpha := (rand.Float64() * (unghiMax - unghiMin)) + unghiMin
		y1 := y + L*math.Sin(alpha)

		if int(y/L) != int(y1/L) {
			m++
		}
		fmt.Println("y=", y, "  alpha=", alpha, "  y1=", y1)

	}

	fmt.Println("Numarul de ace care au cazut pe linie:", m)

	pi := 2.0 * float64(n) / float64(m)
	fmt.Println("Pi=", pi)

}