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

//
