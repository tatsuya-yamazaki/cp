package main

import(
	"fmt"
)

func main() {
	var N, D int
	fmt.Scanf("%d %d", &N, &D)

	primeFactor := [3]int{2,3,5}
	pfNum := make([]int, 3, 3)
	d := D
	for i, pf := range primeFactor {
		for {
			remainder := d % pf
			if remainder == 0 {
				pfNum[i]++
				d = d / pf
			} else {
				break
			}
		}
	}
	if d != 1 {
		fmt.Println(0)
		return
	}

	var dp [][][][]float64
	for i:=0; i<=N; i++ {
		var dpr [][][]float64
		for j:=0; j<=pfNum[0]; j++{
			var dprr [][]float64
			for k:=0; k<=pfNum[1]; k++{
				dprrr := make([]float64, pfNum[2]+1, pfNum[2]+1)
				dprr = append(dprr, dprrr)
			}
			dpr = append(dpr, dprr)
		}
		dp = append(dp, dpr)
	}

	dj := [6]int{0,1,0,2,0,1}
	dk := [6]int{0,0,1,0,0,1}
	dl := [6]int{0,0,0,0,1,0}

	dp[0][0][0][0] = 1
	for i:=1; i<=N; i++ {
		for j:=0; j<=pfNum[0]; j++{
			for k:=0; k<=pfNum[1]; k++{
				for l:=0; l<=pfNum[2]; l++{
					for m:=0; m<6; m++ {
						prevRate := dp[i-1][j][k][l]
						if prevRate == 0 {
							continue
						}
						nj := min(pfNum[0], j+dj[m])
						nk := min(pfNum[1], k+dk[m])
						nl := min(pfNum[2], l+dl[m])
						dp[i][nj][nk][nl] += prevRate / 6
					}
				}
			}
		}
	}

	fmt.Println(dp[N][pfNum[0]][pfNum[1]][pfNum[2]])
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
