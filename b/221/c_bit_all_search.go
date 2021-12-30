package main

import "fmt"
import "sort"

func main() {
	var str string
	fmt.Scanf("%s", &str)
	s := []rune(str)
	sort.Slice(s, func (i, j int) bool { return s[i] > s[j] })
	ans := 0
	for n := 0; n < (1<<len(s)); n++ {
		var l, r int

		for i := 0; i < len(s); i++ {
			if 0 < (n & (1 << i) ) {
				l = l * 10 + int(s[i] - '0')
			} else {
				r = r * 10 + int(s[i] - '0')
			}
		}
		if ans < l*r {
			ans = l*r
		}
	}
	fmt.Printf("%d\n", ans)
}
