package main

import (
	"fmt"
	"gator/internal/config"
)

type commands struct {
	com map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.com[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	err := c.com[cmd.name](s, cmd)
	if err != nil {
		return err
	}

	return nil
}

type command struct {
	name    string
	handler []string
}

type state struct {
	cfg *config.Config
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.handler) == 0 {
		return fmt.Errorf("the login handler expects a single argument")
	}

	err := s.cfg.SetUser(cmd.handler[0])
	if err != nil {
		return err
	}

	fmt.Println("User has been set")

	return nil
}
