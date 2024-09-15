package main

func countApplesAndOranges(s, t, a, b int, apples, oranges []int) {
	acount := 0
	bcount := 0

	for _, apple := range apples {
		temp := a + apple
		if temp >= s && temp <= t {
			acount++
		}
	}

	for _, orange := range oranges {
		temp := b + orange
		if temp >= s && temp <= t {
			bcount++
		}
	}
}