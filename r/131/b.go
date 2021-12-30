package main

import(
	"fmt"
)

func main() {
	var h, w  int

	fmt.Scan(&h)
	fmt.Scan(&w)

	t := make([][]int, h, h)
	for i:=0; i<h; i++ {
		r := make([]int, w, w)
		var s string
		fmt.Scan(&s)
		for j:=0; j<w; j++ {
			if string(s[j]) != "." {
				r[j] = int([]rune(string(s[j]))[0])-48
			}
		}
		t[i] = r
	}

	for c:=-1; c>-6; c-- {
		for i:=0; i<h; i++ {
			for j:=0; j<w; j++ {
				if t[i][j] < 1 {
					canUse := true
					if i != 0 {
						if abs(t[i-1][j]) == abs(c) {
							canUse = false
						}
					}
					if j != 0 {
						if abs(t[i][j-1]) == abs(c) {
							canUse = false
						}
					}
					if i != h-1 {
						if abs(t[i+1][j]) == abs(c) {
							canUse = false
						}
					}
					if j != w-1 {
						if abs(t[i][j+1]) == abs(c) {
							canUse = false
						}
					}
					if canUse && t[i][j] == 0 {
						t[i][j] = c
					}
				}
			}
		}
	}

	for i:=0; i<h; i++ {
		for j:=0; j<w; j++ {
			fmt.Printf("%d",abs(t[i][j]))
		}
		fmt.Printf("\n")
	}
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
