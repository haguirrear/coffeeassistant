package main

import (
	"flag"
	"log"

	"github.com/haguirrear/coffeeassistant/server/internal/server"
)

func main() {
	server := server.NewServer()
	flag.Parse()
	address := flag.Arg(0)
	log.Printf("Args: %#v", flag.Args())
	_ = server.Start(address)
}
