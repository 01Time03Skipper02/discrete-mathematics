package main

import "fmt"

func gcd(a int, b int) int {
	a, b = abs(a), abs(b)
	for b != 0 {
		buf := b
		b = a % b
		a = buf
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Fraction struct {
	num, den int
}

func (fraction Fraction) simplify() Fraction {
	fracGCD := gcd(fraction.num, fraction.den)
	if fracGCD != 0 {
		fraction.num /= fracGCD
		fraction.den /= fracGCD
		if fraction.den < 0 {
			fraction.num *= -1
			fraction.den *= -1
		}
	}
	return fraction
}

func (a Fraction) mul(b Fraction) Fraction {
	var res Fraction
	res.num = a.num * b.num
	res.den = a.den * b.den
	return res.simplify()
}

func (a Fraction) add(b Fraction) Fraction {
	var res Fraction
	k := abs(a.den*b.den) / gcd(abs(a.den), abs(b.den))
	resNum := (a.num * (k / a.den)) + (b.num * (k / b.den))
	resDen := k
	res = Fraction{resNum, resDen}
	return res.simplify()
}

func (a Fraction) diff(b Fraction) Fraction {
	var res Fraction
	k := abs(a.den*b.den) / gcd(abs(a.den), abs(b.den))
	resNum := (a.num * (k / a.den)) - (b.num * (k / b.den))
	resDen := k
	res = Fraction{resNum, resDen}
	return res.simplify()
}

func (a Fraction) div(b Fraction) Fraction {
	var res Fraction
	resNum := a.num * b.den
	resDen := a.den * b.num
	res = Fraction{resNum, resDen}
	return res.simplify()
}

func (a Fraction) findK(b Fraction) Fraction {
	return a.div(b)
}

func diffStrs(matrix [][]Fraction, ans []Fraction, i1, i2, n int, k Fraction) ([][]Fraction, []Fraction) {
	for i := 0; i < n; i++ {
		matrix[i1][i] = matrix[i1][i].diff(matrix[i2][i].mul(k))
	}
	ans[i1] = ans[i1].diff(ans[i2].mul(k))
	return matrix, ans
}

func swapLines(matrix [][]Fraction, ans []Fraction, n, i1, i2 int) ([][]Fraction, []Fraction) {
	for j := 0; j < n; j++ {
		matrix[i1][j], matrix[i2][j] = matrix[i2][j], matrix[i1][j]
	}
	ans[i1], ans[i2] = ans[i2], ans[i1]
	return matrix, ans
}

func fixMatrix(matrix [][]Fraction, ans []Fraction, n int, start, j int) ([][]Fraction, []Fraction, int) {
	if matrix[start][start].num == 0 {
		for i := start + 1; i < n; i++ {
			if matrix[i][0].num != 0 {
				swapLines(matrix, ans, n, i, 0)
			}
		}
		if matrix[start][start].num == 0 {
			return matrix, ans, j + 2
		}
	}
	return matrix, ans, j
}

func gauss(matrix [][]Fraction, ans []Fraction, n int) []Fraction {

	res := make([]Fraction, n)

	for i := 0; i < n; i++ {
		res[i] = Fraction{0, 1}
	}

	for j := 0; j < n; j++ {
		cnt := 0
		for i := j; i < n; i++ {
			if matrix[i][j].num == 0 {
				cnt += 1
			}
		}
		if cnt == n-j {
			return res
		}
		for i := j + 1; i < n; i++ {
			matrix, ans, j = fixMatrix(matrix, ans, n, 0, j)
			k := matrix[i][j].findK(matrix[j][j])
			matrix, ans = diffStrs(matrix, ans, i, j, n, k)
			for z := 0; z < n; z++ {
				check := 0
				for l := 0; l < n; l++ {
					if matrix[z][l].num == 0 {
						check++
					}
				}
				if ans[z].num == 0 {
					check++
				}
				if check == n || check == n+1 {
					return res
				}
			}
			/*for z := 0; z < n; z++ {
				for l := 0; l < n; l++ {
					fmt.Print(matrix[z][l], " ")
				}
				fmt.Println(ans[z])
			}*/
		}
	}

	/*for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println(ans[i])
	}*/

	for i := 0; i < n; i++ {
		check := 0
		for j := 0; j < n; j++ {
			if matrix[i][j].num == 0 {
				check++
			}
		}
		if ans[i].num == 0 {
			check++
		}
		if check == n {
			return res
		}
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			ans[i] = ans[i].diff(matrix[i][j].mul(ans[j]))
		}
		ans[i] = ans[i].div(matrix[i][i])
	}

	for i := 0; i < n; i++ {
		res[i] = ans[i]
	}

	return res
}

func main() {
	var n int
	fmt.Scan(&n)
	matrix := make([][]Fraction, n)
	ans := make([]Fraction, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]Fraction, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			var x int
			fmt.Scan(&x)
			matrix[i][j] = Fraction{x, 1}
		}
		var x int
		fmt.Scan(&x)
		ans[i] = Fraction{x, 1}
	}
	res := gauss(matrix, ans, n)
	check := 0
	for i := 0; i < n; i++ {
		if res[i].num == 0 {
			check++
		}
	}
	if check == n {
		fmt.Println("No solution")
		return
	}
	for i := 0; i < n; i++ {
		fmt.Print(res[i].num)
		fmt.Print("/")
		fmt.Println(res[i].den)
	}
}
