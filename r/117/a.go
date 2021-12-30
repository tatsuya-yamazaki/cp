package main
import "fmt"
func main() {
	var a, b int
	fmt.Scan(&a)
	fmt.Scan(&b)

	ans := f(a,b)

	for i, v := range ans {
		if a < b {
			fmt.Print(v)
		} else {
			fmt.Print(-v)
		}
		if i < len(ans)-1 {
			fmt.Print(" ")
		} else {
			fmt.Print("\n")
		}
	}
}

func f(a, b int) []int {
	var ret []int
	if b < a {
		a, b = b, a
	}
	sum := 0
	xb := 0
	for i:=0; i<b; i++ {
		xb--
		ret = append(ret, xb)
		sum += xb
	}
	xa := 0
	for i:=0; i<a-1; i++ {
		xa++
		ret = append(ret, xa)
		sum += xa
	}
	ret = append(ret, -sum)
	return ret
}
