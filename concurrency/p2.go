package main

import (
	"fmt"
	"sync"
)

func print11() {
	fmt.Println(1)
	fmt.Println(2)
	fmt.Println(3)
	fmt.Println(4)
	fmt.Println(5)
	wg.Done()
}

func print22() {
	fmt.Println("a")
	fmt.Println("b")
	fmt.Println("c")
	fmt.Println("d")
	fmt.Println("e")
	wg.Done()
}

var wg sync.WaitGroup

func main1() {
	wg.Add(2)
	go print1()
	go print2()
	wg.Wait()
}
