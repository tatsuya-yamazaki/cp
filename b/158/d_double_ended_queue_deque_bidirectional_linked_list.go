package main
import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
func main() {

	var s string
	var q int

	fmt.Scan(&s, &q)
	begin, end := newLL(s)
	isReverse := false

	sc := bufio.NewScanner(os.Stdin)
	for i:=0; i<q; i++ {
		sc.Scan()
		qi := strings.Split(sc.Text(), " ")

		switch qi[0] {
		case "1":
			if isReverse {
				isReverse = false
			} else {
				isReverse = true
			}
		case "2":
			c := qi[2]
			l := &LinkedList{nil, nil, c}

			switch qi[1] {
			case "1":
				if isReverse {
					l.prev = end
					end.next = l
					end = l
				} else {
					begin.prev = l
					l.next = begin
					begin = l
				}
			case "2":
				if isReverse {
					begin.prev = l
					l.next = begin
					begin = l
				} else {
					l.prev = end
					end.next = l
					end = l
				}
			}
		}
	}

	if isReverse {
		next := end
		for next != nil {
			// str += next.charとするなど、文字列操作は文字列長nの計算量のため、TLE
			// 文字列の先頭追加、末尾追加いずれも計算量nとなる
			fmt.Print(next.char)
			next = next.prev
		}
	} else {
		next := begin
		for next != nil {
			fmt.Print(next.char)
			next = next.next
		}
	}
	fmt.Println()
}

type LinkedList struct {
	prev, next *LinkedList
	char string
}
func newLL(s string) (begin, end *LinkedList) {
	begin = &LinkedList{nil, nil, string(s[:1])}
	prev := begin
	for i, c := range s {
		if i == 0 {
			continue
		}
		ll := &LinkedList{prev, nil, string(c)}
		prev.next = ll
		prev = ll
	}
	return begin, prev
}
