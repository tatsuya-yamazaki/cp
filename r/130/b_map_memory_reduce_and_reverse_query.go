package main

import(
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"math"
)

func main() {
	var h, w, c, q int

	fmt.Scan(&h)
	fmt.Scan(&w)
	fmt.Scan(&c)
	fmt.Scan(&q)

	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer([]byte{}, math.MaxInt64)

	tl := make([]int, q, q)
	nl := make([]int, q, q)
	cl := make([]int, q, q)
	for i:=0; i<q; i++ {
		sc.Scan()
		t := sc.Text()

		ts := strings.Split(t, " ")
		tr, _ := strconv.Atoi(ts[0])
		nr, _ := strconv.Atoi(ts[1])
		cr, _ := strconv.Atoi(ts[2])

		tl[i] = tr
		nl[i] = nr
		cl[i] = cr

	}

	hl := make(map[int]bool) // わざわざbool型のスライスを用意せずとも、必要分のみmapで追加すればよい。不要メモリ削減
	wl := make(map[int]bool)
	ans := make([]int, c+1, c+1)
	hq, wq := 0, 0
	for i:=q-1; i>=0; i-- {
		if tl[i] == 1 {
			// 通常mapアクセス時value, okを返し、okの値で存在有無を判定するが、
			// 存在しない場合もvalueはゼロ値で帰ってくる。
			// mapのvalueの型はboolなので、存在しない場合はfalseが返却される
			// if文では、第一返り値のみ評価する
			// この場合は存在時必ずtrueを格納するため、
			// okを確認せずとも、valueをチェックすれば存在有無は判定できる
			if !hl[nl[i]] {
				hl[nl[i]] = true
				ans[cl[i]] += w - wq
				hq++
			}
		} else {
			if !wl[nl[i]] {
				wl[nl[i]] = true
				ans[cl[i]] += h - hq
				wq++
			}
		}
	}

	for i:=1; i<=c; i++ {
		fmt.Println(ans[i])
	}
}
