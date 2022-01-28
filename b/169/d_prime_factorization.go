package main

import(
	"fmt"
	"math"
)

func main() {
	n := 0
	fmt.Scan(&n)
	pf := primeFactorize(n)
	ans := 0
	for _, v := range pf {
		d := v[1]
		e := 1
		c := 0
		for d > 0 {
			d -= e
			e++
			if d >= 0 {
				c++
			}
		}
		ans += c
	}
	fmt.Println(ans)
}

func primeFactorize(n int) (ret [][2]int) {
	root := int(math.Sqrt(float64(n)))
	for i:=2; i<=root; i++ {
		var factor [2]int
		for n % i == 0 {
			n /= i
			factor[1]++
		}
		if factor[1] != 0 {
			factor[0] = i
			ret = append(ret, factor)
		}
	}
	if n != 1 {
		ret = append(ret, [2]int{n, 1})
	}
	return
}
