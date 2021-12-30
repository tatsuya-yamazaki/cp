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
	ss[2] = []rune(s[:n])
	ss[0] = []rune(s[n:])
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
			posa := a%n
			if posa == 0 {
				posa += n
			}
			posa -= 1

			if b > n {
				bs *= -1
				b -= n
			}
			posb := b%n
			if posb == 0 {
				posb += n
			}
			posb -= 1

			ss[1+as][posa], ss[1+bs][posb] = ss[1+bs][posb], ss[1+as][posa]

		case 2:
			on *= -1
		}
	}
	fmt.Fprintln(w, string(ss[1+on])+string(ss[1-on]))
}
