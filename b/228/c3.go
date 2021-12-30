package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"math/rand"
)

func main() {
	var n, k int
	var al, bl []int

	fmt.Scan(&n, &k)

	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	for i:=0; i<n; i++ {

		sum := 0
		for j:=0; j<3; j++ {
			sc.Scan()
			t := sc.Text()
			a, _ := strconv.Atoi(t)
			sum += a
		}

		al = append(al, sum)
		bl = append(bl, sum)
	}

	target := getTarget(bl, k)

	for _, v := range al {
		if target <= v + 300 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}

}

func getTarget(a []int, target int) int {
	selects(a, 0, len(a)-1, target-1)
	return a[target-1]
}

func selects(a []int, left, right, target int) {

	pivotIndex := rand.Intn(right-left+1) + left

	swapIndex := partition(a, left, right, pivotIndex)

	if swapIndex > target {
		selects(a, left, swapIndex-1, target)
	} else if swapIndex < target {
		selects(a, swapIndex+1, right, target)
	}
}

func partition(a []int, left, right, pivotIndex int) int {

	pivot := a[pivotIndex]
	a[right], a[pivotIndex] = a[pivotIndex], a[right]
	swapIndex := left

	for i:=left; i<right; i++ {
		if a[i] >= pivot {
			a[swapIndex], a[i] = a[i], a[swapIndex]
			swapIndex++
		}
	}
	a[swapIndex], a[right] = a[right], a[swapIndex]
	return swapIndex

}
