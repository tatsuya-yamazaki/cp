package main

import(
	"fmt"
)
func main() {
	var(
		l, r int
		s string
	)

	fmt.Scan(&l)
	fmt.Scan(&r)
	fmt.Scan(&s)
	l--
	r--

	for i:=0; i<l; i++ {
		fmt.Print(s[i:i+1])
	}
	for i:=r; i>=l; i-- {
		fmt.Print(s[i:i+1])
	}
	for i:=r+1; i<len(s); i++ {
		fmt.Print(s[i:i+1])
	}
	fmt.Println()
}
