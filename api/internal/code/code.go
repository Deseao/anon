package code

import "math/rand"

const CODE_LEN int = 10

//Thanks Icza! https://stackoverflow.com/a/31832326/1162540
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func GenRandCode(length int) string {
	code := make([]rune, length)
	for i := range code {
		code[i] = letters[rand.Intn(len(letters))]
	}
	return string(code)
}
