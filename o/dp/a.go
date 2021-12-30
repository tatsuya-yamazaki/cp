package main

import(
	"fmt"
	"math"
)

func main() {
	var(
		N int
		hl []int
	)
	fmt.Scanf("%d",&N)
	for i:=0; i<N; i++ {
		t := 0
		fmt.Scanf("%d",&t)
		hl = append(hl, t)
	}

	dp := make([]int, N)

	for i, _ := range hl {
		switch i {
		case 0:
			dp[0] = 0
			continue
		case 1:
			dp[1] = int(math.Abs(float64(hl[1] - hl[0])))
			continue
		}

		cost1 := int(math.Abs(float64(hl[i] - hl[i-1]))) + dp[i-1]
		cost2 := int(math.Abs(float64(hl[i] - hl[i-2]))) + dp[i-2]
		if cost1 < cost2 {
			dp[i] = cost1
		} else {
			dp[i] = cost2
		}
	}

	fmt.Println(dp[N-1])
}
