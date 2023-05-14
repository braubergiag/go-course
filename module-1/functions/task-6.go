package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	phrase := readString()
	var abbr []rune
	for _, word := range strings.Fields(phrase) {
		firstLetter := []rune(word)[0]
		if !unicode.IsLetter(firstLetter) {
			continue
		}
		abbr = append(abbr, unicode.ToUpper(firstLetter))
	}

	fmt.Println(string(abbr))
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readString читает строку из `os.Stdin` и возвращает ее
func readString() string {
	rdr := bufio.NewReader(os.Stdin)
	str, _ := rdr.ReadString('\n')
	return str
}
