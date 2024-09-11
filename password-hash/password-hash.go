package main

import (
	"fmt"
	"strconv"
)


const (
	p = 131
	M = 1000000007
)

func authEvents(events [][]string) []int {
	passwordHash := 0

	hashString := func (s string) int {
		hash := 0
		for _, c := range s {
			hash = (hash*p + int(c)) % M
		}
		return hash
	}

	authorize := func (x, currentHash int) int {
		if x == currentHash {
			return 1
		}
		for i := 0; i < 128; i++ {
			newHash := (currentHash*p + i) % M
			if newHash == x {
				return 1
			}
		}
		return 0
	}
	var result []int
	for _, event := range events {
		if event[0] == "setPassword" {
			passwordHash = hashString(event[1])
		} else if event[0] == "authorize" {
			x, _ := strconv.Atoi(event[1])
			result = append(result, authorize(x, passwordHash))
		}
	}
	return result
}

func main()  {
	events := [][]string{
		{"setPassword", "cAr1"},
		{"authorize", "223691457"},
		{"authorize", "303580761"},
	}
	fmt.Println(authEvents(events))
}