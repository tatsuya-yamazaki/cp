package main
import(
	"fmt"
)
func main() {
	var n, k, a int
	fmt.Scan(&n, &k, &a)

	ans := a
	for {
		k--

		if k == 0 {
			break
		}

		a++
		if a > n {
			a = 1
		}
		ans = a

	}

	fmt.Println(ans)
}
