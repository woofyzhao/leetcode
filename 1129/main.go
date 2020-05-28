package main

const (
	MaxN = 200
	Nan  = -1
	Red  = 0
	Blue = 1
)

type Vetex struct {
	id        int
	fromColor int
	len       int
}

//vetex_i -> red|blue -> neighbours
func bfs(graph [][][]int) (result []int) {
	n := len(graph)
	result = make([]int, n)
	for i := range result {
		result[i] = Nan
	}

	result[0] = 0
	seen := make(map[int]map[int]bool)
	seen[0] = make(map[int]bool)
	seen[0][Red], seen[0][Blue] = true, true
	queue := make([]Vetex, 2*MaxN)
	queue = append(queue, Vetex{id: 0, fromColor: Nan, len: 0})

	for i := 0; i < len(queue); i++ {
		v := queue[i]
		for color := Red; color <= Blue; color++ {
			if v.fromColor == color {
				continue
			}
			for _, next := range graph[v.id][color] {
				if seen[next][color] {
					continue
				}
				if seen[next] == nil {
					seen[next] = make(map[int]bool)
				}
				seen[next][color] = true
				newLen := v.len + 1
				queue = append(queue, Vetex{id: next, fromColor: color, len: newLen})
				if result[next] == Nan || newLen < result[next] {
					result[next] = newLen
				}
			}
		}
	}
	return
}

func shortestAlternatingPaths(n int, red_edges [][]int, blue_edges [][]int) []int {
	graph := make([][][]int, n)
	for i := range graph {
		graph[i] = make([][]int, 2)
	}
	for _, e := range red_edges {
		graph[e[0]][Red] = append(graph[e[0]][Red], e[1])
	}
	for _, e := range blue_edges {
		graph[e[0]][Blue] = append(graph[e[0]][Blue], e[1])
	}
	return bfs(graph)
}

func main() {

} // [[0,1],[1,2],[0,4],[4,4],[4,3],[5,3],[5,5],[2,5]]
// [[0,4],[1,1],[4,2],[3,2],[5,3]]

// #shortest_path #graph #bfs #though_for_a_while
// #最短路径 #图 #广度优先搜索
