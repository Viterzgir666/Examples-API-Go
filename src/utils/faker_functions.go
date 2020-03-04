package utils

import (
	"github.com/brianvoe/gofakeit"
)

var GenerateTitle = func() string {
	return gofakeit.HipsterWord()
}

var GenerateDescription = func() string {
	words := GenerateIntegerInRange(2, 10)
	return gofakeit.HipsterSentence(words)
}

var GenerateIntegerInRange = func(min int, max int) int {
	return gofakeit.Number(min, max)
}
