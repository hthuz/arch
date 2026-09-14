package main

import (
	"container/heap"
	"fmt"
	"math"
)

func main() {

	n := 3

	edges := [][]int{{0, 1, 1}, {1, 2, 1}, {0, 2, 3}}

	ans := shortestPathHeap(n, edges)
	fmt.Println(ans)
}

// node 从 0 到 n - 1。该图由二维数组 edges 表示，其中 edges[i] = [ui, vi, wi] 表示一条从节点 ui 指向节点 vi、权重为 wi 的有向边。
// 返回一条从节点 0 到节点 n - 1 的路径的 最小总边权

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
