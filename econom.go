/*//==============================================================================================================
Пусть выражения в польской записи состоят из имён переменных (от a до z), круглых скобок и трёх знаков операций: 
#, $ и @ (смысл операций мы определять не будем).

Выражения могут содержать повторяющиеся подвыражения. Экономное вычисление таких выражений подразумевает, что 
повторяющиеся подвыражения вычисляются только один раз.

Требуется составить программу econom.go, вычисляющую количество операций, которые нужно выполнить для экономного 
вычисления выражения. Примеры работы программы приведены в таблице:

Набор тестов для программы экономного вычисления выражений в польской записи
Выражение				Количество операций
x						0
($xy)						1
($(@ab)c)					2
(#i($jk))					2
(#($ab)($ab))					2
(@(#ab)($ab))					3
(#($a($b($cd)))(@($b($cd))($a($b($cd)))))	5
(#($(#xy)($(#ab)(#ab)))(@z($(#ab)(#ab))))	6
*/ //==============================================================================================================
package main

import (
	"bufio"
	"fmt"
	"os"
)

func countOperations(expression string) int {
	var parenthesis []int
	var subExpression string
	var subExprStart int

	expressionList := map[string]int{}

	for position, token := range expression {
		if token == '(' {
			parenthesis = append(parenthesis, position)
		} else if token == ')' {
			subExprStart = parenthesis[len(parenthesis)-1]
			subExpression = expression[subExprStart:position]
			expressionList[subExpression]++
			parenthesis = parenthesis[:len(parenthesis)-1]
		}
	}

	return len(expressionList)
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	expr := readInput()
	fmt.Println(countOperations(expr))
}
