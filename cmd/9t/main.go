package main

import (
	"log"
	"os"

	ninetail "github.com/admpub/9t"
)

func main() {
	filenames := os.Args[1:]

	runner, err := ninetail.Runner(filenames, ninetail.Config{Colorize: true, ReOpen: true}) // TODO use flags!!
	if err != nil {
		log.Fatal(err)
	}
	runner.Run()
}
