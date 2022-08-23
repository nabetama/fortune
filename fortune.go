package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

const (
	ExitCodeOK        = 0
	ExitCodeError     = 1
	NotMatchedMessage = "Not matched"
)

type CLI struct {
	outStream, errStream io.Writer
}

func randomOne(ss []string) string {
	rand.Seed(time.Now().UnixNano())
	if len(ss) > 0 {
		return ss[rand.Intn(len(ss))]
	} else {
		return NotMatchedMessage
	}
}

func getFortunes(path, pattern string) ([]string, error) {
	lines, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var fortunes []string
	for _, fortune := range strings.Split(string(lines), "\n%\n") {
		if pattern == "" {
			fortunes = append(fortunes, fortune)
		}
		if pattern != "" && strings.Contains(fortune, pattern) {
			fortunes = append(fortunes, fortune)
		}
	}
	return fortunes, nil
}

func (c *CLI) Run(args []string) int {
	var pattern string
	flag.StringVar(&pattern, "m", "", "Print out all fortunes which match the basic regular expression pattern.")
	flag.Parse()

	fortunes, err := getFortunes("./datfiles/cookie", pattern)
	if err != nil {
		fmt.Errorf("%s", err.Error())
		return ExitCodeError
	}

	fmt.Fprintf(c.outStream, randomOne(fortunes))
	return ExitCodeOK
}
