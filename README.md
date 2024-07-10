# Graph Algorithms in Go

This repository contains implementations of various graph algorithms in Go programming language.

## Algorithms Implemented

### 1. Depth-First Search (DFS)

DFS is an algorithm for traversing or searching tree or graph data structures. It starts at the root (selecting some arbitrary node as the root in the case of a graph) and explores as far as possible along each branch before backtracking.

**Implementation:**

File: `dfs/dfs.go`

**Usage:**

    go run dfs/dfs.go

2. Breadth-First Search (BFS)

BFS is an algorithm for traversing or searching tree or graph data structures. It starts at the tree root (or an arbitrary node of a graph), and explores all of the neighbor nodes at the present depth prior to moving on to the nodes at the next depth level.

Implementation:

File: bfs/bfs.go

Usage:


    go run bfs/bfs.go

3. Dijkstra's Algorithm

Dijkstra's algorithm is an algorithm for finding the shortest paths between nodes in a graph, which may represent, for example, road networks.

Implementation:

File: djikstras/djikstras.go

Usage:


    go run djikstras/djikstras.go

4. Prim's Algorithm

Prim's algorithm is a greedy algorithm that finds a minimum spanning tree for a weighted undirected graph. This means it finds a subset of the edges that forms a tree that includes every vertex, where the total weight of all the edges in the tree is minimized.

Implementation:

File: prims/prims.go

Usage:


    go run prims/prims.go

5. Kruskal's Algorithm

Kruskal's algorithm is a minimum-spanning-tree algorithm which finds an edge of the least possible weight that connects any two trees in the forest. It is a greedy algorithm in graph theory.

Implementation:

File: kurskals/kurskals.go

Usage:


    go run kurskals/kurskals.go

Running the Tests

Each algorithm has a corresponding test file. To run the tests, navigate to the directory containing the test files and use the go test command.

Example:

sh

cd dfs
go test

Directory Structure


    .
    |-- README.md
    |-- bfs
    |   |-- bfs.go
    |   `-- bfs_test.go
    |-- dfs
    |   |-- dfs.go
    |   `-- dfs_test.go
    |-- djikstras
    |   |-- djikstars_test.go
    |   `-- djikstras.go
    |-- go.mod
    |-- kurskals
    |   |-- kurskals.go
    |   `-- kurskals_test.go
    `-- prims
        |-- prims.go
        `-- prims_test.go

Contributing

Contributions and improvements are welcome! If you have any suggestions or find issues, please feel free to open an issue or submit a pull request.
