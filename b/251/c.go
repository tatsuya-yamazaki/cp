package main

import(
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type po struct {
	v, i int
}

func main() {
	defer iou.Fl()

	m := make(map[string]int)
	n := iou.I()
	ps := make([]po, 0)
	for i:=0; i<n; i++ {
		s := iou.S()
		t := iou.I()
		_, ok := m[s]
		if ok {
			continue
		}
		m[s]++
		ps = append(ps, po{t, i})
	}
	max := -1
	mi := 9999999
	for _, p := range ps {
		if max < p.v {
			max = p.v
			mi = p.i
		} else if max == p.v {
			if p.i < mi {
				mi = p.i
			}
		}
	}

	iou.Pl(mi+ 1)
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
func (iou *IOUtil) PrintIntSlice(s []int) {
	for _, v := range s {
		fmt.Fprint(iou.Writer, v)
	}
	fmt.Fprintln(iou.Writer)
}
func (iou *IOUtil) Pis(s []int) {
	iou.PrintIntSlice(s)
}
func (iou *IOUtil) PrintIntSliceSpace(s []int) {
	last := len(s) - 1
	for i, v := range s {
		fmt.Fprint(iou.Writer, v)
		if i == last {
			fmt.Fprintln(iou.Writer)
		} else {
			fmt.Fprint(iou.Writer, " ")
		}
	}
}
func (iou *IOUtil) Piss(s []int) {
	iou.PrintIntSliceSpace(s)
}
func (iou *IOUtil) Flush() {
	iou.Writer.Flush()
}
func (iou *IOUtil) Fl() {
	iou.Flush()
}
var iou = NewIOUtil()
