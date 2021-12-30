package main

import "fmt"
import "strconv"

func main() {
	var ( N, P int)
	var anl []string

	fmt.Scanf("%d %d", &N, &P)
	for i := 0 ; i<N; i++ {
		var s string
		fmt.Scanf("%s", &s)
		anl = append(anl, s)
}

	var ans int
	for _, v := range anl {
		i, _ := strconv.Atoi(v)
		if i < P {
			ans++
		}
	}

	fmt.Printf("%d", ans)

}
