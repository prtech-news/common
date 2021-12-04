// Copyright 2021 - present prtech.news. All rights reserved.
package common

import (
	"github.com/mmcdole/gofeed"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var (
	mockFeed      *gofeed.Feed
	rssFeedParser *RSSFeedParser
)

// Init - Prepare mock data for test
func init() {
	fp := gofeed.NewParser()
	mockFeed, _ = fp.ParseString(
		`<rss version="2.0">
		<channel>
		<title>Sample Feed</title>
		</channel>
		</rss>`,
	)
	// Setup mock RSSFeedParser
	rssFeedParser = &RSSFeedParser{func(url string) (*gofeed.Feed, error) {
		return mockFeed, nil
	}}
}

func TestRSSFeedParserSuccess(t *testing.T) {
	feed, err := rssFeedParser.ParseRSSFeed("https://mywebsite.com/feed")
	assert.NotNil(t, feed)
	assert.Nil(t, err)
	assert.EqualValues(t, mockFeed, feed)
}

func TestParseRSSFeedsAsync(t *testing.T) {
	urls := []string{
		"https://mywebsite.com/feed1",
		"https://mywebsite.com/feed2",
		"https://mywebsite.com/feed3",
	}
	feeds := ParseRSSFeedsAsync(rssFeedParser, urls)
	assert.EqualValues(t, 3, len(feeds))
	for _, feed := range feeds {
		assert.EqualValues(t, mockFeed, feed)
	}
}

func TestFromRSSToArticle(t *testing.T) {
	d := time.Date(2021, time.November, 30, 0, 0, 0, 0, time.UTC)
	feeds := []*gofeed.Feed{
		{
			Link:     "https://myfeed1.com",
			FeedLink: "https://myfeed1.com",
			Items: []*gofeed.Item{
				{
					Title:           "My article title 1",
					Link:            "https://myfeed1.com/myarticle1",
					Published:       "2021-11-30",
					PublishedParsed: &d,
					Author: &gofeed.Person{
						Name:  "TestUser",
						Email: "TestUser@testmail.com",
					},
				},
			},
		},
	}
	expectedArticles := []*Article{
		&Article{
			Title:         "My article title 1",
			Link:          "https://myfeed1.com/myarticle1",
			Source:        "https://myfeed1.com",
			PubDateParsed: &d,
			PubDate:       "2021-11-30",
		},
	}
	articles := FromRSSToArticle(feeds)

	for index, a := range articles {
		assert.EqualValues(t, expectedArticles[index], a)
	}
}
