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
        m := iou.I()
        a := iou.Is(n + 1)
        c := iou.Is(n + m + 1)

        var b []int

        for i:=m; i>=0; i-- {
                ci := c[i+n]
                for j:=m; j>i; j-- {
                        k := n - (j - i)
                        l := m - j
                        if k < 0 {
                                continue
                        }
                        ci -= a[k] * b[l]
                }
                b = append(b, ci / a[n])
        }

        for i:=len(b)-1; i>=0; i-- {
                iou.P(b[i])
                if i > 0 {
                        iou.P(" ")
                } else {
                        iou.Pl()
                }
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
