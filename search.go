package searchvietnamese

import (
	"strings"
	"unicode/utf8"
)

var vietnameseToAlphabet = map[string]string{
	"Á": "A",
	"á": "a",
	"À": "A",
	"à": "a",
	"Ả": "A",
	"ả": "a",
	"Ã": "A",
	"ã": "a",
	"Ạ": "A",
	"ạ": "a",

	"Ă": "A",
	"ă": "a",
	"Ắ": "A",
	"ắ": "a",
	"Ằ": "A",
	"ằ": "a",
	"Ẳ": "A",
	"ẳ": "a",
	"Ẵ": "A",
	"ẵ": "a",
	"Ặ": "A",
	"ặ": "a",

	"Â": "A",
	"â": "a",
	"Ấ": "A",
	"ấ": "a",
	"Ầ": "A",
	"ầ": "a",
	"Ẩ": "A",
	"ẩ": "a",
	"Ẫ": "A",
	"ẫ": "a",
	"Ậ": "A",
	"ậ": "a",

	"Đ": "D",
	"đ": "d",

	"É": "E",
	"é": "e",
	"È": "E",
	"è": "e",
	"Ẻ": "E",
	"ẻ": "e",
	"Ẽ": "E",
	"ẽ": "e",
	"Ẹ": "E",
	"ẹ": "e",

	"Ê": "E",
	"ê": "e",
	"Ế": "E",
	"ế": "e",
	"Ề": "E",
	"ề": "e",
	"Ể": "E",
	"ể": "e",
	"Ễ": "E",
	"ễ": "e",
	"Ệ": "E",
	"ệ": "e",

	"Í": "I",
	"í": "i",
	"Ì": "I",
	"ì": "i",
	"Ỉ": "I",
	"ỉ": "i",
	"Ĩ": "I",
	"ĩ": "I",
	"Ị": "I",
	"ị": "i",

	"Ó": "O",
	"ó": "o",
	"Ò": "O",
	"ò": "o",
	"Ỏ": "O",
	"ỏ": "o",
	"Õ": "O",
	"õ": "o",
	"Ọ": "O",
	"ọ": "o",

	"Ô": "O",
	"ô": "o",
	"Ố": "O",
	"ố": "o",
	"Ồ": "O",
	"ồ": "o",
	"Ổ": "O",
	"ổ": "o",
	"Ỗ": "O",
	"ỗ": "o",
	"Ộ": "O",
	"ộ": "o",

	"Ơ": "O",
	"ơ": "o",
	"Ớ": "O",
	"ớ": "o",
	"Ờ": "O",
	"ờ": "o",
	"Ở": "O",
	"ở": "o",
	"Ỡ": "O",
	"ỡ": "o",
	"Ợ": "O",
	"ợ": "o",

	"Ú": "U",
	"ú": "u",
	"Ù": "U",
	"ù": "u",
	"Ủ": "U",
	"ủ": "u",
	"Ũ": "U",
	"ũ": "u",
	"Ụ": "U",
	"ụ": "u",

	"Ư": "U",
	"ư": "u",
	"Ứ": "U",
	"ứ": "u",
	"Ừ": "U",
	"ừ": "u",
	"Ử": "U",
	"ử": "u",
	"Ữ": "U",
	"ữ": "u",
	"Ự": "U",
	"ự": "u",

	"Ý": "Y",
	"ý": "y",
	"Ỳ": "Y",
	"ỳ": "y",
	"Ỷ": "Y",
	"ỷ": "y",
	"Ỹ": "Y",
	"ỹ": "y",
	"Ỵ": "Y",
	"ỵ": "y",
}

