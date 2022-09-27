package main

import "fmt"

func getDividers(x int) []int {
	var res []int
	for i := x; i >= 1; i-- {
		if x%i == 0 {
			res = append(res, i)
		}
	}
	return res
}

func haveEdge(a, b int) bool {
	for i := 2; i < a/b/2; i++ {
		if (a/b)%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	arr := getDividers(x)
	fmt.Println("graph {")
	for i := 0; i < len(arr); i++ {
		fmt.Println("    ", arr[i])
	}
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i]%arr[j] == 0 && haveEdge(arr[i], arr[j]) {
				fmt.Println("    ", arr[i], "--", arr[j])
			}
		}
	}
	fmt.Println("}")
}
