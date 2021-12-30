package main

import(
	"fmt"
	"reflect"
)

func main() {
	var(
		s, t string
	)
	fmt.Scanf("%s", &s)
	fmt.Scanf("%s", &t)

	if test([]rune(s), []rune(t)) {
		fmt.Printf("Yes\n")
	} else {
		fmt.Printf("No\n")
	}
}

func test(s,t []rune) bool {
	if reflect.DeepEqual(s,t) {
		return true
	}

	var diffCount int
	var s1, s2, t1, t2 rune

	for i, _ := range s {
		if s[i] != t[i] {
			if diffCount == 0 {
				s1 = s[i]
				t1 = t[i]
				diffCount++
				continue
			}

			s2 = s[i]
			t2 = t[i]
			diffCount++

			if (s1 == t2) && (s2 == t1) && ((len(s) - 1 == i) || reflect.DeepEqual(s[i+1:] ,t[i+1:])) {
				return true
			} else {
				return false
			}
		}

		if diffCount != 0 {
			break
		}
	}
	return false
}
