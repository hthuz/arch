package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {

	if k == 0 {
		return []int{target.Val}
	}
	if root == target {
		return distanceBelow(target, k)
	}

	height, isLeft := getHeight(root, target)
	fmt.Println(height, isLeft)

	var up []int
	if height > k {
		if isLeft {
			up = distanceBelow(root.Left, height-k-1)
		} else {
			up = distanceBelow(root.Right, height-k-1)
		}
	} else if height < k {
		if isLeft {
			up = distanceBelow(root.Right, k-height-1)
		} else {
			up = distanceBelow(root.Left, k-height-1)
		}
	} else {
		up = []int{root.Val}
	}

	below := distanceBelow(target, k)
	return append(up, below...)

}

func getHeight(root *TreeNode, target *TreeNode) (int, bool) {

	var lefts []*TreeNode
	var rights []*TreeNode
	if root.Left != nil {
		lefts = []*TreeNode{root.Left}
	}
	if root.Right != nil {
		rights = []*TreeNode{root.Right}
	}
	height := 1
	for len(lefts) != 0 {
		next_lefts := make([]*TreeNode, 0)
		for _, node := range lefts {
			if node == target {
				return height, true
			}
			if node.Left != nil {
				next_lefts = append(next_lefts, node.Left)
			}
			if node.Right != nil {
				next_lefts = append(next_lefts, node.Right)
			}

		}
		height += 1
		lefts = next_lefts
	}

	height = 1
	for len(rights) != 0 {
		next_rights := make([]*TreeNode, 0)
		for _, node := range rights {
			if node == target {
				return height, false
			}
			if node.Left != nil {
				next_rights = append(next_rights, node.Left)
			}
			if node.Right != nil {
				next_rights = append(next_rights, node.Right)
			}
		}
		height += 1
		rights = next_rights
	}
	return 0, false

}

func distanceBelow(node *TreeNode, k int) []int {
	if node == nil {
		return []int{}
	}
	if k == 0 {
		return []int{node.Val}
	}
	res := make([]int, 0)
	res = append(res, distanceBelow(node.Left, k-1)...)
	res = append(res, distanceBelow(node.Right, k-1)...)
	return res
}

func main() {

	target := &TreeNode{
		Val: 5,
		Left: &TreeNode{
			Val: 6,
		},
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 7,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}
	root := &TreeNode{
		Val:  3,
		Left: target,
		Right: &TreeNode{
			Val: 1,
			Left: &TreeNode{
				Val: 0,
			},
			Right: &TreeNode{
				Val: 8,
			},
		},
	}

	res := distanceK(root, target, 2)
	fmt.Println(res)

}
