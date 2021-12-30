package main

import(
	"fmt"
	"math"
)

func main() {
	var(
		N float64
	)
	fmt.Scanf("%f", &N)

	fmt.Println(math.Round(N))
}
