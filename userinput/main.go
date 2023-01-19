package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the input: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("Input is : ", input)
}
