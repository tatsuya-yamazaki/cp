package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

func main() {
	var n, k int
	var al, bl []int

	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i:=0; i<n; i++ {

		sum := 0
		for j:=0; j<3; j++ {
			sc.Scan()
			t := sc.Text()
			a, _ := strconv.Atoi(t)
			sum += a
		}

		al = append(al, sum)
		bl = append(bl, sum)
	}

	sort.Ints(bl)

	target := bl[n-k]

	for _, v := range al {
		if target <= v + 300 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}
