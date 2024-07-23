package main

import (
	"flag"

	"github.com/haguirrear/coffeeassistant/server/internal/server"
)

func main() {
	server := server.NewServer()
	flag.Parse()
	address := flag.Arg(0)
	_ = server.Start(address)
}
