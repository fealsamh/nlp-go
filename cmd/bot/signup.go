package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"syscall"

	"github.com/fealsamh/nlp-go/aibot"
	"golang.org/x/term"
)

func signup(cl *aibot.Client) {
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

	fmt.Print("password (again): ")
	b, err = term.ReadPassword(int(syscall.Stdin))
	if err != nil {
		exitWithError(err)
	}
	fmt.Print("\n")
	p2 := string(b)
	if p != p2 {
		exitWithMessage("the passwords don't match")
	}

	fmt.Print("full name: ")
	fn, err := r.ReadString('\n')
	if err != nil {
		exitWithError(err)
	}
	fn = strings.TrimSpace(fn)

	fmt.Print("email: ")
	e, err := r.ReadString('\n')
	if err != nil {
		exitWithError(err)
	}
	e = strings.TrimSpace(e)

	fmt.Print("organisation: ")
	o, err := r.ReadString('\n')
	if err != nil {
		exitWithError(err)
	}
	o = strings.TrimSpace(o)

	if err := cl.Signup(u, fn, e, p, o); err != nil {
		exitWithError(err)
	}
	fmt.Println("user created")
}
