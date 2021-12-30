package main
import "fmt"
func main() {
	var n, k int
	fmt.Scan(&n, &k)
	d := make([][]int, n)
	for i:=0; i<n; i++ {
		di := make([]int, n)
		for j:=0; j<n; j++ {
			fmt.Scan(&di[j])
		}
		d[i] = di
	}
	var other []int
	for i:=1; i<n; i++ {
		other = append(other, i)
	}

	ans := 0
	for {
		var route []int
		route = append(route, 0)
		for _, v := range other {
			route = append(route, v)
		}
		route = append(route, 0)
		sum := 0
		for i:=0; i<n; i++ {
			sum += d[route[i]][route[i+1]]
		}
		if sum == k {
			ans++
		}
		if !nextPermutation(other) {
			break
		}
	}

	fmt.Println(ans)
}

func nextPermutation(s []int) bool {
	l, b := -1, len(s)
	for i:=0; i<len(s)-1; i++ {
		if s[i] < s[i+1] {
			l = i
		}
	}
	if l == -1 {
		return false
	}
	for i:=len(s)-1; i>=0; i-- {
		if s[l] < s[i] {
			b = i
			break
		}
	}
	s[l], s[b] = s[b], s[l]
	for i:=l+1; i<=(l+len(s))/2; i++ {
		s[i], s[len(s)-i+l] = s[len(s)-i+l], s[i]
	}
	return true
}
