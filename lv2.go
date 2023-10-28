package main

import (
	"fmt"
	"math"
)

func s(r float64) float64 {
	return math.Pi * math.Pow(r, 2)
}

func main() {
	var r float64
	fmt.Scanln(&r)
	area := s(r)
	fmt.Println(area)
}
