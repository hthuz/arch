package main

import "fmt"

type Color int

const (
	Red Color = iota
	Black
)

type RBNode struct {
	Key    int
	Val    int
	Left   *RBNode
	Right  *RBNode
	Parent *RBNode
	Color  Color
}

type RBTree struct {
	root *RBNode
}

func (t *RBTree) Insert(key, val int) {
	if t.root == nil {
		t.root = &RBNode{
			Key:   key,
			Val:   val,
			Color: Black,
		}
		return
	}

	// find place to insert
	var parent *RBNode
	cur := t.root
	for cur != nil {
		parent = cur
		if key == cur.Key {
			cur.Val = val
			return
		}
		if key < cur.Key {
			cur = cur.Left
			continue
		}
		if key > cur.Key {
			cur = cur.Right
			continue
		}
	}
	new_node := &RBNode{Key: key, Val: val, Parent: parent, Color: Red}

	if key > parent.Key {
		parent.Right = new_node
	} else {
		parent.Left = new_node
	}
	// if parent is black, insert directly
	if parent.Color == Black {
		return
	}

	cur = new_node
	for cur != nil {
		parent = cur.Parent
		grandparent := parent.Parent

		if grandparent == nil {
			break
		}

		var uncle *RBNode
		if grandparent.Right.Key == parent.Key {
			uncle = grandparent.Left
		} else {
			uncle = grandparent.Right
		}

		if uncle == nil {
			break
		}
		if uncle.Color == Black {
			break
		}

		// if parent is red and uncle is red
		if uncle.Color == Red {
			parent.Color = Black
			uncle.Color = Black
			grandparent.Color = Red
		}

		// set parent as cur and continue
		cur = parent
		if cur.Parent == nil {
			t.root.Color = Black
			return
		}
		if cur.Parent.Color == Black {
			t.root.Color = Black
			return
		}
	}
	// uncle is black or uncle doesn't exist
	grandparent := parent.Parent
	if parent == grandparent.Left && cur == parent.Left {
		t.rotateR(grandparent)
		parent.Color = Black
		grandparent.Color = Red
	}
	if parent == grandparent.Right && cur == parent.Right {
		t.rotateL(grandparent)
		parent.Color = Black
		grandparent.Color = Red
	}
	if parent == grandparent.Left && cur == parent.Right {
		t.rotateL(parent)
		t.rotateR(grandparent)
		cur.Color = Black
		grandparent.Color = Red
	}
	if parent == grandparent.Right && cur == parent.Left {
		t.rotateR(parent)
		t.rotateL(grandparent)
		cur.Color = Black
		grandparent.Color = Red
	}
	t.root.Color = Black

}
func (t *RBTree) Find(key int) *RBNode {
	return nil
}

func (t *RBTree) Delete(key int) {

}

func (t *RBTree) IsValid() bool {
	return false
}

func (t *RBTree) LevelTraverse() {
	leveltraverse([]*RBNode{t.root})
}

func leveltraverse(nodes []*RBNode) {
	if len(nodes) == 0 {
		return
	}
	next := make([]*RBNode, 0)
	for _, node := range nodes {

		if node.Color == Red {
			fmt.Print(node.Key, "(R) ")
		} else {
			fmt.Print(node.Key, "(B) ")
		}
		if node.Left != nil {
			next = append(next, node.Left)
		}
		if node.Right != nil {
			next = append(next, node.Right)
		}
	}
	fmt.Println()
	leveltraverse(next)

}

func (t *RBTree) rotateL(node *RBNode) {
	subR := node.Right
	subRL := subR.Left
	parent := node.Parent

	subR.Parent = parent
	subR.Left = node
	node.Parent = subR
	node.Right = subRL
	if subRL != nil {
		subRL.Parent = node
	}
	if parent == nil {
		t.root = subR
	}
	if node == parent.Left {
		parent.Left = subR
	} else {
		parent.Right = subR
	}

}
func (t *RBTree) rotateR(node *RBNode) {
	subL := node.Left
	subLR := subL.Right
	parent := node.Parent

	subL.Parent = parent
	subL.Right = node
	node.Parent = subL
	node.Left = subLR
	if subLR != nil {
		subLR.Parent = node
	}

	if parent == nil {
		t.root = subL
		return
	}

	// parent is not null
	if node == parent.Left {
		parent.Left = subL
	} else {
		parent.Right = subL
	}

}

func main() {
	tree := &RBTree{}
	tree.Insert(5, 5)
	tree.Insert(2, 2)
	tree.Insert(7, 7)
	tree.Insert(6, 6)
	// tree.Insert(1, 1)
	// tree.Insert(8, 8)
	tree.LevelTraverse()

}
