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

