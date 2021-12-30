package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var(
		n, x int
	)

	fmt.Scan(&n)
	fmt.Scan(&x)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	dp := make(map[int]int)
	dp[1] = 1
	for i:=0; i<n; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		nextDp := make(map[int]int)
		for j:=0; j<k; j++ {
			sc.Scan()
			a, _ := strconv.Atoi(sc.Text())
			for k, v := range dp {
				if k <= x / a {
					nextDp[k*a] += v
				}
			}
		}
		dp = nextDp
	}

	fmt.Println(dp[x])
}
