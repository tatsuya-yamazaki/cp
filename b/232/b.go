package main

import(
	"fmt"
	"reflect"
)

func main() {
	var s, t string
	fmt.Scan(&s)
	fmt.Scan(&t)
	sr := []rune(s)
	tr := []rune(t)

	var ssr []rune
	var d rune
	if tr[0] >= sr[0] {
		d = tr[0]-sr[0]
	} else {
		d = 26 - (sr[0]-tr[0])
	}
	for _, r := range sr {
		n := r+d
		if n > 'z' {
			n -= 26
		}
		ssr = append(ssr, n)
	}

	if reflect.DeepEqual(ssr, tr) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
