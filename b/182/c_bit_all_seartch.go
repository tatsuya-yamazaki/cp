package main
import (
	"fmt"
)
func main() {
	var s string
	fmt.Scan(&s)
	r := []rune(s)

	ans := -1

	for i:=0; i<(1<<len(r)); i++ {
		var list []int
		for j:=0; j<len(r); j++ {
			if i & (1<<j) > 0 {
				list = append(list, j)
			}
		}

		sum := 0
		for _, k := range list {
			sum += int(r[k]-'0')
		}
		num := len(r) - len(list)
		if sum != 0 && sum % 3 == 0 && (ans == -1 || ans > num) {
			ans = num
		}
	}

	fmt.Println(ans)
}

