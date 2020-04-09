package main

import (
	"fmt"
	"strings"
)

func wordCount(s string) map[string]int {
	count := make(map[string]int)
	words := strings.Fields(s)
	for _, val := range words {
		count[val]++
	}

	return count
}

func main() {
	fmt.Println(wordCount("I see you, you see me, we see each other."))
	fmt.Println(wordCount("one two three four five"))
	fmt.Println(wordCount("duplicate duplicate dublicate"))
}
