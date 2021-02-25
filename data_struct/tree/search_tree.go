package tree

import "fmt"

type SNode struct {
	Val   int32
	nodes []*SNode
}

func BuildTree(dic []string) *SNode {
	root := &SNode{
		Val:   0,
		nodes: []*SNode{},
	}

	for _, str := range dic {
		tmp := root
		for _, ch := range str {
			flag := false
			for i, node := range tmp.nodes {
				if node.Val == ch {
					tmp = tmp.nodes[i]
					flag = true
					break
				}
			}

			if !flag {
				tmp.nodes = append(tmp.nodes, &SNode{
					Val:   ch,
					nodes: []*SNode{},
				})
				tmp = tmp.nodes[len(tmp.nodes)-1]
			}
		}
	}

	return root
}

func Match(root *SNode, str string) bool {
	for _, ch := range str {
		flag := false
		for i, node := range root.nodes {
			if node.Val == ch {
				root = root.nodes[i]
				flag = true
				break
			}
		}

		if !flag {
			return false
		}
	}

	return len(root.nodes) == 0
}

func PrintTree(root *SNode) {
	if root == nil {
		return
	}

	for _, node := range root.nodes {
		fmt.Printf("%c, ", node.Val)
		PrintTree(node)
	}
	fmt.Println()
}
