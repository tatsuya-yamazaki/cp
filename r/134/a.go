package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var n, l, w int

	fmt.Scan(&n)
	fmt.Scan(&l)
	fmt.Scan(&w)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	ans := 0
	prev := 0
	for i:=0; i<n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		if prev > l {
			continue
		}

		d := a - prev
		if d > 0 {
			q := d / w
			r := d % w
			ans += q
			if r != 0 {
				ans++
			}
		}

		prev = a + w
	}

	d := l - prev
	if d > 0 {
		q := d / w
		r := d % w
		ans += q
		if r != 0 {
			ans++
		}
	}

	fmt.Println(ans)
}
