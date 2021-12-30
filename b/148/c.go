package main
import(
	"fmt"
)

func main() {
	var a, b int
	fmt.Scan(&a, &b)

	af := factor(a)
	bf := factor(b)
	ans := leastCommonMultiple(af,bf)
	fmt.Println(ans)
}

func factor(num int) (ret []int) {
	n := num
	divisor := 2
	for divisor <= num {
		factorNum := 0
		for n % divisor == 0 {
			n /= divisor
			factorNum++
		}
		ret = append(ret, factorNum)
		divisor++
	}
	return ret
}

func leastCommonMultiple(a, b []int) int {
	if len(a) > len(b) {
		a, b = b, a
	}

	var lcm []int
	for i, bf := range b {
		af := 0
		if len(a) > i {
			af = a[i]
		}
		var factor int
		if af > bf {
			factor = af
		} else {
			factor = bf
		}
		lcm = append(lcm, factor)
	}

	ret, factor := 1, 1
	for _, v := range lcm {
		factor++
		for i:=0; i<v; i++ {
			ret *= factor
		}
	}
	return ret
}
