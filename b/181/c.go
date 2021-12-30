package main
import "fmt"
func main() {
	var n int
	fmt.Scan(&n)
	var x, y []int
	for i:=0; i<n; i++ {
		var xi,yi int
		fmt.Scan(&xi, &yi)
		x = append(x, xi)
		y = append(y, yi)
	}

	for i:=0; i<n; i++ {
		for j:=i+1; j<n; j++ {
			for k:=j+1; k<n; k++ {
				if i == j || j == k || k == i {
					continue
				}
				if area(x[i],y[i],x[j],y[j],x[k],y[k]) == 0 {
					fmt.Println("Yes")
					return
				}
			}
		}
	}
	fmt.Println("No")
}

func area(ax,ay,bx,by,cx,cy int) float64 {
	return float64(abs((ax-cx)*(by-cy)-(bx-cx)*(ay-cy))) / 2
}

func abs(n int) int {
	if n > 0 {
		return n
	}
	return -n
}
