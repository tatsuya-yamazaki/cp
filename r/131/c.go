package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var n int

	fmt.Scan(&n)
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	a := make([]int, n, n)
	x := 0
	for i:=0; i<n; i++ {
		sc.Scan()
		t := sc.Text()
		ar, _ := strconv.Atoi(t)
		a[i] = ar
		x = x ^ ar
	}

	for i:=0; i<n; i++ {
		if x ^ a[i] == 0 {
			fmt.Println("Win")
			return
		}
	}

	if n % 2 == 0 {
		fmt.Println("Lose")
	} else {
		fmt.Println("Win")
	}
}
