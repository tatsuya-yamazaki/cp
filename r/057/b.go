package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

type Node struct {
	Won int
	Game int
}

func main() {
	var(
		N, K int
	)

	fmt.Scanf("%d %d", &N, &K)

	sc := bufio.NewScanner(os.Stdin)

	A := make([]int, N+1, N+1)

	for i:=1; i<N+1; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		A[i] = a + A[i-1]
	}

	var dp [][]Node
	for i:=0; i<N+1; i++ {
		dpr := make([]Node, N+1, N+1)
		dp = append(dp, dpr)
	}

	for i:=1; i<N+1; i++ {
		for j:=0; j<=i; j++{
			if j == 0 {
				dp[i][j].Game = dp[i-1][j].Game + A[i]-A[i-1]
				continue
			}

			if i != 1 && dp[i-1][j-1].Game == 0 {
				continue
			}

			newWon := getGoodMoodWonNum(dp[i-1][j-1].Won, dp[i-1][j-1].Game, A[i]-A[i-1])
			if newWon > A[i]-A[i-1] || newWon + dp[i-1][j-1].Won > K {
				continue
			}

			restWin := K - dp[i-1][j-1].Won - newWon
			restGame := A[N]-A[i]
			isAllWin := false
			if restWin > restGame {
				newWon = newWon + restWin - restGame
				isAllWin = true
			}

			dp[i][j].Game = dp[i-1][j-1].Game + A[i]-A[i-1]
			newTotalWon := newWon + dp[i-1][j-1].Won

			if dp[i-1][j].Game != 0 && dp[i-1][j].Won < newTotalWon && i != 1 && !isAllWin {
				dp[i][j].Won = dp[i-1][j].Won
			} else {
				dp[i][j].Won = newWon + dp[i-1][j-1].Won
			}
		}
	}

	for i:=N; i>=0; i-- {
		if dp[N][i].Game > 0 {
			fmt.Println(i)
			return
		}
	}
}

func getGoodMoodWonNum(totalWon, totalGame, newGame int) int {
	tw, tg, ng := float64(totalWon), float64(totalGame), float64(newGame)

	if totalGame == 0 {
		return 1
	}
	nw := math.Floor((tw*ng) / tg) + 1.0
	return int(nw)
}
