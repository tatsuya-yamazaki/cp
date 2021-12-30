package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	m := make(map[int]struct{})
	for i:=2; i<=100000; i++ {
		a := i
		for {
			a *= i
			if n < a {
				break
			}
			_, ok := m[a]
			if !ok {
				m[a] = struct{}{}
			}
		}
	}
	fmt.Println(n-len(m))
}
