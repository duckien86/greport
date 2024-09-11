package common

import (
	"math/rand"
	"time"
)

var letter = []rune("asdfghjklqwertyuiopzxcvbnmASDFGHJKLQWERTYUIOPZXCVBNM")

func randSequence(n int) string {
	b := make([]rune, n)

	for i := range b {
		b[i] = letter[RandNumber(99999)%len(letter)]
	}
	return string(b)
}

func GenSalt(length int) string {
	if length < 0 {
		length = 50
	}
	return randSequence(length)
}

func RandNumber(numMax int) int {
	if numMax < 0 {
		numMax = 999999
	}
	s := rand.NewSource(time.Now().UnixNano())
	return rand.New(s).Intn(numMax)
}

// GenerateOTP
func GenerateOTP(length int) string {
	if length < 0 {
		length = 6
	}
	digits := "0123456789"
	b := make([]rune, length)
	for i := range b {
		b[i] = rune(digits[RandNumber(len(digits))])
	}
	return string(b)
}
