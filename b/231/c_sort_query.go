package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"sort"
)

type node struct {
	v, i int
}

func main() {
	var n, q int
	fmt.Scan(&n, &q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	var ax []node
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		ax = append(ax, node{ai, 1000000})
	}
	for i:=0; i<q; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		ax = append(ax, node{x, i})
	}

	sort.Slice(ax, func(i, j int) bool { return ax[i].i < ax[j].i })
	sort.SliceStable(ax, func(i, j int) bool { return ax[i].v < ax[j].v })
	ans := make([]int, q, q)
	count := 0
	for i:=n+q-1; i>=0; i-- {
		if ax[i].i != 1000000 {
			ans[ax[i].i] = count
		} else {
			count++
		}
	}

	for _, v := range ans {
		fmt.Println(v)
	}
}
