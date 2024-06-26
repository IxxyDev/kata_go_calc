package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение (например, 1 + 2)")
	expression, _ := reader.ReadString('\n')
	expression = strings.TrimRight(expression, "\n")
	expression = strings.ReplaceAll(expression, " ", "")

	expressionArr := splitExpression(expression)

	if len(expressionArr) < 3 {
		panic("Не является математической операцией")
	} else if len(expressionArr) > 3 {
		panic("Неподходящий формат операции, можно использовать только 2 операнда")
	}

	var isRomanCalculation bool
	var a, b int
	var operator = expressionArr[1]

	// Переведем в арабские, если римские, но запомним, что работаем с римскими для положительного результата
	if isValidRoman(expressionArr[0]) && isValidRoman(expressionArr[2]) {
		isRomanCalculation = true
		a = romanToInt(expressionArr[0])
		b = romanToInt(expressionArr[2])
	} else if !isValidRoman(expressionArr[0]) && !isValidRoman(expressionArr[2]) {
		var err error
		a, err = strconv.Atoi(expressionArr[0])
		if err != nil {
			panic("Неправильный формат числа")
		}
		b, err = strconv.Atoi(expressionArr[2])
		if err != nil {
			panic("Неправильный формат числа")
		}
	} else {
		panic("Ошибка! Можно считать числа только из одной системы счисления")
	}

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Должны быть числа от 1 до 10")
	}

	result := calc(a, b, operator)

	if isRomanCalculation {
		if result < 1 {
			panic("Для римских чисел результат должен быть положительным")
		}
		fmt.Println("Ответ:")
		fmt.Println(intToRoman(result))
	} else {
		fmt.Println("Ответ:")
		fmt.Println(result)
	}

}

func romanToInt(romanNumber string) int {
	result := 0
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	for i := 0; i < len(romanNumber); i++ {
		current := romanToIntMap[rune(romanNumber[i])]
		if i+1 < len(romanNumber) && romanToIntMap[rune(romanNumber[i+1])] > current {
			result -= current
		} else {
			result += current
		}
	}

	return result
}

func isValidRoman(romanNumber string) bool {
	for _, char := range romanNumber {
		// Тут оказался код Unicode, а не символ, нужно привести
		if rune(char) != 'I' && rune(char) != 'V' && rune(char) != 'X' {
			return false
		}
	}
	return true
}

func intToRoman(number int) string {
	arabicNums := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romanNums := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	result := ""
	for i := 0; i < len(arabicNums); i++ {
		for number >= arabicNums[i] {
			result += romanNums[i]
			number -= arabicNums[i]
		}
	}
	return result
}

func calc(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		return a / b
	default:
		panic("Неизвестный оператор, доступны операторы: + – * /")
	}
}

func splitExpression(expression string) []string {
	expression = strings.TrimSpace(expression)
	expression = strings.ReplaceAll(expression, " ", "")

	// Разделяем строку по оператору
	var operator string
	var operands []string

	for _, char := range expression {
		if strings.ContainsRune("+-*/", char) {
			operator = string(char)
			operands = strings.Split(expression, operator)
			break
		}
	}

	if operator == "" || len(operands) != 2 {
		panic("Не является математической операцией")
	}

	return []string{operands[0], operator, operands[1]}
}
