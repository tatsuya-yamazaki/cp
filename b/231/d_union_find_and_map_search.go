package main

import(
	"fmt"
	"bufio"
	"os"
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
	var n, m int
	fmt.Scan(&n, &m)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	uf := make([]int, n, n)
	rank := make([]int, n, n)
	for i:=0; i<n; i++ {
		uf[i] = i
	}
	ans := make(map[int][]int)
	for i:=0; i<n; i++ {
		ans[i+1] = make([]int, 0)
	}
	for i:=0; i<m; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		sc.Scan()
		b, _ := strconv.Atoi(sc.Text())
		if sameRoot(uf, a-1, b-1) {
			fmt.Println("No")
			return
		}
		ans[a] = append(ans[a], b)
		ans[b] = append(ans[b], a)
		unite(uf, rank, a-1, b-1)
	}

	for _, v := range ans {
		if len(v) > 2 {
			fmt.Println("No")
			return
		}
	}

	fmt.Println("Yes")
}

