package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var N int
	fmt.Scanf("%d", &N)

	A := make([]int, N, N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i:=0; i<N; i++{
		sc.Scan()
		t := sc.Text()
		a, _ := strconv.Atoi(t)
		A[i] = a
	}

	dp := make([]int, N, N)

	dp[1] = abs(A[1] - A[0])
	if N == 2 {
		fmt.Println(dp[1])
		return
	}

	for i:=2; i<N; i++ {
		one := dp[i-1] + abs( A[i] - A[i-1] )
		two := dp[i-2] + abs( A[i] - A[i-2] )
		if one < two {
			dp[i] = one
		} else {
			dp[i] = two
		}
	}

	fmt.Println(dp[N-1])
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
