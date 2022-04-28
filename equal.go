package searchvietnamese

import (
	"strings"
)

func equal(targetRune, sourceRune []rune, alphabetTarget, alphabetSource string) bool {
	alphabetKeywordLen := len(alphabetSource)

	if alphabetSource == alphabetTarget {
		for i := 0; i < alphabetKeywordLen; i++ {
			if !isMatchVietnamese(targetRune[i], sourceRune[i]) {
				return false
			}
		}

		return true
	}

	return false
}

// Equal check two Vietnamese string equal
func Equal(target, source string) bool {
	targetRune := []rune(target)
	sourceRune := []rune(source)

	alphabetTarget := strings.ToUpper(ToAlphabetSensitive(targetRune))
	alphabetSource := strings.ToUpper(ToAlphabetSensitive(sourceRune))

	return equal(targetRune, sourceRune, alphabetTarget, alphabetSource)
}

// EqualSensitive check two Vietnamese string equal with case-sensitive.
func EqualSensitive(target, source string) bool {
	targetRune := []rune(target)
	sourceRune := []rune(source)

	alphabetTarget := ToAlphabetSensitive(targetRune)
	alphabetSource := ToAlphabetSensitive(sourceRune)

	return equal(targetRune, sourceRune, alphabetTarget, alphabetSource)
}
