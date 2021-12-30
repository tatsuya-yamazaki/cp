package main
import(
	"fmt"
)
func main() {
	var(
		MW, N, K int
		W, V []int
	)

	fmt.Scanf("%d", &MW)
	fmt.Scanf("%d %d", &N, &K)
	for i:=0; i<N; i++ {
		var w,v int
		fmt.Scanf("%d %d", &w, &v)
		W = append(W, w)
		V = append(V, v)
	}

	dp := make([][][]int, N+1, N+1)
	for i:=0; i<N+1; i++{
		dpr := make([][]int, K+1, K+1)
		for j:=0; j<K+1; j++{
			dprr := make([]int, MW+1, MW+1)
			dpr[j] = dprr
		}
		dp[i] = dpr
	}

	for i:=1; i<N+1; i++{
		w := W[i-1]
		v := V[i-1]
		for j:=1; j<K+1; j++{
			for k:=0; k<MW+1; k++{
				if k < w {
					dp[i][j][k] = dp[i-1][j][k]
					continue
				}
				newValue := v + dp[i-1][j-1][k-w]
				if dp[i-1][j][k] < newValue {
					dp[i][j][k] = newValue
				} else {
					dp[i][j][k] = dp[i-1][j][k]
				}
			}
		}
	}

	fmt.Println(dp[N][K][MW])
}
