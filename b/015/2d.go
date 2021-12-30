package main
import(
	"fmt"
)
func main() {
	var w, n, k, a, b int
	fmt.Scan(&w, &n, &k)

	var dp [51][10001]int

	for i:=0; i<n; i++ {
		fmt.Scan(&a, &b)
		for j:=k; j>0; j-- {
			for l:=w; l>a-1; l-- {
				if dp[j][l] < dp[j-1][l-a] + b {
					dp[j][l] = dp[j-1][l-a] + b
				}
			}
		}
	}

	fmt.Println(dp[k][w])
}
