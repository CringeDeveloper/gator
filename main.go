package main

import (
	"fmt"
	"gator/internal/config"
	"os"
)

func main() {
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

	s := state{&cfg}

	com := commands{make(map[string]func(*state, command) error)}
	com.register("login", handlerLogin)

	err = com.run(&s, command{args[0], args[1:]})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
