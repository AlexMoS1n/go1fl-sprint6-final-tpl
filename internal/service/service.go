package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func DecoderMorseCode(message string) (string, error) {
	code := strings.TrimSpace(message)
	if code == "" {
		return "", errors.New("an empty string is received")
	}

	if getIsMorse(code) {
		return morse.ToText(code), nil
	}
	return morse.ToMorse(code), nil
}

func getIsMorse(code string) bool {
	return !strings.ContainsFunc(code, func(c rune) bool {
		return c != '-' && c != '.' && !unicode.IsSpace(c)
	})
}
