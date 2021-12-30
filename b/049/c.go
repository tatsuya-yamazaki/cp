package main
import "fmt"
import "strings"
func main() {
	var s string
	fmt.Scan(&s)
	eraser := strings.Split(s, "eraser")
	s1 := ""
	for _, r := range eraser {
		s1 += r
	}
	dreamer := strings.Split(s1, "erase")
	s2 := ""
	for _, r := range dreamer {
		s2 += r
	}
	erase := strings.Split(s2, "dreamer")
	s3 := ""
	for _, r := range erase {
		s3 += r
	}
	dream := strings.Split(s3, "dream")
	s4 := ""
	for _, r := range dream {
		s4 += r
	}
	if s4 == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
