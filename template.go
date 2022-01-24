package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
)

type IOUtil struct {
	Scanner bufio.Scanner
}

func NewIOUtil() *IOUtil {
	iou := IOUtil{
		Scanner: bufio.NewScanner(os.Stdin)
	}
	iou.Scanner.Split(bufio.ScanWords)
	return &iou
}

func (*IOUtil) ToInt(s string) int {
	ret, _ := strconv.Atoi(s)
	return ret
}

func (iou *IOUtil) Int() int {
}

func main() {
}
