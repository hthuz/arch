package main

import "fmt"

type Node struct {
	val   int
	left  *Node
	right *Node
}

func (n *Node) Level_traverse() {
	queue := &Queue{}
	queue.push(n)
	level_traverse(queue)
}

func level_traverse(queue *Queue) {
	if queue.empty() {
		fmt.Println()
		return
	}
	node := queue.pop()
	fmt.Print(node.val, " ")
	if node.left != nil {
		queue.push(node.left)
	}
	if node.right != nil {
		queue.push(node.right)
	}
	level_traverse(queue)
}

type Queue struct {
	items []*Node
}

func (q *Queue) push(item *Node) {
	q.items = append(q.items, item)
}

func (q *Queue) pop() (item *Node) {
	item = q.items[0]
	q.items = q.items[1:]
	return
}

func (q *Queue) empty() bool {
	return len(q.items) == 0
}

func main() {
	root := &Node{
		val: 1,
		left: &Node{
			val: 3,
			left: &Node{
				val:   5,
				left:  nil,
				right: nil,
			},
			right: &Node{
				val:   4,
				left:  nil,
				right: nil,
			},
		},
		right: &Node{
			val:   2,
			left:  nil,
			right: nil,
		},
	}
	root.Level_traverse()
}
