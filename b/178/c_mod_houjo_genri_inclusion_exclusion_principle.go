package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	d := 1000000007

	all := pattern(n, 10)

	//0か9を含む整列 = 全数列 - 0か9を含まない整列
	//0か9を含まない数列 = 0を含まない整列 + 9を含まない整列 - 0も9も含まない整列
	no0 := pattern(n, 9)
	no9 := pattern(n, 9)
	no09 := pattern(n, 8)
	//引き算の場合は負になりうるため、割る数を足して無理やり正にしておく
	no0or9 := (((no0 + no9) % d) + d - no09) % d
	ans := (all + d - no0or9) % d
	fmt.Println(ans)
}
func pattern(n, m int) int {
	d := 1000000007
	ret := 1
	for i:=0; i<n; i++ {
		ret *= m
		ret %= d
	}
	return ret
}
