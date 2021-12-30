package main

import(
	"fmt"
	"os"
	"bufio"
)
func main() {
	var(
		n, k int
		s string
	)

	sc := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 600000)
	sc.Buffer(buffer, 600000)
	sc.Scan()
	t := sc.Text()
	a := make([]int, 600000)

	fmt.Println(ans)
}
