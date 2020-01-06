package main

import (
	_ "github.com/lib/pq" // TODO review postgres drivers
	"log"
)

func main() {
	cfg, err := LoadConfig()
	if err != nil {
		log.Fatal("unable to load config:", err)
		return
	}

	s := NewServer(nil)
	s.Start(cfg)
}
