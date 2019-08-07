package main

import (
	"log"
	"os"

	"github.com/admpub/tail"
	ninetail "github.com/cynix/9t"
)

func main() {
	filenames := os.Args[1:]

	runner, err := ninetail.Runner(filenames, ninetail.Config{Config: &tail.Config{Follow: true, ReOpen: true}, Colorize: true}) // TODO use flags!!
	if err != nil {
		log.Fatal(err)
	}
	runner.Run()
}
