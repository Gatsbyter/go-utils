//给定一个仅包含 0 和 1 、大小为 rows x cols 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。
//https://leetcode-cn.com/problems/validate-binary-search-tree/

package leecode

import (
	"testing"
)

func Test_6(t *testing.T) {
	t.Log(isValidBST(&Node{}))
}

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func isValidBST(root *Node) bool {
	if root == nil {
		return true
	}

	ok, _, _ := is(root)

	return ok
}

func is(root *Node) (bool, int, int) {
	var (
		tmpMax int
		tmpMin int
	)

	if root.Left != nil {
		ok, max, min := is(root.Left)
		if !ok || max >= root.Val {
			return false, 0, 0
		}

		tmpMin = Min(min, root.Val)
	} else {
		tmpMin = root.Val
	}

	if root.Right != nil {
		ok, max, min := is(root.Right)
		if !ok || min <= root.Val {
			return false, 0, 0
		}

		tmpMax = Max(max, root.Val)
	} else {
		tmpMax = root.Val
	}

	return true, tmpMax, tmpMin
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
