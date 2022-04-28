package searchvietnamese

import "strings"

func index(textRune, keywordRune []rune, alphabetText, alphabetKeyword string) int {
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
			return index
		}

		index = indexAt(alphabetText, alphabetKeyword, index)
	}

	return -1
}

// Index find keyword in Vietnamese text.
func Index(text, keyword string) int {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := strings.ToUpper(ToAlphabetSensitive(textRune))
	alphabetKeyword := strings.ToUpper(ToAlphabetSensitive(keywordRune))

	return index(textRune, keywordRune, alphabetText, alphabetKeyword)
}

// IndexSensitive find keyword in Vietnamese text with case-sensitive.
func IndexSensitive(text, keyword string) int {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := ToAlphabetSensitive(textRune)
	alphabetKeyword := ToAlphabetSensitive(keywordRune)

	return index(textRune, keywordRune, alphabetText, alphabetKeyword)
}

// StrictIndex find keyword in Vietnamese text.
// This function assume that your text is valid Vietnamese paragraph. It mean there are no other UTF-8 characters than Vietnamese characters.
func StrictIndex(text, keyword string) int {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := strings.ToUpper(StrictToAlphabetSensitive(textRune))
	alphabetKeyword := strings.ToUpper(StrictToAlphabetSensitive(keywordRune))

	return index(textRune, keywordRune, alphabetText, alphabetKeyword)
}

// StrictIndexSensitive find keyword in Vietnamese text with case-sensitive.
// This function assume that your text is valid Vietnamese paragraph. It mean there are no other UTF-8 characters than Vietnamese characters.
func StrictIndexSensitive(text, keyword string) int {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := StrictToAlphabetSensitive(textRune)
	alphabetKeyword := StrictToAlphabetSensitive(keywordRune)

	return index(textRune, keywordRune, alphabetText, alphabetKeyword)
}
