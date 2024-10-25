package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	expr := "(3+5)*2/4"
	result, err := Calc(expr)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}

// Реализовать функцию func Calc(expression string) (float64, error) expression - строка-выражение состоящее 
// из односимвольных идентификаторов и знаков арифметических действий Входящие данные - цифры(рациональные), 
// операции +, -, *, /, операции приоритезации ( и ) В случае ошибки записи выражения функция выдает ошибку.
// Сохраните этот код себе на github. Он понадобится вам при выполнении финальных заданий следующих модулей.
func Calc(expression string) (float64, error) {
	// Удаление пробелов для упрощения анализа
	expression = strings.ReplaceAll(expression, " ", "")
	// Преобразование инфиксного выражения в обратную польскую нотацию (ОПН)
	rpn, err := toRPN(expression)
	if err != nil {
		return 0, err
	}
	// Вычисление выражения в обратной польской нотации
	return evalRPN(rpn)
}

func toRPN(expression string) ([]string, error) {
	var output []string
	var stack []rune

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if unicode.IsDigit(ch) || ch == '.' {
			// Считываем число
			j := i
			for j < len(expression) && (unicode.IsDigit(rune(expression[j])) || expression[j] == '.') {
				j++
			}
			output = append(output, expression[i:j])
			i = j - 1
		} else if ch == '(' {
			stack = append(stack, ch)
		} else if ch == ')' {
			// Выталкиваем из стека до открывающей скобки
			for len(stack) > 0 && stack[len(stack)-1] != '(' {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return nil, errors.New("mismatched parentheses")
			}
			stack = stack[:len(stack)-1]
		} else if isOperator(ch) {
			// Учитываем приоритет оператора
			for len(stack) > 0 && isOperator(stack[len(stack)-1]) && precedence(stack[len(stack)-1]) >= precedence(ch) {
				output = append(output, string(stack[len(stack)-1]))
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, ch)
		} else {
			return nil, errors.New("invalid character in expression")
		}
	}

	// Выталкиваем оставшиеся операторы
	for len(stack) > 0 {
		if stack[len(stack)-1] == '(' {
			return nil, errors.New("mismatched parentheses")
		}
		output = append(output, string(stack[len(stack)-1]))
		stack = stack[:len(stack)-1]
	}

	return output, nil
}

func evalRPN(rpn []string) (float64, error) {
	var stack []float64

	for _, token := range rpn {
		if isOperator(rune(token[0])) && len(token) == 1 {
			if len(stack) < 2 {
				return 0, errors.New("invalid expression")
			}
			b, a := stack[len(stack)-1], stack[len(stack)-2]
			stack = stack[:len(stack)-2]

			var result float64
			switch token {
			case "+":
				result = a + b
			case "-":
				result = a - b
			case "*":
				result = a * b
			case "/":
				if b == 0 {
					return 0, errors.New("division by zero")
				}
				result = a / b
			}
			stack = append(stack, result)
		} else {
			value, err := strconv.ParseFloat(token, 64)
			if err != nil {
				return 0, errors.New("invalid number")
			}
			stack = append(stack, value)
		}
	}

	if len(stack) != 1 {
		return 0, errors.New("invalid expression")
	}

	return stack[0], nil
}

func isOperator(ch rune) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/'
}

func precedence(op rune) int {
	switch op {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

