package main

import (
	"fmt"
	"os"

	"github.com/mr-rambling/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("error: %s", err)
	}

	var s state
	s.cfg = &cfg

	cmds := commands{
		names: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)

	args := os.Args

	if len(args) < 3 {
		fmt.Println("missing argument")
		os.Exit(1)
	}

	cmd := command{
		name: args[1],
		args: args[2:],
	}
	if err := cmds.run(&s, cmd); err != nil {
		fmt.Println(err)
	}
}
