# Binary Search Tree


The minimum value of a BST: going left until a node has no left child
The maximum value of a BST: going right until a node has no right child
The left or right subtree of a node is also a BST

Variants: AVL, Red-Black, B+Tree

Two important properties when solving problems:
- BST is sorted
- BST can be trimmed during traversal


Insertion of BST
```go

func Search(root *Node, val int) bool {
    if root == nil {
        return false
    }
    if val == root.Val {
        return true
    }
    // trim nodes when traversing
    if val < root.Val {
        return Search(root.Left, val)
    }
    if val > root.Val {
        return Search(root.Right, val)
    }
}

func Insert(root *Node, val int) *Node {
    if root == nil {
        return &Node{Val: val}
    }
    if val < root.Val {
        root.Left = Insert(root.Left, val)
    } 
    if val > root.Val {
        root.Right = Insert(root.Right, val)
    }

}

```


If a tree is a valid BST (leetcode 98)
The following typical example shows that you can't simply judge a node and its left child and right child.
Instead, you need to consider the maximum of left subtree and
the minimum of right subtree.

But a simpler way is simply inorder traversing the tree and evaluate if its sorted
```
  5   
 / \  
4   6 
   / \
   3 7
```


```go
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    _, _, valid := subtree(root)
    return valid

}

func subtree(root *TreeNode) (int, int, bool) {
    if root.Left == nil && root.Right == nil {
        return root.Val, root.Val, true
    }
    min_res := root.Val 
    max_res := root.Val
    if root.Left != nil {
        minv, maxv, valid := subtree(root.Left)
        if !valid {
            return 0, 0, false 
        }
        // max value of left subtree should be smaller than cur val
        if maxv >= root.Val {
            return 0, 0, false 
        }
        // only left subtrees may have smaller value
        if minv < min_res {
            min_res = minv
        }
    }
    if root.Right != nil {
        minv, maxv, valid := subtree(root.Right)
        if !valid {
            return 0, 0, false 
        }
        // min value of right subtree should be larger than cur val
        if minv <= root.Val {
            return 0, 0, false
        }
        // only right subtrees may have larger value
        if maxv > max_res {
            max_res = maxv
        }
    }
    return min_res, max_res, true
}

```







