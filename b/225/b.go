package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	var(
		N int
		ans int
	)

	fmt.Scanf("%d", &N)

	nl := make([]int, N+1, N+1)

	sc := bufio.NewScanner(os.Stdin)
	for i:=1; i<N; i++{
		sc.Scan()
		text := sc.Text()
		strs := strings.Split(text, " ")
		a, _ := strconv.Atoi(strs[0])
		b, _ := strconv.Atoi(strs[1])

		nl[a]++
		nl[b]++
	}
	if err := sc.Err(); err != nil {
		panic(err)
	}

	for _, v := range nl {
		if v > ans {
			ans = v
		}
	}
	if ans == N-1 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
