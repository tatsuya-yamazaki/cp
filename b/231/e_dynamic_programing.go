package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var n, x int
	fmt.Scan(&n, &x)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var a []int
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	q := make([]int, n, n)
	for i:=n-1; i>=0; i-- {
		q[i] = x/a[i]
		x %= a[i]
	}

	var dp [61][2]int
	dp[n-1][0] = q[n-1]
	dp[n-1][1] = q[n-1] + 1
	for i:=n-2; i>=0; i-- {
		dp[i][0] = min(dp[i+1][0] + q[i], dp[i+1][1] + a[i+1]/a[i] - q[i])
		dp[i][1] = min(dp[i+1][0] + q[i] + 1, dp[i+1][1] + a[i+1]/a[i] - q[i] - 1)
	}

	fmt.Println(dp[0][0])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
