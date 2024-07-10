package graph

import "testing"

func TestDFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}
	visited := make(map[int]bool)
	DFS(graph, 2, visited)
}
