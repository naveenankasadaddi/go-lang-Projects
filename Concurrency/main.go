package main

import (
	"fmt"
	"time"
)

func main() {
	go greet("Naveen")
	greet("Sonu")
}

func greet(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Millisecond)
		fmt.Println("Hello ", s)
	}

}
