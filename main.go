package main

import (
	"context"
	"database/sql"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
	_ "github.com/lib/pq"
	"os"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	feed, err := FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return
	}

	fmt.Println(feed)

	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		fmt.Println(err)
	}

	dbQueries := database.New(db)

	s := state{dbQueries, &cfg}

	com := commands{make(map[string]func(*state, command) error)}
	com.register("login", handlerLogin)
	com.register("register", handlerRegister)
	com.register("reset", handlerReset)
	com.register("users", handlerUsers)

	err = com.run(&s, command{args[0], args[1:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
