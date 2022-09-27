package main

import "fmt"

type cond struct {
	t    []int
	f    []string
	used bool
	new  int
}

func dfs(i int, A []cond, m, order int) {
	A[i].used = true
	A[i].new = order
	order += 1
	for j := 0; j < m; j++ {
		if !A[A[i].t[j]].used {
			dfs(A[i].t[j], A, m, order)
		}
	}
}

func main() {
	var x int
	var s string
	var n, m, q, order int
	fmt.Scan(&n, &m, &q)
	A := make([]cond, n)
	A1 := make([]cond, n)
	for i := 0; i < n; i++ {
		A[i].t = make([]int, m)
		A1[i].t = make([]int, m)
		//A[i].t.resize(m);
		//A1[i].t.resize(m);
		for j := 0; j < m; j++ {
			fmt.Scan(&x)
			A[i].t[j] = x
		}
	}
	for i := 0; i < n; i++ {
		A[i].f = make([]string, m)
		A1[i].f = make([]string, m)
		//A[i].f.resize(m);
		//A1[i].f.resize(m);
		for j := 0; j < m; j++ {
			fmt.Scan(&s)
			A[i].f[j] = s
		}
	}
	dfs(q, A, m, order)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			A1[A[i].new].t[j] = A[A[i].t[j]].new
			A1[A[i].new].f[j] = A[i].f[j]
		}
	}
	fmt.Println(order)
	fmt.Println(m)
	fmt.Println(0)
	for i := 0; i < order; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(A1[i].t[j], " ")
		}
		fmt.Println()
	}
	for i := 0; i < order; i++ {
		for j := 0; j < m; j++ {
			fmt.Print(A1[i].f[j], " ")
		}
		fmt.Println()
	}
}
