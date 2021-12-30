package main
import(
	"fmt"
	"sort"
	"strconv"
)
func main() {
	var n, k int
	fmt.Scan(&n)
	fmt.Scan(&k)
	for i:=0; i<k; i++ {
		n = f(n)
	}
	fmt.Println(n)
}

func f(n int) int {
	s := strconv.Itoa(n)
	r := []rune(s)
	sort.Slice(r, func(i,j int) bool { return r[i] < r[j] } )
	na, _ := strconv.Atoi(string(r))
	sort.Slice(r, func(i,j int) bool { return r[i] > r[j] } )
	nd, _ := strconv.Atoi(string(r))
	return nd - na
}
