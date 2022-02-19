package main

import (
	"github.com/pocket/cli"
	"os"
)

func main() {
	defer os.Exit(0)

	cmd := cli.CommandLine{}

	cmd.Run()
}
