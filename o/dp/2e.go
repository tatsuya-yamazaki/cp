package main

import(
	"fmt"
	"math"
)

func scan2n() (a,b int) {
	fmt.Scanf("%d %d\n", &a, &b)
	return
}

func main(){
	var vl, wl []int
	N, W := scan2n()
	for i:=0; i<N; i++{
		w, v := scan2n()
		wl = append(wl, w)
		vl = append(vl, v)
	}

	valueSum := 0
	for _, v := range vl {
		valueSum = valueSum + v
	}

	var dp [][]int
	var dpRow []int
	for i:=0; i<valueSum+1; i++{
		dpRow = append(dpRow, math.MaxInt64)
	}
	for i:=0; i<N+1; i++{
		r := make([]int, valueSum+1, valueSum+1)
		copy(r, dpRow)
		dp = append(dp, r)
	}
	dp[0][0] = 0

	for i:=1; i<N+1; i++{
		for v:=0; v<=valueSum; v++{
			if v < vl[i-1] {
				dp[i][v] = dp[i-1][v]
			} else {
				if dp[i-1][v-vl[i-1]] == math.MaxInt64 {
					dp[i][v] = dp[i-1][v]
					continue
				}
				if dp[i-1][v] > dp[i-1][v-vl[i-1]] + wl[i-1] {
					dp[i][v] = dp[i-1][v-vl[i-1]] + wl[i-1]
				} else {
					dp[i][v] = dp[i-1][v]
				}
			}
		}
	}

	ans := 0
	for i:=valueSum; i>=0; i--{
		if dp[N][i] <= W {
			ans = i
			break
		}
	}
	fmt.Println(ans)
}
