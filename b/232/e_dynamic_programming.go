package main

import(
	"fmt"
)

func main() {
	var h, w, k, x1, x2, y1, y2 int
	fmt.Scan(&h, &w, &k)
	fmt.Scan(&x1, &y1)
	fmt.Scan(&x2, &y2)
	d := 998244353

	var dp [1000002][2][2]int

	if x1 == x2 && y1 == y2 {
		dp[0][1][1] = 1
	} else if x1 == x2 {
		dp[0][1][0] = 1
	} else if y1 == y2 {
		dp[0][0][1] = 1
	} else {
		dp[0][0][0] = 1
	}

	for i:=1; i<=k; i++ {
		dp[i][1][0] = dp[i-1][0][0]
		dp[i][1][0] += (dp[i-1][1][0] * (w - 2)) % d
		dp[i][1][0] %= d
		dp[i][1][0] += (dp[i-1][1][1] * (w - 1)) % d
		dp[i][1][0] %= d

		dp[i][0][1] = dp[i-1][0][0]
		dp[i][0][1] += (dp[i-1][0][1] * (h - 2)) % d
		dp[i][0][1] %= d
		dp[i][0][1] += (dp[i-1][1][1] * (h - 1)) % d
		dp[i][0][1] %= d

		dp[i][1][1] = dp[i-1][0][1]
		dp[i][1][1] += dp[i-1][1][0]
		dp[i][1][1] %= d

		dp[i][0][0] = (dp[i-1][1][0] * (h - 1)) % d
		dp[i][0][0] += (dp[i-1][0][1] * (w - 1)) % d
		dp[i][0][0] %= d
		dp[i][0][0] += (dp[i-1][0][0] * (h - 2)) % d
		dp[i][0][0] %= d
		dp[i][0][0] += (dp[i-1][0][0] * (w - 2)) % d
		dp[i][0][0] %= d
	}

	fmt.Println(dp[k][1][1])
}
