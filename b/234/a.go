package main

import(
	"fmt"
)
func main() {
	var(
		n int
	)

	fmt.Scan(&n)

	fmt.Println(f(f(f(n)+n)+f(f(n))))
}
func f(x int) int {
	return x * x + 2 * x + 3
}
