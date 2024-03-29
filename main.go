package main

import (
	"fmt"
	"log"
	"os"

	lib "github.com/mikejke/minigrep/lib"
)

func main() {
	config, err := lib.NewConfig(os.Args)

	if err != nil {
		log.Fatalln("Problem parsing arguments:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Searching for %s\n", config.Query)
	fmt.Printf("In file %s\n", config.Filename)

	err = lib.Run(*config)

	if err != nil {
		log.Fatalln("Application error", err.Error())
		os.Exit(1)
	}
}
