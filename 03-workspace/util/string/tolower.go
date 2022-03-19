package stringutil

import (
	"unicode"
)

func ToLower(s string) string {
	r := []rune(s)
	for i := range r {
		r[i] = unicode.ToLower(r[i])
	}
	return string(r)
}
