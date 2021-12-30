package main

import(
	"fmt"
	"bufio"
	"os"
	"sort"
)

func main() {
	var(
		N int
		A []string
	)
	fmt.Scanf("%d", &N)

	sc := bufio.NewScanner(os.Stdin)
	for i:=0; i<N; i++ {
		sc.Scan()
		at := sc.Text()
		A = append(A, at)
	}

	sort.Strings(A)

	prev := A[0]
	ans := 1
	for _, s := range A {
		if prev != s {
			ans++
		}
		prev = s
	}
	fmt.Println(ans)
}
