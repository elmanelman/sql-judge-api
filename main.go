package main

import (
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
