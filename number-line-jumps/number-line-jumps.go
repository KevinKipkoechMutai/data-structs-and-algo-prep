package main

func kangaroo(x1, v1, x2, v2 int) string {
	if (x1 < x2 && v1 <= v2) || (x2 < x1 && v2 <= v1) {
		return "NO"
	}
	// Check if the kangaroos meet after the same number of jumps
	if (x2-x1)%(v1-v2) == 0 {
		return "YES"
	}
	return "NO"
}