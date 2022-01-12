package main
import "fmt"
func main() {
	var x int
	fmt.Scan(&x)
	// a**5 - b**5 が,bが最大でもxを超えるなら、それ以降のa,bは計算しなくてよい
	//よって、全探索で十分間に合う
	for a:=-120; a<120; a++ {
		for b:=-120; b<120; b++ {
			if a*a*a*a*a - b*b*b*b*b == x {
				fmt.Println(a, b)
				return
			}
		}
	}
}
