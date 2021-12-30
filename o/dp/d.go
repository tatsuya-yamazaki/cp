package main

import(
	"fmt"
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

	var dp [][]int
	for i:=0; i<N+1; i++{
		dp = append(dp, make([]int, W+1, W+1))
	}

	for i:=1; i<N+1; i++{
		for w:=0; w<W+1; w++{
			if w < wl[i-1] {
				dp[i][w] = dp[i-1][w]
			} else {
				newValue := dp[i-1][w-wl[i-1]] + vl[i-1]
				if dp[i-1][w] < newValue {
					dp[i][w] = newValue
				} else {
					dp[i][w] = dp[i-1][w]
				}
			}
		}
	}

	fmt.Println(dp[N][W])
}
