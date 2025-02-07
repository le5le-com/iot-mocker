package utils

import (
	cryptoRand "crypto/rand"
	"fmt"
	"math/rand"
	"time"

	"github.com/dchest/captcha"
)

var numChars = []byte("0123456789")
var chars = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func GetRandHex(n uint) string {
	var size uint
	if n > 0 {
		if n%2 == 1 {
			size = n/2 + 1
		} else {
			size = n / 2
		}
	}

	b := make([]byte, size)
	_, err := cryptoRand.Read(b)
	if err != nil {

		return ""
	}

	if n%2 == 0 {
		return fmt.Sprintf("%x", b)
	}

	text := fmt.Sprintf("%x", b)
	return text[0:n]
}

// GetRandCode 获取一个数据code
func GetRandCode(strLen uint8) string {
	b := captcha.RandomDigits(int(strLen))
	for i, c := range b {
		b[i] = numChars[c]
	}
	return string(b)
}

func GetRandString(n int, allowedChars ...[]rune) string {
	var letters []rune

	if len(allowedChars) == 0 {
		letters = chars
	} else {
		letters = allowedChars[0]
	}

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func GetRandInt(min, max int) int {
	if min == max {
		return min
	}
	if min > max {
		return rand.Intn(min-max) + max
	}
	return rand.Intn(max-min) + min
}

func GetRandFloat(min, max float64) float64 {
	if min == max {
		return min
	}
	if min > max {
		return rand.Float64()*(min-max) + max
	}
	return rand.Float64()*(max-min) + min
}

func GetRandBool() bool {
	return rand.Int()%2 == 1
}

func GetRandDate(min, max time.Time) time.Time {
	delta := max.Unix() - min.Unix()

	// 生成一个随机秒数
	sec := rand.Int63n(delta) + min.Unix()

	// 返回随机时间
	return time.Unix(sec, 0)
}
