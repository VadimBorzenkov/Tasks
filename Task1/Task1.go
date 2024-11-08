package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	// 1)
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var exp float64 = 2.718
	var st string = "Hello"
	var flag bool = false
	var complexNum complex64 = 1 + 1i
	// 2)
	fmt.Println("Типы: ", reflect.TypeOf(numDecimal), reflect.TypeOf(numOctal), reflect.TypeOf(numHexadecimal), reflect.TypeOf(exp), reflect.TypeOf(st), reflect.TypeOf(flag), reflect.TypeOf(complexNum))

	// 3)
	strDecimal := strconv.Itoa(numDecimal)
	strOctal := strconv.Itoa(numOctal)
	strHexadecimal := strconv.Itoa(numHexadecimal)
	strFloat := strconv.FormatFloat(exp, 'f', -1, 64)
	strFlag := strconv.FormatBool(flag)
	strComplex := strconv.FormatComplex(complex128(complexNum), 'f', -1, 64)
	resStr := strDecimal + strOctal + strHexadecimal + strFloat + st + strFlag + strComplex
	fmt.Println("Строка: ", resStr)

	// 4)
	runeSlice := []rune(resStr)
	fmt.Println("Срез рун: ", runeSlice)

	// 5)
	byteSlice := []byte(string(runeSlice))
	salt := []byte("go-2024")

	mid := len(byteSlice) / 2
	withSalt := append(byteSlice[:mid], append(salt, byteSlice[mid:]...)...)

	hash := sha256.Sum256(withSalt)
	fmt.Println("SHA-256 хэш:", hex.EncodeToString(hash[:]))
}
