package main

import (
	"fmt"
	"runtime"
)

type Rect interface {
	PrintArea() int
	PrintPermimeter() int
}

type Dim struct {
	L int
	B int
}

func (d Dim) PrintArea() int {
	return d.L * d.B
}

func (d Dim) PrintPermimeter() int {
	return 2 * (d.L + d.B)
}

func PrintVal(val interface{}) {
	switch val.(type) {
	case int:
		fmt.Println("int passed : ", val)
	case string:
		fmt.Println("string passed : ", val)
	case float64:
		fmt.Println("float64 passed : ", val)
	default:
		fmt.Println("Incorrect input")
	}
}

func main() {
	var R Rect
	R = Dim{10, 20}
	fmt.Println(R.PrintArea())
	fmt.Println(R.PrintPermimeter())

	PrintVal("Hemanth Sai")
	PrintVal(10)
	PrintVal(1.2)

	fmt.Println("Go Version: ", runtime.Version())
}
