package main
import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	var e []float64
	e = append(e, 0)
	for i:=1; i<=n; i++ {
		sc.Scan()
		pi, _ := strconv.Atoi(sc.Text())
		ei := (1+float64(pi))/2
		ei = e[i-1] + ei
		e = append(e, ei)
	}
	ans := 0.0
	for i:=k; i<=n; i++ {
		esum := e[i]-e[i-k]
		if ans < esum {
			ans = esum
		}
	}
	fmt.Println(ans)
}
