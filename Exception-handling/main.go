package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("/home/haveenmahantesh/hh")
	if err != nil {
		fmt.Println("error: ", err)
	}
	println()

}
