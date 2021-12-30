package main
import (
	"fmt"
)

type que struct {
	value int
	prev, next *que
}
type queue struct {
	begin *que
	end *que
}
func (q *queue) push(value int) {
	begin := &que{value, nil, q.begin}
	if q.begin != nil {
		q.begin.prev = begin
	} else {
		q.end = begin
	}
	q.begin = begin
}
func (q *queue) pop() int {
	if q.end == nil {
		return 0
	}
	ret := q.end.value
	if q.end.prev == nil {
		q.begin = nil
	}
	q.end = q.end.prev
	return ret
}
func main() {
	var n, m int
	fmt.Scan(&n, &m)

	r := make(map[int][]int)
	for i:=0; i<m; i++ {
		var a, b int
		fmt.Scan(&a, &b)
		r[a] = append(r[a], b)
	}

	ans := 0
	for i:=1; i<=n; i++ {
		passed := make([]bool, n+1)
		q := queue{nil, nil}
		q.push(i)
		for j:=q.pop(); j!=0; j=q.pop() {
			if passed[j] {
				continue
			}
			passed[j] = true
			ans++
			for _, v := range r[j] {
				q.push(v)
			}
		}
	}

	fmt.Println(ans)
}
