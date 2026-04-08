package main

import (
	"fmt"
	"unicode"
)

func main() {
	var first, last string

	isValid := func(s string) bool {
		for _, r := range s {
			if !unicode.IsLetter(r) {
				return false
			}
		}
		return len(s) > 0
	}

	for !isValid(first) {
		fmt.Print("Enter first name: ")
		fmt.Scanln(&first) // use Scanln instead of Scan
		if !isValid(first) {
			fmt.Println("Error: First name must contain only letters. Try again.")
		}
	}

	for !isValid(last) {
		fmt.Print("Enter last name: ")
		fmt.Scanln(&last)
		if !isValid(last) {
			fmt.Println("Error: Last name must contain only letters. Try again.")
		}
	}

	fmt.Printf("Welcome %s %s, you have made it into the 2 years program\n", first, last)
}
