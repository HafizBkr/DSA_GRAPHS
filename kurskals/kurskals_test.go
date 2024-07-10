package graph

import (
	"testing"
)

func TestKruskal(t *testing.T) {
	g := NewGraph(4)
	g.AddEdge(0, 1, 10)
	g.AddEdge(0, 2, 6)
	g.AddEdge(0, 3, 5)
	g.AddEdge(1, 3, 15)
	g.AddEdge(2, 3, 4)

	mst := Kruskal(g)
	expectedWeight := 19
	totalWeight := 0
	for _, edge := range mst {
		totalWeight += edge.weight
	}
	if totalWeight != expectedWeight {
		t.Errorf("got %v, want %v", totalWeight, expectedWeight)
	}
}
