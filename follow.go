package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func follow(s *state, cmd command) error {
	if len(cmd.handler) < 1 {
		return fmt.Errorf("the follow  expects one arguments")
	}
	url := cmd.handler[0]
	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}
	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
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

	feedRes, err := s.db.CreateFeedFollows(context.Background(), feedFollowsParams)
	if err != nil {
		return err
	}
	fmt.Println(feedRes.FeedName, feedRes.UserName)

	return nil
}

func following(s *state, cmd command) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	for _, v := range feeds {
		fmt.Println(v.FeedsName)
	}

	return nil
}

func unfollow(s *state, cmd command) error {
	if len(cmd.handler) < 1 {
		return fmt.Errorf("the unfollow expects one arguments")
	}
	url := cmd.handler[0]

	user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return err
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), url)
	if err != nil {
		return err
	}

	err = s.db.Unfollow(context.Background(), database.UnfollowParams{FeedID: feed.ID, UserID: user.ID})
	if err != nil {
		return err
	}

	return nil
}
