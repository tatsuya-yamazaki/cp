package main

import "fmt"

func main() {
	var str string
	fmt.Scanf("%s", &str)

	var list []int
	for _, v := range str {
		list = append(list, int(v) - int('0'))
	}

	pmt := Permutation(list)

	var ans int
	length := len(list)

	for _, p := range pmt {
		for i := 0; i < length; i++ {
			var l, r int

			for j := 0; j < i; j++ {
				l = 10*l + p[j]
			}

			for j := i; j < length; j++ {
				r = 10*r + p[j]
			}

			if ans < l*r {
				ans = l*r
			}
		}
	}

	fmt.Println(ans)
}

func Permutation(list []int) (result [][]int) {
	if len(list) == 1 {
		return [][]int{list}
	}

	for i, former := range list {
		for _, latter := range Permutation(GetOtherElementsSlice(list, i)) {
			result = append(result, append([]int{former}, latter...))
		}
	}

	return result
}

func GetOtherElementsSlice (list []int, index int) (other []int) {
	length := len(list)
	switch length {
	case 0:
		fmt.Println(nil)
		return nil
	case 1:
		fmt.Println(nil)
		return nil
	}

	l := make([]int, length)
	_ = copy(l, list)
	max := length -1

	switch index {
	case 0:
		return l[1:]
	case max:
		return l[:max]
	}

	return append(l[:index], l[index+1:]...)
}
