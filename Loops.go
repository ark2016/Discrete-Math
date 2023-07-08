package main

import (
	"fmt"
	"sort"
)

type Stack struct {
	buf []*topOfTheGraph
	top int
}

type topOfTheGraph struct {
	visited bool

	command string

	index int
	move  int

	dom      *topOfTheGraph
	sdom     *topOfTheGraph
	label    *topOfTheGraph
	parent   *topOfTheGraph
	ancestor *topOfTheGraph

	bucket []*topOfTheGraph

	out []*topOfTheGraph
	in  []*topOfTheGraph
}

var (
	stack            Stack
	topOfTheGraphNel int
	topOfTheGraphMap = make(map[int]int)
)

func main() {
	var N int

	fmt.Scanf("%d", &N)
	graph := make([]*topOfTheGraph, 0)
	graph = setGraph(N)
	actionScan(graph)
	actionRecognition(graph, N)
	topOfTheGraphNel++
	dfs(graph[0])
	graph = in(graph)

	sort.Slice(graph, func(i, j int) bool {
		return (graph)[j].index < (graph)[i].index
	})

	N = len(graph)
	graph = dominators(graph, N)
	result := 0

	for _, a := range graph {
		for _, b := range a.in {
			for b != nil && b != a {
				b = b.dom
			}
			if a == b {
				result++
				break
			}
		}
	}
	fmt.Println(result)
}

func (s *Stack) push(x *topOfTheGraph) {
	(*s).buf = append((*s).buf, x)
	(*s).top++
}

func (s *Stack) pop() *topOfTheGraph {
	index := len((*s).buf) - 1
	topElement := (*s).buf[index]
	(*s).buf = (*s).buf[:index]
	(*s).top--
	return topElement
}

func in(graph []*topOfTheGraph) []*topOfTheGraph {
	for i := 0; i < len(graph); i++ {
		if !graph[i].visited {
			l := len(graph) - 1
			graph[i] = graph[l]
			graph[l] = nil
			graph = graph[:l]
			i--
		} else {
			for j := 0; j < len(graph[i].in); j++ {
				if !graph[i].in[j].visited {
					li := len(graph[i].in) - 1
					graph[i].in[j] = graph[i].in[li]
					graph[i].in = graph[i].in[:li]
					j--
				}
			}
		}
	}
	return graph
}

func actionRecognition(graph []*topOfTheGraph, N int) {
	for index := range graph {
		switch graph[index].command {
		case "ACTION":
			if index != N-1 {
				graph[index].out = append(graph[index].out, graph[index+1])
				graph[index+1].in = append(graph[index+1].in, graph[index])
			}
		case "JUMP":
			t := topOfTheGraphMap[graph[index].move]
			graph[index].out = append(graph[index].out, graph[t])
			graph[t].in = append(graph[t].in, graph[index])
		case "BRANCH":
			t := topOfTheGraphMap[graph[index].move]
			graph[index].out = append(graph[index].out, graph[t])
			graph[t].in = append(graph[t].in, graph[index])
			if index != N-1 {
				graph[index].out = append(graph[index].out, graph[index+1])
				graph[index+1].in = append(graph[index+1].in, graph[index])
			}
		default:
			fmt.Println("action error")
		}
	}
}

func actionScan(graph []*topOfTheGraph) {
	var (
		v          int
		w          int
		currentAct string
	)
	for i := range graph {
		fmt.Scan(&v)
		fmt.Scan(&currentAct)
		graph[i].command = currentAct
		if currentAct != "ACTION" {
			fmt.Scan(&w)
			graph[i].move = w
		}
		topOfTheGraphMap[v] = i
	}
}

func setGraph(N int) []*topOfTheGraph {
	var result []*topOfTheGraph
	for i := 0; i < N; i++ {
		var t topOfTheGraph
		t.out = make([]*topOfTheGraph, 0)
		t.in = make([]*topOfTheGraph, 0)
		t.bucket = make([]*topOfTheGraph, 0)
		t.sdom = &t
		t.label = &t
		result = append(result, &t)
	}
	return result
}

func findMin(v *topOfTheGraph) *topOfTheGraph {
	var min *topOfTheGraph
	if (*v).ancestor == nil {
		min = v
	} else {
		u := v
		for (*(*u).ancestor).ancestor != nil {
			stack.push(u)
			u = (*u).ancestor
		}
		for stack.top != 0 {
			v = stack.pop()
			if (*(*(*(*v).ancestor).label).sdom).index < (*(*(*v).label).sdom).index {
				(*v).label = (*(*v).ancestor).label
			}
			v.ancestor = u.ancestor
		}
		min = v.label
	}
	return min
}

func dominators(graph []*topOfTheGraph, N int) []*topOfTheGraph {
	stack.buf = make([]*topOfTheGraph, N)

	for _, w := range graph {
		if (*w).index != 1 {
			for _, v := range (*w).in {
				u := findMin(v)
				if (*(*w).sdom).index > (*(*u).sdom).index {
					(*w).sdom = (*u).sdom
				}
			}
			(*w).ancestor = (*w).parent
			(*(*w).sdom).bucket = append((*(*w).sdom).bucket, w)
			for _, v := range (*(*w).parent).bucket {
				u := findMin(v)
				if (*u).sdom != (*v).sdom {
					(*v).dom = u
				} else {
					(*v).dom = (*v).sdom
				}
			}
			(*(*w).parent).bucket = nil
		}
	}

	for _, w := range graph {
		if ((*w).index != 1) && (*w).dom != (*w).sdom {
			(*w).dom = (*(*w).dom).dom
		}
	}
	graph[len(graph)-1].dom = nil
	return graph
}

func dfs(r *topOfTheGraph) {
	(*r).visited = true
	topOfTheGraphNel++
	(*r).index = topOfTheGraphNel - 1
	for e := range (*r).out {
		if !(*(*r).out[e]).visited {
			(*(*r).out[e]).parent = r
			dfs(r.out[e])
		}
	}
}
