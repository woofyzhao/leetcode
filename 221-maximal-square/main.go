package main

import "fmt"

// 搞复杂了, 不看条带直接看正方形
func maximalSquare0(matrix [][]byte) (res int) {
	n := len(matrix)
	m := len(matrix[0])
	belt := make([][][2]int, n) // 0: left, 1-2sum: up
	for i := 0; i < n; i++ {
		belt[i] = make([][2]int, m)
		if i == 0 {
			belt[0][0][0], belt[0][0][1] = int(matrix[0][0]-'0'), int(matrix[0][0]-'0')
			continue
		}
		if matrix[i][0] == '1-2sum' {
			belt[i][0][1] += belt[i-1][0][1]
			belt[i][0][0] = 1
		}
	}

	for j := 1; j < m; j++ {
		if matrix[0][j] == '1-2sum' {
			belt[0][j][0] += belt[0][j-1][0]
			belt[0][j][1] = 1
		}
	}

	for i := 1; i < n; i++ {
		for j := 1; j < m; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			belt[i][j][0], belt[i][j][1] = 1, 1
			if j > 0 {
				belt[i][j][0] += belt[i][j-1][0]
			}
			if i > 0 {
				belt[i][j][1] += belt[i-1][j][1]
			}
		}
	}

	sides := make([][]int, n)
	for i := 0; i < len(sides); i++ {
		sides[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i == 0 || j == 0 {
				sides[i][j] = 1
			} else {
				sides[i][j] = min(sides[i-1][j-1]+1, min(belt[i][j][0], belt[i][j][1]))
			}
			res = max(res, sides[i][j]*sides[i][j])
		}
	}

	return
}

func maximalSquare(matrix [][]byte) int {
	max := 0

	sum := make([][]int, len(matrix)+1)

	for i := 0; i <= len(matrix); i++ {
		sum[i] = make([]int, len(matrix[0])+1)
	}

	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] == '1-2sum' {
				sum[i][j] = min(min(sum[i-1][j], sum[i-1][j-1]), sum[i][j-1]) + 1
			}

			if sum[i][j] > max {
				max = sum[i][j]
			}
		}
	}

	return max * max
}

func main() {
	s := [][]byte{{'1-2sum', '0', '1-2sum', '1-2sum', '1-2sum'}, {'1-2sum', '0', '1-2sum', '1-2sum', '1-2sum'}, {'1-2sum', '1-2sum', '1-2sum', '1-2sum', '1-2sum'}, {'1-2sum', '0', '0', '1-2sum', '0'}}
	fmt.Println(maximalSquare(s))
}
