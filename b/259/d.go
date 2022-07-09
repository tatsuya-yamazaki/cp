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

	n := iou.i()
	sx := iou.i()
	sy := iou.i()
	tx := iou.i()
	ty := iou.i()

	var x, y, r []int
	c := make([][]int, n)
	for i:=0; i<n; i++ {
		xi := iou.i()
		yi := iou.i()
		ri := iou.i()
		x = append(x, xi)
		y = append(y, yi)
		r = append(r, ri)
	}

	for i:=0; i<n; i++ {
		for j:=i+1; j<n; j++ {
			if f(x[i],y[i],r[i],x[j],y[j],r[j], i, j) {
				c[i] = append(c[i], j)
				c[j] = append(c[j], i)
			}
		}
	}

	si, ti := -1, -1
	for i:=0; i<n; i++ {
		if h(x[i],y[i],r[i],sx,sy) {
			si = i
		}
		if h(x[i],y[i],r[i],tx,ty) {
			ti = i
		}
	}

	a := make([]bool, n)
	q := NewQueue()
	q.Add(value{si})

	for q.Next() {
		p := q.Pop().(value).v
		if a[p] {
			continue
		}
		a[p] = true
		for _, v := range c[p] {
			q.Add(value{v})
		}
	}

	if a[ti] {
		iou.pl("Yes")
	} else {
		iou.pl("No")
	}
}

type value struct {
	v int
}

func h(x, y, r, x2, y2 int) bool {
	x2 = x2 - x
	y2 = y2 - y
	v := g(r) - g(x2)
	if v == g(y2) {
		return true
	}
	return false
}

func f(x1, y1, r1, x2, y2, r2, i, j int) bool {
	l := g(x1-x2) + g(y1-y2)
	m := g(r1 + r2)
	n := g(r1-r2)
	if l == m {
		return true
	}
	if n <= l && l < m {
		return true
	}
	return false
}

func g(i int) int {
	return i * i
}

func Pow(x, n int) int {
	ret := 1
	for i:=0; i<n; i++ {
		ret *= x
	}
	return ret
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
