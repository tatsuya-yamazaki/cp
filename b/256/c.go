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

	h1 := iou.i()
	h2 := iou.i()
	h3 := iou.i()

	w1 := iou.i()
	w2 := iou.i()
	w3 := iou.i()

	ans := 0
	for a11:=1; a11<h1-1; a11++ {
		for a12:=1; a12<h1-a11; a12++ {
			for a21:=1; a21<w1-a11; a21++ {
				for a22:=1; (a22<w2-a12 && a22<h2-a21); a22++ {
					a13 := h1 - a11 - a12
					a31 := w1 - a11 - a21
					a23 := h2 - a21 - a22
					a32 := w2 - a12 - a22
					if a13 <= 0 || a31 <= 0 || a23 <= 0 || a32 <= 0 {
						continue
					}

					a33h := h3 - a31 - a32
					a33w := w3 - a13 - a23

					if a33h <= 0 || a33w <= 0 || a33h != a33w {
						continue
					}

					ans++
				}
			}
		}
	}

	iou.pl(ans)
}

func Pow(x, n int) int {
	ret := 1
	for i:=0; i<n; i++ {
		ret *= x
	}
	return ret
}
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
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
