//给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。
//https://leetcode-cn.com/problems/binary-tree-level-order-traversal/

package leecode

import (
	"testing"
)

func Test_8(t *testing.T) {
	t.Log(levelOrder(&Node{}))
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	var lower []*Node
	tmp := []*Node{root}

	for len(tmp) != 0 {
		var line []int
		for _, node := range tmp {
			line = append(line, node.Val)
			if node.Left != nil {
				lower = append(lower, node.Left)
			}

			if node.Right != nil {
				lower = append(lower, node.Right)
			}
		}
		res = append(res, line)
		tmp = lower
		lower = []*Node{}
	}

	return res
}
