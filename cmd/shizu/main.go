package main

import (
	"log"
	"net"
	"strconv"

	"github.com/vmfunc/shizu/pkg/config"
	"github.com/vmfunc/shizu/pkg/server"
)

func main() {
	config, err := config.LoadConfigFromFile("config.json")
	if err != nil {
		log.Fatalf("Failed to load config: %s\n", err)
	}

	listener, err := net.Listen("tcp", "0.0.0.0:"+strconv.Itoa(config.Port))
	if err != nil {
		log.Fatalf("Failed to listen on port %d: %s\n", config.Port, err)
	} else {
		log.Printf("Listening on port %d\n", config.Port)
	}

	for {
		nConn, err := listener.Accept()
		if err != nil {
			log.Printf("Failed to accept incoming connection: %s\n", err)
			continue
		} else {
			log.Printf("Accepted incoming connection from %s\n", nConn.RemoteAddr())
		}

		go server.HandleServerConn(nConn)
	}
}
