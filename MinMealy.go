/*//==============================================================================================================
Составьте программу, выполняющую минимизацию заданного инициального автомата Мили.

Программа должна считывать из стандартного потока ввода количество состояний автомата $n$, размер входного 
алфавита $m$ ($0<m\leq 26$), номер начального состояния $q_0$ ($0\leq q_0<n$), матрицу переходов $\Delta$ и 
матрицу выходов $\Phi$. Матрицы переходов и выходов имеют размеры $n\times m$. При этом элементами матрицы 
переходов являются номера состояний, а элементами матрицы выходов — выходные сигналы. Каждый выходной сигнал 
представляет собой не содержащую пробелов строку.

Будем считать, что входными сигналами автомата являются $m$ первых строчных букв латинского алфавита. При этом 
первый столбец матриц $\Delta$ и $\Phi$ соответствует букве "a", второй столбец — букве "b", и т.д.

Программа должна выводить в стандартный поток вывода описание диаграммы минимизированного автомата на языке DOT. 
При этом каждое состояние на диаграмме должно быть представлено кружком с каноническим порядковым номером состояния 
внутри, а каждая дуга должна иметь метку вида “x(y)”, где $x$ и $y$ — это входной и выходной сигналы, соответственно.

Например, для входных данных

5
3
0
1 2 3
3 4 1
3 4 2
3 0 4
4 4 3
x x y
y x x
y x x
x x y
x y x
программа должна выводить описание автомата Мили, изображённого на рисунке
*/ //==============================================================================================================
package main

import (
	"fmt"
)
//TRANSITION_SHIFT
const (
	TRANSITION_SHIFT = 97
	INDENT           = "    "
)

type Pair struct {
	num   int
	value string
}

type pair2D [][]Pair
type MIB map[int]bool

func main() {
	var n, m, q0, c int
	fmt.Scan(&n, &m, &q0)
	matrix := makeMatrix(n, m)

	scanMatrixNumAndValue(n, m, matrix)

	classRoot := makeIntArr(n)
	power := makeIntArr(n)
	num := makeIntArr(n)
	reverse := makeIntArr(n)
	visited := makeBoolArr(n)

	newMatrix, usage, piArr := AufenkampHohn(matrix, classRoot, power)

	dfs(usage[piArr[q0]], &c, newMatrix, visited, num, reverse)
	printAutomate(c, newMatrix, reverse, num)
}

func find(v int, classRoot []int) int {
	if v == classRoot[v] {
		return v
	}

	classRoot[v] = find(classRoot[v], classRoot)
	return classRoot[v]
}

func union(a, b int, classRoot, power []int) {
	a = find(a, classRoot)
	b = find(b, classRoot)

	if b != a {
		if power[b] > power[a] {
			a, b = b, a
		}
		classRoot[b] = a
		if power[b] == power[a] {
			power[a]++
		}
	}
}

func split1(matrix pair2D, classRoot, power []int) (m int, piArr []int) {
	m = len(matrix)
	for i := 0; i < len(matrix); i++ {
		classRoot[i] = i
	}

	for q1 := 0; q1 < len(matrix); q1++ {
		for q2 := 0; q2 < len(matrix); q2++ {
			if find(q1, classRoot) != find(q2, classRoot) {
				eq := true
				for x := 0; x < len(matrix[0]); x++ {
					if phi(matrix, q1, x) != phi(matrix, q2, x) {
						eq = false
						break
					}
				}
				if eq {
					union(q1, q2, classRoot, power)
					m--
				}
			}
		}
	}
	piArr = make([]int, len(matrix))

	for q := 0; q < len(matrix); q++ {
		piArr[q] = find(q, classRoot)
	}
	return
}

func phi(matrix pair2D, q int, x int) string {
	return matrix[q][x].value
}

func split(matrixQ pair2D, piArr, classRoot, power []int) int {
	m := len(matrixQ)

	for q := 0; q < len(matrixQ); q++ {
		classRoot[q] = q
	}

	for q1 := 0; q1 < len(matrixQ); q1++ {
		for q2 := 0; q2 < len(matrixQ); q2++ {
			if piArr[q1] == piArr[q2] && find(q1, classRoot) != find(q2, classRoot) {
				eq := true

				for x := 0; x < len(matrixQ[0]); x++ {
					w1, w2 := matrixQ[q1][x].num, matrixQ[q2][x].num
					if piArr[w1] != piArr[w2] {
						eq = false
						break
					}
				}

				if eq {
					union(q1, q2, classRoot, power)
					m--
				}
			}
		}
	}

	for q := 0; q < len(matrixQ); q++ {
		piArr[q] = find(q, classRoot)
	}
	return m
}

func AufenkampHohn(matrix pair2D, classRoot, power []int) (matrixQStroke pair2D, usage []int, piArr []int) {
	var m, cur int

	m, piArr = split1(matrix, classRoot, power)

	for true {
		mStroke := split(matrix, piArr, classRoot, power)
		if m == mStroke {
			break
		}
		m = mStroke
	}

	matrixQStroke = make(pair2D, 0)
	usage = makeIntArr(len(matrix))
	visitedQStroke := make(MIB)

	markup(matrix, usage, piArr)

	for q := 0; q < len(matrix); q++ {
		qStroke := piArr[q]
		if !visitedQStroke[qStroke] {
			visitedQStroke[qStroke] = true
			matrixQStroke = append(matrixQStroke, make([]Pair, len(matrix[0])))
			for x := 0; x < len(matrix[0]); x++ {
				matrixQStroke[cur][x].num = usage[piArr[matrix[q][x].num]]
				matrixQStroke[cur][x].value = matrix[q][x].value
			}
			cur++
		}
	}
	return
}

func markup(matrix pair2D, usage []int, piArr []int) {
	visited := make(MIB)
	var cur int

	for i := 0; i < len(matrix); i++ {
		usage[i] = cur
		if !visited[piArr[i]] {
			visited[piArr[i]] = true
			cur++
		}
	}
}

func dfs(v int, c *int, matrix pair2D, visited []bool, num []int, reverse []int) {
	visited[v] = true
	num[*c] = v
	reverse[v] = *c
	*c++

	for i := 0; i < len(matrix[v]); i++ {
		to := matrix[v][i].num
		if !visited[to] {
			dfs(to, c, matrix, visited, num, reverse)
		}
	}
}

func printAutomate(c int, newMatrix pair2D, reverse []int, num []int) {
	fmt.Println("digraph {")
	fmt.Println(INDENT + "rankdir = LR")

	for i := 0; i < c; i++ {
		for j := 0; j < len(newMatrix[i]); j++ {
			fmt.Printf(INDENT+"%d -> %d [label = \"%c(%s)\"]", i, reverse[newMatrix[num[i]][j].num], (byte)(j+TRANSITION_SHIFT), newMatrix[num[i]][j].value)
			fmt.Println()
		}
	}

	fmt.Println("}")
}

func makeIntArr(n int) []int {
	ints := make([]int, n)
	return ints
}

func makeBoolArr(n int) []bool {
	ints := make([]bool, n)
	return ints
}

func scanMatrixNumAndValue(n int, m int, matrix pair2D) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j].num)
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scan(&matrix[i][j].value)
		}
	}
}

func makeMatrix(n int, m int) pair2D {
	matrix := make(pair2D, n)

	for i := 0; i < n; i++ {
		matrix[i] = make([]Pair, m)
	}
	return matrix
}
