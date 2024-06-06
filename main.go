package main

import (
	"fmt"
)

const (
	lowercaseLetters = "abcdefghijklmnopqrstuvwxyz"
	uppercaseLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits           = "0123456789"
	specialChars     = "!@#$%^&*()-_=+[]{}|;:,.<>?/~`"
)

func main() {
	password, err := generatePassword(12)
	if err != nil {
		fmt.Println("Ошибка в создании пароля:", err)
		return
	}

	fmt.Println("Сгенерированный пароль:", password)
}

func generatePassword(length int) (string, error) {
	//код
	return "", fmt.Errorf("Длина пароля должна быть больше чем 4 символа")
}
/*
Логика кода:
В main вызывается функция generatePassword с длиной пароля 12 символов.
В generatePassword проверяется, что длина пароля не менее 4 символов.
Создается строка allChars, содержащая все возможные символы для пароля.
Создается срез байтов password длиной length.
Первые length-3 символов пароля заполняются случайными символами из allChars.
Последние три символа пароля заполняются случайными символами из нижнего регистра, верхнего регистра и цифр.
Символы пароля перемешиваются функцией shuffle.
randChar возвращает случайный символ из заданного набора charset с использованием криптографически безопасного генератора случайных чисел.
shuffle перемешивает символы в срезе s с использованием криптографически безопасного генератора случайных чисел.
*/
