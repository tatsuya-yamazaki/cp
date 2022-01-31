package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	h := iou.I()
	w := iou.I()
	a := make([][]int, h)
	for i:=0; i<h; i++ {
		a[i] = iou.Is(w)
	}

	for i:=0; i<w; i++ {
		for j:=0; j<h; j++ {
			fmt.Print(a[j][i])
			if j < h - 1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

type IOUtil struct {
	Scanner *bufio.Scanner
}
func NewIOUtil() *IOUtil {
	iou := IOUtil{
		Scanner: bufio.NewScanner(os.Stdin),
	}
	iou.Scanner.Split(bufio.ScanWords)
	bufSize := 600000
	buf := make([]byte, bufSize)
	iou.Scanner.Buffer(buf, bufSize)
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
var iou = NewIOUtil()

