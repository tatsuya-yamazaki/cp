package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type UnionFind []int

func NewUnionFind(n int) *UnionFind {
	t := make(UnionFind, n, n)
	for i:=0; i<n; i++ {
		t[i] = -1
	}
	return &t
}

func (t *UnionFind) Root(n int) int {
	parent := (*t)[n]
	if parent < 0 {
		return n
	}
	(*t)[n] = t.Root(parent)
	return (*t)[n]
}

func (t *UnionFind) Same(a, b int) bool {
	return t.Root(a) == t.Root(b)
}

func (t *UnionFind) Size(n int) int {
	return -(*t)[t.Root(n)]
}

func (t *UnionFind) Unite(a, b int) bool {
	a, b = t.Root(a), t.Root(b)
	if t.Same(a, b) {
		return false
	}
	if (*t).Size(a) < (*t).Size(b) {
		a, b = b, a
	}
	(*t)[a] -= (*t).Size(b)
	(*t)[b] = a
	return true
}

func main() {
	var n, m int

	fmt.Scan(&n, &m)
	sc := bufio.NewScanner(os.Stdin)

	edge := make([][]int, n+1, n+1)
	for i:=0; i<m; i++ {
			sc.Scan()
			t := sc.Text()
			tl := strings.Split(t, " ")
			a, _ := strconv.Atoi(tl[0])
			b, _ := strconv.Atoi(tl[1])
			edge[a] = append(edge[a], b)
	}

	uf := NewUnionFind(n+1)
	ans := 0
	ansl := make([]int, n, n)
	for i:=n; i>=2; i-- {
		ans++
		for _, b := range edge[i] {
			if uf.Unite(i, b) {
				ans--
			}
		}
		ansl[i-2] = ans
	}
	for _, v := range ansl {
		fmt.Println(v)
	}
}
