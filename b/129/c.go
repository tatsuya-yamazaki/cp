package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type Dp struct {
	Pattern int
	IsBroken bool
}

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	sc := bufio.NewScanner(os.Stdin)

	dp := make([]Dp, N+1, N+1)
	for i:=0; i<M; i++{
		sc.Scan()
		t := sc.Text()
		a, _ := strconv.Atoi(t)
		dp[a].IsBroken = true
	}
	if err := sc.Err(); err != nil {
		fmt.Println("scan failed")
		return
	}

	dp[0].Pattern = 1
	if !dp[1].IsBroken {
		dp[1].Pattern = 1
	}

	for i:=2; i<N+1; i++ {
		if dp[i].IsBroken {
			continue
		}

		if !dp[i-1].IsBroken {
			dp[i].Pattern += dp[i-1].Pattern
		}

		if !dp[i-2].IsBroken {
			dp[i].Pattern += dp[i-2].Pattern
		}

		dp[i].Pattern %= 1000000007
	}

	fmt.Println(dp[N].Pattern)
}
