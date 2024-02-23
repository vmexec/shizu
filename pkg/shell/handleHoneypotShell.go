package shell

import (
	"bufio"

	"golang.org/x/crypto/ssh"
)

func HandleHoneypotShell(s ssh.Channel) {
	defer s.Close()

	mw := bufio.NewWriter(s)

	r := bufio.NewReader(s)
	rw := bufio.NewReadWriter(r, mw)

	_, _ = s.SendRequest("shell", true, nil)

	shellSession := NewShellSession(rw, s)

	shellSession.Start()
}
