package main

import(
	"fmt"
	"reflect"
)

func main() {
	var n, m  int
	fmt.Scan(&n, &m)

	mt := make(map[[2]int]struct{})
	for i:=0; i<m; i++ {
		var ai, bi int
		fmt.Scan(&ai, &bi)
		mt[[2]int{ai,bi}] = struct{}{}
	}

	a := make([]int, m)
	b := make([]int, m)
	for i:=0; i<m; i++ {
		var ai, bi int
		fmt.Scan(&ai, &bi)
		a[i] = ai
		b[i] = bi
	}

	s := make([]int, n)
	for i, _ := range s {
		s[i] = i+1
	}
	for {
		ma := make(map[[2]int]struct{})
		for i:=0; i<m; i++ {
			ai := s[a[i]-1]
			bi := s[b[i]-1]
			if bi < ai {
				ai, bi = bi, ai
			}
			ma[[2]int{ai,bi}] = struct{}{}
		}

		if reflect.DeepEqual(mt, ma) {
			fmt.Println("Yes")
			return
		}
		if nextPermutation(s) {
			break
		}
	}

	fmt.Println("No")
}

func nextPermutation(s []int) bool {
	l, b := -1, -1
	for i:=0; i<len(s)-1; i++ {
		if s[i] < s[i+1] {
			l = i
		}
	}
	if l == -1 {
		return true
	}
	for i:=len(s)-1; i>=0; i-- {
		if s[l] < s[i] {
			b = i
			break
		}
	}
	s[l], s[b] = s[b], s[l]
	for i:=1; i<=(len(s)-1-l)/2; i++ {
		s[l+i], s[len(s)-i] = s[len(s)-i], s[l+i]
	}
	return false
}
