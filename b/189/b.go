package main
import "fmt"
func main() {
	var n, x int
	fmt.Scan(&n, &x)
	x *= 100
	ans := -1
	for i:=0; i<n; i++ {
		var v, p int
		fmt.Scan(&v, &p)
		x -= v * p
		if x < 0 && ans == -1 {
			ans = i + 1
		}
	}
	fmt.Println(ans)
}
