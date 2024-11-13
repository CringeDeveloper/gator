package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"gator/internal/database"
	"github.com/google/uuid"
	"html"
	"io"
	"net/http"
	"time"
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

func agg(s *state, cmd command) error {
	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	feed.sanitize()

	fmt.Println(feed)

	return nil
}

func addFeed(s *state, cmd command) error {
	if len(cmd.handler) < 2 {
		return fmt.Errorf("the register handler expects two arguments")
	}
	title := cmd.handler[0]
	url := cmd.handler[1]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      title,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return err
	}

	feedFollowsParams := database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	}
	_, err = s.db.CreateFeedFollows(context.Background(), feedFollowsParams)
	if err != nil {
		return err
	}

	fmt.Println(feed)
	return nil
}

func feeds(s *state, cmd command) error {
	result, err := s.db.GetFeedsWithAuthor(context.Background())
	if err != nil {
		return err
	}

	for _, v := range result {
		fmt.Println(v.Name)
		fmt.Println(v.Url)
		fmt.Println(v.AuthorName.String)
	}

	return nil
}
