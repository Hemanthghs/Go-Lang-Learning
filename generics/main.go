package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Add(a, b float64) float64 {
	return a + b
}

func GAdd[T constraints.Ordered](a, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(10.23, 4.5))
	fmt.Println(GAdd[int32](10, 20))
	FAdd := GAdd[float64]
	fmt.Println(FAdd(234.43, 24.3))
}
