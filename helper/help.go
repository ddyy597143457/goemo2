package helper

import (
	"crypto/sha1"
	"fmt"
	"io"
	"math/rand"
	"time"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func Sha1(str string) string {
	h := sha1.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}

func RandString(len int) string {
	bytes := make([]byte, len)
	for i := 0; i < len; i++ {
		b := rand.Intn(26) + 65
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func JsonEncode() {

}
