package main

import "fmt"

func main() {

	// map literal example
	m := map[string]int{
		"key1": -1,
		"key2": 123,
	}

	arr := []string{"one", "two", "three", "four"}

	mm := make(map[int]string)

	for i, v := range arr {
		mm[i] = v
	}

	fmt.Println(m, mm)

}
