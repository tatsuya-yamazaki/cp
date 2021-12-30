package main
import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)
func main() {
	var n, k int
	var p []int
	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i:=0; i<n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())

		p = append(p, a)
	}

	min := 0
	max := n * 1000000000000 / k
	mid := max / 2

	for {
		sum := 0
		for _, a := range p {
			if a < mid {
				sum += a
				continue
			}
			sum += mid
		}

		if mid*k > sum {
			max = mid - 1
		} else {
			min = mid
		}

		if min == max {
			break
		}
		mid = (max - min) / 2 + min + 1
	}

	fmt.Println(min)
}
