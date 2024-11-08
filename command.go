package main

import (
	"fmt"
	"gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name    string
	handler []string
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
