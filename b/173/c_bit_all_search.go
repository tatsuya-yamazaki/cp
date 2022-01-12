package main
import (
	"fmt"
)
func main() {
	var h, w, k, ans int
	fmt.Scan(&h)
	fmt.Scan(&w)
	fmt.Scan(&k)
	t := make([][]rune, h)
	for i:=0; i<h; i++ {
		var s string
		fmt.Scan(&s)
		t[i] = []rune(s)
	}
	hb := bitAll(h)
	wb := bitAll(w)
	for _, hbr := range hb {
		for _, wbr := range wb {
			c := 0
			for k:=0; k<h; k++ {
				for l:=0; l<w; l++ {
					if t[k][l] == '#' {
						ok := true
						for _, v := range hbr {
							if k == v {
								ok = false
							}
						}
						for _, v := range wbr {
							if l == v {
								ok = false
							}
						}
						if ok {
							c++
						}
					}
				}
			}
			if c == k {
				ans++
			}
		}
	}
	fmt.Println(ans)
}

func bitAll(n int) (ret [][]int) {
	for i:=0; i<(1<<n); i++ {
		var r []int
		for j:=0; j<n; j++ {
			if i & (1<<j) > 0 {
				r = append(r, j)
			}
		}
		ret = append(ret, r)
	}
	return ret
}
