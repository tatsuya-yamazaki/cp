package main
import(
	"fmt"
)

func main() {
	var N int
	fmt.Scanf("%d", &N)

	dp := make([]int, N+1, N+1)
	for i:=1; i<N+1; i++{
		sixExp := getSixExp(i)
		nineExp := getNineExp(i)

		switch i {
		case sixExp:
			dp[i] = 1
			continue
		case nineExp:
			dp[i] = 1
			continue
		}

		timesSix := 1 + dp[i-sixExp]
		timesNine := 1 + dp[i-nineExp]

		if timesSix > timesNine {
			dp[i] = timesNine
		} else {
			dp[i] = timesSix
		}
	}

	fmt.Println(dp[N])
}

func getSixExp(max int) int {
	return getMaxExp(6,max)
}

func getNineExp(max int) int {
	return getMaxExp(9,max)
}

func getMaxExp(n, max int) int {
	if max == 0 || n == 0 {
		return 0
	}

	ret := 1
	for {
		if ret > max {
			ret /= n
			return ret
		} else if ret == max {
			return ret
		}

		ret *= n
	}
}
