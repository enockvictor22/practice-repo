package main

import (
	"strings"
)

/*Write a function applyCase(word, mode string) string that:
func applyCase(word, mode string) string{
}
returns the word uppercased if mode is "up"
returns it lowercased if mode is "low"
returns it capitalized (first letter upper, rest lower) if mode is "cap"
returns it unchanged for anything else
Test:
applyCase("HELLO", "low")  // "hello"
applyCase("hello", "up")   // "HELLO"
applyCase("hELLO", "cap")  // "Hello"
applyCase("hello", "???")  // "hello"
*/

func applyCase(word, mode string) string {
	switch mode {
	case "up":
		strings.ToUpper(word)
	case "low":
		strings.ToLower(word)
	case "cap":

	}
	return word
}

func applyMarkers(word string) string {
	word = strings.ReplaceAll(word, ", ", ",")
	words := strings.Fields(word)
	for i := 0; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			inner := strings.Trim(words[i], "()")
			extract := strings.Split(inner, ",`")
		}

	}
	return word
}

func main() {
	//  var name string = "learn 2 earn @gmail .com"
	//  fmt.Println(strings.ReplaceAll(name, "a", "@"))
	// 	fmt.Println(len(name))
	// 	fmt.Println(name)
	// 	names := strings.Fields(name)
	// 	fmt.Println(len(names))
	// 	fmt.Println(names)
	// 	fmt.Println(strings.HasPrefix("queen", "n"))
	// 	fmt.Println(strings.HasPrefix("learn2earn", "l"))

	// 	fmt.Println(strings.HasSuffix("learn2earn", "earn"))
	// 	fmt.Println(strings.HasSuffix("RuthAndQueen", "earn"))
	// 	fmt.Println(strings.HasSuffix(name, "@gmail.com"))
	// 	fmt.Println(strings.HasPrefix(name, "learn2earn@gmail.com"))

	// 	if strings.HasPrefix(name, "learn2earn") || strings.HasSuffix(name, ".beggar") {
	// 		fmt.Println("email belongs to a fellow")
	// 	} else{
	// 		fmt.Println("email na for beggar")
	// 	}
	// 	//  {
	// 	// } else {
	// 	// 	fmt.Println("email does not belongs to a fellow")

	// 	// }

}
