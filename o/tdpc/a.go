package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"log"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	P := make([]int, N+1, N+1)

	for i:=1; i<N+1; i++ {
		sc.Scan()
		t := sc.Text()
		p, _ := strconv.Atoi(t)
		P[i] = p
	}
	if err := sc.Err(); err != nil {
		log.Fatal(err)
	}

	var dp [][]int
	for i:=0; i<N+1; i++ {
		dpr := make([]int, 10001, 10001)
		dp = append(dp, dpr)
	}
	dp[0][0]++

	count := 1
	for i:=1; i<N+1; i++ {
		for j:=0; j<10001; j++ {
			p := P[i]

			if dp[i-1][j] != 0 {
				dp[i][j] = dp[i-1][j]
				if dp[i-1][j+p] == 0 {
					dp[i][j+p]++
					count++
				}
			}
		}
	}
	fmt.Println(count)
}
