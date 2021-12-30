package main
import (
	"fmt"
)
func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)
	var d []int
	for _, v := range r {
		d = append(d, int(v-'0'))
	}

	sum, one, two := 0, 0, 0
	for _, v := range d {
		sum += v
		switch v % 3 {
		case 1:
			one++
		case 2:
			two++
		}
	}

	l := len(d)
	switch sum % 3 {
	case 0:
		fmt.Println(0)
		return
	case 1:
		if one > 0 && l > 1 {
			fmt.Println(1)
			return
		}
		if two > 1 && l > 2 {
			fmt.Println(2)
			return
		}
	case 2:
		if two > 0 && l > 1 {
			fmt.Println(1)
			return
		}
		if one > 1 && l > 2 {
			fmt.Println(2)
			return
		}
	}

	fmt.Println(-1)
}

