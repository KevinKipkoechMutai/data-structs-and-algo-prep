package main

import "fmt"

func displayDuplicates(nums []int) int {
	duplicates := []int{} // store duplicates
	seen := map[int]bool{} // track seen numbers

	for _, num := range nums {
		if seen[num] {
			if !contains(duplicates, num) {
				duplicates = append(duplicates, num)
			}
		} else {
			seen[num] = true
		}
	}
	return len(duplicates)
}

func contains(slice []int, num int) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

func main()  {
	nums := []int{10, 20, 60, 30, 20, 40, 30, 60, 70, 80}
	fmt.Println(displayDuplicates(nums))
}