package helpers

import (
	"math/rand"
	"time"
)

func RandomNember(n int) int {
	rand.Seed(time.Now().Unix())
	value := rand.Intn(n)

	return value
}