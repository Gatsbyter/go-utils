package tree

import stack "github.com/Gatsbyter/go-utils/data_struct/stack"

//前序遍历：根结点 ---> 左子树 ---> 右子树
//
//中序遍历：左子树---> 根结点 ---> 右子树
//
//后序遍历：左子树 ---> 右子树 ---> 根结点

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

/* ------------------------  递归  ------------------------------ */

func prevOrder(root *Node) []int {
	if root != nil {
		return append(append([]int{root.Val}, prevOrder(root.Left)...), prevOrder(root.Right)...)
	}
	return []int{}
}

func inOrder(root *Node) []int {
	if root != nil {
		return append(append(inOrder(root.Left), root.Val), inOrder(root.Right)...)
	}
	return []int{}
}

func postOrder(root *Node) []int {
	if root != nil {
		return append(append(postOrder(root.Left), postOrder(root.Right)...), root.Val)
	}
	return []int{}
}

/* ------------------------  迭代  ------------------------------ */

func prevOrderLoop(root *Node) (res []int) {
	s := stack.NewStack()
	pNode := root

	for pNode != nil || s.Size() != 0 {
		if pNode != nil {
			res = append(res, pNode.Val)
			s.Push(pNode)
			pNode = pNode.Left
		} else {
			node := s.Pop().(*Node)
			pNode = node.Right
		}
	}

	return res
}

func inOrderLoop(root *Node) (res []int) {
	s := stack.NewStack()
	pNode := root

	for pNode != nil || s.Size() != 0 {
		if pNode != nil {
			s.Push(pNode)
			pNode = pNode.Left
		} else {
			node := s.Pop().(*Node)
			res = append(res, node.Val)
			pNode = node.Right
		}
	}

	return res
}
