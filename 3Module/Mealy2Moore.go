package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		m, start, n int
		inp, out    string
		g           [][]int
		f           [][]string
	)

	Scan := bufio.NewScanner(os.Stdin)
	Scan.Scan()

	m, _ = strconv.Atoi(Scan.Text())
	inp = Scan.Text()
	start, _ = strconv.Atoi(Scan.Text())
	out = Scan.Text()
	n, _ = strconv.Atoi(Scan.Text())

	ainp := strings.Split(inp, " ")
	aout := strings.Split(out, " ")
	start++;

	g = make([][]int, n)
	for i := 0; i < n; i++ {
		g[i] = make([]int, m)
	}

	f = make([][]string, n)
	for i := 0; i < n; i++ {
		f[i] = make([]string, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var a int
			fmt.Scan(&a)
			g[i] = append(g[i], a)
		}
	}
	for i := 0; i < n; i++ {
		var line string
		line = Scan.Text()
		outs := strings.Split(line, " ")
		for j := 0; j < m; j++ {
			f[i] = append(f[i], outs[j])
		}
	}

	set := make(map[int]string)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			_, ok := set[g[i][j]]
			if !ok {
				buf, _ := strconv.Atoi(f[i][j])
				set[g[i][j]] = aout[buf]
			}
		}
	}

	fmt.Println("digraph {")
	fmt.Println("    rankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for k := range set[i] {
				buf, _ := strconv.Atoi(f[i][j])
				fmt.Println("    \"(", i, ",", k, ")\" -> \"(", g[i][j], ",", aout[buf], ")\" [label = \"", ainp[j], "\"]")
			}
		}
	}
	fmt.Println("}")
}
