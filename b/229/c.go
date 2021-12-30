package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

type Ch struct {
	Value int
	Weight int
}

func main() {
	var n, w int
	var ch []Ch

	fmt.Scan(&n, &w)
	sc := bufio.NewScanner(os.Stdin)

	for i:=1; i<=n; i++ {
			sc.Scan()
			t := sc.Text()
			tl := strings.Split(t, " ")
			var c Ch
			c.Value, _ = strconv.Atoi(tl[0])
			c.Weight, _ = strconv.Atoi(tl[1])
			ch = append(ch, c)
	}

	sort.Slice(ch, func(i,j int) bool { return ch[i].Value > ch[j].Value })

	ans := 0
	weight := 0
	for _, c := range ch {
		for c.Weight > 0 && weight < w {
			ans += c.Value
			c.Weight--
			weight++
		}
		if weight == w {
			break
		}
	}
	fmt.Println(ans)

}
