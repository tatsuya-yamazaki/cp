package main

import(
	"fmt"
	"sort"
)

type Person struct{
	Id int
	Point int
	Hand []rune
}

func main() {
	var(
		N, M int
		H string
		pl []Person
	)
	fmt.Scanf("%d %d", &N, &M)
	for i := 0; i < 2*N; i++ {
		fmt.Scanf("%s", &H)
		pl = append(pl, Person{
			Id: i,
			Point: 0,
			Hand: []rune(H),
		})
	}

	for round := 0; round < M; round++ {
		for k := 0; k < N; k++ {
			result := jkp(pl[2*k].Hand[round], pl[2*k+1].Hand[round])
			switch result{
			case 1:
				pl[2*k].Point++
			case 2:
				pl[2*k+1].Point++
			}
		}
		sort.Slice(pl, func (i,j int) bool { return pl[i].Id < pl[j].Id })
		sort.SliceStable(pl, func (i,j int) bool { return pl[i].Point > pl[j].Point})
	}

	for _, p := range pl {
		fmt.Printf("%d\n", p.Id+1)
	}
}

func jkp(a,b rune) int {
	if a == 'G' && b == 'G' {
		return 0
	} else if a == 'G' && b == 'C' {
		return 1
	} else if a == 'G' && b == 'P' {
		return 2
	} else if a == 'C' && b == 'C' {
		return 0
	} else if a == 'C' && b == 'P' {
		return 1
	} else if a == 'C' && b == 'G' {
		return 2
	} else if a == 'P' && b == 'P' {
		return 0
	} else if a == 'P' && b == 'G' {
		return 1
	} else if a == 'P' && b == 'C' {
		return 2
	}
	return 0
}

