package main

import "fmt"

type Vertex struct {
	connections []int
	used        bool
	visited     bool
	comp        int
	ins         int
	min_v       int
}

func refresh(graph []Vertex) {
	for i := 0; i < len(graph); i++ {
		graph[i].used = false
	}
}

func topsort(graph []Vertex, v int, order *[]int) {
	graph[v].used = true
	for i := 0; i < len(graph[v].connections); i++ {
		to := graph[v].connections[i]
		if graph[to].used {
			continue
		}
		topsort(graph, to, order)
	}
	graph[v].visited = true
	*order = append(*order, v)
}

func unite(graph []Vertex, v int, cond []Vertex, comp int) {
	graph[v].used = true
	graph[v].comp = comp
	if v < cond[comp].min_v {
		cond[comp].min_v = v
	}
	for i := 0; i < len(graph[v].connections); i++ {
		to := graph[v].connections[i]
		if graph[to].used {
			if graph[to].comp != comp {
				cond[comp].connections = append(cond[comp].connections, graph[to].comp)
				cond[graph[to].comp].ins++
			}
			continue
		}
		unite(graph, to, cond, comp)
	}
}

func main() {
	var graph []Vertex
	var n, m int
	fmt.Scanf("%d\n %d\n", &n, &m)

	for i := 0; i < n; i++ {
		var ver Vertex
		ver.connections = make([]int, 0)
		graph = append(graph, ver)
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanf("%d %d\n", &a, &b)
		graph[a].connections = append(graph[a].connections, b)
	}

	order := make([]int, 0)
	for i := 0; i < len(graph); i++ {
		if !graph[i].used {
			topsort(graph, i, &order)
		}
	}

	refresh(graph)
	cond := make([]Vertex, 0)
	for i := 0; i < len(order); i++ {
		v := order[i]
		if !graph[v].used {
			var ver Vertex
			ver.connections = make([]int, 0)
			ver.min_v = 100000
			cond = append(cond, ver)
			unite(graph, v, cond, len(cond)-1)
		}
	}

	for i := 0; i < len(cond); i++ {
		if cond[i].ins == 0 {
			fmt.Println(cond[i].min_v)
		}
	}
}
