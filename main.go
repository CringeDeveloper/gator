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

	c, _ := config.Read()
	s := state{&c}

	com := commands{make(map[string]func(*state, command) error)}
	com.register("login", handlerLogin)

	pass := command{args[0], args[1:]}

	err := com.run(&s, pass)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
