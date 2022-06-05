package main

import(
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	defer iou.fl()

	_ = iou.i()
	m := iou.i()
	route := make(map[int][]int)
	for i:=0; i<m; i++ {
		a := iou.i()
		b := iou.i()
		route[a] = append(route[a], b)
		route[b] = append(route[b], a)
	}
	query := iou.i()
	for i:=0; i<query; i++ {
		x := iou.i()
		k := iou.i()
		b := make(map[int]bool)
		q := NewQueue()
		q.Add(node{x, 0})
		b[x] = true
		ans := 0
		for q.Next() {
			ni := q.Pop().(node)
			ans += ni.n
			if ni.d == k {
				continue
			}
			for _, v := range route[ni.n] {
				if b[v] {
					continue
				}
				b[v] = true
				q.Add(node{v, ni.d+1})
			}
		}
		iou.pl(ans)
	}
}

type node struct {
        n, d int
}

type Queue struct {
	begin *queueLinkedList
	end *queueLinkedList
}

func NewQueue() *Queue {
	return &Queue{}
}

func (q *Queue) Next() bool {
	if q.begin == nil {
		return false
	}
	return true
}

func (q *Queue) Add(value QueueValue) {
	ll := &queueLinkedList{q.end, nil, value}
	if q.end == nil {
		q.begin = ll
	} else {
		q.end.next = ll
	}
	q.end = ll
}

func (q *Queue) Pop() QueueValue {
	value := q.begin.value
	if q.begin == q.end {
		q.begin = nil
		q.end = nil
	} else {
		q.begin.next.prev = nil
		q.begin = q.begin.next
	}
	return value
}

type QueueValue interface {
}

type queueLinkedList struct {
        prev, next *queueLinkedList
	value QueueValue
}

func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func Min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
type ioUtil struct {
	scanner *bufio.Scanner
	reader *bufio.Reader
	writer *bufio.Writer
}
func newIOUtil() *ioUtil {
	iou := ioUtil{
		scanner: bufio.NewScanner(os.Stdin),
		writer: bufio.NewWriter(os.Stdout),
	}
	iou.scanner.Split(bufio.ScanWords)
	iou.scanner.Buffer(make([]byte, 1024), math.MaxInt64)
	return &iou
}
func (*ioUtil) toInt(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}
func (iou *ioUtil) int() int {
	iou.scanner.Scan()
	return iou.toInt(iou.scanner.Text())
}
func (iou *ioUtil) str() string {
	iou.scanner.Scan()
	return iou.scanner.Text()
}
func (iou *ioUtil) i() int {
	return iou.int()
}
func (iou *ioUtil) s() string {
	return iou.str()
}
func (iou *ioUtil) ints(n int) []int {
	ret := make([]int, n)
	for i:=0; i<n; i++ {
		ret[i] = iou.int()
	}
	return ret
}
func (iou *ioUtil) is(n int) []int {
	return iou.ints(n)
}
func (iou *ioUtil) ints2(n int) (a, b []int) {
	a = make([]int, n)
	b = make([]int, n)
	for i:=0; i<n; i++ {
		a[i] = iou.int()
		b[i] = iou.int()
	}
	return a, b
}
func (iou *ioUtil) is2(n int) (a, b []int) {
	return iou.ints2(n)
}
func (iou *ioUtil) cumulativeSum(n int) (cumulative, ints []int) {
	cumulative = append(cumulative, 0)
	for i:=0; i<n; i++ {
		ai := iou.int()
		cumulative = append(cumulative, cumulative[i] + ai)
		ints = append(ints, ai)
	}
	return
}
func (iou *ioUtil) cms(n int) (cumulative, ints []int) {
	return iou.cumulativeSum(n)
}
func (iou *ioUtil) print(a ...interface{}) {
	fmt.Fprint(iou.writer, a...)
}
func (iou *ioUtil) p(a ...interface{}) {
	iou.print(a...)
}
func (iou *ioUtil) println(a ...interface{}) {
	fmt.Fprintln(iou.writer, a...)
}
func (iou *ioUtil) pl(a ...interface{}) {
	iou.println(a...)
}
func (iou *ioUtil) printIntSlice(s []int) {
	for _, v := range s {
		fmt.Fprint(iou.writer, v)
	}
	fmt.Fprintln(iou.writer)
}
func (iou *ioUtil) pis(s []int) {
	iou.printIntSlice(s)
}
func (iou *ioUtil) printIntSliceSpace(s []int) {
	last := len(s) - 1
	for i, v := range s {
		fmt.Fprint(iou.writer, v)
		if i == last {
			fmt.Fprintln(iou.writer)
		} else {
			fmt.Fprint(iou.writer, " ")
		}
	}
}
func (iou *ioUtil) piss(s []int) {
	iou.printIntSliceSpace(s)
}
func (iou *ioUtil) flush() {
	iou.writer.Flush()
}
func (iou *ioUtil) fl() {
	iou.flush()
}
var iou = newIOUtil()
