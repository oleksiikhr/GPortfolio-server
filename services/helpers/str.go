package helpers

import "math/rand"

// Rnd generate a random of symbols specified length
func Rnd(n int) string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	letterRunesLen := len(letterRunes)
	b := make([]rune, n)

	for i := range b {
		b[i] = letterRunes[rand.Intn(letterRunesLen)]
	}

	return string(b)
}
