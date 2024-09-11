package main

import (
	"fmt"
	"strconv"
)

func distinctStockPairs(stocksProfit []int, target int) int {
	seen := make(map[int]bool)
	distinctPairs := make(map[string]bool)

	for _, profit := range stocksProfit {
		complement := target - profit
		if seen[complement] {
			minVal, maxVal := min(profit, complement), max(profit, complement)
			pairKey := strconv.Itoa(minVal) + "," + strconv.Itoa(maxVal)
			distinctPairs[pairKey] = true
		}
		seen[profit] = true
	}

	return len(distinctPairs)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	stocksProfit := []int{5, 7, 9, 13, 11, 6, 6, 3, 3}
	target := 12
	fmt.Println(distinctStockPairs(stocksProfit, target))  // Output: 3
}
