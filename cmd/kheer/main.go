package main

import (
	// "fmt"
	"os"

	"github.com/kheer/kheer/pkg/version"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	kheer     = kingpin.New("kheer", "Automation on top of kubernetes")
	bootstrap = kheer.Command("bootstrap", "Does initial setup.")
	server    = kheer.Command("server", "Runs kheer server.")
)

func main() {

	kheer.Version(version.Version)
	args := os.Args[1:]

	switch kingpin.MustParse(kheer.Parse(args)) {

	case bootstrap.FullCommand():
		println("bootstraping...")
	case server.FullCommand():
		println("starting server...")
	default:
		kheer.Usage(args)
	}
}
