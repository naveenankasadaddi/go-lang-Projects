package main

import "fmt"

type Student struct {
	srn   int
	name  string
	marks float64
}

func main() {
	var s1 = Student{1, "Naveen", 90}
	fmt.Println(s1.name)
	fmt.Println(s1.marks)
	fmt.Println(s1.srn)

}
