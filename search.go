// Package searchvietnamese provide tool to search Vietnamese text more pleasant.
package searchvietnamese

import (
	"strings"
)

// ToAlphabetSensitive convert slice of rune to a string that replace all Vietnamese characters with alphabet characters, respectively.
func ToAlphabetSensitive(text []rune) string {
	var sb strings.Builder
	sb.Grow(len(text))

	for _, character := range text {
		if alphabet, ok := vietnameseToAlphabetSensitive[character]; ok {
			sb.Write([]byte{string(alphabet)[0]})
		} else {
			sb.Write([]byte{string(character)[0]})
		}
	}

	return sb.String()
}

// StrictToAlphabetSensitive convert slice of rune to a string that replace all Vietnamese characters with alphabet characters, respectively.
// This function assume that your text is valid Vietnamese paragraph. It mean there are no other UTF-8 characters than Vietnamese characters.
func StrictToAlphabetSensitive(text []rune) string {
	var sb strings.Builder
	sb.Grow(len(text))

	for _, character := range text {
		if alphabet, ok := vietnameseToAlphabetSensitive[character]; ok {
			sb.Write([]byte{string(alphabet)[0]})
		} else {
			sb.WriteByte(byte(character))
		}
	}

	return sb.String()
}

func isMatchVietnamese(textCharacter, keywordCharacter rune) bool {
	if _, ok := vietnameseToAlphabetSensitive[keywordCharacter]; !ok {
		return true
	}

	if textCharacter == keywordCharacter {
		return true
	} else if singleTextCharacter, ok := mapCombineVietnameseToVietnameseSensitive[textCharacter]; ok && singleTextCharacter == keywordCharacter {
		return true
	}

	return false
}

// Credit https://stackoverflow.com/questions/25837030/find-index-of-a-substring-in-a-string-with-start-index-specified
func indexAt(text, sep string, position int) int {
	index := strings.Index(text[position:], sep)
	if index > -1 {
		index += position
	}

	return index
}
