package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

func main() {
	n := iou.I()
	s := iou.S()

	d := NewDeque()

	d.AddLeft(n)
	for i:=len(s)-1; i>=0; i-- {
		n--
		if s[i] == 'L' {
			d.AddRight(n)
		} else {
			d.AddLeft(n)
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for {
		fmt.Fprint(w, d.PopLeft())
		if d.Next() {
			fmt.Fprint(w, " ")
		} else {
			fmt.Fprintln(w)
			break
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

func (d *Deque) AddLeft(value int) {
	ll := &LinkedList{nil, d.begin, value}
	if d.begin == nil {
		d.end = ll
	} else {
		d.begin.prev = ll
	}
	d.begin = ll
}

func (d *Deque) AddRight(value int) {
	ll := &LinkedList{d.end, nil, value}
	if d.end == nil {
		d.begin = ll
	} else {
		d.end.next = ll
	}
	d.end = ll
}

func (d *Deque) PopLeft() int {
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

func (d *Deque) PopRight() int {
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

type LinkedList struct {
        prev, next *LinkedList
	value int
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

