package main

import(
	"fmt"
)

func main() {
	var a, b int

	fmt.Scan(&a)
	fmt.Scan(&b)

	x := 0
	if b % 2 == 0 {
		x = b / 2
	} else {
		x = b*5
	}

	xe := 0
	tmp := x * 2
	for {
		tmp /= 10
		if tmp == 0 {
			break
		}
		xe++
	}

	x += pow(10, xe+1) * a

	fmt.Println(x)
}

func pow(a, b int) int {
	ans := 1
	for i:=0; i<b; i++ {
		ans *= a
	}
	return ans
}
