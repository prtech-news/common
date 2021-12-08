// Copyright 2021 - present prtech.news. All rights reserved.
package common

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	sampleArticles []*Article
	t1             time.Time
	t2             time.Time
	t3             time.Time
)

func init() {
	t1 = time.Date(2021, time.November, 21, 0, 0, 0, 0, time.UTC)
	t2 = time.Date(2021, time.November, 30, 0, 0, 0, 0, time.UTC)
	t3 = time.Date(2021, time.November, 25, 0, 0, 0, 0, time.UTC)
	sampleArticles = []*Article{
		&Article{
			Title:         "My article title 1",
			Link:          "https://myfeed1.com/myarticle1",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t1,
			PubDate:       "2021-11-21",
		},
		&Article{
			Title:         "Hola titulo differente",
			Link:          "https://myfeed1.com/myarticle2",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t2,
			PubDate:       "2021-11-30",
		},
		&Article{
			Title:         "Puerto Rico is booming in tech",
			Link:          "https://myfeed1.com/puertorico-is-booming-in-tech",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-25",
		},
		&Article{
			Title:         "Puertorican blockchain scales to new highs",
			Link:          "https://myfeed1.com/puertorican-blockchain-scales",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-25",
		},
		&Article{
			Title:         "Ethereum 2.0 blockchain scales to new highs",
			Link:          "https://myfeed1.com/eth-blockchain-scales",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-25",
		},
		&Article{
			Title:         "aTh Móvil collapsa en todo el país",
			Link:          "https://myfeed1.com/ath-movil",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-25",
		},
	}
}

func TestFilterByTitleSpanish(t *testing.T) {
	expectedArticles := []*Article{
		&Article{
			Title:         "Hola titulo differente",
			Link:          "https://myfeed1.com/myarticle2",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t2,
			PubDate:       "2021-11-30",
		},
		&Article{
			Title:         "aTh Móvil collapsa en todo el país",
			Link:          "https://myfeed1.com/ath-movil",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-25",
		},
	}
	phrases := map[string]bool{"hola": true, "ATH Móvil": true}
	results := FilterByTitle(sampleArticles, phrases)
	assert.EqualValues(t, 2, len(results))
	for index, a := range results {
		assert.EqualValues(t, expectedArticles[index], a)
	}
}

func TestFilterByTitleInEnglish(t *testing.T) {
	expectedArticles := []*Article{
		&Article{
			Title:         "Puerto Rico is booming in tech",
			Link:          "https://myfeed1.com/puertorico-is-booming-in-tech",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-25",
		},
		&Article{
			Title:         "Puertorican blockchain scales to new highs",
			Link:          "https://myfeed1.com/puertorican-blockchain-scales",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-25",
		},
		&Article{
			Title:         "Puertorican caca is not in list",
			Link:          "https://myfeed1.com/puertorican-caca-is-not-in-list",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-26",
		},
	}
	phrases := map[string]bool{"blockchain": true, "tech": true}
	results := FilterByTitle(sampleArticles, phrases)
	assert.EqualValues(t, 2, len(results))
	for index, artcl := range results {
		assert.EqualValues(t, expectedArticles[index], artcl)
	}
}

func TestIsPhraseCaseInsensitiveMatch(t *testing.T) {
	text := "Puertorican blockchain scales to new highs"
	var matched bool

	matched = isPhraseCaseInsensitiveMatch(text, "blockchain scales")
	assert.True(t, matched, fmt.Sprintf("'%s' Should have found a match\n", text))

	matched = isPhraseCaseInsensitiveMatch(text, "block")
	assert.True(t, matched, fmt.Sprintf("'%s' Should have found a match\n", text))

	matched = isPhraseCaseInsensitiveMatch(text, " to new")
	assert.True(t, matched, fmt.Sprintf("'%s' Should have found a match\n", text))

	matched = isPhraseCaseInsensitiveMatch(text, "caca")
	assert.False(t, matched, fmt.Sprintf("'%s' Should no have found a match\n", text))

	matched = isPhraseCaseInsensitiveMatch(text, "ScaLes tO ")
	assert.True(t, matched, fmt.Sprintf("'%s' Should have found a match\n", text))

	matched = isPhraseCaseInsensitiveMatch(text, "highs")
	assert.True(t, matched, fmt.Sprintf("'%s' Should have found a match\n", text))
}

func TestIsEnglish(t *testing.T) {
	var text string
	var description string
	text = "Puertorican blockchain scales to new highs"
	description = "blockchain in the land of chains"
	assert.True(t, isEnglish(text, description), fmt.Sprintf("'%s' should be classified as english\n", text))

	text = "Juan del pueblo es un cabron"
	description = "Un tipo ordinario y muy caballeroso"
	assert.False(t, isEnglish(text, description), fmt.Sprintf("'%s' should not be classified as english\n", text))

	text = "Sunrise brief: SEC investigating Tesla over alleged solar system fire negligence"
	description = "Also on the rise: Clean energy bonds issued in California, EV charging stations with integrated advertising space, DART mission uses solar power to redirect Earthbound asteroids, and the Bezos Earth Fund supports tribal-led program."
	assert.True(t, isEnglish(text, description), fmt.Sprintf("'%s' should be classified as english\n", text))

	text = "Next-generation solar technologies will drive a lower LCOE"
	description = "Also on the rise: Clean energy bonds issued in California, EV charging stations with integrated advertising space, DART mission uses solar power to redirect Earthbound asteroids, and the Bezos Earth Fund supports tribal-led program."
	assert.True(t, isEnglish(text, description), fmt.Sprintf("'%s' should be classified as english\n", text))
}

func TestIsPuertoRicoInTitle(t *testing.T) {
	var title string = "Puerto Rico power grid is destroyed"
	assert.True(t, isPrMentionedInTitle(title))
	title = "Puertorican singer top spotify artist world wide"
	assert.True(t, isPrMentionedInTitle(title))
	title = "Boricuas in the Bronx rebel against the system"
	assert.True(t, isPrMentionedInTitle(title))
	title = "SEC: Tesla system is down"
	assert.False(t, isPrMentionedInTitle(title))
	title = "Goldman is ripping off the U.S. national debt"
	assert.False(t, isPrMentionedInTitle(title))
}
