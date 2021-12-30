package main
import(
	"fmt"
	"math"
)

func main() {
	var n, ans int
	fmt.Scan(&n)

	for a:=1; a<=int(math.Cbrt(float64(n))); a++ {
		for b:=a; b<=int(math.Sqrt(float64(n/a))); b++ {
			ans += int(math.Floor(float64((n/(a*b))-b+1)))
		}
	}

	fmt.Println(ans)
}
