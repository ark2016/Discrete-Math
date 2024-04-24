/*//==============================================================================================================
На строительном участке нужно создать телефонную сеть, соединяющую все бытовки. Для того, чтобы телефонные линии 
не мешали строительству, их решили проводить вдоль дорог. Составьте программу, реализующую алгоритм Прима для 
вычисления минимальной общей длины телефонных линий для указанной конфигурации участка. Граф конфигурации участка 
должен быть представлен в программе в виде списка инцидентности.

Программа должна считывать со стандартного потока ввода количество бытовок $N$, количество дорог $M$, соединяющих 
бытовки, и информацию об этих дорогах. При этом каждая дорога задаётся тремя целыми числами $u$, $v$ и $len$, где 
$u$ и $v$ — номера соединяемых дорогой бытовок ($0\leq u,v<N$), а $len$ — длина дороги.

Программа должна выводить в стандартный поток вывода минимальную общую длину телефонных линий.

Например, для входных данных

7
10
0 1 200
1 2 150
0 3 100
1 4 170
1 5 180
2 5 100
3 4 240
3 6 380
4 6 210
5 6 260
программа должна выводить число 930.
*/ //==============================================================================================================
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
