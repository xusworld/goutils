package goutils

import (
	"math/rand"
	"strings"
	"time"
)

var charset = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func RandomStringFromCharset(length int) string {
	rand.Seed(time.Now().UnixNano())

	var builder strings.Builder
	for idx := 0; idx < length; idx++ {
		builder.WriteRune(charset[rand.Intn(len(charset))])
	}
	return builder.String()
}

func RandInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Int()
}

func RandIntn(n int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(n)
}

func RandIntSlice(length int, upperBound int) []int {
	rand.Seed(time.Now().UnixNano())
	slice := make([]int, length)
	for idx := 0; idx < length; idx++ {
		slice[idx] = RandIntn(upperBound)
	}
	return slice
}

func RandByteSlice(length int) []byte {
	rand.Seed(time.Now().UnixNano())
	slice := make([]byte, length)
	rand.Read(slice)
	return slice
}

func RandStringSlice(sliceLen, strMaxLen int) []string {
	rand.Seed(time.Now().UnixNano())
	slice := make([]string, sliceLen)
	for idx := 0; idx < sliceLen; idx++ {
		slice[idx] = RandomStringFromCharset(strMaxLen)
	}
	return slice
}
