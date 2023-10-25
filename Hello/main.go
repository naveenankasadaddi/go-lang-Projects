package main

import "fmt"

func main() {
	var num int32
	fmt.Println("Enter the number:")
	fmt.Scanf("%d", &num)
	EveOdd(int(num))
	fmt.Println(factorial(int(num)))
	fmt.Println(fibonacci(int(num)))

}

func factorial(num int) int {
	if num <= 0 {
		return 1
	} else {
		return num * factorial(num-1)
	}

}

func EveOdd(num int) {
	if num%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("Odd")
	}

}

func fibonacci(num int) int32 {
	var num1 int32 = 0
	var num2 int32 = 1
	var sum int32
	//var i int =0

	if num == 0 {
		return num1
	}
	for i := 0; i < num; i++ {
		sum = num1 + num2
		num1 = num2
		num2 = sum

	}
	return num2

}
