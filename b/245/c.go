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
        k := iou.I()
        a := iou.Is(n)
        b := iou.Is(n)

        var dp [200001][2]bool
        dp[0][0] = true
        dp[0][1] = true

        for i:=0; i<n-1; i++ {
                if dp[i][0] && Abs(a[i] - a[i+1]) <= k {
                        dp[i+1][0] = true
                } else if dp[i][1] && Abs(b[i] - a[i+1]) <= k {
                        dp[i+1][0] = true
                }
                if dp[i][0] && Abs(a[i] - b[i+1]) <= k {
                        dp[i+1][1] = true
                } else if dp[i][1] && Abs(b[i] - b[i+1]) <= k {
                        dp[i+1][1] = true
                }
        }


        if dp[n-1][0] || dp[n-1][1] {
                iou.Pl("Yes")
        } else {
                iou.Pl("No")
        }
}

func Abs(a int) int {
        if a < 0 {
                return -a
        }
        return a
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
