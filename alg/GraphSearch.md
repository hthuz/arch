# GraphSearch



Basic

```python

u: node
Explore(G, u):
	visited[u] = True
	toExplore.push(u)
	while (toExplore is not empty):
        x <- toExplore.pop()
        for y in Adj(x):
            if !visited[y]:
                visited[y] = True // 在这里就应该标记为visited 为true,而不是在toExplore.pop时标记为true,否则可能有重复
                toExplore.push(y)
                
# toExplore stack -> DFS
# toExplore queue -> BFD

# Recursive way

Explore(G,u):
    visited[u] = True
   	for y in Adj(u):
        if !visited[y]:
            visited[y] = True
            Explore(G,y)
	
```



## DFS 获取path



递归的时候用回溯

```go

path := make([]int, 0)
func dfs(node int) {
    
    visited[node] = true
    
    path = append(path, node)
    
    for _, next := range graph[node] {
        if visited[next] {
            continue
        }
        dfs(next)
    }
    
    path = path[:len(path)-1]
    
}
```

用stack的时候，记录每个节点的parent, 然后parent遍历，反转



```go

parent := make(map[int]int)

stack := make([]int, 0)
stack = append(stack, 0)
parent[0] = 0

for len(stack) != 0 {
    cur := stack[len(stack)-1]
    stack = stack[:len(stack)-1]
    visited[cur] = true
    
    for _, next := graph[cur] {
        if visited[next] {
            continue
        }
        parent[next] = cur
        stack = append(stack, next)
    }
}

path := make([]int, 0)
for node := target; node != 0; node = parent[node] {
    path = append(path, node)
}

for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
    path[i], path[j] = path[j], path[i]
}
```









## Leetcode 934 shortest bridge

对于宽度搜索，朴素方法是，对一个island的每一个节点，都进行宽搜，找到到另一个岛屿的距离，取最小值, 但是效率很低

```go
minBridge := n * n

for r := 0; r < n; r++ {
	for c := 0; c < n; c++ {
		if grid[r][c] != -1 {
			continue
		}
		node := pair{r, c}
		dists := make(map[pair]int)
		dists[node] = 0
		queue := make([]pair, 0)
		queue = append(queue, node)

		bridge := 0
		for len(queue) != 0 {
			node := queue[0]

			if grid[node.r][node.c] == 1 {
				bridge = dists[node] - 1
				break
			}
			for _, dir := range dirs {
				next_r, next_c := node.r+dir.r, node.c+dir.c
				if next_r >= 0 && next_r < n && next_c >= 0 && next_c < n {
					next_node := pair{next_r, next_c}
					if _, exists := dists[next_node]; exists {
						continue
					}
					dists[next_node] = dists[node] + 1
					queue = append(queue, next_node)
				}
			}
			queue = queue[1:]
		}

		minBridge = min(bridge, minBridge)
	}
}

```



对此有两个trick

1. BFS开始的不一定是一个点，而是一整个岛的集合
2. 通过记录queue, next_queue， 来记录每一层layer的距离，而不是直接对queue更新
2. 可以单独设立一个visited二维数组来记录每个位置是否被visited到，但是另一种方式是直接对grid进行修改，修改为-1,标记为visited,这样就不需要额外的visited数据声明。

```go
	queue = island
	dist := 0
	for len(queue) != 0 {
		next_queue := make([]pair, 0)
		for _, node := range queue {
			for _, d := range dirs {
				next_r, next_c := node.r+d.r, node.c+d.c
				if next_r >= 0 && next_r < n && next_c >= 0 && next_c < n && grid[next_r][next_c] != -1 {
					if grid[next_r][next_c] == 1 {
						return dist
					}
					grid[next_r][next_c] = -1
					next_queue = append(next_queue, pair{next_r, next_c})
				}
			}
		}
		queue = next_queue
		dist += 1
	}

```





## 将树所有节点都能改成访问节点0的改边方向次数





情况比较差的时候slices.Contain 仍可能O(n2)

```go
func main() {

	n := 0
	fmt.Scan(&n)

	graph := make([][]int, n)
	ungraph := make([][]int, n)
	for i := 0; i < n-1; i++ {
		u, v := 0, 0
		fmt.Scan(&u, &v)
		graph[u] = append(graph[u], v)
		ungraph[u] = append(ungraph[u], v)
		ungraph[v] = append(ungraph[v], u)
	}

	visited := make([]bool, n)
	prev := make([]int, n)
	var dfs func(int)

	dfs = func(node int) {

		visited[node] = true

		for _, next := range ungraph[node] {
			if visited[next] {
				continue
			}
			prev[next] = node
			dfs(next)
		}

	}

	dfs(0)

	ans := 0
	for i := 1; i < n; i++ {
		if slices.Contains(graph[prev[i]], i) {
			ans += 1
		}
	}
	fmt.Println(ans)

}

```



更好的方法

```go
package main

import "fmt"

type Edge struct {
	to  int
	dir int // 1: 当前节点 -> to，0: to -> 当前节点
}

func main() {
	var n int
	fmt.Scan(&n)

	graph := make([][]Edge, n)

	for i := 0; i < n-1; i++ {
		var u, v int
		fmt.Scan(&u, &v)

		// 原始方向：u -> v
		graph[u] = append(graph[u], Edge{
			to:  v,
			dir: 1,
		})

		// 从 v 看：v <- u
		graph[v] = append(graph[v], Edge{
			to:  u,
			dir: 0,
		})
	}

	visited := make([]bool, n)
	ans := 0

	var dfs func(int)
	dfs = func(node int) {
		visited[node] = true

		for _, edge := range graph[node] {
			if visited[edge.to] {
				continue
			}

			// 当前 node 是父节点，edge.to 是子节点。
			// 最终应该是：
			// edge.to -> node
			//
			// 如果原来是 node -> edge.to，则需要反转。
			if edge.dir == 1 {
				ans++
			}

			dfs(edge.to)
		}
	}

	dfs(0)

	fmt.Println(ans)
}
```

