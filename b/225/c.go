package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main() {
	var(
		N, M int
	)

	fmt.Scanf("%d %d", &N, &M)

	nl := make([][]int, N, N)
	for i:=0; i<N; i++{
		r := make([]int,M ,M)
		nl[i] = r
	}

	sc := bufio.NewScanner(os.Stdin)
	for i:=0; i<N; i++{
		sc.Scan()
		text := sc.Text()
		strs := strings.Split(text, " ")
		for j:=0; j<M; j++{
			num, _ := strconv.Atoi(strs[j])
			nl[i][j] = num
		}
	}

	if err := sc.Err(); err != nil {
		panic(err)
	}

	ans := true
	for i:=0; i<N; i++{
		if i != 0 {
			if (nl[i][0] - nl[i-1][0]) != 7 {
				ans = false
			}
		}
		for j:=0; j<M-1; j++{
			if M == 1 {
				break
			}
			if (nl[i][j+1] - nl[i][j]) != 1 {
				ans = false
			}
		}
		if !ans {
			break
		}
	}

	num := 1
	end := false
	if !ans {
		end = true
	}
	for {
		if num > nl[0][0] {
			ans = false
			end = true
		}
		for j:=0; j<7; j++{
			if nl[0][0] == (num + j) && (j + M) <= 7 {
				end = true
			}
		}

		if end {
			break
		}
		num += 7
	}

	if ans {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
