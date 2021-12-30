package main
import(
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	ans := leastCommonMultiple(a, b)
	fmt.Println(ans)
}

func greatestCommonDivisor(a, b int) int {
	if b > a {
		a, b = b, a
	}
	return gcd(a,b)
}

func gcd(a, b int) int {
	if a % b == 0 {
		return b
	}
	return gcd(b, a%b)
}

func leastCommonMultiple(a, b int) int {
	return a*b / greatestCommonDivisor(a,b)
}
