package main

import (
	"log"
	"os"

	"github.com/Yishen1011/blog_aggregator/internal/config"
	"github.com/Yishen1011/blog_aggregator/internal/database"
	_ "github.com/lib/pq"
	"database/sql"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {

	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error failed to read config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)

	s := &state{
		cfg: &cfg,
		db: dbQueries,
	}

	cmds := &commands{
		registered: map[string]func(*state, command) error{},
	}

	cmds.register("login", handlerLogin)
	cmds.register("register", handlerRegister)
	cmds.register("reset", handlerReset)
	cmds.register("users", handlerUsers)
	cmds.register("agg", handlerAgg)
	cmds.register("addfeed", handlerAddFeed)
	cmds.register("feeds", handlerFeeds)

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