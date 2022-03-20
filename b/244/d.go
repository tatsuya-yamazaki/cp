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

	s1 := iou.S()
	s2 := iou.S()
	s3 := iou.S()
	t1 := iou.S()
	t2 := iou.S()
	t3 := iou.S()

	rs,gs,bs := 0, 0, 0
	rt,gt,bt := 0, 0, 0

	if s1 == "R" {
		rs++
	}
	if s2 == "R" {
		rs++
	}
	if s3 == "R" {
		rs++
	}
	if s1 == "G" {
		gs++
	}
	if s2 == "G" {
		gs++
	}
	if s3 == "G" {
		gs++
	}
	if s1 == "B" {
		bs++
	}
	if s2 == "B" {
		bs++
	}
	if s3 == "B" {
		bs++
	}

	if t1 == "R" {
		rt++
	}
	if t2 == "R" {
		rt++
	}
	if t3 == "R" {
		rt++
	}
	if t1 == "G" {
		gt++
	}
	if t2 == "G" {
		gt++
	}
	if t3 == "G" {
		gt++
	}
	if t1 == "B" {
		bt++
	}
	if t2 == "B" {
		bt++
	}
	if t3 == "B" {
		bt++
	}

	if rs != rt || gs != gt || bs != bt {
		iou.Pl("No")
		return
	}

	if rs == 3 || gs == 3 || bs == 3 {
		iou.Pl("Yes")
		return
	} else if rs == 2 || gs == 2 || bs == 2 {
		iou.Pl("Yes")
		return
	}

	diff := 0

	if s1 != t1 {
		diff++
	}

	if s2 != t2 {
		diff++
	}

	if s3 != t3 {
		diff++
	}

	switch diff {
	case 0:
		iou.Pl("Yes")
	case 2:
		iou.Pl("No")
	case 3:
		iou.Pl("Yes")
	}
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
