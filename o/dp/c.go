package main

import(
	"fmt"
	tls "github.com/tatsuya-yamazaki/cptools"
)

func main() {
	var A,B,C []int
	var ABC [][]int
	N := tls.ScanInt()
	for i:=0; i<N; i++ {
		a, b, c := tls.Scan3Ints()
		A = append(A, a)
		B = append(B, b)
		C = append(C, c)
	}
	ABC = append(ABC, A)
	ABC = append(ABC, B)
	ABC = append(ABC, C)
	var dp [][]int
	for i:=0; i<N+1; i++ {
		dp = append(dp, make([]int, 3, 3))
	}

	for i:=1; i<N+1; i++ {
		for j:=0; j<3; j++{
			for k:=0; k<3; k++{
				if j == k {
					continue
				}
				point := dp[i-1][j] + ABC[k][i-1]
				if dp[i][k] < point {
					dp[i][k] = point
				}
			}
		}
	}

	max := 0
	for i:=0; i<3; i++ {
		if max < dp[N][i] {
			max = dp[N][i]
		}
	}
	fmt.Println(max)
}
