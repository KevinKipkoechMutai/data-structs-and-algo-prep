package main

import (
	"fmt"
	"strings"
)

func main()  {
	rows := 5
	x := 0

	for i := rows; i > 0; i-- {
		x++
		fmt.Println(strings.Repeat(fmt.Sprintf("%d ", x), i))
	}
}