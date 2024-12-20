package main

func numIslands(grid [][]byte) (res int) {
	n := len(grid)
	m := len(grid[0])

	var dfs func(i int, j int)
	dfs = func(i int, j int) {
		if i < 0 || i >= n || j < 0 || j >= m || grid[i][j] == '0' {
			return
		}
		grid[i][j] = '0'
		dfs(i-1, j)
		dfs(i+1, j)
		dfs(i, j-1)
		dfs(i, j+1)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if grid[i][j] == '0' {
				continue
			}
			res += 1
			dfs(i, j)
		}
	}

	return
}
