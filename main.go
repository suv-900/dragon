package main

import (
	"flag"
	"log"
	"path/filepath"

	"github.com/suv-900/dragon/workers"
)

func main() {
	flag.Parse()
	filename := flag.Arg(0)

	if filename == "" {
		log.Fatal("not enough arguments")

	}
	ext := filepath.Ext(filename)
	if ext == "" {
		filename += ".txt"
	}
	workers.AddtoFile(filename)
}
