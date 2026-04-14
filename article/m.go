package main

import (
	"fmt"
	"strings"
)

func fixArticle(sentence string) string {
	words := strings.Fields(sentence)
	for i := 0; i < len(words); i++ {
		if strings.ToLower(words[i]) == "a" {
			first := strings.ToLower(string(words[i+1][0]))
			if strings.ContainsAny(first, "aeiouh") {
				if words[i] == "A" {
					words[i] = "An"
				} else {
					if !strings.ContainsAny(first, "aeiouh") {
						if words[i] == "A" {
							words[i] = "An"
						}
					}
				}
			}

		}
	}
	return strings.Join(words, " ")
}


// func fixArticle(sentence string) string {
// 	words := strings.Fields(sentence)
// 	for i := 0; i < len(words)-1; i++ {
// 		nextWord := words[i+1][:1]
// 		vowels := "aeiouhAEIOUH"
// 	switch {
// 	case words[i] == "a" && strings.ContainsAny(nextWord, vowels):
// 		words[i] = "an"
// 	case words[i] == "A" && strings.ContainsAny(nextWord, vowels):
// 		words[i] = "An"
// 	case words[i] == "an" && !strings.ContainsAny(nextWord, vowels):
// 		words[i] = "a"
// 	case words[i] == "An" && !strings.ContainsAny(nextWord, vowels):
// 		words[i] = "A"
// 	}
// }
// 	return strings.Join(words, " ")
// }
// }

func main() {
	sentence := "logic overflow a amazing group, an coding mentor said"
	fmt.Println(fixArticle(sentence))

}
