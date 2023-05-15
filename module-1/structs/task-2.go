package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

// validator проверяет строку на соответствие некоторому условию
// и возвращает результат проверки
type validator func(s string) bool

// digits возвращает true, если s содержит хотя бы одну цифру
// согласно unicode.IsDigit(), иначе false
func digits(s string) bool {
	for sym := range s {
		if unicode.IsDigit(rune(sym)) {
			return true
		}
	}
	return false
	// ...
}

// letters возвращает true, если s содержит хотя бы одну букву
// согласно unicode.IsLetter(), иначе false
func letters(s string) bool {
	for sym := range s {
		if unicode.IsLetter(rune(sym)) {
			return true
		}
	}
	return false
}

// minlen возвращает валидатор, который проверяет, что длина
// строки согласно utf8.RuneCountInString() - не меньше указанной
func minlen(length int) validator {
	var val validator = func(s string) bool { return utf8.RuneCountInString(s) >= length }
	return val
}

// and возвращает валидатор, который проверяет, что все
// переданные ему валидаторы вернули true
func and(funcs ...validator) validator {


}

// or возвращает валидатор, который проверяет, что хотя бы один
// переданный ему валидатор вернул true
func or(funcs ...validator) validator {
	// ...
}

// password содержит строку со значением пароля и валидатор
type password struct {
	value string
	validator
}

// isValid() проверяет, что пароль корректный, согласно
// заданному для пароля валидатору
func (p *password) isValid() bool {
	if p.
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

func main() {
	var s string
	fmt.Scan(&s)
	// валидатор, который проверяет, что пароль содержит буквы и цифры,
	// либо его длина не менее 10 символов
	validator := or(and(digits, letters), minlen(10))
	p := password{s, validator}
	fmt.Println(p.isValid())
}
