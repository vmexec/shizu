package server

import (
	"log"
	"net"
	"os"
	"sync"

	"github.com/vmfunc/shizu/pkg/auth"
	"github.com/vmfunc/shizu/pkg/shell"
	"golang.org/x/crypto/ssh"
)

func HandleServerConn(nConn net.Conn) {
	privateKeyPath := os.Getenv("HOME") + "/.ssh/id_rsa"
	privateBytes, err := os.ReadFile(privateKeyPath)
	if err != nil {
		log.Fatal("Failed to load private key")
	}

	private, err := ssh.ParsePrivateKey(privateBytes)
	if err != nil {
		log.Fatal("Failed to parse private key")
	}

	config := &ssh.ServerConfig{
		PasswordCallback: func(c ssh.ConnMetadata, pass []byte) (*ssh.Permissions, error) {
			log.Printf("Login attempt: user=%s, pass=%s\n", c.User(), string(pass))
			if auth.ValidateUser(c.User(), string(pass)) {
				log.Printf("User %s authenticated\n", c.User())
				return nil, nil
			} else {
				log.Printf("User %s non-authenticated: wrong password\n", c.User())
				return nil, ssh.ErrNoAuth
			}
		},
	}

	config.AddHostKey(private)

	_, chans, reqs, err := ssh.NewServerConn(nConn, config)

	if err != nil {
		log.Printf("Failed to establish SSH connection: %s\n", err)
		return
	}

	var wg sync.WaitGroup
	defer wg.Wait()

	wg.Add(1)
	go func() {
		ssh.DiscardRequests(reqs)
		wg.Done()
	}()

	for newChannel := range chans {
		if newChannel.ChannelType() != "session" {
			newChannel.Reject(ssh.UnknownChannelType, "unknown channel type")
			continue
		}
		channel, requests, err := newChannel.Accept()
		if err != nil {
			log.Fatalf("Could not accept channel: %v", err)
		}

		wg.Add(1)
		go func(in <-chan *ssh.Request) {
			for req := range in {
				req.Reply(req.Type == "shell", nil)
			}
			wg.Done()
		}(requests)

		wg.Add(1)
		go func() {
			defer func() {
				term := shell.NewShellSession(channel, channel)
				term.Start()
				channel.Close()
				wg.Done()
			}()
		}()
	}
}
