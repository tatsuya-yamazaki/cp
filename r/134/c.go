package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var n, l, w int

	fmt.Scan(&n)
	fmt.Scan(&l)
	fmt.Scan(&w)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
	}
}
