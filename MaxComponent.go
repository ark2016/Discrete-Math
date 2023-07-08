package main

import "fmt"

const (
	minVertex = 1000000
)

type topOfTheGraph struct {
	neighbors []int
	visited   bool
}

type Component struct {
	numOfGraphVertex     int
	numOfGraphEdges      int
	componentVerticesMap map[int]bool
}

func graphScan(m int, graph []topOfTheGraph) {
	var first, second int
	for i := 0; i < m; i++ {
		fmt.Scanf("%d", &first)
		fmt.Scanf("%d\n", &second)
		graph[first].neighbors = append(graph[first].neighbors, second)
		graph[second].neighbors = append(graph[second].neighbors, first)
	}
}

func dfs(v int, j int, graph []topOfTheGraph, components []Component) {
	graph[v].visited = true
	components[j].numOfGraphVertex++
	components[j].componentVerticesMap[v] = true

	for i := 0; i < len(graph[v].neighbors); i++ {
		components[j].numOfGraphEdges++
		to := graph[v].neighbors[i]
		if !graph[to].visited {
			dfs(to, j, graph, components)
		}
	}
}

func vertexMin(component Component) int {
	vMin := minVertex
	for vertex, _ := range component.componentVerticesMap {
		if vMin > vertex {
			vMin = vertex
		}
	}
	return vMin
}

func vertexMinCompare(a Component, b Component) bool {
	return vertexMin(a) > vertexMin(b)
}

func compareComponents(a Component, b Component) bool {
	var res bool
	if a.numOfGraphVertex != b.numOfGraphVertex {
		res = a.numOfGraphVertex < b.numOfGraphVertex
	} else if a.numOfGraphEdges != b.numOfGraphEdges {
		res = a.numOfGraphEdges < b.numOfGraphEdges
	} else {
		res = vertexMinCompare(a, b)
	}
	return res
}

func findMaxComponent(components []Component) int {
	compMax := 0
	for i := 1; i < len(components); i++ {
		if compareComponents(components[compMax], components[i]) {
			compMax = i
		}
	}
	return compMax
}

func main() {
	var graph []topOfTheGraph
	var n, m int
	fmt.Scanf("%d\n", &n)
	fmt.Scanf("%d\n", &m)

	for i := 0; i < n; i++ {
		var vertex topOfTheGraph
		vertex.neighbors = make([]int, 0)
		graph = append(graph, vertex)
	}

	graphScan(m, graph)

	components := make([]Component, 0)
	for i, j := 0, 0; i < n; i++ {
		if !graph[i].visited {
			var component Component
			component.componentVerticesMap = make(map[int]bool)
			components = append(components, component)
			dfs(i, j, graph, components)
			j++
		}
	}

	maxComponent := findMaxComponent(components)

	fmt.Println("graph {")
	for i := 0; i < len(graph); i++ {
		fmt.Print("    ")
		fmt.Print(i)
		if components[maxComponent].componentVerticesMap[i] {
			fmt.Print(" [color = red]")
		}
		fmt.Println("")
	}
	for i := 0; i < len(graph); i++ {
		for j := 0; j < len(graph[i].neighbors); j++ {
			k := graph[i].neighbors[j]
			if i > k {
				continue
			}
			fmt.Print("    ")
			fmt.Print(i)
			fmt.Print(" -- ")
			fmt.Print(k)
			if components[maxComponent].componentVerticesMap[i] {
				fmt.Print(" [color = red]")
			}
			fmt.Println("")
		}
	}
	fmt.Println("}")
}
