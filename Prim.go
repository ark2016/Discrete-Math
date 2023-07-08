package main

import (
	"container/heap"
	"fmt"
)

type Edge struct {
	node, weight int
}

type EdgeHeap []Edge

func (h EdgeHeap) Len() int {
	return len(h)
}

func (h EdgeHeap) Less(i, j int) bool {
	return h[i].weight < h[j].weight
}

func (h EdgeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *EdgeHeap) Push(x interface{}) {
	*h = append(*h, x.(Edge))
}

func (h *EdgeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

func prims(graph [][]Edge, n int) int {
	totalWeight := 0
	visited := make([]bool, n)
	current := 0
	h := &EdgeHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		for _, edge := range graph[current] {
			heap.Push(h, edge)
		}
		visited[current] = true

		for h.Len() > 0 {
			minEdge := heap.Pop(h).(Edge)
			if !visited[minEdge.node] {
				totalWeight += minEdge.weight
				current = minEdge.node
				break
			}
		}
	}

	return totalWeight
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	graph := make([][]Edge, N)
	for i := 0; i < M; i++ {
		var u, v, len int
		fmt.Scan(&u, &v, &len)
		graph[u] = append(graph[u], Edge{v, len})
		graph[v] = append(graph[v], Edge{u, len})
	}

	fmt.Println(prims(graph, N))
}
