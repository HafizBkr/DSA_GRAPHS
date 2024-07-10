// depth_first_search.go
package graph

import "fmt"

func DFS(graph map[int][]int, start int, visited map[int]bool) {
	if visited[start] {
		return
	}
	visited[start] = true
	fmt.Println(start)
	for _, neighbor := range graph[start] {
		if !visited[neighbor] {
			DFS(graph, neighbor, visited)
		}
	}
}
