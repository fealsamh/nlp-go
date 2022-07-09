package main

import (
	"fmt"
	"strings"

	"github.com/fealsamh/nlp-go/aibot"
)

func whoami(cl *aibot.Client) {
	u, err := cl.Whoami()
	if err != nil {
		exitWithError(err)
	}
	fmt.Println("username:    ", u.Username)
	fmt.Println("full name:   ", u.Fullname)
	fmt.Println("email:       ", u.Email)
	fmt.Println("organisation:", u.OrgName)
	fmt.Println("roles:       ", strings.Join(u.Roles, ", "))
}
