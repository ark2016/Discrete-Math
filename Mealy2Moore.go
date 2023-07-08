package main

import (
	"fmt"
	"strconv"
)

type pair struct {
	i     int
	value string
}

type MPI map[pair]int
type str2D [][]string
type int2D [][]int

func main() {
	var (
		n, m, q int
		pairs   []pair
	)

	fmt.Scan(&m)
	stringArray1 := scanStringArray(m)

	fmt.Scan(&q)
	stringArray2 := scanStringArray(q)

	fmt.Scan(&n)
	intMatrix := scanIntMatrix(n, m)
	stringMatrix := scanStringMatrix(n, m)

	mpi := make(MPI)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			ij, _ := strconv.Atoi(stringMatrix[i][j])
			p := pair{intMatrix[i][j], stringArray2[ij]}
			_, value := mpi[p]
			if !value {
				pairs = append(pairs, p)
				mpi[pairs[len(mpi)]] = len(mpi)
			}
		}
	}

	printout(pairs, m, stringMatrix, intMatrix, stringArray2, mpi, stringArray1)
}

func printout(pairs []pair, m int, stringMatrix str2D, intMatrix int2D, stringArray2 []string, mpi MPI, stringArray1 []string) {
	indent := "    "
	fmt.Println("digraph {")
	fmt.Println(indent + "rankdir = LR")
	for k, v := range pairs {
		fmt.Printf(indent+"%d [label = \"(%d,%s)\"]", k, v.i, v.value)
		fmt.Println()
	}
	for k, v := range pairs {
		for j := 0; j < m; j++ {
			ij, _ := strconv.Atoi(stringMatrix[v.i][j])
			p := pair{intMatrix[v.i][j], stringArray2[ij]}
			fmt.Printf(indent+"%d -> %d [label = \"%s\"]\n", k, mpi[p], stringArray1[j])
		}
	}
	fmt.Println("}")
}

func scanStringMatrix(n int, m int) str2D {
	arr := make(str2D, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]string, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	return arr
}

func scanStringArray(n int) []string {
	arr := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}
	return arr
}

func scanIntMatrix(n int, m int) int2D {
	arr := make(int2D, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, m)
		for j := 0; j < m; j++ {
			fmt.Scan(&arr[i][j])
		}
	}
	return arr
}
