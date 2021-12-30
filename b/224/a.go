package main

import(
	"fmt"
	"reflect"
)

func main() {

	var(
		S string
	)

	fmt.Scanf("%s", &S)
	ss := []rune(S)
	l := len(ss)

	er := []rune("er")

	if reflect.DeepEqual(er, ss[l-2:]) {
		fmt.Println("er")
	} else {
		fmt.Println("ist")
	}
}
