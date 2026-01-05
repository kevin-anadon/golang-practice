package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// HTTP response example {"page":"words","input":"","words":["word1"]}

type Page struct {
	Name string `json:"page"`
}

type Occurrence struct {
	Words map[string]int `json:"words"`
}

type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

func main() {
	args := os.Args
	hostname := args[1]
	if len(args) < 2 {
		fmt.Printf("Usage: ./http-get <url>\n")
		os.Exit(1)
	}

	// Ignore the first argument with underscore _
	if _, err := url.ParseRequestURI(hostname); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", hostname)
		os.Exit(1)
	}

	response, err := http.Get(hostname)

	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	// You could execute it like this: go run main.go https://pokeapi.co/api/v2/pokemon/arcanine
	// You could execute it like this with custom API: go run main.go http://localhost:8080/words

	if response.StatusCode != 200 {
		fmt.Printf("Invalid output: status code %d\n%s\n", response.StatusCode, response.Body)
		os.Exit(1)
	}

	var page Page
	err = json.Unmarshal(body, &page)
	if err != nil {
		log.Fatal(err)
	}

	switch page.Name {
	case "words":
		var words Words
		err = json.Unmarshal(body, &words)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("HTTP Status Code: %d\nJSON Parsed Body\nPage: %s\nInput: %s\nWords: %s", response.StatusCode, words.Page, words.Input, strings.Join(words.Words, ", "))
	case "occurrence":
		var occurrence Occurrence
		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("HTTP Status Code: %d\nJSON Parsed Body\n==Ocurrences==\n", response.StatusCode)
		for key, value := range occurrence.Words {
			fmt.Printf("%s: %d\n", key, value)
		}
	default:
		fmt.Printf("Page not found\n")
	}
}
