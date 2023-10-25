package main

import (
	"fmt"
)

func main() {
	createArr()
}

func createArr() {
	arr := [5]int{1, 2, 3, 4, 5}
	for i, value := range arr {
		fmt.Printf("Element at index %d is %d\n", i, value)
	}
}
