package main

import(
	"fmt"
	"bufio"
	"os"
)

func main() {
	var(
		N int
	)

	ans := make(map[string]struct{})

	fmt.Scanf("%d", &N)
	sc := bufio.NewScanner(os.Stdin)
	for i:=0; i<N; i++ {
		sc.Scan()
		ans[sc.Text()] = struct{}{}
	}

	fmt.Println(len(ans))
}
