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
}

func (c *chess) check(row, column int) bool {
	if row < 0 || row >= c.n || column < 0 || column >= c.n || c.s[row][column] == '#' {
		return false
	}
	return true
}

const (
	NE = iota
	NW
	SE
	SW
	OO
)

func main() {
	defer iou.Fl()

	n := iou.I()
	ax := iou.I() - 1
	ay := iou.I() - 1
	bx := iou.I() - 1
	by := iou.I() - 1
	c := new(chess)
	c.n = n

	times := make([][][]int, n)
	for i:=0; i<n; i++ {
		infs := make([][]int, n)
		for j:=0; j<n; j++ {
			inf := make([]int, 4)
			for k:=0; k<4; k++ {
				inf[k] = math.MaxInt64
			}
			infs[j] = inf
		}
		times[i] = infs
	}

	for i:=0; i<n; i++ {
		c.s = append(c.s, []rune(iou.S()))
	}

	d := NewDeque()
	for i:=0; i<4; i++ {
		times[ax][ay][i] = 0
	}

	d.PushLeft(Value{ax, ay, OO})
	times[ax][ay] = append(times[ax][ay], 0)

	for d.Next() {
		value := d.PopLeft()

		f(NE, value, c, d, times)
		f(NW, value, c, d, times)
		f(SE, value, c, d, times)
		f(SW, value, c, d, times)

	}

	ans := math.MaxInt64
	for i:=0; i<4; i++ {
		ans = Min(ans, times[bx][by][i])
	}
	if ans == math.MaxInt64 {
		iou.Pl(-1)
	} else {
		iou.Pl(ans)
	}
}

func f(nextDirection int, value Value, c *chess, d *Deque, times [][][]int) {
	row, column, preDirection := value.row, value.column, value.preDirection
	t := times[row][column][preDirection]
	rd, cd := 0, 0
	switch nextDirection {
	case NE:
		rd, cd = 1, 1
	case NW:
		rd, cd = 1, -1
	case SE:
		rd, cd = -1, 1
	case SW:
		rd, cd = -1, -1
	}
	if c.check(row+rd, column+cd) {
		tn := times[row+rd][column+cd][nextDirection]
		if preDirection == nextDirection {
			if tn > t {
				times[row+rd][column+cd][nextDirection] = t
				d.PushLeft(Value{row+rd, column+cd, nextDirection})
			}
		} else {
			if tn > t + 1 {
				times[row+rd][column+cd][nextDirection] = t + 1
				d.PushRight(Value{row+rd, column+cd, nextDirection})
			}
		}
	}
}


type Deque struct {
        begin *LinkedList
        end *LinkedList
}

func NewDeque() *Deque {
        return &Deque{}
}

func (d *Deque) Next() bool {
        if d.begin == nil {
                return false
        }
        return true
}

func (d *Deque) PushLeft(value Value) {
        ll := &LinkedList{nil, d.begin, value}
        if d.begin == nil {
                d.end = ll
        } else {
                d.begin.prev = ll
        }
        d.begin = ll
}

func (d *Deque) PushRight(value Value) {
        ll := &LinkedList{d.end, nil, value}
        if d.end == nil {
                d.begin = ll
        } else {
                d.end.next = ll
        }
        d.end = ll
}

func (d *Deque) PopLeft() Value {
        value := d.begin.value
        if d.begin == d.end {
                d.begin = nil
                d.end = nil
        } else {
                d.begin.next.prev = nil
                d.begin = d.begin.next
        }
        return value
}

func (d *Deque) PopRight() Value {
        value := d.end.value
        if d.begin == d.end {
                d.begin = nil
                d.end = nil
        } else {
                d.end.prev.next = nil
                d.end = d.end.prev
        }
        return value
}

type Value struct {
	row, column, preDirection int
}

type LinkedList struct {
        prev, next *LinkedList
        value Value
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
