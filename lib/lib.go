package lib

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Run(config Config) error {
	content, err := os.ReadFile(config.Filename)

	if err != nil {
		return err
	}

	var result []string

	if config.CaseSensitive {
		result = Search(config.Query, string(content))
	} else {
		result = SearchCaseInsensitive(config.Query, string(content))
	}

	for _, line := range result {
		fmt.Println(line)
	}

	return nil
}

type Config struct {
	Query         string
	Filename      string
	CaseSensitive bool
}

func NewConfig(args []string) (*Config, error) {
	if len(os.Args) < 3 {
		return nil, errors.New("not enough arguments")
	}

	query := os.Args[1]
	filename := os.Args[2]

	_, caseSensitive := os.LookupEnv("CASE_INSENSITIVE")

	return &Config{query, filename, !caseSensitive}, nil
}

func Search(query string, content string) []string {
	result := []string{}

	for _, line := range strings.Split(content, "\n") {
		if strings.Contains(line, query) {
			result = append(result, line)
		}
	}

	return result
}

func SearchCaseInsensitive(query string, content string) []string {
	result := []string{}
	query = strings.ToLower(query)

	for _, line := range strings.Split(content, "\n") {
		if strings.Contains(strings.ToLower(line), query) {
			result = append(result, line)
		}
	}

	return result
}
