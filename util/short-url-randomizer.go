package util

import "math/rand"

var runes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func NewRandomShortUrl(len int) string {
	short_url := make([]rune, len)
	for i := range short_url {
		short_url[i] = runes[rand.Intn(len)]
	}

	return string(short_url)
}
