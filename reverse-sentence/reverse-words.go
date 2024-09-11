package main

import (
	"fmt"
	"strings"
)

func reverseWords(sentence string) string {
	words := strings.Split(sentence, " ")
	var newWordList []string

	for _, word := range words {
		reversedWord := reverseString(word)
		newWordList = append(newWordList, reversedWord)
	}
	return strings.Join(newWordList, " ")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main()  {
	fmt.Println(reverseWords("My Name is Kevin"))
}