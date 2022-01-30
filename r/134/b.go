package main

import(
	"fmt"
)

func main() {
	var n int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&s)
	runes := []rune(s)

	l, r := 0, len(s)-1
	for l < r {
		nr := r
		target := -1
		for l < nr {
			if runes[l] > runes[nr] && (target == -1 || runes[nr] < runes[target]) {
				target = nr
			}
			nr--
		}

		if target > -1 {
			runes[l], runes[target] = runes[target], runes[l]
			r = target
		}

		l++
	}
	ans := ""
	for _, v := range runes {
		ans += string(v)
	}

	fmt.Println(ans)
}
