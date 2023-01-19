package main

import (
	"fmt"
	"time"
)

func print1() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
}

func print2() {
	fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
	fmt.Println("d")
	fmt.Println("e")
}

func main2() {
	go print1()
	go print2()
	time.Sleep(1 * time.Second)
}
