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
	var a, b, c []int
	for i:=0; i<n; i++ {
		a = append(a, iou.i())
		b = append(b, iou.i())
		c = append(c, iou.i())
	}

	iou.pl()
}


type node struct {
        a int
}

func (n node) Less(a *HeapNode) bool {
        v := (*a).(node)
        return n.a < v.a
}

func (n node) Greater(a *HeapNode) bool {
        v := (*a).(node)
        return n.a > v.a
}

type HeapNode interface {
        Less(a *HeapNode) bool
        Greater(a *HeapNode) bool
}

// Heap is the binary heap structure.
// To use it, the Node interface must be implemented.
// Its indexes are 0-origin.
// It can use ascending or descending order.
// TODO It may need to be refactored, expecially Pop().
// TODO It may need to be devided into min heap and max heap. Then remove isChild, use Less or Greater
type Heap struct {
        n []*HeapNode
        isChild func(parent, child int) bool
}

const(
        ASCENDING = true
        DESCENDING = false
)

func NewHeap(ascending bool) *Heap {
        h := &Heap{make([]*HeapNode, 0), nil}
        if ascending {
                h.isChild = func(parent, child int) bool { return (*h.n[parent]).Less(h.n[child]) }
        } else {
                h.isChild = func(parent, child int) bool { return (*h.n[parent]).Greater(h.n[child]) }
        }
        return h
}

func parent(i int) int {
        return (i - 1) / 2
}

func left(i int) int {
        return i * 2 + 1
}

func right(i int) int {
        return (i + 1) * 2
}

func (h *Heap) Add(value HeapNode) {
        h.n = append(h.n, &value)
        i := len(h.n) - 1
        for i != 0 {
                p := parent(i)
                if h.isChild(p, i) {
                        break
                }
                h.n[p], h.n[i] = h.n[i], h.n[p]
                i =p
        }
}

func (h *Heap) Top() *HeapNode {
        return h.n[0]
}

func (h *Heap) Pop() HeapNode {
        ret := h.n[0]
        last := len(h.n)-1
        h.n[0] = h.n[last]
        h.n = h.n[:last]
        i := 0
        for last > left(i) {
                c, r := left(i), right(i)
                if len(h.n) > r && h.isChild(r, c) {
                        c = r
                }
                if h.isChild(i, c) {
                        break
                } else {
                        h.n[i], h.n[c] = h.n[c], h.n[i]
                        i = c
                }
        }
        return *ret
}

func (h *Heap) Next() bool {
        return len(h.n) != 0
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
