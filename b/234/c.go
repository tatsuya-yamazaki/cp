package main

import(
	"fmt"
)
func main() {
	var(
		n int
	)

	fmt.Scan(&n)

	s := ""
	d := 1
	for d <= n {
		if d & n > 0 {
			s += "2"
		} else {
			s += "0"
		}
		d *= 2
	}

	for i:=len(s)-1; i>=0; i-- {
		fmt.Print(s[i:i+1])
	}
	fmt.Println()
}
