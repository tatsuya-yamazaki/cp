package main
import(
	"fmt"
	"sort"
	"math"
)
func main() {
	var n int
	var ans []int
	fmt.Scan(&n)

	a := 1
	b := n
	s := int(math.Sqrt(float64(n)))
	for {
		b = n / a
		if a > s {
			break
		}
		if n % a != 0 {
			a++
			continue
		}
		ans = append(ans, a)
		if a != b {
			ans = append(ans, b)
		}
		a++
	}

	sort.Ints(ans)

	for _, v := range ans {
		fmt.Println(v)
	}
}
