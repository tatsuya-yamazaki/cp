package main

import(
	"fmt"
)
func main() {
	var(
		n, ans int
	)

	fmt.Scan(&n)
	a := make([]int, 0)
	for i:=0; i<n; i++ {
		var x int
		fmt.Scan(&x)
		a = append(a, x)
	}

	for i:=0; i<n; i++ {
		if ans >= a[i] {
			break
		}
		ans = max(ans, a[i])
	}

	fmt.Println(ans)
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
