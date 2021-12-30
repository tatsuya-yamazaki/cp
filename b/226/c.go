package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	var(
		N int
		T []int
		D []bool
		A [][]int
	)
	fmt.Scanf("%d", &N)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	for i:=0; i<N; i++ {
		sc.Scan()
		tx := sc.Text()
		t, _ := strconv.Atoi(tx)
		T = append(T, t)
		D = append(D, true)
		sc.Scan()
		tx2 := sc.Text()
		k, _ := strconv.Atoi(tx2)
		var Ar []int
		for j:=0; j<k; j++{
			sc.Scan()
			tx3 := sc.Text()
			a, _ := strconv.Atoi(tx3)
			Ar = append(Ar, a)
		}
		A = append(A, Ar)
	}

	ans := T[N-1]
	T[N-1] = 0
	var task []int
	task = append(task, N)
	nt := task
	for {
		nnt := make([]int, 0)
		for _, v := range nt {
			if D[v-1] {
				for _, a := range A[v-1] {
					nnt = append(nnt, a)
					ans += T[a-1]
					T[a-1] = 0
				}
			}
			D[v-1] = false
		}
		if len(nnt) == 0 {
			break
		}
		nt = nnt
	}

	fmt.Println(ans)
}
