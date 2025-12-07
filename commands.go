package main

import (
	"errors"
	"fmt"
	"github.com/mr-rambling/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerLogin(s *state, cmd command) error {
	if cmd.name != "login" {
		return errors.New("wrong command type")
	}
	if len(cmd.args) == 0 {
		return errors.New("no username provided")
	}

	if err := config.SetUser(*s.cfg, cmd.args[0]); err != nil {
		return err
	}
	fmt.Println("username set")
	return nil
}

type commands struct {
	names map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {
	val, ok := c.names[cmd.name]
	if !ok {
		return errors.New("command does not exist")
	}
	return val(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.names[name] = f
}
