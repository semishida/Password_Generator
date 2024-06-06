package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
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
	if length < 4 {
		return "", fmt.Errorf("Длина пароля должна быть больше чем 4 символа")
	}

	//Создаем срез, содержащий все символы, из которых будет генерироваться пароль
	allChars := lowercaseLetters + uppercaseLetters + digits + specialChars

	//Создаем результируюзий срез, в который будем помещать символы пароля
	password := make([]byte, length)

	//Выбираем по одному случайному символу из каждой категории и добавляем их в пароль
	for i := 0; i < length-3; i++ {
		char, err := randChar(allChars)
		if err != nil {
			return "", err
		}
		password[i] = char
	}

	//Выбираем по одному случайному символу из каждой категории и добавляем их в пароль
	password[length-3], _ = randChar(lowercaseLetters)
	password[length-2], _ = randChar(uppercaseLetters)
	password[length-1], _ = randChar(digits)

	//Перемешиваем символы пароля
	shuffle(password)

	return string(password), nil
}

func randChar(charset string) (byte, error) {
	max := big.NewInt(int64(len(charset)))
	n, err := rand.Int(rand.Reader, max)
	if err != nil {
		return 0, err
	}
	return charset[n.Int64()], nil
}

func shuffle(s []byte) {
	n := len(s)
	for i := n - 1; i > 0; i-- {
		j, err := rand.Int(rand.Reader, big.NewInt(int64(i+1)))
		if err != nil {
			return
		}
		s[i], s[j.Int64()] = s[j.Int64()], s[i]
	}
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
