package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var a []int
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
	}
	sort.Ints(a)
	for i:=0; i<q; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		fmt.Println(search(a, x))
	}
}

func search(a []int, x int) int {
	l, r := -1, len(a)
	m := (l+r) / 2
	for r - l > 1 {
		if a[m] >= x {
			r = m
		} else {
			l = m
		}
		m = (l + r) / 2
	}
	return len(a) - r
}
