// Copyright 2021 - present prtech.news. All rights reserved.
package common

import (
	"github.com/abadojack/whatlanggo"
	"strings"
)

func FilterByTitle(articles []*Article, phrases map[string]bool) []*Article {
	filtered := []*Article{}
	for _, article := range articles {
		if isEnglish(article.Title) {
			if isPrMentionedInTitle(article.Title) && anyPhraseMatch(article.Title, phrases) {
				filtered = append(filtered, article)
			}
		} else {
			// Assumes articles from urls in spanish are always about puerto rico
			// since the sources should be about PR.
			// If generalized to other topics this block should be updated
			if anyPhraseMatch(article.Title, phrases) {
				filtered = append(filtered, article)
			}
		}
	}
	return filtered
}

func isPrMentionedInTitle(title string) bool {
	phrases := map[string]bool{"Puerto Rico": true, "puertorican": true}
	return anyPhraseMatch(title, phrases)
}

func isEnglish(text string) bool {
	info := whatlanggo.Detect(text)
	return "english" == strings.ToLower(info.Lang.String())
}

func anyPhraseMatch(text string, phrases map[string]bool) bool {
	for phrase, _ := range phrases {
		if isPhraseCaseInsensitiveMatch(text, phrase) {
			return true
		}
	}
	return false
}

// Scans a string in order to find a matching substring phrase (can include whitepace)
// case insensitive
func isPhraseCaseInsensitiveMatch(text string, phrase string) bool {
	if len(phrase) > len(text) {
		return false
	}
	n := len(text)
	windowSize := len(phrase)
	for i := 0; i <= n-windowSize; i++ {
		if strings.ToLower(text[i:i+windowSize]) == strings.ToLower(phrase) {
			return true
		}
	}
	return false
}
