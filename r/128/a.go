package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"log"
	"math/big"
)

type DpCell struct {
	Value *big.Rat
	Route string
}

func main() {
	var(
		A []*big.Rat
		dp []DpCell
	)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Buffer(make([]byte, 1024), int(1e9+1))
	scanner.Scan()
	nStr := scanner.Text()
	N, _ := strconv.Atoi(nStr)
	var strs []string
	for i:=0; i<N; i++ {
		scanner.Scan()
		t := scanner.Text()
		strs = append(strs, t)
	}

	for _, s := range strs {
		a, _ := strconv.Atoi(s)
		A = append(A, big.NewRat(int64(a),1))
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return
	}

	dp = append(dp, DpCell{Value:big.NewRat(1,1), Route:""})
	dp = append(dp, DpCell{Value:big.NewRat(0,1), Route:""})

	for i:=0; i<N; i++{
		tmp1 := big.NewRat(0,1)
		tmp2 := big.NewRat(0,1)
		newGold := tmp1.Quo(dp[1].Value, A[i])
		newSilver := tmp2.Mul(dp[0].Value, A[i])
		oldGoldRoute := dp[0].Route
		oldSilverRoute := dp[1].Route

		if newGold.Cmp(dp[0].Value) > 0 {
			dp[0].Value = newGold
			dp[0].Route = oldSilverRoute + "1"
		} else {
			dp[0].Route = oldGoldRoute + "0"
		}
		if newSilver.Cmp(dp[1].Value) > 0 {
			dp[1].Value = newSilver
			dp[1].Route = oldGoldRoute + "1"
		} else {
			dp[1].Route = oldSilverRoute + "0"
		}
	}

	for i, c := range dp[0].Route{
		fmt.Printf("%s", string(c))
		if i == N - 1 {
			fmt.Printf("\n")
			break
		}
		fmt.Printf(" ")
	}
}
