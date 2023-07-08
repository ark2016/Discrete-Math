package main

import "fmt"

type topOfTheGraph struct {
	neighbors []int
	visited   bool
	tin, fup  int
}

func min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func dfs(v int, p int, graph []topOfTheGraph, timer *int, count *int) {
	graph[v].visited = true
	*timer++
	graph[v].tin = *timer
	graph[v].fup = *timer

	for i := 0; i < len(graph[v].neighbors); i++ {
		to := graph[v].neighbors[i]
		if to == p {
			continue
		}
		if graph[to].visited {
			graph[v].fup = min(graph[v].fup, graph[to].tin)
		} else {
			dfs(to, v, graph, timer, count)
			graph[v].fup = min(graph[v].fup, graph[to].fup)
			if graph[to].fup > graph[v].tin {
				*count++
			}
		}
	}
}

func main() {
	var graph []topOfTheGraph
	var N, M, first, second, count, timer int

	fmt.Scanf("%d\n", &N)
	fmt.Scanf("%d\n", &M)

	for i := 0; i < N; i++ {
		var vertex topOfTheGraph
		vertex.neighbors = make([]int, 0)
		graph = append(graph, vertex)
	}

	for i := 0; i < M; i++ {
		fmt.Scanf("%d", &first)
		fmt.Scanf("%d\n", &second)
		graph[first].neighbors = append(graph[first].neighbors, second)
		graph[second].neighbors = append(graph[second].neighbors, first)
	}

	for i := 0; i < N; i++ {
		if !graph[i].visited {
			dfs(i, -1, graph, &timer, &count)
		}
	}

	fmt.Println(count)
}