var mapCombineVietnameseToVietnamese = map[string]string{
	"Ắ": "Ă",
	"ắ": "ă",
	"Ằ": "Ă",
	"ằ": "ă",
	"Ẳ": "Ă",
	"ẳ": "ă",
	"Ẵ": "Ă",
	"ẵ": "ă",
	"Ặ": "Ă",
	"ặ": "ă",

	"Ấ": "Â",
	"ấ": "â",
	"Ầ": "Â",
	"ầ": "â",
	"Ẩ": "Â",
	"ẩ": "â",
	"Ẫ": "Â",
	"ẫ": "â",
	"Ậ": "Â",
	"ậ": "â",

	"Ế": "Ê",
	"ế": "ê",
	"Ề": "Ê",
	"ề": "ê",
	"Ể": "Ê",
	"ể": "ê",
	"Ễ": "Ê",
	"ễ": "ê",
	"Ệ": "Ê",
	"ệ": "ê",

	"Ố": "Ô",
	"ố": "ô",
	"Ồ": "Ô",
	"ồ": "ô",
	"Ổ": "Ô",
	"ổ": "ô",
	"Ỗ": "Ô",
	"ỗ": "ô",
	"Ộ": "Ô",
	"ộ": "ô",

	"Ớ": "Ơ",
	"ớ": "ơ",
	"Ờ": "Ơ",
	"ờ": "ơ",
	"Ở": "Ơ",
	"ở": "ơ",
	"Ỡ": "Ơ",
	"ỡ": "ơ",
	"Ợ": "Ơ",
	"ợ": "ơ",

	"Ứ": "Ư",
	"ứ": "ư",
	"Ừ": "Ư",
	"ừ": "ư",
	"Ử": "Ư",
	"ử": "ư",
	"Ữ": "Ư",
	"ữ": "ư",
	"Ự": "Ư",
	"ự": "ư",
}

func ToAlphabet(text string) string {
	var sb strings.Builder

	for _, character := range text {
		character := string(character)
		if alphabet, ok := vietnameseToAlphabet[character]; ok {
			sb.WriteString(alphabet)
		} else {
			sb.WriteString(character)
		}
	}

	return sb.String()
}

func isMatchVietnamese(textCharacter, keywordCharacter string) bool {
	if _, ok := vietnameseToAlphabet[keywordCharacter]; ok {
		if textCharacter == keywordCharacter {
			return true
		} else if singleTextCharacter, ok := mapCombineVietnameseToVietnamese[textCharacter]; ok {
			if singleTextCharacter == keywordCharacter {
				return true
			}
		}

		return false
	}

	return true
}

// Credit https://stackoverflow.com/questions/25837030/find-index-of-a-substring-in-a-string-with-start-index-specified
func indexAt(text, sep string, position int) int {
	index := strings.Index(text[position:], sep)
	if index > -1 {
		index += position
	}

	return index
}

func Contains(text, keyword string) bool {
	alphabetText := strings.ToUpper(ToAlphabet(text))
	alphabetKeyword := strings.ToUpper(ToAlphabet(keyword))
	alphabetKeywordLen := len(alphabetKeyword)

	textRune := []rune(text)
	keywordRune := []rune(keyword)
	index := indexAt(alphabetText, alphabetKeyword, 0)
	for index > -1 {
		isMatch := true
		for i := index; i < index+alphabetKeywordLen; i++ {
			textCharacter := string(textRune[i])
			keywordCharacter := string(keywordRune[i-index])

			if !isMatchVietnamese(textCharacter, keywordCharacter) {
				isMatch = false
			}
		}

		if isMatch {
			return true
		}

		oldIndex := index
		index = indexAt(alphabetText, alphabetKeyword, index)
		if index == oldIndex {
			break
		}
	}

	return false
}

