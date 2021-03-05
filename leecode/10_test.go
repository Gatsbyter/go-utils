//给你一个二维整数数组 envelopes ，其中 envelopes[i] = [wi, hi] ，表示第 i 个信封的宽度和高度
//当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样
//请计算 最多能有多少个 信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）
//https://leetcode-cn.com/problems/russian-doll-envelopes/

package leecode

import (
	"sort"
	"testing"
)

var envelops = [][]int{
	{5, 4},
	{6, 4},
	{7, 6},
	{2, 3},
}

func Test_10(t *testing.T) {
	t.Log(maxEnvelopes(envelops))
}

func maxEnvelopes(envelopes [][]int) int {
	max := 0

	points := NewPoints(envelopes)
	sort.Sort(points)

	for i := 0; i < len(points); i++ {
		tempMax := 1
		for j := 0; j < i; j++ {
			if points[i].y > points[j].y && points[i].x > points[j].x && points[j].max+1 > tempMax {
				tempMax = points[j].max + 1
			}
		}

		points[i].max = tempMax

		if points[i].max > max {
			max = points[i].max
		}
	}

	return max
}

type Points []Point

type Point struct {
	x, y, max int
}

func NewPoints(envelopes [][]int) Points {
	var points Points

	for _, en := range envelopes {
		points = append(points, Point{
			x:   en[0],
			y:   en[1],
			max: 1,
		})
	}

	return points
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
