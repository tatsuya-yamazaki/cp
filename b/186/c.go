package main
import (
	"fmt"
	"sort"
	"strconv"
)
func main() {
	var n int
	fmt.Scan(&n)
	ans := 0
	for i:=1; i<=n; i++ {
		if !f(strconv.Itoa(i)) {
			continue
		}
		if !f(e(i)) {
			continue
		}
		ans++
	}
	fmt.Println(ans)
}

func e(n int) string {
	var r []rune
	for n != 0 {
		r = append(r, rune(n%8+'0'))
		n /= 8
	}
	sort.Slice(r, func(i, j int) bool { return true } )
	return string(r)
}

func f(n string) bool {
	for _, c := range n {
		if c == '7' {
			return false
		}
	}
	return true
}
