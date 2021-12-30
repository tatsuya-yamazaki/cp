package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var n int
	var a []int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	m := 1000000007
	az, _ := strconv.Atoi(sc.Text())
	a = append(a, az % m)
	for i:=1; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		//足し算において、毎回剰余をとっても、最終的な剰余は一致する
		a = append(a, a[i-1]+ai % m)
	}

	//掛け算において、各項の剰余を毎回とっても、最終的な積の剰余は一致する
	sum := (a[0] % m) * ((a[len(a)-1]-a[0]) % m)
	for i:=1; i<n; i++ {
		sum %= m
		//ただし、引き算の場合は負になりうるため、割る数を足して無理やり正にしておく
		sum += ((a[i]+m - a[i-1]) % m) * ((a[len(a)-1]-a[i]) % m)
	}

	fmt.Println(sum%m)
}
