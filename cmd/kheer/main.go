package main

import (
	"fmt"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	kheer = kingpin.New("kheer", "automation on top of kubernetes")
)

func main() {
	kingpin.Parse()
	fmt.Println("hello kheer")

}
