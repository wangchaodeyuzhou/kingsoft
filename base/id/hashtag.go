package id

import (
	"fmt"
	"strings"
)

var (
	ALPHABET = "0289PYLQGRJCUV"
	BASE     = len(ALPHABET)
)

func Id2HashTag(num int64) string {
	var sb []byte
	for num > 0 {
		index := int(num % int64(BASE))
		num /= int64(BASE)
		sb = append(sb, ALPHABET[index])
	}

	return reverse(sb)
}

func HashTag2Id(s string) int64 {
	var num int64
	for _, char := range s {
		index := strings.IndexRune(ALPHABET, char)
		if index == -1 {
			panic(fmt.Sprintf("Invalid char: %c", char))
		} else {
			num *= int64(BASE)
			num += int64(index)
		}
	}
	return num
}

func reverse(s []byte) string {
	n := len(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return string(s)
}
