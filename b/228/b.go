package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var n, x, ans int
	var al []int

	fmt.Scan(&n, &x)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	al = append(al, 0)
	for i:=0; i<n; i++ {
		sc.Scan()
		t := sc.Text()
		a, _ := strconv.Atoi(t)

		al = append(al, a)
	}

	ans = 1
	next := al[x]
	al[x] = 0
	for {
		if al[next] == 0 {
			break
		}
		ans++
		nnext := al[next]
		al[next] = 0
		next = nnext
	}

	fmt.Println(ans)
}
