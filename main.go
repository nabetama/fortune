package main

import (
	"os"
)

const name = "fortune"

const version = "0.0.1"

var revision = "HEAD"

func main() {
	cli := &CLI{
		outStream: os.Stdout,
		errStream: os.Stdout,
	}

	os.Exit(cli.Run(os.Args))
}