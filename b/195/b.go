package main
import "fmt"
func main() {
	var a, b, w, max, min int
	fmt.Scan(&a, &b, &w)
	w *= 1000

	max = w / a
	for max > 0 {
		max--
		rest := w - a*max
		count := 1
		for {
			weight := float64(rest) / float64(count)
			if float64(a) <= weight && weight <= float64(b) {
				break
			}
			if weight < float64(a) {
				break
			}
			count++
		}
		weight := float64(rest) / float64(count)
		if float64(a) <= weight && weight <= float64(b) {
			max += count
			break
		}
	}
	if max == 0 {
		fmt.Println("UNSATISFIABLE")
		return
	}

	min = w / b
	for min < max {
		rest := w - b*min
		count := 1
		for {
			weight := float64(rest) / float64(count)
			if float64(a) <= weight && weight <= float64(b) {
				break
			}
			if weight < float64(a) {
				break
			}
			count++
		}
		weight := float64(rest) / float64(count)
		if float64(a) <= weight && weight <= float64(b) {
			min += count
			break
		}
		min--
	}
	fmt.Println(min, max)
}
