package main

import (
	"log"
	"net"

	"github.com/vmfunc/shizu/pkg/server"
)

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:2222")
	if err != nil {
		log.Fatalf("Failed to listen on port 2222: %s\n", err)
	} else {
		log.Println("Listening on port 2222")
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
