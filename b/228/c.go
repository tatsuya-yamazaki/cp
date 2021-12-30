package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var n, k int
	var al []int

	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	points := make([]int, 1201, 1201)
	for i:=0; i<n; i++ {

		sum := 0
		for j:=0; j<3; j++ {
			sc.Scan()
			t := sc.Text()
			a, _ := strconv.Atoi(t)
			sum += a
		}

		al = append(al, sum)
		points[sum]++
	}

	pl := make([]int, 1201, 1201)
	pl[1200] = 1
	for i:=1199; i>=0; i-- {
		pl[i] += pl[i+1]
		if points[i] != 0 && i != 0 {
				pl[i-1] += points[i]
		}
	}

	for _, v := range al {
		if pl[v+300] <= k {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
