package main

import(
	"fmt"
	"reflect"
)

func main() {
	var s string
	var a []string
	b := []string{".#","#."}
	c := []string{"#.",".#"}

	fmt.Scan(&s)
	a = append(a, s)
	fmt.Scan(&s)
	a = append(a, s)

	if reflect.DeepEqual(a,b) {
		fmt.Println("No")
		return
	}
	if reflect.DeepEqual(a,c) {
		fmt.Println("No")
		return
	}


	fmt.Println("Yes")
}
