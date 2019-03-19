package utils

import (
	"math/rand"
	"time"
)

const tokens = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandomString(min, max int) string {
	size := rand.Intn((max-min)+1) + min
	buffer := make([]byte, size)

	for i := range buffer {
		buffer[i] = tokens[rand.Intn(len(tokens))]
	}

	return string(buffer)
}
