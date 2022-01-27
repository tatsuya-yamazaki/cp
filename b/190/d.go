package main
import (
	"fmt"
	"math"
)
func main() {
	var n int
	fmt.Scan(&n)
	r := int(math.Sqrt(float64(2*n)))
	var d []int
	for i:=1; i<=r; i++ {
		q := 2*n / i
		if 2*n % i == 0 && (((i % 2 == 0) && (q % 2 != 0)) || ((i % 2 != 0) && (q % 2 == 0 ))) {
			d = append(d, i)
			d = append(d, q)
		}
	}
	fmt.Println(len(d))
}
