package main

import(
	"fmt"
)

func main(){
	var ll [][]int
	n := 5
	for i:=0; i<(1<<n); i++{
		var l []int
		for j:=0; j<n; j++{
			if (i & (1<<j)) > 0 {
				l = append(l, j)
			}
		}
		ll = append(ll, l)
	}
	fmt.Println(ll)
}
