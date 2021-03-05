package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	p := Points{
		{1, 5},
		{6, 2},
		{4, 5},
		{2, 8},
		{9, 7},
	}

	fmt.Println(p)
	sort.Sort(p)
	fmt.Println(p)
}

type Points []Point

type Point struct {
	x, y int
}

func (p Points) Len() int {
	return len(p)
}

func (p Points) Less(i, j int) bool {
	return p[i].x < p[j].x
}

func (p Points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}
