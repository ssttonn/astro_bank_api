package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomInt(min, max int) int {
	return min + int(rand.Int63n(int64(max-min+1)))
}

func RandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[RandomInt(0, len(letter)-1)]
	}
	return string(b)
}

func RandomMoney() int64 {
	return int64(RandomInt(0, 1000))
}

func RandomOwner() string {
	return RandomString(6)
}

func RandomCurrency() string {
	currencies := []string{"USD", "EUR", "CAD"}
	n := len(currencies)
	return currencies[RandomInt(0, n-1)]
}
