package main

func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	a, b, c := 1, 2, 0
	for k := 3; k <= n; k++ {
		c = a + b
		a, b = b, c
	}
	return c
}

func main() {

}
