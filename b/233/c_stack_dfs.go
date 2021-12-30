package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var(
		n, x int
		a [][]int
	)

	fmt.Scan(&n)
	fmt.Scan(&x)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i:=0; i<n; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		var ar []int
		for j:=0; j<k; j++ {
			sc.Scan()
			ai, _ := strconv.Atoi(sc.Text())
			ar = append(ar, ai)
		}
		a = append(a, ar)
	}

	var stack [][]int
	for i:=0; i<n; i++ {
	}

	fmt.Println(dfs(a, x, 0, 1))
}

func dfs(bags [][]int, x, n, productOfSequences int) int {
	if len(bags) == n {
		if productOfSequences == x {
			return 1
		} else {
			return 0
		}
	}
	ret := 0
	for _, v := range bags[n] {
		if productOfSequences > x / v {
			continue
		}
		ret += dfs(bags, x, n+1, productOfSequences * v)
	}
	return ret
}
