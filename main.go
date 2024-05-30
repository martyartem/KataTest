package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Словари для преобразования римских чисел в арабские и обратно
var romanNums = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
}

var arabicNums = map[int]string{
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
}

// Является ли строка римским числом
func isRoman(s string) bool {
	_, ok := romanNums[s]
	return ok
}

// Является ли строка арабским числом
func isArabic(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

// Преобразование римского числа в арабское
func romanToArabic(s string) (int, error) {
	if val, ok := romanNums[s]; ok {
		return val, nil
	}
	return 0, errors.New("Неверная римская цифра")
}

// Преобразование арабского числа в римское
func arabicToRoman(num int) (string, error) {
	if num <= 0 {
		return "", errors.New("Римские цифры должны быть положительными")
	}
	if val, ok := arabicNums[num]; ok {
		return val, nil
	}
	return "", errors.New("Число вне диапазона")
}

// Арифметические операции
func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, errors.New("Нелья делить на 0")
		}
		return a / b, nil
	default:
		return 0, errors.New("Некорректная операция")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение")

	// Чтение ввода пользователя
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разделение ввода на части
	parts := strings.Split(input, " ")
	if len(parts) != 3 {
		panic("Неверный формат ввода")
	}

	num1 := parts[0]
	num2 := parts[2]
	op := parts[1]

	var result int
	var err error

	//Являются ли оба числа римскими
	if isRoman(num1) && isRoman(num2) {
		a, err1 := romanToArabic(num1)
		b, err2 := romanToArabic(num2)
		if err1 != nil || err2 != nil {
			panic("Неверное римское число")
		}
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("Римские числа должны быть в диапазоне от 1 до 10")
		}
		result, err = calculate(a, b, op)
		if err != nil {
			panic(err)
		}
		if result < 1 {
			panic("результат меньше единицы для римских чисел")
		}
		romanResult, err := arabicToRoman(result)
		if err != nil {
			panic(err)
		}
		fmt.Println(romanResult)
	} else if isArabic(num1) && isArabic(num2) {
		// Являются ли оба числа арабскими
		a, _ := strconv.Atoi(num1)
		b, _ := strconv.Atoi(num2)
		if a < 1 || a > 10 || b < 1 || b > 10 {
			panic("Арабаские числа должны быть в диапазоне от 1 до 10")
		}
		result, err = calculate(a, b, op)
		if err != nil {
			panic(err)
		}
		fmt.Println(result)
	} else {
		//Ошибка, если числа смешанного типа
		panic("Числа должны быть либо арабскими, либо римскими")
	}
}