func Index(text, keyword string) int {
	alphabetText := strings.ToUpper(ToAlphabet(text))
	alphabetKeyword := strings.ToUpper(ToAlphabet(keyword))
	alphabetKeywordLen := len(alphabetKeyword)

	textRune := []rune(text)
	keywordRune := []rune(keyword)
	index := indexAt(alphabetText, alphabetKeyword, 0)
	for index > -1 {
		isMatch := true
		for i := index; i < index+alphabetKeywordLen; i++ {
			textCharacter := string(textRune[i])
			keywordCharacter := string(keywordRune[i-index])

			if !isMatchVietnamese(textCharacter, keywordCharacter) {
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

// Case sensitive
func ContainsSensitive(text, keyword string) bool {
	alphabetText := ToAlphabet(text)
	alphabetKeyword := ToAlphabet(keyword)
	alphabetKeywordLen := len(alphabetKeyword)

	textRune := []rune(text)
	keywordRune := []rune(keyword)
	index := indexAt(alphabetText, alphabetKeyword, 0)
	for index > -1 {
		isMatch := true
		for i := index; i < index+alphabetKeywordLen; i++ {
			textCharacter := string(textRune[i])
			keywordCharacter := string(keywordRune[i-index])

			if !isMatchVietnamese(textCharacter, keywordCharacter) {
				isMatch = false
			}
		}

		if isMatch {
			return true
		}

		oldIndex := index
		index = indexAt(alphabetText, alphabetKeyword, index)
		if index == oldIndex {
			break
		}
	}

	return false
}

func IndexSensitive(text, keyword string) int {
	alphabetText := ToAlphabet(text)
	alphabetKeyword := ToAlphabet(keyword)
	alphabetKeywordLen := len(alphabetKeyword)

	textRune := []rune(text)
	keywordRune := []rune(keyword)
	index := indexAt(alphabetText, alphabetKeyword, 0)
	for index > -1 {
		isMatch := true
		for i := index; i < index+alphabetKeywordLen; i++ {
			textCharacter := string(textRune[i])
			keywordCharacter := string(keywordRune[i-index])

			if !isMatchVietnamese(textCharacter, keywordCharacter) {
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

func Equal(target, source string) bool {
	alphabetTarget := strings.ToUpper(ToAlphabet(target))
	alphabetSource := strings.ToUpper(ToAlphabet(source))

	alphabetKeywordLen := len(alphabetSource)

	if alphabetSource == alphabetTarget {
		targetRune := []rune(target)
		sourceRune := []rune(source)

		for i := 0; i < alphabetKeywordLen; i++ {
			targetCharacter := string(targetRune[i])
			sourceCharacter := string(sourceRune[i])

			if !isMatchVietnamese(targetCharacter, sourceCharacter) {
				return false
			}
		}

		return true
	}

	return false
}

// Credit https://github.com/lithammer/fuzzysearch/blob/master/fuzzy/fuzzy.go
func FuzzyMatch(text, keyword string) bool {
	lenDiff := len(text) - len(keyword)

	if lenDiff < 0 {
		return false
	}

	if lenDiff == 0 && Equal(text, keyword) {
		return true
	}

	keywordRune := []rune(keyword)
	keywordLength := len(keywordRune)

Outer:
	for i := 0; i < keywordLength; i++ {

		textRune := []rune(text)

		textLength := len(textRune)
		for j := 0; j < textLength; j++ {
			if Equal(string(textRune[j]), string(keywordRune[i])) {
				text = text[j+utf8.RuneLen(textRune[j]):]
				continue Outer
			}
		}

		return false
	}

	return true
}

func EqualSensitive(target, source string) bool {
	alphabetTarget := ToAlphabet(target)
	alphabetSource := ToAlphabet(source)

	alphabetKeywordLen := len(alphabetSource)

	if alphabetSource == alphabetTarget {
		targetRune := []rune(target)
		sourceRune := []rune(source)

		for i := 0; i < alphabetKeywordLen; i++ {
			targetCharacter := string(targetRune[i])
			sourceCharacter := string(sourceRune[i])

			if !isMatchVietnamese(targetCharacter, sourceCharacter) {
				return false
			}
		}

		return true
	}

	return false
}

// Credit https://github.com/lithammer/fuzzysearch/blob/master/fuzzy/fuzzy.go
func FuzzyMatchSensitive(text, keyword string) bool {
	lenDiff := len(text) - len(keyword)

	if lenDiff < 0 {
		return false
	}

	if lenDiff == 0 && EqualSensitive(text, keyword) {
		return true
	}

	keywordRune := []rune(keyword)
	keywordLength := len(keywordRune)

Outer:
	for i := 0; i < keywordLength; i++ {

		textRune := []rune(text)

		textLength := len(textRune)
		for j := 0; j < textLength; j++ {
			if EqualSensitive(string(textRune[j]), string(keywordRune[i])) {
				text = text[j+utf8.RuneLen(textRune[j]):]
				continue Outer
			}
		}

		return false
	}

	return true
}
