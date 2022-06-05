package main

import(
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	defer iou.Fl()

	q := iou.I()
	var s, x, k, a []int
	for i:=0; i<q; i++ {
		qi := iou.S()
		s = append(s, qi)

		if qi == "1" {
			xi = iou.I()
			a = append(a, xi)
			x = append(x, xi)
			k = append(k, 0)
		} else if qi == "2" {
			x = append(x, iou.I())
			k = append(k, iou.I())
		} else {
			x = append(x, iou.I())
			k = append(k, iou.I())
		}

	}

	w := NewWaveletMatrix(a)

	for i:=0; i<q; i++ {
		si, xi, ki := s[i], x[i], k[i]
	}
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

// Node is the interface a node of Heap.
// Less returns Node is less than a or not.
// Greater returns Node is greater than a or not.
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
type IOUtil struct {
	Scanner *bufio.Scanner
	Reader *bufio.Reader
	Writer *bufio.Writer
}
func NewIOUtil() *IOUtil {
	iou := IOUtil{
		Scanner: bufio.NewScanner(os.Stdin),
		Writer: bufio.NewWriter(os.Stdout),
	}
	iou.Scanner.Split(bufio.ScanWords)
	iou.Scanner.Buffer(make([]byte, 1024), math.MaxInt64)
	return &iou
}
func (*IOUtil) ToInt(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}
func (iou *IOUtil) Int() int {
	iou.Scanner.Scan()
	return iou.ToInt(iou.Scanner.Text())
}
func (iou *IOUtil) Str() string {
	iou.Scanner.Scan()
	return iou.Scanner.Text()
}
func (iou *IOUtil) I() int {
	return iou.Int()
}
func (iou *IOUtil) S() string {
	return iou.Str()
}
func (iou *IOUtil) Ints(n int) []int {
	ret := make([]int, n)
	for i:=0; i<n; i++ {
		ret[i] = iou.Int()
	}
	return ret
}
func (iou *IOUtil) Is(n int) []int {
	return iou.Ints(n)
}
func (iou *IOUtil) Ints2(n int) (a, b []int) {
        a = make([]int, n)
        b = make([]int, n)
        for i:=0; i<n; i++ {
                a[i] = iou.I()
                b[i] = iou.I()
        }
	return a, b
}
func (iou *IOUtil) Is2(n int) (a, b []int) {
	return iou.Ints2(n)
}
func (iou *IOUtil) CumulativeSum(n int) (cumulative, ints []int) {
	cumulative = append(cumulative, 0)
	for i:=0; i<n; i++ {
		ai := iou.Int()
		cumulative = append(cumulative, cumulative[i] + ai)
		ints = append(ints, ai)
	}
	return
}
func (iou *IOUtil) Cms(n int) (cumulative, ints []int) {
	return iou.CumulativeSum(n)
}
func (iou *IOUtil) Print(a ...interface{}) {
	fmt.Fprint(iou.Writer, a...)
}
func (iou *IOUtil) P(a ...interface{}) {
	iou.Print(a...)
}
func (iou *IOUtil) Println(a ...interface{}) {
	fmt.Fprintln(iou.Writer, a...)
}
func (iou *IOUtil) Pl(a ...interface{}) {
	iou.Println(a...)
}
func (iou *IOUtil) Flush() {
	iou.Writer.Flush()
}
func (iou *IOUtil) Fl() {
	iou.Flush()
}
var iou = NewIOUtil()
