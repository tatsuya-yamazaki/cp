package main
import(
	"fmt"
)
func main() {
	var n, k, r, s, p, ans int
	var t string
	fmt.Scan(&n)
	fmt.Scan(&k)
	fmt.Scan(&r)
	fmt.Scan(&s)
	fmt.Scan(&p)
	fmt.Scan(&t)
	tr := []rune(t)

	for i:=0; i<k; i++ {
		var dp []bool
		c := -1
		for j:=i; j<n; j += k {
			c++
			if j >= k {
				if tr[j] == tr[j-k] && dp[c-1] {
					dp = append(dp, false)
					continue
				}
			}
			switch tr[j] {
			case 'r':
				ans += p
			case 's':
				ans += r
			case 'p':
				ans += s
			}
			dp = append(dp, true)
		}
	}
	fmt.Println(ans)
}
