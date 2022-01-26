package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)
func main() {
	var n int

	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	p := make([]int, n)
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		p[i] = ai
	}

	qpos := make([]int, n+1)
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		qpos[ai] = i
	}

	dp := make([]int, n)
	const INF = 99999999999999999
	for i, _ := range dp {
		dp[i] = INF
	}

	for i:=0; i<n; i++ {
		var pairs [][2]int
		for j:=1; j<=n; j++ {
			a := p[i] * j
			if a > n {
				break
			}
			pairs = append(pairs, [2]int{i, qpos[a]})
		}

		//decending order q pos
		sort.Slice(pairs, func(i, j int) bool { return pairs[i][1] > pairs[j][1] })

		for _, v := range pairs {
			index := lowerBound(dp, v[1])
			dp[index] = v[1]
		}
	}

	ans := 0
	for _, v := range dp {
		if v == INF {
			break
		}
		ans++
	}
	fmt.Println(ans)
}

func lowerBound(s []int, value int) int {
	l, r := 0, len(s)
	for l != r {
		m := (l + r) / 2
		if value > s[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

