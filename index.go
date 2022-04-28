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

func Index(text, keyword string) int {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := strings.ToUpper(ToAlphabetSensitive(textRune))
	alphabetKeyword := strings.ToUpper(ToAlphabetSensitive(keywordRune))

	return index(textRune, keywordRune, alphabetText, alphabetKeyword)
}

func IndexSensitive(text, keyword string) int {
	textRune := []rune(text)
	keywordRune := []rune(keyword)

	alphabetText := ToAlphabetSensitive(textRune)
	alphabetKeyword := ToAlphabetSensitive(keywordRune)

	return index(textRune, keywordRune, alphabetText, alphabetKeyword)
}
