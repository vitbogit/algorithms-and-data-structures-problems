package main

import "fmt"

func main() {
	// Parse graph size and edges count
	var N, M int
	fmt.Scanf("%d %d\n", &N, &M)

	// Create empty graph as adjacency list (array)
	graph := make([][]int, N)
	for i := 0; i < N; i++ {
		graph[i] = make([]int, 0)
	}

	// Create colours array
	// 0 - not visited
	// 1 - red
	// 2 - blue
	colours := make([]int, N)

	// Fill graph
	for i := 0; i < M; i++ {
		// Read edge
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)

		// Add edge
		graph[a-1] = append(graph[a-1], b-1)
		graph[b-1] = append(graph[b-1], a-1)
	}

	var dfs func(int, int) bool
	dfs = func(v, colour int) bool {
		colours[v] = colour
		neighbours := graph[v]
		for _, neighbour := range neighbours {
			if colours[neighbour] == colour {
				return false
			}
			if colours[neighbour] == 0 {
				newColour := (colour % 2) + 1
				res := dfs(neighbour, newColour)
				if !res {
					return false
				}
			}
		}
		return true
	}

	for i := 0; i < N; i++ {
		if colours[i] == 0 {
			res := dfs(i, 1)
			if !res {
				fmt.Println("NO")
				return
			}
		}
	}

	fmt.Println("YES")
}
