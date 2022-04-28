package searchvietnamese

import "unicode/utf8"

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
