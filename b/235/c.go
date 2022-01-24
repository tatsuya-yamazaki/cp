package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var(
		n, q int
	)

	fmt.Scan(&n)
	fmt.Scan(&q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	m := make(map[int][]int)
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		m[ai] = append(m[ai], i+1)
	}

	for i:=0; i<q; i++ {
		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		_, ok := m[x]
		if ok {
			if len(m[x]) < k {
				fmt.Println(-1)
			} else {
				fmt.Println(m[x][k-1])
			}
		} else {
			fmt.Println(-1)
		}
	}
}
