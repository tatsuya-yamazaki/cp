package main

import(
	"fmt"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)
	var ss []string
	for i:=0; i<h; i++ {
		var s string
		fmt.Scan(&s)
		ss = append(ss, s)
	}

	var dp [][]int
	for j:=0; j<h; j++ {
		dpr := make([]int, w)
		dp = append(dp, dpr)
	}
	dp[0][0] = 1

	for i:=0; i<h+w-2; i++ {
		for j:=0; j<h; j++ {
			for k:=0; k<w; k++ {
				if dp[j][k] == 0 {
					continue
				}
				if j < h-1 {
					if dp[j+1][k] <= dp[j][k] && string(ss[j+1][k]) == "." {
						dp[j+1][k] = dp[j][k] + 1
					}
				}
				if k < w-1 {
					if dp[j][k+1] <= dp[j][k] && string(ss[j][k+1]) == "." {
						dp[j][k+1] = dp[j][k] + 1
					}
				}
			}
		}
	}

	ans := 0
	for _, r := range dp {
		for _, c := range r {
			if c > ans {
				ans = c
			}
		}
	}
	fmt.Println(ans)
}
