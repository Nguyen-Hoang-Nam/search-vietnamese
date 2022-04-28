package searchvietnamese

import "strings"

func contains(textRune, keywordRune []rune, alphabetText, alphabetKeyword string) bool {
	alphabetKeywordLen := len(alphabetKeyword)

	index := indexAt(alphabetText, alphabetKeyword, 0)
	for index > -1 {
		isMatch := true
		for i := index; i < index+alphabetKeywordLen; i++ {
			if !isMatchVietnamese(textRune[i], keywordRune[i-index]) {
				isMatch = false
			}
		}

		if isMatch {
			return true
		}

		oldIndex := index
		index = indexAt(alphabetText, alphabetKeyword, index)
		if index == oldIndex {
			return false
		}
	}

	return false
}

// Contains check keyword in Vietnamese text.
func Contains(text, keyword string) bool {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := strings.ToUpper(ToAlphabetSensitive(textRune))
	alphabetKeyword := strings.ToUpper(ToAlphabetSensitive(keywordRune))

	return contains(textRune, keywordRune, alphabetText, alphabetKeyword)
}

// ContainsSensitive check keyword in Vietnamese text with case-sensitive.
func ContainsSensitive(text, keyword string) bool {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := ToAlphabetSensitive(textRune)
	alphabetKeyword := ToAlphabetSensitive(keywordRune)

	return contains(textRune, keywordRune, alphabetText, alphabetKeyword)
}

// StrictContains check keyword in Vietnamese text.
// This function assume that your text is valid Vietnamese paragraph. It mean there are no other UTF-8 characters than Vietnamese characters.
func StrictContains(text, keyword string) bool {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := strings.ToUpper(StrictToAlphabetSensitive(textRune))
	alphabetKeyword := strings.ToUpper(StrictToAlphabetSensitive(keywordRune))

	return contains(textRune, keywordRune, alphabetText, alphabetKeyword)
}

// StrictContainsSensitive check keyword in Vietnamese text with case-sensitive.
// This function assume that your text is valid Vietnamese paragraph. It mean there are no other UTF-8 characters than Vietnamese characters.
func StrictContainsSensitive(text, keyword string) bool {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := StrictToAlphabetSensitive(textRune)
	alphabetKeyword := StrictToAlphabetSensitive(keywordRune)

	return contains(textRune, keywordRune, alphabetText, alphabetKeyword)
}
