package main

import(
	"fmt"
	"reflect"
)

func main() {
	var n, m  int
	fmt.Scan(&n, &m)

	tam := NewAdjacencyMatrix(n)
	for i:=0; i<m; i++ {
		var ai, bi int
		fmt.Scan(&ai, &bi)
		ai--
		bi--
		tam[ai][bi] = true
		tam[bi][ai] = true
	}

	a := make([]int, m)
	b := make([]int, m)
	for i:=0; i<m; i++ {
		var ai, bi int
		fmt.Scan(&ai, &bi)
		a[i] = ai - 1
		b[i] = bi - 1
	}

	s := make([]int, n)
	for i, _ := range s {
		s[i] = i
	}
	for {
		nam := NewAdjacencyMatrix(n)
		for i:=0; i<m; i++ {
			ai := s[a[i]]
			bi := s[b[i]]
			nam[ai][bi] = true
			nam[bi][ai] = true
		}

		if reflect.DeepEqual(tam, nam) {
			fmt.Println("Yes")
			return
		}
		if nextPermutation(s) {
			break
		}
	}

	fmt.Println("No")
}

func NewAdjacencyMatrix(n int) [][]bool {
	m := make([][]bool, n)
	for i:=0; i<n; i++ {
		m[i] = make([]bool, n)
	}
	return m
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
