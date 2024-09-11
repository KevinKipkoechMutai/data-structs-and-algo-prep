package main

import "fmt"

func main()  {
	asciiDict := map[string]int{"A": 65, "B": 66, "C": 67, "D": 68}
	reversedDict := make(map[int]string)

	for key, value := range asciiDict {
		reversedDict[value] = key
	}

	fmt.Println(reversedDict)
}