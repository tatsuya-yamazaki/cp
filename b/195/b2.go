package main
import "fmt"
func main() {
	var a, b, w, max, min int
	fmt.Scan(&a, &b, &w)
	w *= 1000
	min = 1000001

	for i:=1; i<=1000000; i++ {
		if a*i <= w && w <= b*i {
			if i < min {
				min = i
			}
			if i > max {
				max = i
			}
		}
	}

	if max == 0 {
		fmt.Println("UNSATISFIABLE")
		return
	} else {
		fmt.Println(min, max)
	}
}
