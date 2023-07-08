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
