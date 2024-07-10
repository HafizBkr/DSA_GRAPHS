package graph

import (
	"container/heap"
)

type Edge struct {
	node   int
	weight int
}

type PriorityQueue []Edge

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Edge)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Prim(graph map[int]map[int]int, start int) []Edge {
	mst := []Edge{}
	visited := make(map[int]bool)
	pq := &PriorityQueue{}
	heap.Push(pq, Edge{node: start, weight: 0})

	for pq.Len() > 0 {
		edge := heap.Pop(pq).(Edge)
		if visited[edge.node] {
			continue
		}
		visited[edge.node] = true
		if edge.weight != 0 {
			mst = append(mst, edge)
		}

		for neighbor, weight := range graph[edge.node] {
			if !visited[neighbor] {
				heap.Push(pq, Edge{node: neighbor, weight: weight})
			}
		}
	}

	return mst
}
