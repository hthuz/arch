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



## 二叉树迭代遍历



```go

// 前序
func traverse(root *TreeNode) {
    if root == nil {
        return
    }
    stack := make([]*TreeNode, 0)
    stack.push(root)
    for len(stack) != 0 {
        cur := stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        fmt.Println(cur.Val)
        
        if cur.Right != nil {
            stack = append(stack, cur.Right)
        }
        if cur.Left != nil {
            stack = append(stack, cur.Left)
        }
    }
}

// 中序
func traverse(root *TreeNode) {
    if root == nil {
        return 
    }
    stack := make([]*TreeNode, 0)
    cur := root
    
    for len(stack) != 0 || cur != nil{
        // 一直往左前进
        
        for cur != nil {
            stack = append(stack)
            cur = cur.Left
        }
        cur = stack[len(stack)-1]
        stack = stack[:len(stack)-1]
        
        // 访问当前，然后访问右边
        fmt.Println(stack.Val)
        cur = cur.Right
              
    }
}
                  
// 后序
func traverse(root *TreeNode) {
    stack1 := make([]*TreeNode, 0)
    stack2 := make([]*TreeNode, 0)
    stack1 = append(stack1, root)
    for len(stack1) > 0 {
        cur := stack1[len(stack1)-1]
        stack1 = stack1[:len(stack1)-1]
        stack2 = append(stack2, cur)
        if cur.Left != nil {
            stack1 = append(stack1, cur.Left)
        }
        if cur.Right != nil {
            stack2 = append(stack2, cur.Right)
        }
    }
    for len(stack2) > 0 {
        cur := stack2[len(stac2)-1]
        stack2 = stack2[:len(stack2)-1]   
        fmt.Println(stack2)

    }
    
}
```







## Level order traversal

以下两种方法都很重要，

using queueOrder

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






## adding parent to Binary Tree

给定一个Binary Tree,把他变成一个带parent指针的binary tree

```go
type TreeNode struct {
    Val   int
    Left  *TreeNode
    Right *TreeNode
}


```



## 二叉树的最近公共祖先

我的解法： 记录root 到节点的path, 然后逐渐对path比较，先把长的去掉，知道path的最后一个相同，那个即是root

但是需要两次遍历树

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    var stack1 []*TreeNode
    var stack2 []*TreeNode
    traverse(root, p, make([]*TreeNode, 0), &stack1)
    traverse(root, q, make([]*TreeNode, 0), &stack2)

    for stack1[len(stack1)-1] != stack2[len(stack2)-1] {
        if len(stack1) > len(stack2) {
            stack1 = stack1[:len(stack1)-1]
        } else if len(stack1) < len(stack2) {
            stack2 = stack2[:len(stack2)-1]
        } else {
            stack1 = stack1[:len(stack1)-1]
            stack2 = stack2[:len(stack2)-1]

        }
    }
 
    return stack1[len(stack1)-1]
  
}

func traverse(root, node *TreeNode, path []*TreeNode, stack *[]*TreeNode, )  {
    if root == nil {
        return
    }
    if root == node {
        path = append(path, root)
        *stack = append([]*TreeNode{}, path...)
        return
    }
    path = append(path, root)
    traverse(root.Left, node, path, stack)
    traverse(root.Right, node, path, stack)
    path = path[:len(path)-1]

}
```

另一种解法，记录每个节点的父节点，然后不断往父节点遍历，找到第一个被p和q都遍历到过的父节点

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    parent := map[int]*TreeNode{}
    visited := map[int]bool{}

    var dfs func(*TreeNode)
    dfs = func(r *TreeNode) {
        if r == nil {
            return
        }
        if r.Left != nil {
            parent[r.Left.Val] = r
            dfs(r.Left)
        }
        if r.Right != nil {
            parent[r.Right.Val] = r
            dfs(r.Right)
        }
    }
    dfs(root)

    for p != nil {
        visited[p.Val] = true
        p = parent[p.Val]
    }
    for q != nil {
        if visited[q.Val] {
            return q
        }
        q = parent[q.Val]
    }

    return nil
}

```

最巧妙的解法

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    if root == p || root == q {
        return root
    }

    left := lowestCommonAncestor(root.Left, p, q)
    right := lowestCommonAncestor(root.Right, p, q)
    
    if left != nil && right != nil {
        return root
    }
    if left != nil && right == nil {
        return left
    } 
    if left == nil && right != nil {
        return right
    }
    return nil
    
    //////////////////////////////////////////
    // 后面这一段也可以简化为
    if left != nil && right != nil {
        return root
    }
    if left == nil {
        return right
    }
    return left
    
  
}
```







