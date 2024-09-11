package main

import (
	"fmt"
	"os"
	"log"
	"strings"
)

func main()  {
	data, err := os.ReadFile("sample.txt")
	if err != nil {
		log.Fatal(err)
	}

	modifiedData := strings.ReplaceAll(string(data), "\n", " ")
	fmt.Println(modifiedData)
}