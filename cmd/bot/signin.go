package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"syscall"

	"github.com/fealsamh/nlp-go/aibot"
	"golang.org/x/term"
)

func signin(cl *aibot.Client) {
	r := bufio.NewReader(os.Stdin)
	fmt.Print("username: ")
	u, err := r.ReadString('\n')
	if err != nil {
		exitWithError(err)
	}
	u = strings.TrimSpace(u)
	fmt.Print("password: ")
	b, err := term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		exitWithError(err)
	}
	fmt.Print("\n")
	p := string(b)
	sk, on, err := cl.Signin(u, p)
	if err != nil {
		exitWithError(err)
	}
	fmt.Println("organisation:", on)
	dir, err := os.UserHomeDir()
	if err != nil {
		exitWithError(err)
	}
	dir = filepath.Join(dir, ".aibot")
	if _, err := os.Stat(dir); err != nil {
		if !os.IsNotExist(err) {
			exitWithError(err)
		}
		if err := os.Mkdir(dir, 0750); err != nil {
			exitWithError(err)
		}
	}
	f, err := os.OpenFile(filepath.Join(dir, "sk"), os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		exitWithError(err)
	}
	fmt.Fprintln(f, sk)
	if err := f.Close(); err != nil {
		exitWithError(err)
	}
	fmt.Println("signed in")
}
