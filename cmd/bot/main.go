package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/fealsamh/nlp-go/aibot"
)

func exitWithMessage(msg string, args ...interface{}) {
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
	var serviceUrl string
	flag.StringVar(&serviceUrl, "s", aibot.DefaultServiceURL, "service URL")
	flag.Parse()
	if flag.NArg() == 0 {
		exitWithMessage("no command provided")
	}
	cmd := flag.Arg(0)
	cl := &aibot.Client{ServiceURL: serviceUrl}
	processed := true
	switch cmd {
	case "signin":
		signin(cl)
	case "signup":
		signup(cl)
	case "check":
		if flag.NArg() == 1 {
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
	case "get":
		if flag.NArg() == 1 {
			exitWithMessage("no bot ID provided")
		}
		getBot(cl, flag.Arg(1))
	case "deploy":
		if flag.NArg() == 1 {
			exitWithMessage("no input file provided")
		}
		deploy(cl, flag.Arg(1))
	default:
		exitWithMessage("unknown command '%s'", cmd)
	}
}
