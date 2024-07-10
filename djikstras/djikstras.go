// dijkstra.go
package graph

import (
	"container/heap"
	"math"
)

type Item struct {
	node   int
	weight int
}

type PriorityQueue []Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func Dijkstra(graph map[int]map[int]int, start int) map[int]int {
	dist := make(map[int]int)
	for node := range graph {
		dist[node] = math.MaxInt32
	}
	dist[start] = 0

	pq := &PriorityQueue{}
	heap.Push(pq, Item{node: start, weight: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(Item)
		node := item.node

		for neighbor, weight := range graph[node] {
			alt := dist[node] + weight
			if alt < dist[neighbor] {
				dist[neighbor] = alt
				heap.Push(pq, Item{node: neighbor, weight: alt})
			}
		}
	}

	return dist
}
