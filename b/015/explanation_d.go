package main
import(
	"fmt"
)
func main() {
	var(
		MW = 6
		N = 3
		K = 3
		W = [3]int{2,2,2}
		V = [3]int{7,1,10}
	)

	fmt.Println(MW)
	fmt.Println(N, K)
	for i:=0; i<N; i++{
		fmt.Println(W[i], V[i])
	}

	var dp [4][4][7]int

	for i:=1; i<N+1; i++{
		w := W[i-1]
		v := V[i-1]
		for j:=1; j<K+1; j++{
			for k:=0; k<MW+1; k++{
				if k < w {
					dp[j][i][k] = dp[j][i-1][k]
					continue
				}
				newValue := v + dp[j-1][i-1][k-w]
				if dp[j][i-1][k] < newValue {
					dp[j][i][k] = newValue
				} else {
					dp[j][i][k] = dp[j][i-1][k]
				}
			}
		}
	}

	for j:=0; j<K+1; j++{
		fmt.Println(dp[j])
	}
	fmt.Println("answer:", dp[K][N][MW])
	fmt.Println("dpテーブル各行ごとに最大使用枚数が増える")
	fmt.Println("前の行はその行よりも1枚最大枚数が少ない場合の最大値なので、")
	fmt.Println("その値に新たに取得した値を足せば、自動的に前の行より最大枚数が1枚多い行ができる。")
	fmt.Println("初期状態のデータから段階的に処理していく考え方がポイント。")
	fmt.Println("これは動的計画法に共通する考え方。")
	fmt.Println("最大使用枚数が0の行をまず作り、1の行は0の行を、2の行は1の行を利用することで実現している。")
	fmt.Println("他のコード例だと、dpテーブルの添字の順番が違うが、考え方や処理は同じなので結果も一致する。")
}
