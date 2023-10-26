package main

import "fmt"

func main() {
	i := 100
	p := &i
	fmt.Println(&i)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)

}
