package main

import (
	"fmt"
	"math"
)

func add(a, b []int32, p int) []int32 {
	var res []int32
	var buf int32
	var i int
	for ; i < int(math.Min(float64(len(a)), float64(len(b)))); i++ {
		res = append(res, ((a[i]+b[i])+buf)%int32(p))
		buf = (a[i] + b[i]) / int32(p)
	}
	if len(a) == len(b) {
		res = append(res, buf)
	}
	for ; i < int(math.Max(float64(len(a)), float64(len(b)))); i++ {
		if len(a) > len(b) {
			res = append(res, (a[i]+buf)%int32(p))
			buf = (a[i] + buf) / int32(p)
			if i == len(a)-1 && buf != 0 {
				res = append(res, buf)
			}
		}
		if len(b) > len(a) {
			res = append(res, b[i]+buf)
			buf = (b[i] + buf) / int32(p)
			if i == len(b)-1 && buf != 0 {
				res = append(res, buf)
			}
		}
	}
	return res
}

func main() {
	var (
		c, d int32
		a, b []int32
		p    int32
	)
	fmt.Scan(&c, &d, &p)
	for c > 0 {
		a = append(a, c%p)
		c /= p
	}
	for d > 0 {
		b = append(b, d%p)
		d /= p
	}
	res := add(a, b, int(p))

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(res)
}
