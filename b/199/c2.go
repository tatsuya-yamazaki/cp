package main
import (
	"fmt"
	"os"
	"bufio"
)
func main() {
	var n, q int
	var s string
	fmt.Scan(&n, &s, &q)

	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	ss := make([][]rune, 3)
	ss[2] = []rune("0" + s[:n])
	ss[0] = []rune("0" + s[n:])
	on := 1

	for i:=0; i<q; i++ {
		var t, a, b int
		fmt.Fscan(r, &t, &a, &b)

		switch t {
		case 1:
			as, bs := on, on
			if a > n {
				as *= -1
				a -= n
			}
			if b > n {
				bs *= -1
				b -= n
			}
			a, b := a%(n+1), b%(n+1)
			ss[1+as][a], ss[1+bs][b] = ss[1+bs][b], ss[1+as][a]

		case 2:
			on *= -1
		}
	}
	fmt.Fprintln(w, string(ss[1+on][1:])+string(ss[1-on][1:]))
}
