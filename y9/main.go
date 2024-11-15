package main

import (
	"encoding/base64"
	"errors"
	"flag"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func main() {
	// Define flags
	key := flag.String("k", "", "Key to store in YAML")
	value := flag.String("v", "", "Value to store in YAML")
	encode := flag.Bool("e", false, "Base64 encode the value")
	decode := flag.Bool("d", false, "Base64 decode the value (prints to stdout only)")
	printKeyValue := flag.Bool("r", false, "Print key:value to stdout")
	file := flag.String("file", "output.yaml", "YAML file to write to")

	// Parse flags
	flag.Parse()

	// Validate input
	if *key == "" && !*decode && !*printKeyValue {
		fmt.Println("Error: -k flag is required unless using -d or -r")
		flag.Usage()
		os.Exit(1)
	}

	if *value == "" && !*decode && !*printKeyValue {
		fmt.Println("Error: -v flag is required unless using -d or -r")
		flag.Usage()
		os.Exit(1)
	}

	// Handle -d flag
	if *decode {
		decodedValue, err := base64Decode(*value)
		if err != nil {
			fmt.Printf("Error decoding value: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(decodedValue)
		return
	}

	// Handle -r flag
	if *printKeyValue {
		fmt.Printf("%s:%s\n", *key, *value)
		return
	}

	// Handle -e flag
	finalValue := *value
	if *encode {
		finalValue = base64Encode(*value)
	}

	// Write to YAML file
	if err := writeToYAML(*file, *key, finalValue); err != nil {
		fmt.Printf("Error writing to YAML: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Successfully wrote %s:%s to %s\n", *key, finalValue, *file)
}

// Base64 encode function
func base64Encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}

// Base64 decode function
func base64Decode(encoded string) (string, error) {
	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	return string(decoded), nil
}

// Write key-value pair to YAML file
func writeToYAML(filename, key, value string) error {
	// Read existing data if file exists
	data := make(map[string]string)
	if _, err := os.Stat(filename); err == nil {
		file, err := os.Open(filename)
		if err != nil {
			return err
		}
		defer file.Close()
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&data); err != nil && !errors.Is(err, os.ErrNotExist) {
			return err
		}
	}

	// Add or update key-value pair
	data[key] = value

	// Write data back to file
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := yaml.NewEncoder(file)
	defer encoder.Close()
	return encoder.Encode(data)
}

