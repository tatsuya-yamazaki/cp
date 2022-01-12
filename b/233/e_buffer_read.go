package main

import(
	"fmt"
	"os"
	"bufio"
)
func main() {
	sc := bufio.NewScanner(os.Stdin)
	buffer := make([]byte, 600000)
	sc.Buffer(buffer, 600000)
	sc.Scan()
	t := sc.Text()
	s := make([]int, len(t)+1)
	s[0] = 0
	for i, v := range t {
		s[i+1] = s[i] + int(v-'0')
	}

	rest := 0
	var stack []int
	for i:=len(s)-1; i>0; i-- {
		sum := s[i] + rest
		stack = append(stack, sum % 10)
		rest = sum / 10
	}
	for rest != 0 {
		stack = append(stack, rest % 10)
		rest /= 10
	}
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	for i:=len(stack); i>0; i-- {
		fmt.Fprint(wr, stack[i-1])
	}
	fmt.Fprintln(wr)
}
