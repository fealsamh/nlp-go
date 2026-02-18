package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/fealsamh/nlp-go/aibot"
)

func exitWithMessage(msg string, args ...any) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}

func exitWithError(err error) {
	fmt.Fprintln(os.Stderr, err.Error())
	os.Exit(1)
}

func getSecretKey() string {
	if sk := os.Getenv("AIBOT_SECRET_KEY"); sk != "" {
		return sk
	}
	dir, err := os.UserHomeDir()
	if err != nil {
		exitWithError(err)
	}
	fp := filepath.Join(dir, ".aibot", "sk")
	f, err := os.Open(fp)
	if err != nil {
		if os.IsNotExist(err) {
			exitWithMessage("no secret key provided")
		}
		exitWithError(err)
	}
	defer f.Close()
	b, err := ioutil.ReadAll(f)
	if err != nil {
		exitWithError(err)
	}
	return strings.TrimSpace(string(b))
}

func main() {
	var serviceURL, regExp string
	flag.StringVar(&serviceURL, "s", aibot.DefaultServiceURL, "service URL")
	flag.StringVar(&regExp, "r", "", "regular expression")
	flag.Parse()

	if flag.NArg() == 0 {
		exitWithMessage("no command provided")
	}
	cmd := flag.Arg(0)
	cl := &aibot.Client{ServiceURL: serviceURL}
	processed := true

	var re *regexp.Regexp
	if regExp != "" {
		var err error
		re, err = regexp.Compile(regExp)
		if err != nil {
			exitWithError(err)
		}
	}

	switch cmd {
	case "signin":
		signin(cl)
	case "signup":
		signup(cl)
	case "check":
		if flag.NArg() <= 1 {
			exitWithMessage("no input file provided")
		}
		check(flag.Arg(1), false)
	default:
		processed = false
	}
	if processed {
		return
	}

	cl.SecretKey = getSecretKey()
	switch cmd {
	case "whoami":
		whoami(cl)
	case "list":
		listBots(cl, re)
	case "get":
		if flag.NArg() <= 1 {
			exitWithMessage("no bot ID provided")
		}
		getBot(cl, flag.Arg(1))
	case "intent":
		if flag.NArg() <= 2 {
			exitWithMessage("a bot ID & an input sentence must be provided")
		}
		recogniseIntent(cl, flag.Arg(1), flag.Arg(2))
	case "deploy":
		if flag.NArg() <= 1 {
			exitWithMessage("no input file provided")
		}
		deploy(cl, flag.Arg(1))
	default:
		exitWithMessage("unknown command '%s'", cmd)
	}
}
