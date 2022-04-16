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
	k := iou.I()

	var dp [51][2501]*ModInt

	mod := 998244353
	dp[0][k] = NewModInt(mod)
	dp[0][k].Add(1)
	for i:=1; i<=n; i++ {
		for j:=1; j<=k; j++ {
			if dp[i-1][j] == nil || dp[i-1][j].Get() == 0 {
				continue
			}
			for l:=1; l<=m; l++ {
				if j - l < 0 {
					continue
				}
				if dp[i][j-l] == nil {
					dp[i][j-l] = NewModInt(mod)
				}
				dp[i][j-l].Add(dp[i-1][j].Get())
			}
		}
	}

	ans := NewModInt(mod)
	for i:=0; i<=k; i++ {
		if dp[n][i] == nil {
			continue
		}
		ans.Add(dp[n][i].Get())
	}
	iou.Pl(ans.Get())
}

type ModInt struct {
        mod, value int
}

func NewModInt(mod int) *ModInt {
        return &ModInt{mod: mod}
}

func (m *ModInt) Get() int {
        return m.value
}

func (m *ModInt) Add(x int) {
        m.value = (m.value + x % m.mod) % m.mod
}

func (m *ModInt) Mul(x int) {
        m.value = (m.value * (x % m.mod)) % m.mod
}

func (m *ModInt) Sub(x int) {
        m.value = m.value - x % m.mod
        if m.value < 0 {
                m.value += m.mod
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
