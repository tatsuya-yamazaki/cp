package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var n int
	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	a := make([]int, n, n)
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a[i] = ai
	}

	ans := 0
	for l:=0; l<n; l++ {
		x := a[l]
		for r:=l; r<n; r++ {
			if a[r] < x {
				x = a[r]
			}
			if ans < x*(r-l+1) {
				ans = x*(r-l+1)
			}
		}
	}
	fmt.Println(ans)
}
