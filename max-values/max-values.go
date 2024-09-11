package main

import (
	"fmt"
)

func maxValue(n int, index int, maxSum int) int {
	// Helper function to calculate the total sum
	calculateSum := func(mid int) int {
		leftLen := index
		rightLen := n - index - 1
		
		// Calculate sum for the left side
		leftSum := 0
		if mid > leftLen {
			leftSum = (mid - 1 + mid - leftLen) * leftLen / 2
		} else {
			leftSum = (mid - 1 + 1) * mid / 2 + (leftLen - mid + 1)
		}

		// Calculate sum for the right side
		rightSum := 0
		if mid > rightLen {
			rightSum = (mid - 1 + mid - rightLen) * rightLen / 2
		} else {
			rightSum = (mid - 1 + 1) * mid / 2 + (rightLen - mid + 1)
		}

		return leftSum + rightSum + mid
	}

	low, high := 1, maxSum
	for low < high {
		mid := (low + high + 1) / 2
		if calculateSum(mid) <= maxSum {
			low = mid
		} else {
			high = mid - 1
		}
	}

	return low
}

func main() {
	n := 3
	index := 1
	maxSum := 6
	fmt.Println(maxValue(n, index, maxSum))  // Output: 2
}
