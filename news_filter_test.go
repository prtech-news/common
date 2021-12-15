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

func TestFilterByTitle(t *testing.T) {
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
			Title:         "aTh Móvil collapsa en todo el país",
			Link:          "https://myfeed1.com/ath-movil",
			Source:        "https://myfeed1.com",
			PubDateParsed: &t3,
			PubDate:       "2021-11-25",
		},
	}
	phrases := map[string]bool{"ATH Móvil": true, "tech": true, "scale": true}
	results := FilterByTitle(sampleArticles, phrases)
	assert.EqualValues(t, 3, len(results))
	for index, a := range results {
		assert.EqualValues(t, expectedArticles[index], a)
	}
}

func TestFilterByTitleSpanish(t *testing.T) {
	expectedArticles := []*Article{
		&Article{
			Title:         "Puerto Rico is booming in tech",
			Link:          "https://myfeed1.com/puertorico-is-booming-in-tech",
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
	phrases := map[string]bool{"ATH Móvil": true, "puerto rico": true}
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

func TestIsPhraseCaseInsensitiveMatchAnscii(t *testing.T) {
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

	matched = isPhraseCaseInsensitiveMatch(text, "pUerTorIcan")
	assert.True(t, matched, fmt.Sprintf("'%s' Should have found a match\n", text))
}

func TestIsPhraseCaseInsensitiveMatchUTF8(t *testing.T) {
	text := "Startup de logística Nuvocargo recauda $20.5 mdd"
	text2 := "Startup de logistica Nuvocargo recauda $20.5 mdd"
	assert.EqualValues(t, 49, len(text))
	assert.EqualValues(t, 48, len(text2))
	var matched bool

	matched = isPhraseCaseInsensitiveMatch(text, "logística")
	assert.True(t, matched, fmt.Sprintf("'%s' Should have found a match\n", text))

	matched = isPhraseCaseInsensitiveMatch(text2, "logistica")
	assert.True(t, matched, fmt.Sprintf("'%s' Should have found a match\n", text))
}

func TestAnyPhraseMatch(t *testing.T) {
	arr := []struct {
		title string
		res   bool
	}{
		{title: "Platzi anuncia ronda de inversión Serie B por $62 mdd", res: false},
		{title: "La fintech Clara se convierte en el quinto unicornio mexicano", res: true},
		{title: "EVLO launches 1 MWh storage system", res: false},
		{title: "La proptech colombiana Aptuno levanta ronda semilla de $5,1 mdd", res: false},
		{title: "Blockchain.com adquiere a la firma de crypto SeSocio y se expande en América Latina", res: true},
		{title: "Haz tus pagos a Popular desde Mi Banco Puerto Rico", res: true},
	}
	phrases := map[string]bool{
		"fintech": true, "blockchain": true, "puErto Rico": true,
	}
	for _, obj := range arr {
		assert.EqualValues(t, obj.res, anyPhraseMatch(obj.title, phrases))
	}
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

	text = "EVLO launches 1 MWh storage system"
	description = "The 1 MWh system is set to make its debut suppling power to Haute-Mauricie during transmission line upgrades.</span></p><p>EVLO, a Hydro-Quebec-owned turnkey energy storage systems provider, announced the launch of the compan"
	assert.True(t, isEnglish(text, description), fmt.Sprintf("'%s' should be classified as english\n", text))
}

func TestIsNotSpanish(t *testing.T) {
	var text string
	var description string
	text = "Puertorican blockchain scales to new highs"
	description = "blockchain in the land of chains"
	assert.True(t, isNotSpanish(text, description), fmt.Sprintf("'%s' should be classified as english\n", text))

	text = "Juan del pueblo es un cabron"
	description = "Un tipo ordinario y muy caballeroso"
	assert.False(t, isNotSpanish(text, description), fmt.Sprintf("'%s' should not be classified as english\n", text))

	text = "Sunrise brief: SEC investigating Tesla over alleged solar system fire negligence"
	description = "Also on the rise: Clean energy bonds issued in California, EV charging stations with integrated advertising space, DART mission uses solar power to redirect Earthbound asteroids, and the Bezos Earth Fund supports tribal-led program."
	assert.True(t, isNotSpanish(text, description), fmt.Sprintf("'%s' should be classified as english\n", text))

	text = "Next-generation solar technologies will drive a lower LCOE"
	description = "Also on the rise: Clean energy bonds issued in California, EV charging stations with integrated advertising space, DART mission uses solar power to redirect Earthbound asteroids, and the Bezos Earth Fund supports tribal-led program."
	assert.True(t, isNotSpanish(text, description), fmt.Sprintf("'%s' should be classified as english\n", text))

	text = "EVLO launches 1 MWh storage system"
	description = "The 1 MWh system is set to make its debut suppling power to Haute-Mauricie during transmission line upgrades.</span></p><p>EVLO, a Hydro-Quebec-owned turnkey energy storage systems provider, announced the launch of the compan"
	assert.True(t, isNotSpanish(text, description), fmt.Sprintf("'%s' should be classified as english\n", text))
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
