package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCreateHtmlFromArticles(t *testing.T) {
	d := time.Date(2021, time.November, 30, 0, 0, 0, 0, time.UTC)
	articles := []*Article{
		&Article{
			Title:         "My article title 1",
			Link:          "https://myfeed1.com/myarticle1",
			Source:        "https://myfeed1.com",
			PubDateParsed: &d,
			PubDate:       "2021-11-30",
		},
		&Article{
			Title:         "My different title 2",
			Link:          "https://myfeed1.com/myarticle2",
			Source:        "https://myfeed1.com",
			PubDateParsed: &d,
			PubDate:       "2021-11-30",
		},
	}
	bytez, err := CreateHtmlFromArticles(articles)
	assert.Nil(t, err)
	assert.NotEqual(t, len(bytez), 0)
}
