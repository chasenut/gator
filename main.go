package main

import (
	"fmt"
	"log"

	"github.com/chasenut/rss-feed-aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)
	err = cfg.SetUser("chasenut")
	if err != nil {
		log.Fatalf("error writing to config: %v", err)
	}
	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("Read config: %+v\n", cfg)

}
