package main

import "fmt"
import "time"

func main() {
	l := []int{0,1,2,3,4,5,6,7,8,9,0}
	now := time.Now()
	ans := Permutation(l)
	fmt.Println(ans)
	fmt.Printf("passed: %vms\n", time.Since(now).Milliseconds())
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

func GetOtherElementsSlice (list []int, index int) []int {
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
