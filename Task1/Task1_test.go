package main

import (
	"crypto/sha256"
	"encoding/hex"
	"strconv"
	"testing"
)

func TestTypesToString(t *testing.T) {
	// Проверка конкатенации различных типов в строку
	numDecimal := 42
	numOctal := 052
	numHexadecimal := 0x2A
	exp := 2.718
	st := "Hello"
	flag := false
	complexNum := complex64(1 + 1i)

	// Ожидаемое значение после конкатенации
	expectedStr := "4242422.718Hellofalse(1+1i)"

	strDecimal := strconv.Itoa(numDecimal)
	strOctal := strconv.Itoa(numOctal)
	strHexadecimal := strconv.Itoa(numHexadecimal)
	strFloat := strconv.FormatFloat(exp, 'f', -1, 64)
	strFlag := strconv.FormatBool(flag)
	strComplex := strconv.FormatComplex(complex128(complexNum), 'f', -1, 64)
	result := strDecimal + strOctal + strHexadecimal + strFloat + st + strFlag + strComplex

	if result != expectedStr {
		t.Errorf("Ожидалось %s, получено %s", expectedStr, result)
	}
}

func TestHashWithSalt(t *testing.T) {
	// Проверка корректности хэширования с солью
	inputStr := "4242422.718Hellofalse(1+1i)"
	expectedHash := "a928366fd0e2af438b9aed30797ab4dbcc2313163d606c75f2efb0b9f82f1f7c"

	runeSlice := []rune(inputStr)
	byteSlice := []byte(string(runeSlice))
	salt := []byte("go-2024")

	mid := len(byteSlice) / 2
	withSalt := append(byteSlice[:mid], append(salt, byteSlice[mid:]...)...)

	hash := sha256.Sum256(withSalt)
	resultHash := hex.EncodeToString(hash[:])

	if resultHash != expectedHash {
		t.Errorf("Ожидался хэш %s, получен %s", expectedHash, resultHash)
	}
}
