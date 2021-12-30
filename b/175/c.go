package main
import "fmt"
func main() {
	var x, k, d int
	fmt.Scan(&x, &k, &d)
	x = abs(x)
	r := k - x/d
	if r < 0 {
		fmt.Println(x - d*k)
		return
	}
	m := x % d
	mm := -(m - d)
	if r % 2 == 0 {
		fmt.Println(m)
	} else {
		fmt.Println(mm)
	}
}
func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
