/*//==============================================================================================================
Составьте программу, выполняющую визуализацию заданного инициального автомата Мили через graphviz.

Программа должна считывать из стандартного потока ввода количество состояний автомата $n$, размер входного алфавита 
$m$ ($0<m\leq 26$), номер начального состояния $q_0$ ($0\leq q_0<n$), матрицу переходов $\Delta$ и матрицу выходов 
$\Phi$. Матрицы переходов и выходов имеют размеры $n\times m$. При этом элементами матрицы переходов являются номера 
состояний, а элементами матрицы выходов — выходные сигналы. Каждый выходной сигнал представляет собой не содержащую 
пробелов строку.

Будем считать, что входными сигналами автомата являются $m$ первых строчных букв латинского алфавита. При этом первый 
столбец матриц $\Delta$ и $\Phi$ соответствует букве "a", второй столбец — букве "b", и т.д.

Программа должна выводить в стандартный поток вывода описание диаграммы автомата на языке DOT. При этом каждое состояние
на диаграмме должно быть представлено кружком с порядковым номером состояния внутри, а каждая дуга должна иметь метку 
вида “x(y)”, где $x$ и $y$ — это входной и выходной сигналы, соответственно.

Например, для входных данных

4
3
0
1 3 3
1 1 2
2 2 2
1 2 3
x y y
y y x
x x x
x y y
программа должна выводить

digraph {
    rankdir = LR
    0 -> 1 [label = "a(x)"]
    0 -> 3 [label = "b(y)"]
    0 -> 3 [label = "c(y)"]
    1 -> 1 [label = "a(y)"]
    1 -> 1 [label = "b(y)"]
    1 -> 2 [label = "c(x)"]
    2 -> 2 [label = "a(x)"]
    2 -> 2 [label = "b(x)"]
    2 -> 2 [label = "c(x)"]
    3 -> 1 [label = "a(x)"]
    3 -> 2 [label = "b(y)"]
    3 -> 3 [label = "c(y)"]
}
*/ //==============================================================================================================
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
