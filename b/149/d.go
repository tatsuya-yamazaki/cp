package main
import(
	"fmt"
)
func main() {
	var n, k, r, s, p, ans int
	var t string
	fmt.Scan(&n)
	fmt.Scan(&k)
	fmt.Scan(&r)
	fmt.Scan(&s)
	fmt.Scan(&p)
	fmt.Scan(&t)
	tr := []rune(t)

	h := make([]rune, 0)

	for i:=0; i<n; i++ {
		switch tr[i] {
		case 'r':
			if i >= k {
				if 'p' == h[i-k] {
					h = append(h, 'x')
					continue
				}
			}
			h = append(h, 'p')
			ans += p
		case 's':
			if i >= k {
				if 'r' == h[i-k] {
					h = append(h, 'x')
					continue
				}
			}
			h = append(h, 'r')
			ans += r
		case 'p':
			if i >= k {
				if 's' == h[i-k] {
					h = append(h, 'x')
					continue
				}
			}
			h = append(h, 's')
			ans += s
		}
	}
	fmt.Println(ans)
}
