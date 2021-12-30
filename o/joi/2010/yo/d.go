package main
import(
	"fmt"
)

func main() {
	var(
		n, num int
		dp [101][21]int
	)

	fmt.Scan(&n)

	fmt.Scan(&num)
	dp[1][num] = 1
	for i:=2; i<n; i++ {
		fmt.Scan(&num)
		for j:=0; j<21; j++ {
			if dp[i-1][j] == 0 {
				continue
			}
			if j + num < 21 {
				dp[i][j+num] += dp[i-1][j]
			}
			if j - num >= 0 {
				dp[i][j-num] += dp[i-1][j]
			}
		}
	}
	fmt.Scan(&num)

	fmt.Println(dp[n-1][num])
}
