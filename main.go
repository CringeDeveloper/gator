package main

import (
	"fmt"
	"gator/internal/config"
)

func main() {
	c, _ := config.Read()
	c.SetUser("Vlad") // login
	c, _ = config.Read()
	fmt.Println(c)
}
