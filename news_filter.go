// Copyright 2021 - present prtech.news. All rights reserved.
package common

import (
	"github.com/abadojack/whatlanggo"
	"strings"
)

// Values should be limited to feed / rss .xml file listed in the config.json 
// and look for top level <Link> value
var feedUrlWhiteList []string

func init() {
	feedUrlWhiteList = []string{
		"joinbased.com",
	}
}

func FilterByTitle(articles []*Article, phrases map[string]bool) []*Article {
	var filtered []*Article = []*Article{}
	seenMap := make(map[string]bool)
	for _, article := range articles {
		if isNotSpanish(article.Title, article.Description) {
			if shouldAppendNonSpanishArticle(article, phrases) {
				filtered = doAppend(filtered, seenMap, article)
			}
		} else {
			// Assumes articles from urls in spanish are always about puerto rico
			// since the sources should be about PR.
			// If generalized to other topics this block should be updated
			if anyPhraseMatch(article.Title, phrases) {
				filtered = doAppend(filtered, seenMap, article)
			}
		}
		seenMap[article.Title] = true
	}
	return filtered
}

func doAppend(arr []*Article, seenMap map[string]bool, article *Article) []*Article {
	_, seen := seenMap[article.Title]
	if seen {
		return arr
	}
	return append(arr, article)
}

func shouldAppendNonSpanishArticle(article *Article, phrases map[string]bool) bool {
	return isWhiteListedSource(article.Source) ||
		(isPrMentionedInTitle(article.Title) && anyPhraseMatch(article.Title, phrases))
}

func isWhiteListedSource(source string) bool {
	for _, domain := range feedUrlWhiteList {
		if strings.ToLower(domain) == strings.ToLower(source) {
			return true
		}
	}
	return false
}

func isPrMentionedInTitle(title string) bool {
	phrases := map[string]bool{
		"puerto rico": true,
		"puertorican": true,
		"boricua":     true,
	}
	return anyPhraseMatch(title, phrases)
}

func isNotSpanish(text string, description string) bool {
	info := whatlanggo.Detect(text + " " + description)
	if "spanish" != strings.ToLower(info.Lang.String()) {
		return true
	}
	if info.Confidence*100 < 60.0 {
		return true
	}
	return false
}

func isEnglish(text string, description string) bool {
	info := whatlanggo.Detect(text + " " + description)
	return "english" == strings.ToLower(info.Lang.String()) &&
		info.Confidence*100 > 60.0
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
