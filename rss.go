// Copyright 2021 - present prtech.news. All rights reserved.
package common

import (
	"github.com/mmcdole/gofeed"
	"log"
	"strings"
	"time"
)

type Article struct {
	Title         string     `json:"title,omitempty"`
	Description   string     `json:description,omitempty"`
	Link          string     `json:"link,omitempty"`
	Source        string     `json:"source,omitempty"`
	PubDateParsed *time.Time `json:"pubDateParsed,omitempty"`
	PubDate       string     `json:"pubDate,omitempty"`
}

var (
	fp *gofeed.Parser
)

type RSSFeedParser struct {
	ParseRSSFeedFromUrl func(string) (*gofeed.Feed, error)
}

func (p *RSSFeedParser) ParseRSSFeed(url string) (*gofeed.Feed, error) {
	if p.ParseRSSFeedFromUrl != nil {
		return p.ParseRSSFeedFromUrl(url)
	}
	// Default to gofeed.ParseUrl
	return fp.ParseURL(url)
}

func init() {
	fp = gofeed.NewParser()
	fp.UserAgent = "prtech.news crawler"
}

// Parse RSS feeds from a list of urls asynchronously
// Takes a &RSSFeedParser{nil} in order to default to fp.ParseURL
func ParseRSSFeedsAsync(parser *RSSFeedParser, urls []string) []*gofeed.Feed {

	feedChan := make(chan *gofeed.Feed)
	done := make(chan bool)

	n := len(urls)

	for i := 0; i < n; i++ {
		go func(url string) {

			feed, err := parser.ParseRSSFeed(url)
			if err != nil {
				log.Printf("Error: %s\n", err)
				done <- true
				return
			}

			feedChan <- feed
			done <- true
		}(urls[i])
	}
	// Wait for async RSS parsing is complete
	go func() {
		for i := 0; i < n; i++ {
			<-done
		}
		close(feedChan)
	}()

	acc := []*gofeed.Feed{}
	for feed := range feedChan {
		acc = append(acc, feed)
	}

	log.Printf("A total of %d/%d successfully RSS feeds parsed.\n", len(acc), n)
	return acc
}

// Converts a list of RSS feed structs into a list of article structs
func FromRSSToArticle(feeds []*gofeed.Feed) []*Article {
	articles := []*Article{}

	for _, feed := range feeds {
		var source string = feed.Link
		if source == "" {
			source = feed.FeedLink
		}
		count := 0
		for _, item := range feed.Items {
			articles = append(
				articles,
				&Article{
					Title:         item.Title,
					Description:   item.Description,
					Link:          item.Link,
					Source:        urlSourceForHtml(source),
					PubDateParsed: item.PublishedParsed,
					PubDate:       FormatDateFromTime(item.PublishedParsed),
				},
			)
			count++
		}
		log.Printf("RSS url: %s; Article count: %d\n", source, count)
	}

	return articles
}

func urlSourceForHtml(url string) string {
	parts := strings.Split(url, "//")
	if len(parts) == 2 {
		return parts[1]
	}
	return url
}

// TimeIn returns the time in UTC if the name is "" or "UTC".
// It returns the local time if the name is "Local".
// Otherwise, the name is taken to be a location name in
// the IANA Time Zone database, such as "Africa/Lagos".
func TimeIn(t time.Time, name string) (time.Time, error) {
	loc, err := time.LoadLocation(name)
	if err == nil {
			t = t.In(loc)
	}
	return t, err
}

// https://cs.opensource.google/go/go/+/go1.17.5:src/time/format.go;l=91
func FormatDateFromTime(t *time.Time) string {
	pubTime, _ := TimeIn(*t, "America/Puerto_Rico")
	return pubTime.Format(time.RFC1123)
}
