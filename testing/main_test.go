package main

import (
	"fmt"
	"testing"
)

func TestReturnHello(t *testing.T) {
	actualString := ReturnHello()
	expectedString := "Hello"
	if actualString != expectedString {
		t.Errorf("Expected String(%s) is not same as"+"actual String(%s)", expectedString, actualString)
	}
	fmt.Println("TestReturnHello passed")
}

func TestAdd(t *testing.T) {
	actualResult := Add(10, 25)
	expectedResult := 35
	if actualResult != expectedResult {
		t.Errorf("Expected Int(%d) is not same as "+"actual Int(%d)", expectedResult, actualResult)
	}
	fmt.Println("TestAdd passed")
}
