package main

import(
	"fmt"
)
func main() {
	var(
		n, k int
	)

	fmt.Scan(&n)
	fmt.Scan(&k)

	d := k - n
	if d < 0 {
		fmt.Println(0)
	} else {
		q := d / 10
		r := d % 10

		if r > 0 {
			fmt.Println(q+1)
		} else {
			fmt.Println(q)
		}
	}
}
