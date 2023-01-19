package main

import (
	"encoding/json"
	"fmt"
)

type Emp struct {
	Name string
	Id   int32
	Age  int32
}

func main() {
	emp1 := Emp{"Hemanth", 20, 21}
	emp1_enc, _ := json.Marshal(emp1)
	fmt.Println(string(emp1_enc))

	emp2 := []Emp{
		{"sai", 1, 2},
		{"hemanthsai", 2, 3},
		{"ghs", 3, 4},
	}

	emp2_enc, _ := json.Marshal(emp2)
	fmt.Println(string(emp2_enc))

	var emp3 Emp
	Data := []byte(`
	
	{
		"Name":"saihemanth",
		"Id":12,
		"Age":30
	}
	`)

	data2 := json.Unmarshal(Data, &emp3)
	fmt.Println(data2)
}
