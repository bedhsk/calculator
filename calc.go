package mycalculator

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Calc struct{}

func (Calc) Operate(input, operator string) int {
	cleanInput := strings.Split(input, operator)
	operator1 := convert(cleanInput[0])
	operator2 := convert(cleanInput[1])
	switch operator {
	case "+":
		return operator1 + operator2
	case "-":
		return operator1 - operator2
	case "*":
		return operator1 * operator2
	case "/":
		return operator1 / operator2
	default:
		return 0
	}
}

func convert(input string) int {
	operator, _ := strconv.Atoi(input)
	return operator
}

func ReadInput() string {
	// Declarar el scaner
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text() // leer los datos ingresados
}
