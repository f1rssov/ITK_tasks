package main

import(
	"testing"
	"strings"
)

func TestMyToString(t *testing.T) {
	res := myToString(42, 0x2A, 052, 3.14, "Go", true, 1+2i)
	expectedParts := []string{"42", "42", "42", "3.14", "Go", "true", "(1+2i)"}

	for _, part := range expectedParts {
		if !strings.Contains(res, part) {
			t.Errorf("Ожидалось, что результат содержит %s, но не содержит", part)
		}
	}
}

func TestHashRunes(t *testing.T) {
	input := []rune("HelloWorld")
	salt := "go-2024"

	hash1 := hashRunes(input, salt)
	hash2 := hashRunes(input, salt)

	if hash1 != hash2 {
		t.Errorf("Хеш должен быть детерминированным, но hash1 != hash2")
	}

	if len(hash1) != 64 {
		t.Errorf("Ожидалась hex строка длиной 64, но получено %d", len(hash1))
	}
}

// func TestGetTypePrint(t *testing.T) {
// 	res := getTypePrint()
// }