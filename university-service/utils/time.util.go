package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomTime() string {
	now := time.Now()
	min := now.Unix()
	max := now.AddDate(0, 1, 0).Unix()
	delta := max - min

	sec := rand.Int63n(delta) + min
	randomTime := time.Unix(sec, 0)
	return randomTime.Format("2006-01-02 15:04:05")
}
