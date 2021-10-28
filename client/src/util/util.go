package util

import (
	"math/rand"
	"strings"
	"time"
)

const CHARSET = "ABCDEFGHIKLMNOPQRSTUVWXYZ1234567890"

func GenerateString(length int) string {
	rand.Seed(time.Now().Unix() * int64(rand.Intn(1337)))
	var str = make([]string, length)

	for i := 0; i < length; i++ {
		str[i] = string(CHARSET[rand.Intn(len(CHARSET))])
	}

	return strings.Join(str, "")
}
