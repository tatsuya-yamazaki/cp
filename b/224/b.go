package main

import(
	"fmt"
)

func main () {

	var(
		H, W int
		A [][]int
	)

	fmt.Scanf("%d %d", &H, &W)
	for i:=0; i<H; i++{
		var ar []int
		for j:=0; j<W; j++{
			var a int
			fmt.Scanf("%d", &a)
			ar = append(ar, a)
		}
		A = append(A, ar)
	}

	for i1:=0; i1<H; i1++{
		for i2:=i1; i2<H; i2++{
			for j1:=0; j1<W; j1++{
				for j2:=j1; j2<W; j2++{
					if i1 > i2 || j1 > j2 {
						continue
					}
					if (A[i1][j1] + A[i2][j2]) > (A[i2][j1] + A[i1][j2]) {
						fmt.Println("No")
						return
					}
				}
			}
		}
	}
	fmt.Println("Yes")
}
