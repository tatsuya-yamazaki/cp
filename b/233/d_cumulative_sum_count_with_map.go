package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var(
		n, k int
		a []int
	)

	fmt.Scan(&n)
	fmt.Scan(&k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	a = append(a, 0)
	for i:=1; i<=n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, a[i-1] + ai)
	}

	ans := 0
	m := make(map[int]int)
	for r:=1; r<len(a); r++ {
		m[a[r-1]]++
		ans += m[a[r]-k]
	}
	fmt.Println(ans)
}
