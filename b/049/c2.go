package main
import "fmt"
import "strings"
func main() {
	var s string
	fmt.Scan(&s)
	s1 := ""
	for i:=len(s)-1; i>=0; i-- {
		s1 += string(s[i])
	}

	eraser := strings.Split(s1, "resare")
	s2 := ""
	for _, r := range eraser {
		s2 += r
	}

	dreamer := strings.Split(s2, "esare")
	s3 := ""
	for _, r := range dreamer {
		s3 += r
	}

	erase := strings.Split(s3, "remaerd")
	s4 := ""
	for _, r := range erase {
		s4 += r
	}

	dream := strings.Split(s4, "maerd")
	s5 := ""
	for _, r := range dream {
		s5 += r
	}
	if s5 == "" {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
