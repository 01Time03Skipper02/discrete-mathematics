package main

import "fmt"

type Vertex struct {
	neighbors []int
	dists     []int
}

func bfs(graph []Vertex, root, id int) {
	queue := make([]int, 0)
	begin := 0
	queue = append(queue, root)
	graph[root].dists[id] = 0
	for begin < len(queue) {
		v := queue[begin]
		begin++
		for i := 0; i < len(graph[v].neighbors); i++ {
			to := graph[v].neighbors[i]
			if graph[to].dists[id] < graph[v].dists[id]+1 {
				continue
			}
			graph[to].dists[id] = graph[v].dists[id] + 1
			queue = append(queue, to)
		}
	}
}

func main() {
	var graph []Vertex
	var n, m int
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%d\n", &m)

	for i := 0; i < n; i++ {
		var ver Vertex
		ver.neighbors = make([]int, 0)
		ver.dists = make([]int, 0)
		graph = append(graph, ver)
	}

	for i := 0; i < m; i++ {
		var a, b int
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d\n", &b)
		if a != b {
			graph[a].neighbors = append(graph[a].neighbors, b)
		}
		graph[b].neighbors = append(graph[b].neighbors, a)
	}

	var k int
	fmt.Scanf("%d\n", &k)

	for i := 0; i < k; i++ {
		var v int
		fmt.Scanf("%d", &v)
		for j := 0; j < n; j++ {
			graph[j].dists = append(graph[j].dists, n)
		}
		bfs(graph, v, i)
	}

	count := 0
	for i := 0; i < n; i++ {
		equal := true
		for j := 1; j < k; j++ {
			if graph[i].dists[j] != graph[i].dists[j-1] || graph[i].dists[j] == n {
				equal = false
				break
			}
		}
		if equal {
			fmt.Printf("%d ", i)
			count++
		}
	}
	if count == 0 {
		fmt.Printf("-")
	}
}
