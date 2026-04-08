package main

import (
	"fmt"
)

var (
	a = 10
	b = 27
	c = 25
)

func main() {

	LargestNum()
}

func LargestNum() int {
	var a, b, c int
	fmt.Print("input first number : ")
	fmt.Scanln(&a)

	fmt.Print("input second number: ")
	fmt.Scanln(&b)

	fmt.Print("input third number: ")
	fmt.Scanln(&c)

	if a >= b && a >= c {
		fmt.Println(a)
		return a
	} else if b >= a && b >= c {
		fmt.Println(b)
		return b
	} else {
		fmt.Println(c)
		return c
	}
}

func greatestNum() {
	if a >= b && a >= c {
		fmt.Println(a)
	} else if b >= a && b >= c {
		fmt.Println(b)
	} else {
		fmt.Println(c)
	}
}
