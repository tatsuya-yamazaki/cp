package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func root(uf []int, index int) int {
	if uf[index] == index {
		return index
	} else {
		uf[index] = root(uf, uf[index])
		return uf[index]
	}
}

func sameRoot(uf []int, a, b int) bool {
	return root(uf, a) == root(uf, b)
}

func unite(parent, rank []int, a, b int) {
	a = root(parent, a)
	b = root(parent, b)

	if a == b {
		return
	}

	if rank[a] < rank[b] {
		parent[a] = b
	} else {
		parent[b] = a
		if rank[a] == rank[b] {
			rank[a]++
		}
	}
}

func main() {
	var q int

	fmt.Scan(&q)

	const n = 1048576
	a := make([]int, n, n)
	for i:=0; i<n; i++ {
		a[i] = -1
	}

	ufParent := make([]int, n, n)
	for i:=0; i<n; i++ {
		ufParent[i] = i
	}

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i:=0; i<q; i++ {
		sc.Scan()
		t, _ := strconv.Atoi(sc.Text())

		sc.Scan()
		x, _ := strconv.Atoi(sc.Text())

		if t == 1 {

			pos := x % n
			r := root(ufParent, pos)
			a[r] = x
			next := 0
			for i:=r+1; i<=n; i++ {
				if i == n {
					i = 0
				}

				next = root(ufParent, i)
				if r != next {
					ufParent[r] = next
					break
				}
			}

		} else {
			fmt.Println(a[x%n])
		}

	}

}
