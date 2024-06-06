package main

import (
	"testing"
	"unicode"
)

func TestGeneratePassword(t *testing.T) {
	length := 12
	password, err := generatePassword(length)
	if err != nil {
		t.Fatalf("Ожидалась ошибка без кода, получено %v", err)
	}

	if len(password) != length {
		t.Errorf("Ожидалась длинна пароля %d, получено %d", length, len(password))
	}

	var hasLower, hasUpper, hasDigit, hasSpecial bool
	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsDigit(char):
			hasDigit = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	if !hasLower {
		t.Error("ожидалось хотя бы одно строчное буквенное значение")
	}
	if !hasUpper {
		t.Error("ожидалось хотя бы одно заглавное буквенное значение")
	}
	if !hasDigit {
		t.Error("ожидалась хотя бы одна цифра")
	}
	if !hasSpecial {
		t.Error("ожидался хотя бы один специальный символ")
	}
}

func TestRandChar(t *testing.T) {
	charset := lowercaseLetters + uppercaseLetters + digits + specialChars
	for i := 0; i < 100; i++ {
		char, err := randChar(charset)
		if err != nil {
			t.Fatalf("ожидалась ошибка отсутствия, получено %v", err)
		}
		if !isInCharset(char, charset) {
			t.Errorf("символ %q не входит в набор символов", char)
		}
	}
}

func isInCharset(char byte, charset string) bool {
	for i := 0; i < len(charset); i++ {
		if char == charset[i] {
			return true
		}
	}
	return false
}
