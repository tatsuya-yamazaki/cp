package main
import(
	"fmt"
	"sort"
)
func main() {
	var n, es int
	var sl []int

	for i:=1; i<151; i++ {
		for j:=i; j<151; j++ {
			sl = append(sl, 4*i*j+3*i+3*j)
		}
	}
	sort.Ints(sl)

	fmt.Scan(&n)

	ans := n
	for i:=0; i<n; i++ {
		fmt.Scan(&es)
		for _, s := range sl {
			if es == s {
				ans--
				break
			}

			if s > es {
				break
			}
		}
	}

	fmt.Println(ans)
}
