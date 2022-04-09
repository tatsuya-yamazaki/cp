package main

import(
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type chess struct {
	n int
	s [][]rune
	done [][]bool
}

func (c *chess) check(row, column int) bool {
	if row < 0 || row >= c.n || column < 0 || column >= c.n || c.s[row][column] == '#' {
		return false
	}
	return true
}

func main() {
	defer iou.Fl()

	n := iou.I()
	ax := iou.I() - 1
	ay := iou.I() - 1
	bx := iou.I() - 1
	by := iou.I() - 1
	c := new(chess)
	c.n = n
	for i:=0; i<n; i++ {
		c.s = append(c.s, []rune(iou.S()))
		c.done = append(c.done, make([]bool, n))
	}
	c.done[ax][ay] = true

	q := NewQueue()
	q.Push(ax, ay, 0)

	for q.Next() {
		row, column, times := q.Pop()
		if row == bx && column == by {
			iou.Pl(times)
			return
		}

		for i:=1; c.check(row+i, column+i); i++ {
			if c.done[row+i][column+i] {
				continue
			}
			c.done[row+i][column+i] = true
			q.Push(row+i, column+i, times+1)
		}

		for i:=1; c.check(row+i, column-i); i++ {
			if c.done[row+i][column-i] {
				continue
			}
			c.done[row+i][column-i] = true
			q.Push(row+i, column-i, times+1)
		}

		for i:=1; c.check(row-i, column+i); i++ {
			if c.done[row-i][column+i] {
				continue
			}
			c.done[row-i][column+i] = true
			q.Push(row-i, column+i, times+1)
		}

		for i:=1; c.check(row-i, column-i); i++ {
			if c.done[row-i][column-i] {
				continue
			}
			c.done[row-i][column-i] = true
			q.Push(row-i, column-i, times+1)
		}

	}

	iou.Pl(-1)
}


type Queue struct {
        begin *LinkedList
        end *LinkedList
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

func (q *Queue) Push(row, column, times int) {
        ll := &LinkedList{q.end, nil, row, column, times}
        if q.end == nil {
                q.begin = ll
        } else {
                q.end.next = ll
        }
        q.end = ll
}

func (q *Queue) Pop() (row, column, times int) {
        row = q.begin.row 
        column = q.begin.column
        times = q.begin.times
        if q.begin == q.end {
                q.begin = nil
                q.end = nil
        } else {
                q.begin.next.prev = nil
                q.begin = q.begin.next
        }
        return row, column, times
}

type LinkedList struct {
        prev, next *LinkedList
        row, column, times int
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
