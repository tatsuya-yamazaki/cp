package main
import "fmt"
func main() {
	var n, y int
	fmt.Scan(&n, &y)
	a, b, c := -1, -1, -1

	for i:=0; i<=n; i++ {
		for j:=0; j<=n-i; j++ {
			if (i*10000 + j*5000 + (n-i-j)*1000) == y {
				a, b, c = i, j, n-i-j
			}
		}
		if a != -1 || b != -1 || c != -1 {
			break
		}
	}

	fmt.Println(a,b,c)
}
