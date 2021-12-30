package main

import(
	"fmt"
)

func main() {
	var n, ans int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&s)
	cs := []rune(s)

	l, r := 0, 0
	var current rune
	count := 0
	for r < n {
		if l == r {
			current = cs[r]
			count = 0
		}
		if current == cs[r] {
			count++
		} else {
			ans += combi(count)
			l = r
			continue
		}
		r++
	}
	ans += combi(count)

	fmt.Println(ans)
}

func combi(n int) int {
	if n > 1 {
		return n * (n-1) / 2
	}
	return 0
}
