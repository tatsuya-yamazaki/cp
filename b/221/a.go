package main

import(
	"fmt"
	"math"
)

func main() {
	var(
		a, b float64
	)
	fmt.Scanf("%g %g", &a, &b)

	ans := int(math.Pow(32, a-b))
	//fmt.Printf("%g\n", math.Pow(32, a-b))

	fmt.Printf("%d\n", ans)
}
