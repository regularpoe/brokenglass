package main

import (
	"database/sql"
	"encoding/base64"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type LanguageCommands map[string]map[string]string

func getDBPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Printf("Error getting home directory: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(homeDir, "bento.db")
}

func initDB(dbPath string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS snippets (
			language TEXT NOT NULL,
			name TEXT NOT NULL,
			content TEXT NOT NULL,
			content_encoded TEXT,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			PRIMARY KEY (language, name)
		)
	`)
	return db, err
}

func addSnippet(db *sql.DB, language, name, content string, isEncoded bool) error {
	var decodedContent, encodedContent string

	if isEncoded {
		// If content is base64 encoded, decode it and store both versions
		decoded, err := base64.StdEncoding.DecodeString(content)
		if err != nil {
			return fmt.Errorf("error decoding base64 content: %v", err)
		}
		decodedContent = string(decoded)
		encodedContent = content
	} else {
		// If content is plain text, encode it and store both versions
		decodedContent = content
		encodedContent = base64.StdEncoding.EncodeToString([]byte(content))
	}

	_, err := db.Exec(`
		INSERT OR REPLACE INTO snippets
		(language, name, content, content_encoded, updated_at)
		VALUES (?, ?, ?, ?, CURRENT_TIMESTAMP)`,
		strings.ToLower(language),
		strings.ToLower(name),
		decodedContent,
		encodedContent,
	)
	return err
}

func exportJSON(db *sql.DB, outputFile string, exportEncoded bool) error {
	var query string
	if exportEncoded {
		query = "SELECT language, name, content_encoded as content FROM snippets"
	} else {
		query = "SELECT language, name, content FROM snippets"
	}

	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	commands := make(LanguageCommands)
	for rows.Next() {
		var language, name, content string
		if err := rows.Scan(&language, &name, &content); err != nil {
			return err
		}

		if commands[language] == nil {
			commands[language] = make(map[string]string)
		}
		commands[language][name] = content
	}

	file, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "    ")
	return encoder.Encode(commands)
}

func listSnippets(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT language, name, content, content_encoded, created_at, updated_at
		FROM snippets
		ORDER BY language, name`)
	if err != nil {
		return err
	}
	defer rows.Close()

	currentLang := ""
	for rows.Next() {
		var lang, name, content, contentEncoded, createdAt, updatedAt string
		if err := rows.Scan(&lang, &name, &content, &contentEncoded, &createdAt, &updatedAt); err != nil {
			return err
		}

		if lang != currentLang {
			if currentLang != "" {
				fmt.Println()
			}
			fmt.Printf("Language: %s\n", lang)
			fmt.Println(strings.Repeat("-", 40))
			currentLang = lang
		}

		fmt.Printf("  Snippet: %s\n", name)
		fmt.Printf("  Created: %s\n", createdAt)
		fmt.Printf("  Updated: %s\n", updatedAt)
		fmt.Printf("  Content: %s\n", content)
		fmt.Println()
	}
	return nil
}

func main() {
	exportFlag := flag.String("export", "", "Export snippets to JSON file")
	exportEncodedFlag := flag.Bool("export-encoded", false, "Export base64 encoded snippets")
	encodedFlag := flag.Bool("encoded", false, "Input is base64 encoded")
	listFlag := flag.Bool("list", false, "List all snippets")
	flag.Parse()
	args := flag.Args()

	dbPath := getDBPath()
	db, err := initDB(dbPath)
	if err != nil {
		fmt.Printf("Error initializing database: %v\n", err)
		os.Exit(1)
	}
	defer db.Close()

	if *listFlag {
		if err := listSnippets(db); err != nil {
			fmt.Printf("Error listing snippets: %v\n", err)
			os.Exit(1)
		}
		return
	}

	if *exportFlag != "" {
		if err := exportJSON(db, *exportFlag, *exportEncodedFlag); err != nil {
			fmt.Printf("Error exporting to JSON: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Snippets exported to %s\n", *exportFlag)
		return
	}

	if len(args) < 3 {
		fmt.Println("Usage:")
		fmt.Println("  Add snippet:           bentoi <language> <snippet_name> <snippet_content>")
		fmt.Println("  Add encoded snippet:   bentoi -encoded <language> <snippet_name> <base64_content>")
		fmt.Println("  Export to JSON:        bentoi -export <filename.json>")
		fmt.Println("  Export encoded:        bentoi -export <filename.json> -export-encoded")
		fmt.Println("  List all snippets:     bentoi -list")
		fmt.Println("\nExample:")
		fmt.Println("  bentoi ruby read_file 'def read_file(path); end'")
		fmt.Println("  bentoi -encoded python write_file 'ZGVmIHdyaXRlX2ZpbGUocGF0aCwgY29udGVudCk6Cg=='")
		os.Exit(1)
	}

	language := args[0]
	name := args[1]
	content := strings.Join(args[2:], " ")

	if err := addSnippet(db, language, name, content, *encodedFlag); err != nil {
		fmt.Printf("Error adding snippet: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Added snippet '%s' for language '%s'\n", name, language)
}

