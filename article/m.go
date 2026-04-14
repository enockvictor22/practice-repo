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

func main() {
	sentence := "logic overflow a amazing group, an coding mentor said"
	fmt.Println(fixArticle(sentence))

}
