package util

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefgijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max-min+1) // 0->max-min
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomOwner generates a random owner name
func RandomOwner() string {
	return RandomString(6)
}

// RandomMoney  generates a random money
func RandomMoney() int64 {
	return RandomInt(0, 1000)
}

// RandomCurrency  generates a random currency
func RandomCurrency() string {
	currencies := []string{"EUR", "BDT", "USD"}
	n := len(currencies)
	return currencies[rand.Intn(n)]
}

// RandomEmail generates a random email address
func RandomEmail() string {
	return fmt.Sprintf("%s@gmail.com", RandomString(6))
}
