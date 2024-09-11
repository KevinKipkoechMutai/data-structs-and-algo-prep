package main

import "fmt"

func main()  {
	numberList := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	i := 0
	n := len(numberList)

	for i < n {
		if numberList[i] > 50 {
			numberList = append(numberList[:i], numberList[i+1:]...)
			n--
		} else {
			i++
		}
	}
	fmt.Println(numberList)
}