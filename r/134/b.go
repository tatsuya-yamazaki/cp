package main

import(
	"fmt"
	"sort"
)

func main() {
	var n int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&s)
	r := make([]int, len(s))
	for i, v := range s {
		r[i] = int(v)
	}

	sort.Slice(r, func(i, j int) bool { return true }) //reverse

	dp := make([]int, len(s))

	sort.Slice(r, func(i, j int) bool { return true }) //reverse
}
