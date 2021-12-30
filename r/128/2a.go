package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	A := make([]int, N, N)
	for i:=0; i<N; i++{
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		A[i] = a
	}

	ans := make([]int, N, N)
	for i:=0; i<N-1; i++ {
		if A[i] > A[i+1] {
			ans[i] ^= 1
			ans[i+1] ^= 1
		}
	}
	for i, v := range ans {
		fmt.Printf("%d", v)
		if i != N-1 {
			fmt.Printf(" ")
		} else {
			fmt.Printf("\n")
		}
	}
}
