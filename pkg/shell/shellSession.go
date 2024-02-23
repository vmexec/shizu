package shell

import (
	"io"
	"log"

	"golang.org/x/crypto/ssh"
	"golang.org/x/term"
)

type ShellSession struct {
	term *term.Terminal
	s    ssh.Channel
}

func NewShellSession(w io.ReadWriter, s ssh.Channel) *ShellSession {
	term := term.NewTerminal(w, "> ")
	return &ShellSession{term: term, s: s}
}

func (ss *ShellSession) Start() {
	for {
		line, err := ss.term.ReadLine()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Printf("Failed to read line: %s\n", err)
			return
		}

		_, err = ss.s.Write([]byte(line + "\r\n"))
		if err != nil {
			log.Printf("Failed to write to SSH channel: %s\n", err)
			return
		}
	}
}
