package main

import (
	"context"
	"fmt"
	"gator/internal/database"
	"github.com/google/uuid"
	"log"
	"time"
)

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}

func (c *commands) run(s *state, com command) error {
	err := c.handlers[com.name](s, com)
	if err != nil {
		return err
	}

	return nil
}

type command struct {
	name    string
	handler []string
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.handler) != 1 {
		return fmt.Errorf("the login handler expects a single argument")
	}
	userName := cmd.handler[0]

	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		log.Fatal(err)
	}

	err = s.cfg.SetUser(cmd.handler[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")
	return nil
}

func handlerRegister(s *state, cmd command) error {
	if len(cmd.handler) != 1 {
		return fmt.Errorf("the register handler expects a single argument")
	}
	userName := cmd.handler[0]

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      userName}

	user, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		log.Fatal(err)
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Println("User was created")
	log.Println(user)

	return nil
}

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteAllUsers(context.Background())
	return err
}

func handlerUsers(s *state, cmd command) error {
	u, err := s.db.GetAllUsers(context.Background())
	if err != nil {
		return err
	}

	for _, v := range u {
		if v.Name == s.cfg.CurrentUserName {
			fmt.Println(v.Name, "(current)")
		} else {
			fmt.Println(v.Name)
		}
	}

	return nil
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

	fmt.Println(feed)
	return nil
}
