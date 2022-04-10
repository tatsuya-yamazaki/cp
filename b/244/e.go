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
	s := iou.I()
	t := iou.I()
	x := iou.I()
	u, v := iou.Is2(m)

	r := make(map[int][]int)
	var dp [2001][2001][2]*ModInt

	for i:=0; i<m; i++ {
		ui, vi := u[i], v[i]

		r[ui] = append(r[ui], vi)
		r[vi] = append(r[vi], ui)
	}

	d := 998244353
	dp[0][s][0] = NewModInt(d)
	dp[0][s][0].Add(1)

	for i:=0; i<k; i++ {
		for j:=1; j<=n; j++ {
			for _, v := range r[j] {
				if dp[i][j][0] == nil {
					dp[i][j][0] = NewModInt(d)
				}
				if dp[i][j][1] == nil {
					dp[i][j][1] = NewModInt(d)
				}
				if dp[i+1][v][0] == nil {
					dp[i+1][v][0] = NewModInt(d)
				}
				if dp[i+1][v][1] == nil {
					dp[i+1][v][1] = NewModInt(d)
				}

				if v == x {
					dp[i+1][v][0].Add(dp[i][j][1].Get())
					dp[i+1][v][1].Add(dp[i][j][0].Get())
				} else {
					dp[i+1][v][0].Add(dp[i][j][0].Get())
					dp[i+1][v][1].Add(dp[i][j][1].Get())
				}
			}
		}
	}

	if dp[k][t][0] == nil {
		iou.Pl(0)
	} else {
		iou.Pl(dp[k][t][0].Get())
	}
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
