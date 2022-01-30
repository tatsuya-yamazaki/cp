package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	a := iou.I()
	b := iou.I()
	x := iou.I()

	max := 1000000000
	l, r := 0, max
	for l != r {
		m := (l + r) / 2 + 1
		n := f(a,b,m)
		if n > x {
			r = m - 1
		} else {
			l = m
		}
	}
	if l > max {
		l = max
	}
	fmt.Println(l)
}

func f(a, b, n int) int {
	return a * n + b * d(n)
}

func d(n int) int {
	ret := 1
	for n / 10 != 0 {
		ret++
		n /= 10
	}
	return ret
}

type IOUtil struct {
	Scanner *bufio.Scanner
}
func NewIOUtil() *IOUtil {
	iou := IOUtil{
		Scanner: bufio.NewScanner(os.Stdin),
	}
	iou.Scanner.Split(bufio.ScanWords)
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
func (iou *IOUtil) Is(n int) []int {
	ret := make([]int, n)
	for i:=0; i<n; i++ {
		ret[i] = iou.Int()
	}
	return ret
}
var iou = NewIOUtil()
