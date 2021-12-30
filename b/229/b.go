package main

import(
	"fmt"
)

func main() {
	var a, b string

	fmt.Scan(&a, &b)

	al := []rune(a)
	bl := []rune(b)

	l := len(al)
	la := len(al)
	lb := len(bl)
	if len(al) > len(bl) {
		l = len(bl)
	}

	for i:=0; i<l; i++ {
		ar := al[la-i-1]
		br := bl[lb-i-1]
		if (int(ar) + int(br)-96) > 9 {
			fmt.Println("Hard")
			return
		}
	}

	fmt.Println("Easy")
}
