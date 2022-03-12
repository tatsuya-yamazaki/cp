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

	t := NewAvl()
	q := iou.I()
	m := make(map[int]int)
	for i:=0; i<q; i++ {
		qi := iou.S()
		if qi == "1" {
			x := iou.I()
			if t.Add(x) {
				m[x] = 1
			} else {
				m[x]++
			}
		} else if qi == "2" {
			x := iou.I()
			k := iou.I()
			n, route := t.FindSmaller(x)
				iou.Pl(-1)
		} else {
			x := iou.I()
			k := iou.I()
			n, route := t.FindBigger(x)
		}
	}
}

type Avl struct {
	root *Node
}

func NewAvl () *Avl {
	return &Avl{}
}

type Node struct {
	value int
	height int
	right *Node
	left *Node
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
		height: 1,
	}
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Height() int {
	if n != nil {
		return n.height
	}
	return 0
}

func (n *Node) getBalance() int {
	return n.right.Height() - n.left.Height()
}

func (n *Node) updateHeight() int {
	left := n.left.Height() + 1
	right := n.right.Height() + 1
	if left > right {
		n.height = left
	} else {
		n.height = right
	}
	return n.height
}

func (t *Avl) FindBigger(value int) (n *Node, route []*Node) {
	n = t.root
	for n != nil {
		route = append(route, n)
		if value == n.value {
			return n, route
		} else if value > n.value {
			return n, route
		} else {
			n = n.left
		}
	}
	return nil, route
}

func (t *Avl) FindSmaller(value int) (n *Node, route []*Node) {
	n = t.root
	for n != nil {
		route = append(route, n)
		if value == n.value {
		} else if value > n.value {
			n = n.right
		} else {
			return n, route
		}
	}
	return nil, route
}

func (t *Avl) Find(value int) (n *Node, route []*Node) {
	n = t.root
	for n != nil {
		route = append(route, n)
		if value == n.value {
			return n, route
		} else if value > n.value {
			n = n.right
		} else {
			n = n.left
		}
	}
	return nil, route
}

func (t *Avl) Add(value int) bool {
	if t.root == nil {
		t.root = NewNode(value)
		return true
	}
	_, route := t.Find(value)
	parent := route[len(route)-1]
	if parent.value == value {
		return false
	}
	nn := NewNode(value)
	if value > parent.value {
		parent.right = nn
	} else {
		parent.left = nn
	}
	t.balance(route, true)
	return true
}

func (t *Avl) Remove(value int) bool {
	if t.root == nil {
		return false
	}
	n, route := t.Find(value)
	if n == nil {
		return false
	}
	parent := t.getParentFromRoute(route)
	if n.left == nil && n.right == nil {
		t.removeNodeHasNoChild(n, parent)
		route = route[:len(route)-1]
	} else if n.right == nil {
		t.removeNodeHasChild(n, parent, n.left)
		route = route[:len(route)-1]
	} else if n.left == nil {
		t.removeNodeHasChild(n, parent, n.right)
		route = route[:len(route)-1]
	} else {
		additionalRoute := t.removeNodeHasChildren(n)
		for i:=0; i<len(additionalRoute)-1; i++ {
			route = append(route, additionalRoute[i])
		}
	}
	t.balance(route, false)
	return true
}

func (*Avl) getParentFromRoute(route []*Node) *Node {
	if len(route) > 1 {
		return route[len(route)-2]
	}
	return nil
}

func (t *Avl) removeNodeHasChildren(n *Node) []*Node {
	leftMax, route := t.Max(n.left)
	leftMaxParent := n
	if leftMax != n.left {
		leftMaxParent = t.getParentFromRoute(route)
	}
	n.value = leftMax.value
	if leftMax.left == nil {
		t.removeNodeHasNoChild(leftMax, leftMaxParent)
	} else {
		t.removeNodeHasChild(leftMax, leftMaxParent, leftMax.left)
	}
	return route
}

