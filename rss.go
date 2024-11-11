package main

import (
	"context"
	"encoding/xml"
	"html"
	"io"
	"net/http"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedUrl string) (*RSSFeed, error) {
	request, err := http.NewRequestWithContext(ctx, http.MethodGet, feedUrl, nil)
	if err != nil {
		return nil, err
	}
	request.Header.Add("User-Agent", "gator")
	resp, err := http.DefaultClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var rss RSSFeed

	if err := xml.Unmarshal(data, &rss); err != nil {
		return nil, err
	}

	return &rss, nil
}

func (rss *RSSFeed) sanitize() {
	html.UnescapeString(rss.Channel.Title)
	html.UnescapeString(rss.Channel.Description)
	for _, v := range rss.Channel.Item {
		html.UnescapeString(v.Title)
		html.UnescapeString(v.Description)
	}
}
