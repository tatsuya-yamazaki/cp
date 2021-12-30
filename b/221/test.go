package main

import "fmt"
import "reflect"

func main() {
	var list []int

	for i := 0; i < 10; i++ {
		list = append(list, i)

		for j := 0; j < len(list); j++ {
			a1 := GetOtherElementsSlice(list, j)
			a2 := GetOtherElementsSlice2(list, j)
			if !reflect.DeepEqual(a1,a2) {
				fmt.Println("list")
				fmt.Println(list)
				fmt.Println("j")
				fmt.Println(j)
				fmt.Println("a1")
				fmt.Println(a1)
				fmt.Println("a2")
				fmt.Println(a2)
			}
		}
	}
}

func GetOtherElementsSlice (list []int, index int) []int {

	length := len(list)
	switch length {
	case 0:
		return nil
	case 1:
		return nil
	}

	l := make([]int, len(list))
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

