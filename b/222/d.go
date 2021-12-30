package main

import(
	"fmt"
	"math/big"
	"crypto/rand"
)

func main () {
	var(
		N int
		A,B []int
	)

	fmt.Scanf("%d", &N)
	for i:=0; i<N; i++{
		var a int
		fmt.Scanf("%d", &a)
		A = append(A, a)
	}

	for i:=0; i<N; i++{
		var b int
		fmt.Scanf("%d", &b)
		B = append(B, b)
	}

	ans, _ := fastMod(N,A,B)
	fmt.Println(ans)
}

func genDp(N int, B []int) (dp [][]int) {
	for i:=0; i<N; i++{
		max := B[len(B)-1] + 1
		dp = append(dp, make([]int, max, max))
	}
	return
}

func fastMod(N int, A,B []int) (int, [][]int) {
	mod := 998244353

	dp := genDp(N,B)
	for i:=0; i<B[len(B)-1]+1; i++{
		if i >= A[0] && i <= B[0] && i == 0 {
			dp[0][i] = 1
		} else if i >= A[0] && i <= B[0] {
			dp[0][i] = dp[0][i-1] + 1
		} else if i > B[0] {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i:=1; i<N; i++ {
		for j:=A[i]; j<=B[len(B)-1]; j++{
			if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > B[i] {
				dp[i][j] = dp[i][j-1]
			} else if A[i-1] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][A[i-1]-1]
			}
			dp[i][j] %= mod
		}
	}

	ans := dp[N-1][B[len(B)-1]]
	return ans, dp
}

func mAB() (A,B []int) {
	for i:=0; i<3000; i++{
		a,b := 0,3000
		A = append(A, a)
		B = append(B, b)
	}
	return
}

func randN(n int) int {
	r, _ := rand.Int(rand.Reader, big.NewInt(int64(n)))
	ret := int(r.Int64()) + 1
	return ret
}

func randAB(N int) ( A,B []int ) {
	//maxNum := 3001
	maxNum := 30
	preva, prevb := 0, 0
	for i:=0; i<N; i++{
		ra, _ := rand.Int(rand.Reader, big.NewInt(int64(maxNum-preva)))
		a := int(ra.Int64()) + preva
		minb := prevb
		if a > prevb {
			minb = a
		}
		rb, _ := rand.Int(rand.Reader, big.NewInt(int64(maxNum-minb)))
		b := int(rb.Int64()) + minb
		A = append(A, a)
		B = append(B, b)
		preva = a
		prevb = b
	}
	return
}

func genDpBig(N int, B []int) (dp [][]*big.Int) {
	for i:=0; i<N; i++{
		var dpr []*big.Int
		for j:=0; j<B[len(B)-1]+1; j++{
			dpr = append(dpr, big.NewInt(0))
		}
		dp = append(dp, dpr)
	}
	return
}

func fastBig(N int, A,B []int) (int, [][]*big.Int) {
	dp := genDpBig(N,B)

	for i:=0; i<B[len(B)-1]+1; i++{
		if i >= A[0] && i <= B[0] && i == 0 {
			dp[0][i] = big.NewInt(1)
		} else if i >= A[0] && i <= B[0] {
			dp[0][i].Add(dp[0][i-1], big.NewInt(1))
		} else if i > B[0] {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i:=1; i<N; i++ {
		for j:=A[i]; j<=B[len(B)-1]; j++{
			if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > B[i] {
				dp[i][j] = dp[i][j-1]
			} else if A[i-1] == 0 {
				dp[i][j].Add(dp[i][j-1], dp[i-1][j])
			} else {
				tmp := big.NewInt(0)
				tmp.Sub(dp[i-1][j], dp[i-1][A[i-1]-1])
				dp[i][j].Add(dp[i][j-1], tmp)
			}
		}
	}

	ans := big.NewInt(0)
	ans.Mod(dp[N-1][B[len(B)-1]], big.NewInt(998244353))
	return int(ans.Int64()), dp
}

func fast(N int, A,B []int) (int, [][]int) {
	dp := genDp(N,B)
	for i:=0; i<B[len(B)-1]+1; i++{
		if i >= A[0] && i <= B[0] && i == 0 {
			dp[0][i] = 1
		} else if i >= A[0] && i <= B[0] {
			dp[0][i] = dp[0][i-1] + 1
		} else if i > B[0] {
			dp[0][i] = dp[0][i-1]
		}
	}
	for i:=1; i<N; i++ {
		for j:=A[i]; j<=B[len(B)-1]; j++{
			if j == 0 {
				dp[i][j] = dp[i-1][j]
			} else if j > B[i] {
				dp[i][j] = dp[i][j-1]
			} else if A[i-1] == 0 {
				dp[i][j] = dp[i][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i][j-1] + dp[i-1][j] - dp[i-1][A[i-1]-1]
			}
		}
	}
	ans := dp[N-1][B[len(B)-1]]
	ans = ans % 998244353
	return ans, dp
}

func slow(N int, A,B []int) (int, [][]int) {
	dp := genDp(N,B)

	for i:=0; i<B[N-1]+1; i++{
		if A[0] <= i && i <= B[0] {
			dp[0][i] = 1
		}
	}

	for i:=1; i<N; i++{
		for j:=A[i]; j<=B[i]; j++{
			for k:=A[i-1]; k<=j && k<=B[i-1]; k++{
				dp[i][j] = dp[i][j] + dp[i-1][k]
			}
		}
	}

	for i:=0; i<N; i++{
		for j, v := range dp[i]{
			if j == 0 {
				continue
			}
			dp[i][j] = dp[i][j-1] + v
		}
	}

	ans := dp[N-1][B[len(B)-1]]
	ans = ans % 998244353
	return ans, dp
}
