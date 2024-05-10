package main

import (
	"fmt"
	"strings"
)

func main() {
	var expression string
	fmt.Print("Введите выражение (например, 1 + 2): ")
	fmt.Scanln(&expression)

	expression_arr := strings.Split(expression, " ")
	if len(expression_arr) != 3 {
		print("Не является математической операцией")
	}

	var isRoman bool
	var a, b int
	var operator = expression_arr[1]

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Должны быть числа от 1 до 10")
	}

	result := calc(a, b, operator)

}

func romanToInt(romanNumber string) int {
	result := 0
	romanToIntMap := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
	}

	for i := 0; i < len(romanNumber); i++ {
		current := int(romanNumber[rune(romanNumber[i])])

		if i+1 < len(romanNumber) && romanToIntMap[rune(romanNumber[i+1])] > current {
			result -= current
		} else {
			result += current
		}
	}

	return result
}

func intToRoman(number int) string {
	arabicNums := []int{10, 9, 5, 4, 1}
	romanNums := []string{"X", "IX", "V", "IV", "I"}
	result := ""

	for i := 0; i < len(arabicNums); i++ {
		for arabicNums[i] > number {
			i++
			if i >= len(arabicNums) {
				break
			}
		}
		result += romanNums[i]
		number -= arabicNums[i]
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
