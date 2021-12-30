package main

import(
	"fmt"
	"math"
)

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
}

func scanInts(times int) (ret []int) {
	var n int
	for i := 0; i < times; i++ {
		fmt.Scanf("%d", &n)
		ret = append(ret, n)
	}
	return
}

func scan2i() (a,b int) {
	fmt.Scanf("%d %d", &a, &b)
	return a, b
}

func main() {
	N, K := scan2i()
	Hl := scanInts(N)
	for i := 0; i < K; i++ {
		Hl = append(Hl, 0)
	}

	dp := make([]int, N + K, N + K)
	for i, _ := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	dp[1] = abs(Hl[1] - Hl[0])

	for i := 0; i < N; i++ {
		for j := 1; j <= K; j++ {
			cost := dp[i] + abs(Hl[i+j] - Hl[i])
			if cost < dp[i+j] {
				dp[i+j] = cost
			}
		}
	}

	fmt.Println(dp[N-1])
}
