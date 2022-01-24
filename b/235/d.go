package main

import(
	"fmt"
)

func main() {
	var(
		a, n int
	)

	fmt.Scan(&a)
	fmt.Scan(&n)

	m := make(map[int]int)

	q := &Queue{}
	q.add(1)

	limit := 1
	for limit <= n {
		limit *= 10
	}

	for q.next() {
		value := q.pop()
		times := m[value] + 1

		av := a * value
		if av < limit {
			_, ok := m[av]
			if !ok {
				q.add(av)
				m[av] = times
			}
		}

		rv, revok := rev(value)
		if revok && rv < limit {
			_, ok := m[rv]
			if !ok {
				q.add(rv)
				m[rv] = times
			}
		}
	}

	ans, ok := m[n]
	if ok {
		fmt.Println(ans)
	} else {
		fmt.Println(-1)
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func rev(x int) (int, bool) {
	if x < 10 {
		return x, false
	}

	ret := x / 10
	rem := x % 10

	if rem == 0 {
		return x, false
	}

	ten := 10
	for x > ten {
		ten *= 10
	}
	ten /= 10
	ret = ret + rem * ten

	return ret, true
}

type Queue struct {
	begin *LinkedList
	end *LinkedList
}

func (q *Queue) next() bool {
	if q.begin == nil {
		return false
	}
	return true
}

func (q *Queue) add(value int) {
	ll := &LinkedList{q.end, nil, value}
	if q.begin == nil {
		q.begin = ll
	}
	if q.end != nil {
		q.end.next = ll
	}
	q.end = ll
}

func (q *Queue) pop() int {
	value := q.begin.value
	if q.begin == q.end {
		q.end = nil
	}
	q.begin = q.begin.next
	return value
}

type LinkedList struct {
        prev, next *LinkedList
	value int
}
