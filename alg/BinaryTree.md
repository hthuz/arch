# Binary Tree

Two ways of traversal

- Recursive (DFS)
- Level Order (BFS)



```
        5         
       / \        
      4     8     
     /     / \    
  11     13    4  
 /  \         / \ 
7    2       5   1
```



```
if node == nil {
	return
}
// 
traverse(node.left)
//
traverse(node.right)
//
```



Preorder: 5 4 11 7 2 8 13 4 5 1

Inorder: 7 11 2 4 5 13 8 5 4 1

Postorder: 7 2 11 4 13 5 1 4 8 5



preorder: execute before you enter a node

inoder: execute after you exit left subtree and going to enter right subtree

postorder: execute after you exit a node





**Inorder Traversal of binary search tree is sorted**



## Level order traversal

using queue

```go
func levelOrder(root *TreeNode) [][]int {
    res := make([][]int, 0)
    queue := make([]*TreeNode, 0)

    if root == nil {
        return res
    }
    queue = append(queue, root)

    for len(queue) > 0 {
        num := len(queue)
        row := make([]int, 0)
        for range num {
            head := queue[0]
            queue = queue[1:]
            row = append(row, head.Val)
            
            if head.Left != nil {
                queue = append(queue, head.Left)
            }
            if head.Right != nil {
                queue = append(queue, head.Right)
            }
        }
        res = append(res, row)
    }
    return res
}
```



or traverse layer by layer

```go
var res [][]int
func levelOrderBottom(root *TreeNode) [][]int {
    res = make([][]int, 0)
    if root == nil {
        return res
    }
    traverse([]*TreeNode{root})
    return res
}

func traverse(nodes []*TreeNode) {

    if len(nodes) == 0 {
        return
    }
    vals := make([]int, 0)
    nextNodes := make([]*TreeNode, 0)
    for _, node := range nodes {
        vals = append(vals, node.Val)
        if node.Left != nil {
            nextNodes = append(nextNodes, node.Left)
        }
        if node.Right != nil {
            nextNodes = append(nextNodes, node.Right)
        }
    }
	
    // travere from top to bottom
    traverse(nextNodes)
    // traverse from bottom to top
    res = append(res, vals)
}
```











## Two ways of thinking for binary tree

1. Whether you can solve the problem by traversing the tree

2. Whether you can solve the problem by solving subproblem (subtree) usually using postorder traversal

   problem = problem of current node + problem of left subtree + problem of right subtree

   
