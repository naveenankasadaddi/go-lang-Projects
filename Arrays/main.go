package main

import (
	"fmt"
)

func main() {
	createArr()
}

func createArr() {
	arr := [5]int{1, 2, 3, 4, 5}
	//var i int=0
	fmt.Println("Array elements :")
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
		fmt.Printf("\t")

	}
}
