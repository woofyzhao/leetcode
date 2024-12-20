package main

func maxAreaOfIsland(grid [][]int) (res int) {
	n := len(grid)
	m := len(grid[0])
	for i, row := range grid {
		for j, c := range row {
			if c == 0 {
				continue
			}
			var area int
			var travel func(int, int)
			travel = func(x, y int) {
				if x < 0 || x >= n || y < 0 || y >= m || grid[x][y] == 0 {
					return
				}
				area += 1
				grid[x][y] = 0
				travel(x-1, y)
				travel(x+1, y)
				travel(x, y-1)
				travel(x, y+1)
			}
			travel(i, j)
			res = max(res, area)
		}
	}
	return
}

func main() {

}
