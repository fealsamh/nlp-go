package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	fmt.Printf("%s '%s'\n", runtime.GOOS, os.Getenv("TEST_ARG"))
}
