package main

import "fmt"

var arr []int

func less(i, j int) bool {
	if arr[i] < arr[j] {
		return true
	}
	if arr[i] > arr[j] {
		return false
	}
	return false
}

func swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}

func partition(less func(i, j int) bool, swap func(i, j int), start int, end int) int {
	i := start
	j := start
	for j < end {
		if less(j, end) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, end)
	return i
}

func qsortrec(less func(i, j int) bool, swap func(i, j int), start int, end int) {
	if start < end {
		q := partition(less, swap, start, end)
		qsortrec(less, swap, start, q-1)
		qsortrec(less, swap, q+1, end)
	}
}

func qsort(n int, less func(i, j int) bool, swap func(i, j int)) {
	qsortrec(less, swap, 0, n-1)
}

func main() {
	var n int
	fmt.Scan(&n)
	arr = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	qsort(n, less, swap)
	fmt.Println(arr)
}
