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
	mid := (len(a)-1) / 2
	left, right := 0, len(a)-1
	for left != right {
		if a[mid] >= x {
			right = mid
		} else {
			left = mid + 1
		}
		mid = (left + right) / 2
	}
	ret := len(a) - left
	if left == len(a) -1 && x > a[left] {
		ret = 0
	} else if left == 0 && x < a[left] {
		ret = len(a)
	}
	return ret
}
