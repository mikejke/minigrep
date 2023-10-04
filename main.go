package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	config, err := NewConfig(os.Args)

	if err != nil {
		println("Problem parsing arguments:", err.Error())
		os.Exit(1)
	}

	fmt.Printf("Searching for %s\n", config.query)
	fmt.Printf("In file %s\n", config.filename)

	err = Run(*config)

	if err != nil {
		println("Application error", err.Error())
		os.Exit(1)
	}
}

func Run(config Config) error {
	content, err := os.ReadFile(config.filename)

	if err != nil {
		return err
	}

	fmt.Printf("With text:\n%s", string(content))

	return nil
}

type Config struct {
	query    string
	filename string
}

func NewConfig(args []string) (*Config, error) {
	if len(os.Args) < 3 {
		return nil, errors.New("not enough arguments")
	}

	query := os.Args[1]
	filename := os.Args[2]

	return &Config{query, filename}, nil
}
