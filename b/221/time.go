package main

import "fmt"
import "time"

func main() {
	list := make([]int, 9999)

	now1 := time.Now()
	for i, _ := range list {
		_ = GetOtherElementsSlice(list, i)
	}
	fmt.Printf("経過1: %vms\n", time.Since(now1).Milliseconds())

	now2 := time.Now()
	for i, _ := range list {
		_ = GetOtherElementsSlice2(list, i)
	}
	fmt.Printf("経過2: %vms\n", time.Since(now2).Milliseconds())
}

func GetOtherElementsSlice (list []int, index int) []int {
	length := len(list)

	switch length {
	case 0:
		return nil
	case 1:
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

func GetOtherElementsSlice2 (l []int, index int) (other []int) {
	for i, v := range l {
		if i == index {
			continue
		}
		other = append(other, v)
	}
	return other
}
