package main

import(
	"fmt"
)

func main() {
	var(
		N int
		T [][]int
	)

	fmt.Scanf("%d", &N)

	for i:=0; i<N; i++{
		var x, y int
		var tr []int
		fmt.Scanf("%d %d", &x, &y)
		tr = append(tr, x)
		tr = append(tr, y)
		T = append(T, tr)
	}

	ans := 0
	for i:=0; i<N; i++{
		for j:=i+1; j<N; j++{
			for k:=j+1; k<N; k++{
				if (T[i][0]-T[k][0])*(T[j][1]-T[k][1]) != (T[j][0]-T[k][0])*(T[i][1]-T[k][1]) {
					ans++
				}
			}
		}
	}
	fmt.Println(ans)
}
