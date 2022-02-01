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

	n := iou.I()
	m := iou.I()
	a := iou.Is(n)

	h := NewHeap(DECENDING)

	for _, v := range a {
		h.Push(v)
	}

	for m > 0 {
		p := h.Pop()
		h.Push(p / 2)
		m--
	}

	ans := 0
	for h.Next() {
		ans += h.Pop()
	}

	iou.Pl(ans)
}

type Heap struct {
	list []int
	isChild func(parent, child int) bool
}

const(
        ASCENDING = true
        DECENDING = false
)

func NewHeap(ascending bool) *Heap {
	h := &Heap{make([]int, 0), nil}
	if ascending {
		h.isChild = func(parent, child int) bool { return h.list[parent] < h.list[child] }
	} else {
		h.isChild = func(parent, child int) bool { return h.list[parent] > h.list[child] }
	}
	return h
}

func (*Heap) parent(index int) int {
	// index is zero base
	return (index - 1) / 2
}

func (*Heap) left(index int) int {
	// index is zero base
	return index * 2 + 1
}

func (*Heap) right(index int) int {
	// index is zero base
	return (index + 1) * 2
}

func (h *Heap) Push(value int) {
	h.list = append(h.list, value)
	index := len(h.list) - 1
	for index != 0 {
		parent := h.parent(index)
		if h.isChild(parent, index) {
			break
		}
		h.list[parent], h.list[index] = h.list[index], h.list[parent]
		index = parent
	}
}

func (h *Heap) Top() int {
	return h.list[0]
}

func (h *Heap) Pop() int {
	ret := h.list[0]
	h.list[0] = h.list[len(h.list)-1]
	h.list = h.list[:len(h.list)-1]
	index := 0
	for len(h.list) > h.left(index) {
		child, right := h.left(index), h.right(index)
		if len(h.list) > right && h.isChild(right, child) {
			child = right
		}
		if h.isChild(index, child) {
			break
		} else {
			h.list[index], h.list[child] = h.list[child], h.list[index]
			index = child
		}
	}
	return ret
}

func (h *Heap) Next() bool {
	return len(h.list) != 0
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
