/*//==============================================================================================================
Польская запись — это форма записи арифметических, логических и алгебраических выражений, в которой операция 
располагается слева от операндов. Выражения в польской записи могут обходиться без скобок, однако мы оставим скобки 
для наглядности.

Например, выражение $5\cdot(3+4)$ в польской записи выглядит как

(* 5 (+ 3 4))
Пусть в нашем случае выражения состоят из чисел от 0 до 9, круглых скобок и трёх знаков операций: плюс, минус и 
звёздочка (умножить). Требуется составить программу polish.go, вычисляющую значение выражения.
*/ //==============================================================================================================
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func evaluatePolish(tokens []string) (int, error) {
	stack := make([]int, 0, len(tokens))
	for _, token := range tokens {
		switch token {
		case "+", "-", "*":
			val1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			val2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch token {
			case "+":
				stack = append(stack, val1+val2)
			case "-":
				stack = append(stack, val1-val2)
			case "*":
				stack = append(stack, val1*val2)
			}
		default:
			num, err := strconv.Atoi(token)
			if err != nil {
				return 0, fmt.Errorf("некорректный токен: %s", token)
			}
			stack = append(stack, num)
		}
	}
	if len(stack) != 1 {
		return 0, fmt.Errorf("некорректное количество токенов")
	}
	return stack[0], nil
}

func stringToArray(str string) []string {
	str = strings.ReplaceAll(str, "(", " ")
	str = strings.ReplaceAll(str, ")", " ")
	str = strings.ReplaceAll(str, " ", "")
	return strings.Split(str, "")
}

func main() {
	//"(* 5 (+ 3 4))"
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	expr := myscanner.Text()
	tokens := stringToArray(expr)
	for i, j := 0, len(tokens)-1; i < j; i, j = i+1, j-1 {
		tokens[i], tokens[j] = tokens[j], tokens[i]
	}

	result, err := evaluatePolish(tokens)

	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
	} else {
		fmt.Println(result)
	}
}
