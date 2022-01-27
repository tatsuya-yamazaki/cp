package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	uf := NewUnionFind(n)

	for i:=0; i<m; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		uf.Unite(a-1, b-1)
	}

	g := make([]int, n)
	ans := 0
	for i:=0; i<n; i++ {
		r := uf.Root(i)
		g[r]++
		if g[r] > ans {
			ans = g[r]
		}
	}

	fmt.Println(ans)
}

type UnionFind struct {
	parent []int
	rank []int
}

func NewUnionFind(length int) *UnionFind {
	parent := make([]int, length, length)
	rank := make([]int, length, length)
	for i:=0; i<length; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (u *UnionFind) Root(index int) int {
	if u.parent[index] == index {
		return index
	} else {
		u.parent[index] = u.Root(u.parent[index])
		return u.parent[index]
	}
}

func (u *UnionFind) SameRoot(a, b int) bool {
	return u.Root(a) == u.Root(b)
}

func (u *UnionFind) Unite(a, b int) {
	a = u.Root(a)
	b = u.Root(b)

	if a == b {
		return
	}

	if u.rank[a] < u.rank[b] {
		u.parent[a] = b
	} else {
		u.parent[b] = a
		if u.rank[a] == u.rank[b] {
			u.rank[a]++
		}
	}
}
