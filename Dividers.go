/*//==============================================================================================================
Составьте программу, выполняющую построение графа делителей натурального числа $x$ такого, что $0<x<2^{32}$.

Вершинами графа делителей являются все делители числа $x$. Ребро соединяет вершины $u$ и $v$ в том случае, если 
$u$ делится на $v$, и не существует такого $w$, что $u$ делится на $w$, и $w$ делится на $v$. Пример графа делителей 
изображён на рисунке:

Граф делителей числа 18

Число $x$ должно считываться со стандартного потока ввода. Граф делителей, описанный на языке DOT, должен 
выводиться в стандартный поток вывода.
*/ //==============================================================================================================
package main

import (
	"fmt"
	"math"
	"sort"
)

func getDividers(x int) []int {
	size := int(math.Max(float64(x/4), 100))
	//if x > 3940325370{
	if x > 300000000 {
		size = 1000
	}
	a := make([]int, size)
	j := 0
	for i := 1; i < int(math.Sqrt(float64(x)))+1; i++ {
		if 0 == x%i {
			a[j] = i
			j++
			a[j] = x / i
			j++
		}
	}
	return a
}

func isConnected(i int, j int) bool {
	x := i / j
	for a := 2; a <= x/2; a++ {
		if 0 == x%a {
			return false
		}
	}
	return true
}

func main() {
	var x int
	fmt.Scan(&x)
	if 1 == x {
		fmt.Printf("graph {\n    %d\n}", x)
		return
	}
	dividers := getDividers(x)
	sort.Ints(dividers)
	for i, j := 0, len(dividers)-1; i < j; i, j = i+1, j-1 {
		dividers[i], dividers[j] = dividers[j], dividers[i]
	}
	fmt.Println("graph {")
	//ans := "graph {\n"

	for _, i := range dividers {
		if 0 == i {
			break
		}
		//ans += "    " + strconv.Itoa(i) + "\n"
		fmt.Print("\t")
		fmt.Println(i)
	}

	for _, i := range dividers {
		if 0 == i {
			break
		}
		for _, j := range dividers {
			if 0 == j {
				break
			}
			if i != j && 0 == i%j && isConnected(i, j) {
				//ans += "    " + strconv.Itoa(i) + " -- " + strconv.Itoa(j) + "\n"
				fmt.Print("\t")
				fmt.Print(i)
				fmt.Print(" -- ")
				fmt.Println(j)
			}
		}

	}

	fmt.Print("}")
	//ans += "}"
	//fmt.Print(ans)
}
