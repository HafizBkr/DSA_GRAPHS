package graph

import "fmt"

func BFS(graph map[int][]int, start int) {
	visited := make(map[int]bool)
	queue := []int{start}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if !visited[node] {
			visited[node] = true
			fmt.Println(node)
			queue = append(queue, graph[node]...)
		}
	}
}
