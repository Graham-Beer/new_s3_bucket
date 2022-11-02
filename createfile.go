package main

import (
	"fmt"
	"os"
)

func CreateLocalFile(file string, content string) (*os.File, error) {
	fmt.Printf("Creating file '%s'...\n", file)

	f, err := os.Create(file)
	if err != nil {
		return nil, err
	}

	fmt.Println("Adding content to our file...")

	if _, err := f.Write([]byte(content)); err != nil {
		return nil, err
	}

	if err := f.Close(); err != nil {
		return nil, err
	}

	fmt.Println("File and content completed.")
	return f, nil
}
