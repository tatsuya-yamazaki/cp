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
	d := iou.Is(n)
	a, b := iou.Is2(m)

	u := NewUnionFind(n+1)

	for i:=0; i<n; i++ {
		u.Unite(a[i], b[i])
		d[a[i]]--
		d[b[i]]--
	}
	for i:=0; i<n; i++ {
		d[i]
	}

	for i:=1; i<=n; i++ {
		if !u.SameRoot(1, i) {
			iou.Pl(-1)
			return
		}
	}
	iou.Pl()
}

type UnionFind struct {
	parent []int
	rank []int
}

func NewUnionFind(length int) *UnionFind {
	parent := make([]int, length)
	rank := make([]int, length)
	for i:=0; i<length; i++ {
		parent[i] = i
	}
	return &UnionFind{parent, rank}
}

func (u *UnionFind) Root(index int) int {
	if u.parent[index] == index {
		return index
	} else {
		u.parent[index] = u.Root(u.parent[index])
		return u.parent[index]
	}
}

func (u *UnionFind) SameRoot(a, b int) bool {
	return u.Root(a) == u.Root(b)
}

func (u *UnionFind) Unite(a, b int) {
	a = u.Root(a)
	b = u.Root(b)

	if a == b {
		return
	}

	if u.rank[a] < u.rank[b] {
		u.parent[a] = b
	} else {
		u.parent[b] = a
		if u.rank[a] == u.rank[b] {
			u.rank[a]++
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
