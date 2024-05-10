package main

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
		panic("Unknown operator")
	}
}
