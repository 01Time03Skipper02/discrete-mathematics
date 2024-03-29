package main

import "fmt"

type Edge struct {
	vertex int
	id     int
}

type Connection struct {
	dest   int
	weight int
}

type Vertex struct {
	connections []Connection
	visited     bool
}

func main() {
	var graph []Vertex
	var n, m int
	fmt.Scanf("%d\n %d\n", &n, &m)

	for i := 0; i < n; i++ {
		var ver Vertex
		ver.connections = make([]Connection, 0)
		ver.visited = false
		graph = append(graph, ver)
	}

	for i := 0; i < m; i++ {
		var a, b, w int
		fmt.Scanf("%d %d %d\n", &a, &b, &w)
		graph[a].connections = append(graph[a].connections, Connection{b, w})
		graph[b].connections = append(graph[b].connections, Connection{a, w})
	}

	//
	edges := make([]Edge, 0)
	for i := 0; i < len(graph[0].connections); i++ {
		edges = append(edges, Edge{0, i})
	}

	sum := 0
	for len(edges) > 0 {
		edge_min := 0
		for i := 0; i < len(edges); i++ {
			if graph[edges[i].vertex].connections[edges[i].id].weight <
				graph[edges[edge_min].vertex].connections[edges[edge_min].id].weight {
				edge_min = i
			}
		}
		v := edges[edge_min].vertex
		graph[v].visited = true
		to := graph[v].connections[edges[edge_min].id].dest
		//fmt.Println(v, to, graph[v].connections[edges[edge_min].id].weight)

		if !graph[to].visited {
			graph[to].visited = true
			sum += graph[v].connections[edges[edge_min].id].weight
			for i := 0; i < len(graph[to].connections); i++ {
				if graph[graph[to].connections[i].dest].visited {
					continue
				}
				edges = append(edges, Edge{to, i})
			}
		}

		edges = append(edges[:edge_min], edges[edge_min+1:]...)
	}

	fmt.Println(sum)
}
