package main

import "fmt"

func totalNQueens(n int) (result int) {
	col := make([]bool, n)
	dl := make([]bool, 2*n)
	dr := make([]bool, 2*n) //offset+n

	var place func(int)

	place = func(i int) {
		if i == n {
			result += 1
			return
		}
		for j := 0; j < n; j++ {
			if col[j] || dl[i+j] || dr[i-j+n] {
				continue
			}
			col[j], dl[i+j], dr[i-j+n] = true, true, true
			place(i + 1)
			col[j], dl[i+j], dr[i-j+n] = false, false, false
		}
	}

	place(0)
	return
}

func main() {
	fmt.Println(totalNQueens(8))
}

// #nqueens
// #N皇后
