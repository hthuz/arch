# Greedy

利用局部最优，达到全局最优



-8, -5, -5, -3, -2, 3



-4 -3 -2





## dijkstra



注意的几点

1. next下一轮的，不是所有当前访问节点的unvisited edge,不是next edge,是剩下所有未访问的节点，所以一开始应该把所有节点distance记录为infinite,每一轮更新距离即可
2. 因此记录的变量就两个，dist和to_explore, dist记录一开始距离均为infinite, to_explore可以用heap实现
3. 循环可以用for range n 实现，因为最多也就循环n轮
4. 使用heap时，可以一开始把所有节点放进 heap（每次都需要reorganize heap），也可以只有发现更短距离时才 push进入heap （第二次发现更短距离时，一个节点有可能多次被放入 heap，比如（1,dist 5），（1, dist 10），不做 decrease-key，而是直接 push 一个新的状态。有意义的实际上就是距离更小的那一个）



### 不使用heap

```go
func shortestPathNaive(n int, edges [][]int) int {

	// 1. build graph
	// src -> [src1, dst1, weight1], [src2, dst2, weight2]
	graph := make([][][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge)
	}

	// 2. init
	dist := make([]int, n) // 每个节点到node 0的最短距离
	to_explore := make([]int, n)
	for i := range n {
		dist[i] = math.MaxInt32 // not reached yet
		to_explore[i] = i
	}

	// 3. node 0
	dist[0] = 0

	// visit a new node each time
	for range n {

		// get min
		min_dist := math.MaxInt32
		min_i := -1
		for i, node := range to_explore {
			if dist[node] < min_dist {
				min_dist = dist[node]
				min_i = i
			}
		}
		if min_i == -1 {
			return -1
		}

		next := to_explore[min_i]
		to_explore = append(to_explore[:min_i], to_explore[min_i+1:]...)
		if next == n-1 {
			return dist[n-1]
		}

		// update
		for _, edge := range graph[next] {
			dist[edge[1]] = min(dist[edge[1]], dist[next]+edge[2])
		}

	}

	return dist[n-1]
}

```



### 使用heap

```go
type MinHeap []HeapNode

func (h MinHeap) Len() int {
	return len(h)
}

func (h MinHeap) Less(i, j int) bool {
	return h[i].Dist < h[j].Dist

}

func (h *MinHeap) Pop() any {
	tail := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return tail
}

func (h *MinHeap) Push(x any) {
	*h = append(*h, x.(HeapNode))

}

func (h MinHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

type HeapNode struct {
	Node int
	Dist int
}

func shortestPathHeap(n int, edges [][]int) int {

	graph := make([][][]int, n)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge)
	}

	// init
	dist := make([]int, n)
	for i := range n {
		dist[i] = math.MaxInt32
	}

	h := &MinHeap{}
	heap.Init(h)

	// node 0
	heap.Push(h, HeapNode{
		Node: 0,
		Dist: 0,
	})
	dist[0] = 0

	for h.Len() > 0 {
		next := heap.Pop(h).(HeapNode)

		if dist[next.Node] < next.Dist {
			continue
		}

		if next.Node == n-1 {
			return dist[n-1]
		}

		for _, edge := range graph[next.Node] {
			u, v, w := edge[0], edge[1], edge[2]
			if dist[u]+w < dist[v] {
				dist[v] = dist[u] + w

				heap.Push(h, HeapNode{
					Node: v,
					Dist: dist[v],
				})

			}
		}

	}

	return -1

}

```



### leetcode 3970

简单理解，每个节点不仅仅只有一个dist状态，还有count的状态，在不同count的数目下，都可能有不同的最小dist，因此不能只用count []int{} 来简单理解各个节点的连续letter数目, 因此不仅仅需要对每个node记录距离，还要拓展到node x count做状态记录

```go
    dist := make([][]int, n)
    for i := range n {
        dist[i] = make([]int, k)
        for j := range k {
            dist[i][j] = math.Int32
        }
    }
```

```
状态 = (Node, Count)

dist[Node][Count]
    ↓
到达 Node
且当前连续 label 数量为 Count
时的最短距离
```



假设：

```
k = 2

0(A) -> 1(A) -> 3(A)
       \-> 2(B) -> 3(A)
```

可能出现：

```
到达 3 的最短路径：
0(A) -> 1(A) -> 3(A)
连续 A = 3
```

这条路径违反 `k=2`。

但另一条：

```
0(A) -> 2(B) -> 3(A)
```

可能更长，但是合法。



### leetcode 1293

给你一个 `m * n` 的网格，其中每个单元格不是 `0`（空）就是 `1`（障碍物）。每一步，您都可以在空白单元格中上、下、左、右移动。

如果您 **最多** 可以消除 `k` 个障碍物，请找出从左上角 `(0, 0)` 到右下角 `(m-1, n-1)` 的最短路径，并返回通过该路径所需的步数。如果找不到这样的路径，则返回 `-1` 。

解法可以和3970一样，就是按照(r,c, chance)的状态来走dijkstra. 不过也可以直接用bfs

dijkstra: 优先队列

BFS： 队列

因为权重都是1, 所以队列获取到的元素，默认就是下一个最近的元素，BFS即是特殊情况下的dijkstra
