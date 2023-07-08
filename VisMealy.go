package main

import "fmt"

const (
	transitionShift = 97
)

func main() {
	var n, m, qo int

	fmt.Scan(&n, &m, &qo)

	capitalDelta := make([][]int, n)
	capitalFi := make([][]string, n)
	for i := 0; i < n; i++ {
		capitalDelta[i] = make([]int, m)
		capitalFi[i] = make([]string, m)
	}

	scanMatrix(n, m, capitalDelta, capitalFi)
	displayMatrix(n, m, capitalDelta, capitalFi, qo)
}

func scanMatrix(n int, m int, capitalDelta [][]int, capitalFi [][]string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&capitalDelta[i][j])
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&capitalFi[i][j])
		}
	}
}

func displayMatrix(n int, m int, capitalDelta [][]int, capitalFi [][]string, qo int) {
	indent := "    "
	fmt.Println("digraph {")
	fmt.Println(indent + "rankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf(indent+"%d -> %d [label = \"%c(%s)\"]", i, capitalDelta[i][j], (byte)(j+transitionShift), capitalFi[i][j])
			fmt.Println()
		}
		qo++
	}
	fmt.Println("}")
}
