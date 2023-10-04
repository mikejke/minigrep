package lib

import (
	"errors"
	"fmt"
	"os"
)

func Run(config Config) error {
	content, err := os.ReadFile(config.Filename)

	if err != nil {
		return err
	}

	fmt.Printf("With text:\n%s", string(content))

	return nil
}

type Config struct {
	Query    string
	Filename string
}

func NewConfig(args []string) (*Config, error) {
	if len(os.Args) < 3 {
		return nil, errors.New("not enough arguments")
	}

	query := os.Args[1]
	filename := os.Args[2]

	return &Config{query, filename}, nil
}
