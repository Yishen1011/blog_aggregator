package main

import (
	"log"
	"os"

	"github.com/Yishen1011/blog_aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error failed to read config: %v", err)
	}

	s := &state{
		cfg: &cfg,
	}

	cmds := &commands{
		registered: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)

	arguments := os.Args

	if len(arguments) < 2 {
		log.Fatalf("Usage: cli <command> [args...]")
	}

	cmd := command{
		Name: arguments[1],
		Args: arguments[2:],
	}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}