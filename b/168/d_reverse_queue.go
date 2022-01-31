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

	aisle := make([][]int, n+1)
	for i:=0; i<m; i++ {
		ai, bi := iou.I(), iou.I()
		aisle[ai] = append(aisle[ai], bi)
		aisle[bi] = append(aisle[bi], ai)
	}

	guide, ans := f(aisle, n)

	if ans {
		iou.Pl("Yes")
		for i, v := range guide {
			if i < 2 {
				continue
			}
			iou.Pl(v)
		}
	} else {
		iou.Pl("No")
	}
}

func f(aisle [][]int, n int) ([]int, bool) {
	guide := make([]int, n+1)
	guide[1] = -1
	var q Queue
	q.Add(1)
	for q.Next() {
		pos := q.Pop()
		for _, next := range aisle[pos] {
			if guide[next] == 0 {
				guide[next] = pos
			} else {
				continue
			}
			q.Add(next)
		}
	}
	for i:=2; i<=n; i++ {
		if guide[i] == 0 {
			return guide, false
		}
	}
	return guide, true
}

type Queue struct {
	begin *LinkedList
	end *LinkedList
}

func (q *Queue) Next() bool {
	if q.begin == nil {
		return false
	}
	return true
}

func (q *Queue) Add(value int) {
	ll := &LinkedList{q.end, nil, value}
	if q.end == nil {
		q.begin = ll
	} else {
		q.end.next = ll
	}
	q.end = ll
}

func (q *Queue) Pop() int {
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

type LinkedList struct {
        prev, next *LinkedList
	value int
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