func (t *Avl) removeNodeHasChild(n, parent, child *Node) {
	t.replaceNode(n, parent, child)
	child = nil
}

func (t *Avl) removeNodeHasNoChild(n, parent *Node) {
	t.replaceNode(n, parent, nil)
}

func (t *Avl) replaceNode(n, parent, newNode *Node) {
	if parent != nil {
		if parent.left == n {
			parent.left = newNode
		} else {
			parent.right = newNode
		}
	}
	if n == t.root {
		t.root = newNode
	}
}

func (t *Avl) Max(n *Node) (max *Node, route []*Node) {
	if t.root == nil {
		return nil, nil
	}
	if n == nil {
		n = t.root
	}
	route = append(route, n)
	for n.right != nil {
		n = n.right
		route = append(route, n)
	}
	return n, route
}

func (t *Avl) Min(n *Node) (min *Node, route []*Node) {
	if t.root == nil {
		return nil, nil
	}
	if n == nil {
		n = t.root
	}
	route = append(route, n)
	for n.left != nil {
		n = n.left
		route = append(route, n)
	}
	return n, route
}

func (t *Avl) Echo() {
	if t.root == nil {
		fmt.Println("nil")
		return
	}
	t.echo(t.root, "")
}

func (t *Avl) echo(n *Node, space string) {
	space += "    "
	if n.right != nil {
		t.echo(n.right, space)
	}
	fmt.Println(space, n.value)
	if n.left != nil {
		t.echo(n.left, space)
	}
}

func (t *Avl) balance(route []*Node, isAdd bool) {
	for i:=len(route)-1; i>=0; i-- {
		n := route[i]
		var parent *Node
		if i != 0 {
			parent = route[i-1]
		}
		n.updateHeight()
		switch n.getBalance() {
		case -1:
			if !isAdd { return }
		case -2:
			if n.left.getBalance() > 0 {
				t.rotateLR(n, parent, n.left)
			} else {
				t.rotateR(n, parent, n.left)
			}
			if isAdd { return }
		case 0:
			if isAdd { return }
		case 1:
			if !isAdd { return }
		case 2:
			if n.right.getBalance() < 0 {
				t.rotateRL(n, parent, n.right)
			} else {
				t.rotateL(n, parent, n.right)
			}
			if isAdd { return }
		}
	}
}

func (t *Avl) setPivotAsParentsChild(n, parent, pivot *Node) {
	if parent == nil {
		t.root = pivot
		return
	}
	if parent.right == n {
		parent.right = pivot
	} else {
		parent.left = pivot
	}
}

func (t *Avl) rotateL(n, parent, pivot *Node) {
	n.right = pivot.left
	pivot.left = n
	n.updateHeight()
	pivot.updateHeight()
	t.setPivotAsParentsChild(n, parent, pivot)
}

func (t *Avl) rotateR(n, parent, pivot *Node) {
	n.left = pivot.right
	pivot.right = n
	n.updateHeight()
	pivot.updateHeight()
	t.setPivotAsParentsChild(n, parent, pivot)
}

func (t *Avl) rotateLR(n, parent, pivot *Node) {
	t.rotateL(pivot, n, pivot.right)
	t.rotateR(n, parent, n.left)
}

func (t *Avl) rotateRL(n, parent, pivot *Node) {
	t.rotateR(pivot, n, pivot.left)
	t.rotateL(n, parent, n.right)
}

func (t *Avl) Repl() {
	for {
		s := ""
		n := 0
		fmt.Scan(&s)
		switch s {
		case "a":
			fmt.Scan(&n)
			t.Add(n)
			t.Echo()
		case "r":
			fmt.Scan(&n)
			t.Remove(n)
			t.Echo()
		case "f":
			fmt.Scan(&n)
			fmt.Println(t.Find(n))
		case "p":
			t.Echo()
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
