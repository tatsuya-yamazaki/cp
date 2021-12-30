package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int

	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)

	maxL := -1
	minR := 1000000001
	for i:=0; i<n; i++ {
		sc.Scan()
		t := sc.Text()

		ts := strings.Split(t, " ")
		l, _ := strconv.Atoi(ts[0])
		r, _ := strconv.Atoi(ts[1])

		if maxL < l {
			maxL = l
		}
		if minR > r {
			minR = r
		}

		if maxL <= minR {
			fmt.Println(0)
		} else {
			x := (maxL + minR) / 2
			fmt.Println(maxL - x)
		}
	}
}
