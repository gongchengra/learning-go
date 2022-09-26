package main

import (
	"math/rand"
	"os"
	"strconv"
	"time"
)

const letters = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"

func RandomString(n int) string {
	rand.Seed(time.Now().UnixNano())
	s := make([]byte, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

var src = rand.NewSource(time.Now().UnixNano())

const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// Better performance:
// https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-go/22892986#22892986

func RandStringBytesMaskImprSrc(n int) string {
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letters) {
			b[i] = letters[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}

func main() {
	pwdLen := 10
	if len(os.Args) > 1 {
		pwdLen, _ = strconv.Atoi(os.Args[1])
	}
	println(RandomString(pwdLen))
	println(RandStringBytesMaskImprSrc(pwdLen))
	if len(os.Args) > 2 {
		c, _ := strconv.Atoi(os.Args[2])
		for i := 0; i < c; i++ {
			println(RandomString(pwdLen))
		}
	}
}
