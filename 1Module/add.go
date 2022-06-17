package main

import (
	"fmt"
)

func add(a, b []int32, p int) []int32 {
	var res []int32
	var q int32
	i := 0
	if len(a) > len(b) {
		for ; i < len(b); i++ {
			res = append(res, (a[i]+b[i]+q)%int32(p))
			q = (a[i] + b[i] + q) / int32(p)
		}
		for j := len(b); j < len(a); j++ {
			res = append(res, (a[j]+q)%int32(p))
			q = (a[j] + q) / int32(p)
		}
		if q > 0 {
			res = append(res, q)
		}
	} else if len(a) < len(b) {
		for ; i < len(a); i++ {
			res = append(res, (a[i]+b[i]+q)%int32(p))
			q = (a[i] + b[i] + q) / int32(p)
		}
		for j := len(a); j < len(b); j++ {
			res = append(res, (b[j]+q)%int32(p))
			q = (b[j] + q) / int32(p)
		}
		if q > 0 {
			res = append(res, q)
		}
	} else {
		for ; i < len(a); i++ {
			res = append(res, (a[i]+b[i]+q)%int32(p))
			q = (a[i] + b[i] + q) / int32(p)
		}
		if q > 0 {
			res = append(res, q)
		}
	}
	return res
}

func main() {
	a := []int32{0, 1, 0, 0, 1}
	b := []int32{0, 1, 1, 0, 1}
	fmt.Println(add(a, b, 2))
}
