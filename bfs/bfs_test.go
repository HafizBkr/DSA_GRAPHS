package graph

import "testing"

func TestBFS(t *testing.T) {
	graph := map[int][]int{
		0: {1, 2},
		1: {2},
		2: {0, 3},
		3: {3},
	}
	BFS(graph, 2)
}
