package main

import "fmt"

func spiralOrder(matrix [][]int) (res []int) {
	n := len(matrix)
	m := len(matrix[0])
	if n == 0 || m == 0 {
		return
	}

	dir := [4][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	shift := [4][]int{{1, 0, 0, 0}, {0, -1, 0, 0}, {0, 0, -1, 0}, {0, 0, 0, 1}}
	bounds := [4]int{0, m - 1, n - 1, 0}
	row, col, i := 0, 0, 0

	outOfBounds := func(row, col int) bool {
		return row < bounds[0] || row > bounds[2] || col < bounds[3] || col > bounds[1]
	}

	res = append(res, matrix[0][0])
	for len(res) < n*m {
		idx := i % 4
		dr, dc := dir[idx][0], dir[idx][1]
		if outOfBounds(row+dr, col+dc) {
			for k := 0; k < 4; k++ {
				bounds[k] += shift[idx][k]
			}
			i += 1
			continue
		}
		row, col = row+dr, col+dc
		res = append(res, matrix[row][col])
	}
	return
}

func main() {
	m := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	fmt.Println(spiralOrder(m))
}
