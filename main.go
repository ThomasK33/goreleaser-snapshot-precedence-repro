package main

import (
	_ "embed"
	"log"
)

var (
	Version string = "0.0.0"
)

//go:embed version.txt
var file string

func main() {
	log.Printf("Version: %s", Version)

	log.Printf("File version: %s", file)
}
