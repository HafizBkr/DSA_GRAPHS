package graph

import (
	"testing"
)

func TestPrim(t *testing.T) {
	graph := map[int]map[int]int{
		0: {1: 1, 2: 4},
		1: {0: 1, 2: 2, 3: 5},
		2: {0: 4, 1: 2, 3: 1},
		3: {1: 5, 2: 1},
	}
	mst := Prim(graph, 0)
	expectedWeight := 4
	totalWeight := 0
	for _, edge := range mst {
		totalWeight += edge.weight
	}
	if totalWeight != expectedWeight {
		t.Errorf("got %v, want %v", totalWeight, expectedWeight)
	}
}
