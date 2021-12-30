package main

import(
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)
	a, _ := strconv.Atoi(string(s[0]))
	b, _ := strconv.Atoi(string(s[2]))
	fmt.Println(a * b)
}
