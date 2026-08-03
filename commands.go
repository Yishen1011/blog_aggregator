package main

import "errors"

type command struct {
	Name  string
	Args  []string 
}

type commands struct {
	registered map[string]func(*state, command) error
}

func (c *commands) run(s *state, cmd command) error {

	function, exist := c.registered[cmd.Name]
	if !exist {
		return errors.New("command does not exist")
	}

	return function(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registered[name] = f
}