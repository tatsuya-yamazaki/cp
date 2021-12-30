package main

import(
	"fmt"
)

func main() {
	var(
		str string
		ans int
	)
	fmt.Scanf("%s", &str)
	s := []rune(str)

	num := 4
	for i:=0; i<3; i++{
		r := s[i]
		numr := 0
		for _, c := range str {
			if r == c {
				numr++
			}
		}
		newNum := 4 - numr

		if num > newNum {
			num = newNum
		}
	}

	switch num {
	case 1:
		ans = 1
	case 2:
		ans = 3
	case 3:
		ans = 6
	}

	fmt.Println(ans)
}
