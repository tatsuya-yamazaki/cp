package main
import "fmt"
func main() {
	var s string
	fmt.Scan(&s)

	dp := make([]bool, len(s)+1, len(s)+1)
	dp[0] = true

	for i:=1; i<=len(s); i++ {
		if i >= 5 {
			switch s[i-5:i] {
				case "dream":
					if dp[i-5] {
						dp[i] = true
					}
				case "erase":
					if dp[i-5] {
						dp[i] = true
					}
			}
		}
		if i >= 6 {
			switch s[i-6:i] {
				case "eraser":
					if dp[i-6] {
						dp[i] = true
					}
			}
		}
		if i >= 7 {
			switch s[i-7:i] {
				case "dreamer":
					if dp[i-7] {
						dp[i] = true
					}
			}
		}
	}

	if dp[len(s)] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
