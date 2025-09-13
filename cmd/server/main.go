package main

import (
	"countdown/internal/server"
	"flag"
)

const DEFAULT_PORT = 8080

type Options struct {
	Port int
}

func options() *Options {
	var port int
	flag.IntVar(&port, "port", DEFAULT_PORT, "The port on which the server should run")
	flag.Parse()
	return &Options{
		Port: port,
	}
}

func main() {
	var options = options()
	server.StartServer(options.Port)
}
