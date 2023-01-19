package main

import (
	"encoding/json"
	"fmt"
)

type Emp struct {
	Id   int
	Name string
	Age  int
}

func main() {
	emp1 := Emp{1, "Hemanth", 21}
	fmt.Println(emp1)
	emp1_enc, _ := json.Marshal(emp1)
	fmt.Println(string(emp1_enc))

	var emp2 Emp
	Data := []byte(`
		{
			"Id":20,
			"Name":"Sai",
			"Age":20
		}
	`)
	json.Unmarshal(Data, &emp2)
	fmt.Println(emp2)
	// a := strings.HasSuffix("hemanthsai2", "sai")
	fmt.Println(a)
}
