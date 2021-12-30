package main

import(
	"fmt"
	"math"
	"math/big"
)

func main() {
	var(
		N int
		T [][]int64
		R [][]*big.Rat
	)

	fmt.Scanf("%d", &N)
	all := (N * (N-1) * (N-2)) / 6

	for i:=0; i<N; i++{
		var x, y int64
		var tr []int64
		fmt.Scanf("%d %d", &x, &y)
		tr = append(tr, x)
		tr = append(tr, y)
		T = append(T, tr)
	}

	for i:=0; i<N; i++{
		rr := make([]*big.Rat, N, N)
		for j:=0; j<N; j++{
			if i == j {
				continue
			}
			rate := big.NewRat(0,1)
			ydif := T[j][1] - T[i][1]
			xdif := T[j][0] - T[i][0]
			if T[j][0] == T[i][0] {
				rate = big.NewRat(math.MaxInt64,1)
			} else if T[j][1] == T[i][1] {
				rate = big.NewRat(0,1)
			} else {
				ydifRat := big.NewRat(ydif, 1)
				xdifRat := big.NewRat(xdif, 1)
				rate.Quo(ydifRat, xdifRat)
			}
			rr[j] = rate
		}
		R = append(R, rr)
	}

	reject := 0
	for i:=0; i<N; i++{
		for j:=i+1; j<N; j++{
			for k:=j+1; k<N; k++{
				if R[i][j].Cmp(R[i][k]) == 0 {
					reject += 1
				}
			}
		}
	}
	fmt.Println(all - reject)
}
