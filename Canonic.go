package main

import (
	"fmt"
)

type Pair struct {
	to    int
	value string
}

//type Node struct {
//	visited []bool
//	pint1   []int
//	pint2   []int
//}

func main() {
	var n, m, q0 int

	fmt.Scan(&n, &m, &q0)

	matrix := make([][]Pair, n)
	//node := newNode(n)
	var (
		visited []bool
		pint1   []int
		pint2   []int
	)

	pint1 = make([]int, n)
	pint2 = make([]int, n)
	visited = make([]bool, n)
	c := 0

	makeMatrix(&n, &matrix, &m)
	scanMatrix(&n, &m, &matrix)
	dfs(&q0, &c, &matrix, &visited, &pint1, &pint2)
	print(&c, &m, &pint1, &pint2, &matrix)
	//dfs(q0, &c, matrix, node)
	//print(c, m, node, matrix)
}

func dfs(v *int, count *int, matrix *[][]Pair, visited *[]bool, pint1, pint2 *[]int) {
	//node.visited[v] = true
	//node.pint1[*count] = v
	//node.pint2[v] = *count
	(*visited)[*v] = true
	(*pint1)[*count] = *v
	(*pint2)[*v] = *count
	*count++

	for i := 0; i < len((*matrix)[*v]); i++ {
		to := (*matrix)[*v][i].to
		if (*visited)[to] {
			//if node.visited[to] {
			continue
		}
		dfs(&to, count, matrix, visited, pint1, pint2)
	}
}

//func newNode(n int) Node {
//	var m Node
//
//	m.pint1 = make([]int, n)
//	m.pint2 = make([]int, n)
//	m.visited = make([]bool, n)
//
//	return m
//}

func print(c *int, m *int, pint1, pint2 *[]int, matrix *[][]Pair) {
	//stdout := bufio.NewWriter(os.Stdout)
	//stdout.WriteString(strconv.Itoa(c) + "\n")
	//stdout.WriteString(strconv.Itoa(m) + "\n" + "0" + "\n")

	fmt.Println(*c)
	fmt.Println(*m)
	fmt.Println(0)

	for i := 0; i < *c; i++ {
		for j := 0; j < *m; j++ {
			//fmt.Print(node.pint2[matrix[node.pint1[i]][j].to], " ")
			fmt.Print((*pint2)[(*matrix)[(*pint1)[i]][j].to], " ")
			//stdout.WriteString(strconv.Itoa(node.pint2[matrix[node.pint1[i]][j].to]) + " ")
		}
		fmt.Println()
		//stdout.WriteString("\n")
	}

	for i := 0; i < *c; i++ {
		for j := 0; j < *m; j++ {
			//fmt.Print(matrix[node.pint1[i]][j].value + " ")
			fmt.Print((*matrix)[(*pint1)[i]][j].value + " ")
			//stdout.WriteString(matrix[node.pint1[i]][j].value + " ")
		}
		fmt.Println()
		//stdout.WriteString("\n")
	}
	//stdout.Flush()
}

func makeMatrix(n *int, matrix *[][]Pair, m *int) {
	for i := 0; i < *n; i++ {
		(*matrix)[i] = make([]Pair, *m)
	}
}

func scanMatrix(n *int, m *int, matrix *[][]Pair) {
	for i := 0; i < *n; i++ {
		for j := 0; j < *m; j++ {
			fmt.Scan(&(*matrix)[i][j].to)
		}
	}

	for i := 0; i < *n; i++ {
		for j := 0; j < *m; j++ {
			fmt.Scan(&(*matrix)[i][j].value)
		}
	}
}
