package main

import(
	"fmt"
)
func main() {
	var n, m, x int
	fmt.Scan(&n, &m, &x)

	c := make([]int, n)
	a := make([][]int, n)
	for i:=0; i<n; i++ {
		fmt.Scan(&c[i])
		ai := make([]int, m)
		for j:=0; j<m; j++ {
			fmt.Scan(&ai[j])
		}
		a[i] = ai
	}

	ans := 1000000000

	for i:=0; i<1<<n; i++ {
		next := 0
		s := make([]int, m)
		for j:=0; j<n; j++ {
			if i & (1<<j) > 0 {
				next += c[j]
				for k:=0; k<m; k++ {
					s[k] += a[j][k]
				}
			}
		}
		isMaster := true
		for k:=0; k<m; k++ {
			if s[k] < x {
				isMaster = false
				break
			}
		}
		if isMaster && next < ans {
			ans = next
		}
	}

	if ans == 1000000000 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
