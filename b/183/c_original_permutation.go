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
	for _, os := range permutation(other) {
		var route []int
		route = append(route, 0)
		for _, o := range os {
			route = append(route, o)
		}
		route = append(route, 0)

		sum := 0
		for i:=0; i<n; i++ {
			sum += d[route[i]][route[i+1]]
		}
		if sum == k {
			ans++
		}
	}

	fmt.Println(ans)
}

func permutation(ns []int) (ret [][]int) {
	if len(ns) == 1 {
		ret = append(ret, ns)
		return ret
	}
	for i, v := range ns {
		ns2 := make([]int, len(ns))
		copy(ns2[:i], ns[:i])
		copy(ns2[i:], ns[i+1:])
		ns2 = ns2[:len(ns2)-1]
		for _, v2 := range permutation(ns2) {
			var r []int
			r = append(r, v)
			for _, v3 := range v2 {
				r = append(r, v3)
			}
			ret = append(ret, r)
		}
	}
	return ret
}
