package main

import "testing"

type addTest struct {
	arg1, arg2, expected int
}

var addTests = []addTest{
	addTest{10, 20, 30},
	addTest{30, 40, 70},
	addTest{10, 10, 20},
	addTest{5, 5, 10},
	addTest{100, 300, 40},
}

func TestAdd(t *testing.T) {
	for _, test := range addTests {
		if output := Add(test.arg1, test.arg2); output != test.expected {
			t.Errorf("got %d want %d", output, test.expected)
		}
	}

}
