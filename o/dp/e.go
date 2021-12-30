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

	valueSum := 0
	for _, v := range vl {
		valueSum = valueSum + v
	}

	var dp [][]int
	for i:=0; i<N+1; i++{
		dpRow := make([]int, valueSum+1, valueSum+1)
		dp = append(dp, dpRow)
	}

	for i:=1; i<N+1; i++{
		for v:=0; v<=valueSum; v++{
			if vl[i-1] > v {
				dp[i][v] = dp[i-1][v]

			} else if vl[i-1] == v {
				if dp[i-1][v] == 0 || dp[i-1][v] > wl[i-1] {
					dp[i][v] = wl[i-1]
				} else {
					dp[i][v] = dp[i-1][v]
				}

			} else {
				restWeight := dp[i-1][v-vl[i-1]]
				newWeight := restWeight + wl[i-1]
				prevWeight := dp[i-1][v]

				if prevWeight == 0 && restWeight != 0 {
					dp[i][v] = newWeight
					continue
				}

				if prevWeight > newWeight && restWeight != 0 {
					dp[i][v] = newWeight
				} else {
					dp[i][v] = prevWeight
				}
			}
		}
	}

	ans := 0
	for v:=valueSum; v>=0; v--{
		if dp[N][v] <= W && dp[N][v] > 0 {
			ans = v
			break
		}
	}
	fmt.Println(ans)
}
