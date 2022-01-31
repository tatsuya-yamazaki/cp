package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math"
)

func main() {
	a := iou.I()
	a++
	s := iou.S()

	z := newll(0)
	var p *LinkedList
	p = z
	for i, v := range s {
		n := newll(i+1)
		if v == 'L' {
			if p.left != nil {
				p.left.right = n
				n.left = p.left
			}
			n.right = p
			p.left = n
		} else {
			if p.right != nil {
				p.right.left = n
				n.right = p.right
			}
			n.left = p
			p.right = n
		}
		p = n
	}
	for p.left != nil {
		p = p.left
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for p.right != nil {
		fmt.Fprint(w, p.value)
		p = p.right
		if p.right != nil {
			fmt.Fprint(w, " ")
		}
	}
	fmt.Fprint(w, " ")
	fmt.Fprintln(w, p.value)
}

func newll(value int) *LinkedList {
	return &LinkedList{nil, nil, value}
}

type LinkedList struct {
	left, right *LinkedList
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

