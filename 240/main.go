package main

func searchMatrix0(matrix [][]int, target int) bool {
	if len(matrix) == 0 {
		return false
	}
	var search func([][]int, int, int) bool
	search = func(matrix [][]int, n int, m int) bool {
		i := 0
		for i < n && target > matrix[i][m-1] {
			i++
		}
		if i >= n {
			return false
		}
		j := m - 1
		for j >= 0 && target < matrix[i][j] {
			j--
		}
		if j < 0 {
			return false
		}
		return matrix[i][j] == target || search(matrix[i:], n-i, j+1)
	}
	return search(matrix, len(matrix), len(matrix[0]))
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	i := m - 1
	j := 0
	for i >= 0 && j < n {
		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] > target {
			i--
		} else {
			j++
		}
	}
	return false
}

func main() {

}
