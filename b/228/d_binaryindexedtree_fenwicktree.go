package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func add(fwtree []int, index, inc int) {
	for i:=index+1; i<=len(fwtree); i += i & -i {
		fwtree[i-1] += inc
	}
}

func sum(fwtree []int, index int) (ret int) {
	for i:=index+1; i>0; i -= i & -i {
		ret += fwtree[i-1]
	}
	return
}

func sumRange(fwtree []int, l, r int) (ret int) {
	if l == 0 {
		return sum(fwtree, r)
	}
	if l == r {
		return value(fwtree, l)
	}
	return sum(fwtree, r) - sum(fwtree, l-1)
}

func value(fwtree []int, index int) int {
	if index == 0 {
		return fwtree[0]
	}
	return sum(fwtree, index) - sum(fwtree, index-1)
}

func getTarget(fwtree []int, index int) (int, bool) {
	right, ok := getNewIndex(fwtree, index, len(fwtree)-1)
	if ok {
		return right, ok
	}
	if index == 0 {
		return 0, false
	}

	left, ok := getNewIndex(fwtree, 0, index-1)
	if ok {
		return left, ok
	}
	return 0, false
}

func getNewIndex(fwtree []int, l, r int) (int, bool) {
	if l == r {
		ret := value(fwtree, l)
		if ret == 0 {
			return 0, false
		}
		return l, true
	}

	mid := (l+r) / 2
	left := sumRange(fwtree, l, mid)
	if left > 0 {
		return getNewIndex(fwtree, l, mid)
	}

	right := sumRange(fwtree, mid+1, r)
	if right > 0 {
		return getNewIndex(fwtree, mid+1, r)
	}

	return 0, false
}

func main() {
	var q int

	fmt.Scan(&q)

	const n = 1048576
	a := make([]int, n, n)
	fwtree := make([]int, n, n)
	for i:=0; i<n; i++ {
		a[i] = -1
		add(fwtree, i, 1)
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
			index, ok := getTarget(fwtree, pos)
			if ok {
				a[index] = x
				add(fwtree, index, -1)
			}
		} else {
			fmt.Println(a[x%n])
		}

	}

}
