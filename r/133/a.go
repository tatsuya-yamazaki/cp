package main
import (
	"fmt"
	"os"
	"bufio"
	"strconv"
)
func main() {
	var n int

	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	max := 0
	var a []int
	change := true
	for i:=0; i<n; i++ {
		sc.Scan()
		ai, _ := strconv.Atoi(sc.Text())
		a = append(a, ai)
		if ai < max {
			change = false
		}
		if change && max < ai {
			max = ai
		}
	}

	for i, ai := range a {
		if max != ai {
			fmt.Print(ai)
			if i != n-1 {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println()
}
