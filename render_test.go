package common

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestCreateHtmlFromArticles(t *testing.T) {
	d := time.Date(2021, time.November, 30, 0, 0, 0, 0, time.UTC)
	articles := []*Article{
		&Article{
			Title:         "My article title 1",
			Link:          "https://myfeed1.com/myarticle1",
			Source:        "myfeed1.com",
			PubDateParsed: &d,
			PubDate:       "Fri, 03 Dec 2021 20:23:30",
		},
		&Article{
			Title:         "My different title 2",
			Link:          "https://myfeed1.com/myarticle2",
			Source:        "myfeed1.com",
			PubDateParsed: &d,
			PubDate:       "Fri, 43 Dec 2021 14:23:30",
		},
	}
	htmlBytes, err := CreateHtmlFromArticles(articles)
	assert.Nil(t, err)
	assert.NotEqual(t, len(htmlBytes), 0)
	writeToLocal(false)
}

func writeToLocal(b bool) {
	if !b {
		return
	}
	f, ex := os.Create("out.html")
	assert.Nil(t, ex)
	_, ex2 := f.Write(htmlBytes)
	assert.Nil(t, ex2)
}
