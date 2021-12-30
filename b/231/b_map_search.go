package main

import(
	"fmt"
)

func main() {
	var n  int
	fmt.Scan(&n)
	ans := make(map[string]int)
	for i:=0; i<n; i++ {
		var s string
		fmt.Scan(&s)
		_, ok := ans[s]
		if ok {
			ans[s]++
		} else {
			ans[s] = 1
		}
	}
	maxName := ""
	max := 0
	for k, v := range ans {
		if v > max {
			max = v
			maxName = k
		}
	}
	fmt.Println(maxName)
}
