/*//==============================================================================================================
Составьте программу, реализующую алгоритм Крускала для вычисления минимальной суммарной длины дорожек в парке 
аттракционов. Дорожки должны быть проложены таким образом, чтобы между любыми двумя аттракционами существовал маршрут.

Программа должна считывать со стандартного потока ввода количество аттракционов и их координаты. При этом координаты 
каждого аттракциона задаются парой целых чисел (в декартовой системе).

Программа должна выводить в стандартный поток вывода минимальную суммарную длину дорожек с точностью до двух знаков 
после запятой.

Например, для входных данных

12
2 4
2 5
3 4
3 5
6 5
6 6
7 5
7 6
5 1
5 2
6 1
6 2
программа должна выводить число 14.83.
*/ //==============================================================================================================
package main

import (
	"fmt"
	"math"
	"sort"
)

type edge struct {
	u, v   int
	weight float64
}

func findRoot(v int, roots []int) int {
	if roots[v] != v {
		roots[v] = findRoot(roots[v], roots)
	}
	return roots[v]
}

func kruskal(edges []edge, n int) float64 {
	sort.Slice(edges, func(i, j int) bool {
		return edges[i].weight < edges[j].weight
	})
	roots := make([]int, n)
	for i := range roots {
		roots[i] = i
	}

	totalWeight := 0.0
	edgesInTree := 0
	for _, e := range edges {
		if edgesInTree == n-1 {
			break
		}
		uRoot := findRoot(e.u, roots)
		vRoot := findRoot(e.v, roots)
		if uRoot != vRoot {
			roots[vRoot] = uRoot
			totalWeight += e.weight
			edgesInTree++
		}
	}

	return totalWeight
}

func distance(x1 int, y1 int, x2 int, y2 int) float64 {
	return math.Sqrt(float64((x1-x2)*(x1-x2) + (y1-y2)*(y1-y2)))
}

func main() {
	var n int
	fmt.Scan(&n)

	points := make([][2]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&points[i][0], &points[i][1])
	}

	var edges []edge
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			edges = append(edges, edge{i, j,
				distance(points[i][0], points[i][1], points[j][0], points[j][1])})
		}
	}

	minTotalWeight := kruskal(edges, n)
	fmt.Printf("%.2f\n", minTotalWeight)
}
