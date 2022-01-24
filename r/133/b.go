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

	var p []int
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		p = append(p, ai)
	}

	var q []int
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		q = append(q, ai)
	}

	fmt.Println()
}
