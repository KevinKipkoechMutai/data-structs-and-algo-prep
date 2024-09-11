package main

import (
	"fmt"
	"strings"
)

func main()  {
	fmt.Println(reverseWords("My name is Kevin"))
}

func reverseWords(sentence: string) string {
	split_sentence = strings.Split(sentence, " ")
}