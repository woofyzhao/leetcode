package main

func mySqrt(x int) int {
	a, b := 0, x
	for a <= b {
		m := (a + b) / 2
		if (m+1)*(m+1) <= x {
			a = m + 1
		} else if m*m > x {
			b = m - 1
		} else {
			return m
		}
	}
	return -1
}
