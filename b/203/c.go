package main

import (
	"fmt"
	"os"
	"bufio"
	"strconv"
	"sort"
)

type Friend struct {
	A int
	B int
}
func main() {
	var f []Friend

	var n, k int

	fmt.Scan(&n, &k)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for i:=0; i<n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		f = append(f, Friend{a,b})
	}

	sort.Slice(f, func(i,j int) bool { return f[i].A < f[j].A })

	pos := 0
	for _, v := range f {
		cost := v.A - pos
		if cost > k {
			break
		}
		pos = v.A
		k = k + v.B - cost
	}
	fmt.Println(pos+k)
}
