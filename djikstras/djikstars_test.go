package graph

import (
	"reflect"
	"testing"
)

func TestDijkstra(t *testing.T) {
	graph := map[int]map[int]int{
		0: {1: 1, 2: 4},
		1: {2: 2, 3: 5},
		2: {3: 1},
		3: {},
	}
	expected := map[int]int{0: 0, 1: 1, 2: 3, 3: 4}
	dist := Dijkstra(graph, 0)
	if !reflect.DeepEqual(dist, expected) {
		t.Errorf("got %v, want %v", dist, expected)
	}
}
