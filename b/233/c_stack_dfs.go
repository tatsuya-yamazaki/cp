package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var(
		n, x int
		a [][]int
	)

	fmt.Scan(&n)
	fmt.Scan(&x)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i:=0; i<n; i++ {
		sc.Scan()
		k, _ := strconv.Atoi(sc.Text())
		var ar []int
		for j:=0; j<k; j++ {
			sc.Scan()
			ai, _ := strconv.Atoi(sc.Text())
			ar = append(ar, ai)
		}
		a = append(a, ar)
	}

	var stack [][]int
	ans := 0
	for _, ai := range a[0] {
		s := []int{1, ai}
		stack = append(stack, s)
	}
	for len(stack) != 0 {
		task := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if task[0] == n {
			if task[1] == x {
				ans++
			}
			continue
		}
		for _, ai := range a[task[0]] {
			if ai > x / task[1] {
				continue
			}
			s := []int{task[0]+1, task[1]*ai}
			stack = append(stack, s)
		}
	}

	fmt.Println(ans)
}
