package main
import (
	"fmt"
)
func main() {
	var n int
	fmt.Scan(&n)
	var t, x, y []int
	for i:=0; i<n; i++ {
		var tr, xr, yr int
		fmt.Scan(&tr, &xr, &yr)
		t = append(t, tr)
		x = append(x, xr)
		y = append(y, yr)
	}

	nx, ny, nt := 0, 0, 0
	for i:=0; i<n; i++ {
		d := abs(x[i]-nx) + abs(y[i]-ny)
		rest := d - (t[i] - nt)
		if rest > 0 || rest % 2 != 0 {
			fmt.Println("No")
			return
		}
		nx, ny, nt = x[i], y[i], t[i]
	}
	fmt.Println("Yes")
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
