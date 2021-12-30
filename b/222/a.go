package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s", &str)

	nl := []rune(str)

	length := len(nl)
	rest := 4 - length

	var ans []rune
	for i := 0; i<rest; i++ {
		ans = append(ans, '0')
	}
	for _, v := range nl {
		ans = append(ans, v)
	}

	fmt.Printf("%s", string(ans))
}
