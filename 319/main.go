package main

import (
	"fmt"
	"math"
)

//直接模拟, TLE
func bulbSwitch_1(n int) int {
	bulbs := make([]int, n)
	for i := 0; i < n; i++ {
		for next := i; next < n; next += (i + 1) {
			bulbs[next] = 1 - bulbs[next]
		}
	}
	res := 0
	for i := 0; i < n; i++ {
		if bulbs[i] > 0 {
			res++
		}
	}
	return res
}

//看了一些数据, 发现好像就是sqrt(n)
//提示: 数字的因子个数奇偶性, 完全平方数
func bulbSwitch(n int) int {
	return int(math.Sqrt(float64(n)))
}

func main() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d: %d\n", i, bulbSwitch(i))
	}
}

// #数学 #杂题 #数论
// #math #adhoc #number_theory
