package main

import(
	"fmt"
)

func main() {
	var s string
	var k int

	fmt.Scan(&s)
	fmt.Scan(&k)

	dots := make([]int, len(s)+1, len(s)+1)
	sumDots := 0
	xs := make([]int, len(s)+1, len(s)+1)
	sumXs := 0
	for i, r := range s {
		if r == '.' {
			sumDots++
		} else {
			sumXs++
		}
		dots[i+1] = sumDots
		xs[i+1] = sumXs
	}

	ans := 0
	l, r := 1, 1
	for l<=len(s) {
		for r<=len(s) {
			d := dots[r] - dots[l-1]
			x := xs[r] - xs[l-1]
			if d > k {
				break
			}
			if x + d > ans {
				ans = x + d
			}
			r++
		}
		l++
	}

	fmt.Println(ans)
}
