package main

import (
	"fmt"

	"github.com/Yishen1011/blog_aggregator/internal/config"
)

func main() {
	cfg, read_err := config.ReadConfig()
	if read_err != nil {
		fmt.Printf("Failed to read config\n")
	}
	fmt.Printf("%s\n",cfg.DB_URL)

	write_err := config.WriteConfig(cfg)
	if write_err != nil {
		fmt.Printf("Failed to write username onto config\n")
	}
	fmt.Println("Writing config")

	cfg2, read2_err := config.ReadConfig()
	if read2_err != nil {
		fmt.Printf("Failed to read config\n")
	}
	fmt.Printf("%s\n",cfg2.DB_URL)
	fmt.Printf("%s\n",cfg2.Username)
}

