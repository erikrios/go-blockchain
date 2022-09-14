package utils

import (
	"math/rand"
	"time"
)

func RandStr(len uint8) string {
	rand.Seed(time.Now().UnixNano())

	res := make([]byte, len)

	for i := range res {
		randInt := rand.Intn(26)

		randByte := 'a' + byte(randInt)

		res[i] = randByte
	}

	return string(res)
}
