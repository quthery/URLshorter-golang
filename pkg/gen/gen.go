package gen

import (
	"strings"
	"time"

	"golang.org/x/exp/rand"
)

func GenerateAlias(length int) string {
	sec := rand.New(rand.NewSource(uint64(time.Now().UnixNano())))
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[sec.Intn(len(chars))])
	}
	str := b.String()

	return str
}
