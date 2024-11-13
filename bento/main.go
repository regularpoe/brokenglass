package main

import (
	_ "embed"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strings"
)

//go:embed data.json
var jsonData []byte

type LanguageCommands map[string]map[string]string

func main() {
	rawOutput := flag.Bool("raw", false, "Output only the raw snippet content")
	flag.Parse()
	args := flag.Args()

	if len(args) < 1 {
		fmt.Println("Usage: bento [--raw] <language> [snippet]")
		fmt.Println("Example: bento ruby read_file")
		fmt.Println("         bento --raw ruby read_file")
		os.Exit(1)
	}

	var commands LanguageCommands
	if err := json.Unmarshal(jsonData, &commands); err != nil {
		fmt.Printf("Error parsing JSON: %v\n", err)
		os.Exit(1)
	}

	language := strings.ToLower(args[0])
	langData, exists := commands[language]

	if !exists {
		fmt.Printf("Language '%s' not found\n", language)
		fmt.Println("Available languages:")
		for lang := range commands {
			fmt.Printf("- %s\n", lang)
		}
		os.Exit(1)
	}

	if len(args) == 1 {
		if *rawOutput {
			fmt.Println("Error: --raw flag requires an exact snippet name")
			os.Exit(1)
		}
		fmt.Printf("Available snippets for %s:\n", language)
		for cmd, content := range langData {
			fmt.Printf("\n%s:\n%s\n", cmd, content)
		}
		return
	}

	searchTerm := args[1]

	if *rawOutput {
		if content, exists := langData[searchTerm]; exists {
			fmt.Print(content)
			return
		}

		fmt.Printf("No exact match found for '%s'\n", searchTerm)
		fmt.Println("\nAvailable snippets:")
		for cmd := range langData {
			fmt.Printf("- %s\n", cmd)
		}
		os.Exit(1)
	} else {
		found := false
		var partialMatches []string

		if content, exists := langData[searchTerm]; exists {
			fmt.Printf("\n%s:\n%s\n", searchTerm, content)
			return
		}

		for cmd := range langData {
			if strings.Contains(strings.ToLower(cmd), strings.ToLower(searchTerm)) {
				partialMatches = append(partialMatches, cmd)
				found = true
			}
		}

		if found {
			fmt.Printf("No exact match found for '%s', but found similar snippets:\n", searchTerm)
			for _, cmd := range partialMatches {
				fmt.Printf("- %s\n", cmd)
			}
			return
		}

		fmt.Printf("No snippets found matching '%s' for language '%s'\n", searchTerm, language)
		fmt.Println("\nAvailable snippets:")
		for cmd := range langData {
			fmt.Printf("- %s\n", cmd)
		}
		os.Exit(1)
	}
}

