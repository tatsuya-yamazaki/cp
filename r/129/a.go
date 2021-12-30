package main

import(
	"fmt"
)

func main() {
	var n, l, r int

	fmt.Scan(&n)
	fmt.Scan(&l)
	fmt.Scan(&r)

	k := 0
	num := 1
	for n >= num {
		num *= 2
		k++
	}

	var ks []int
	rest := n
	for num != 0 && rest != 0 {
		num /= 2
		k--
		if rest >= num {
			ks = append(ks, k)
			rest -= num
		}
	}

	ans := 0
	for _, v := range ks {
		maxX := min(r, pow(2,v+1)-1)
		minX := max(l, pow(2,v))
		if maxX >= minX {
			ans += maxX - minX + 1
		}
	}

	fmt.Println(ans)
}

func pow(num, degree int) int {
	n := 1
	for d:=0; d<degree; d++ {
		n *= num
	}
	return n
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}
