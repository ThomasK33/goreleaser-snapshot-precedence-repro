package main

import (
	"flag"
	"log"
	"os"
)

func main() {
	version := flag.String("version", "0.0.0", "Version")
	flag.Parse()

	log.Printf("Version: %s", *version)

	os.WriteFile("./version.txt", []byte(*version), 0644)
}
