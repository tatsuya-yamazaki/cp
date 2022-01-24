package main

import(
	"fmt"
)
func main() {
	var(
		n string
	)

	fmt.Scan(&n)
	r := []rune(n)
	ans := 0
	for _, v := range r {
		ans += int(v - '0')
	}

	fmt.Println(ans*111)
}
