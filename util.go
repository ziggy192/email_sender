package email_sender

import (
	"crypto/rand"
	"fmt"
	"log"
)

// Some alphabets
const (
	UpperEnglishLetters    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	LowerEnglishLetters    = "abcdefghijklmnopqrstuvwxyz"
	DecimalDigits          = "0123456789"
	EnglishLetters         = UpperEnglishLetters + LowerEnglishLetters
	AlphaNumericCharacters = EnglishLetters + DecimalDigits
)

func LogErr(v ...any) {
	args := make([]any, 0, 1+len(v))
	args = append(append(args, "[error] "), v...)
	_ = log.Output(2, fmt.Sprint(args...))
}

// RandomString generates a random string consisting of characters in the provided alphabet
func RandomString(n int, alphabet string) string {
	if n <= 0 {
		return ""
	}

	r := []rune(alphabet)
	k := byte(len(r))

	s := make([]rune, n)
	b := make([]byte, n)

	_, _ = rand.Read(b)
	for i := range b {
		s[i] = r[b[i]%k]
	}

	return string(s)
}
