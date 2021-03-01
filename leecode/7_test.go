//给定一个二叉树的根节点 root ，返回它的 中序 遍历。
//https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

package leecode

import (
	"testing"
)

func Test_7(t *testing.T) {
	t.Log(inorderTraversal(&Node{}))
}

func inorderTraversal(root *Node) []int {
	if root == nil {
		return []int{}
	}

	return inorder(root)
}

func inorder(node *Node) []int {
	var left, right []int

	if node.Left != nil {
		left = inorder(node.Left)
	}

	if node.Right != nil {
		right = inorder(node.Right)
	}

	return append(append(left, node.Val), right...)
}
