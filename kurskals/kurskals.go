// kruskal.go
package graph

import (
	"sort"
)

type Edge struct {
	u, v, weight int
}

type Graph struct {
	edges []Edge
	nodes int
}

func NewGraph(nodes int) *Graph {
	return &Graph{
		edges: []Edge{},
		nodes: nodes,
	}
}

func (g *Graph) AddEdge(u, v, weight int) {
	g.edges = append(g.edges, Edge{u, v, weight})
}

func find(parent []int, i int) int {
	if parent[i] == i {
		return i
	}
	return find(parent, parent[i])
}

func union(parent, rank []int, x, y int) {
	xroot := find(parent, x)
	yroot := find(parent, y)

	if rank[xroot] < rank[yroot] {
		parent[xroot] = yroot
	} else if rank[xroot] > rank[yroot] {
		parent[yroot] = xroot
	} else {
		parent[yroot] = xroot
		rank[xroot]++
	}
}

func Kruskal(g *Graph) []Edge {
	result := []Edge{}
	sort.Slice(g.edges, func(i, j int) bool {
		return g.edges[i].weight < g.edges[j].weight
	})

	parent := make([]int, g.nodes)
	rank := make([]int, g.nodes)

	for i := range parent {
		parent[i] = i
	}

	e := 0
	i := 0
	for e < g.nodes-1 && i < len(g.edges) {
		u := g.edges[i].u
		v := g.edges[i].v
		weight := g.edges[i].weight
		i++
		x := find(parent, u)
		y := find(parent, v)

		if x != y {
			e++
			result = append(result, Edge{u, v, weight})
			union(parent, rank, x, y)
		}
	}

	return result
}
