package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

func main() {
	defer iou.Fl()
	n := iou.I()
	k := iou.I()
	a := iou.Is(n)

	cnt := 1
	pos := 1
	m := make(map[int]int)
	m[1] = 0
	for {
		t := a[pos-1]
		_, ok := m[t]
		if ok {
			break
		} else {
			m[t] = cnt
			pos = t
			cnt++
		}
	}

	loopFirst := a[pos-1]
	cntBeforeLoop := m[loopFirst]
	cntLoopEnd := m[pos]
	loopLen := cntLoopEnd - cntBeforeLoop + 1

	if k <= cntBeforeLoop {
		for key, v := range m {
			if v == k {
				iou.Pl(key)
				return
			}
		}
	} else {
		k -= cntBeforeLoop
		r := k % loopLen
		pos := loopFirst
		cntp := 0
		for {
			if cntp == r {
				iou.Pl(pos)
				break
			}
			pos = a[pos-1]
			cntp++
		}
	}

}

type IOUtil struct {
	Scanner *bufio.Scanner
	Writer *bufio.Writer
}
func NewIOUtil() *IOUtil {
	iou := IOUtil{
		Scanner: bufio.NewScanner(os.Stdin),
		Writer: bufio.NewWriter(os.Stdout),
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

