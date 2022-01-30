package main

import(
	"fmt"
)

func main() {
	var n int
	var s string

	fmt.Scan(&n)
	fmt.Scan(&s)
	runes := []rune(s)
	alp := make([][]int, 26)
	for i, v := range runes {
		alp[v-'a'] = append(alp[v-'a'], i)
	}

	l, r := 0, len(s)-1
	for l < r {
		for i, pos := range alp {
			if i >= int(runes[l]-'a') {
				break
			}
			if len(pos) == 0 {
				continue
			}
			upper := UpperBound(pos, r)
			targetPos := -1
			if pos[upper] <= r {
				targetPos = upper
			} else if pos[upper] > r && upper != 0 {
				targetPos = upper - 1
			} else {
				continue
			}
			target := pos[targetPos]
			if l < target && target <= r {
				runes[l], runes[target] = runes[target], runes[l]
				r = target - 1
				alp[i] = alp[i][:targetPos]
				break
			}
		}
		l++
	}

	ans := ""
	for _, v := range runes {
		ans += string(v)
	}

	fmt.Println(ans)
}

func UpperBound(s []int, value int) int {
        l, r := 0, len(s)-1
        for l != r {
                m := (l + r) / 2
                if value < s[m] {
			r = m
                } else {
			l = m + 1
                }
        }
        return l
}
