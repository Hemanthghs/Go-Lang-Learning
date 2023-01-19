package main

import (
	"fmt"
	"net/http"
	"sync"
)

func getStatusCode(endpoint string) {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("Error")
	} else {
		endpoints = append(endpoints, endpoint)
		fmt.Printf("%d for %s\n", res.StatusCode, endpoint)

	}

}

var endpoints = []string{"test"}

var wg sync.WaitGroup

func main() {
	websiteslist := []string{
		"https://lco.com",
		"https://go.dev",
		"https://fb.com",
		"https://google.com",
	}

	for _, web := range websiteslist {
		go getStatusCode(web)
		go getStatusCode(web)
		wg.Add(2)
	}

	wg.Wait()
	fmt.Println(endpoints)
}
