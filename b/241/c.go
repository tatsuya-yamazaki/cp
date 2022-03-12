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
	s := make([][]rune, n)
	for i:=0; i<n; i++ {
		s[i] = []rune(iou.S())
	}

	ans := false

	for i:=0; i<n; i++ {
		if i + 5 >= n {
			continue
		}
		for j:=0; j<n; j++ {
			if f(s[j][i:i+6]) {
				ans = true
			}
		}
	}

	for i:=0; i<n; i++ {
		if i + 5 >= n {
			continue
		}
		for j:=0; j<n; j++ {
			si := make([]rune, 0)
			for k:=i; k<i+6; k++ {
				si = append(si, s[k][j])
			}
			if f(si) {
				ans = true
			}
		}
	}

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			if i - 5 < 0 || j + 5 >= n {
				continue
			}
			si := make([]rune, 0)
			for k:=0; k<6; k++ {
				si = append(si, s[i-k][j+k])
			}
			if f(si) {
				ans = true
			}
		}
	}

	for i:=0; i<n; i++ {
		for j:=0; j<n; j++ {
			if i + 5 >= n || j + 5 >= n {
				continue
			}
			si := make([]rune, 0)
			for k:=0; k<6; k++ {
				si = append(si, s[i+k][j+k])
			}
			if f(si) {
				ans = true
			}
		}
	}

	if ans {
		iou.Pl("Yes")
	} else {
		iou.Pl("No")
	}
}

func f(si []rune) bool {
	cnt := 0
	for _, r := range si {
		if r == '#' {
			cnt++
		}
	}
	if cnt >= 4 {
		return true
	}
	return false
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
