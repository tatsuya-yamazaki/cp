package main

import(
	"fmt"
	"math"
)
func main() {
	var(
		n, ans int
	)

	fmt.Scan(&n)
	var a, b []int
	for i:=0; i<n; i++ {
		var x, y int
		fmt.Scan(&x)
		fmt.Scan(&y)
		a = append(a, x)
		b = append(b, y)
	}

	for i:=0; i<n; i++ {
		for j:=i+1; j<n; j++ {
			ans = max(ans, (a[i] - a[j]) * (a[i] - a[j]) + (b[i] - b[j]) * (b[i] - b[j]))
		}
	}
	fmt.Println(math.Sqrt(float64(ans)))
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
